package finances

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/shav/telegram-bot/internal/caching"
	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/db"
	"github.com/shav/telegram-bot/internal/common/multi_error"
	"github.com/shav/telegram-bot/internal/common/serialization"
	"github.com/shav/telegram-bot/internal/modules"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/client"
	"github.com/shav/telegram-bot/internal/modules/finances/caches"
	"github.com/shav/telegram-bot/internal/modules/finances/clients"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/add_spending_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/change_currency_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/set_spend_limit_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/show_currency_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/show_spend_limit_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/spendings_report_command"
	"github.com/shav/telegram-bot/internal/modules/finances/config"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/reports/spendings"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases"
	finance_services_currency_convert "github.com/shav/telegram-bot/internal/modules/finances/services/currency/convert"
	finance_services_currency_rates "github.com/shav/telegram-bot/internal/modules/finances/services/currency/rates"
	finance_services_currency_settings "github.com/shav/telegram-bot/internal/modules/finances/services/currency/settings"
	"github.com/shav/telegram-bot/internal/modules/finances/services/spendings"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/currency_rates"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/spendings"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/user_settings/currencies"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/user_settings/spend_limits"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue/handlers"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/transport/message_queue"
)

// useCases описывает пользовательские сценарии модуля финансов.
type useCases interface {
	// GetUserCurrency возвращает текущую валюту пользователя userId.
	GetUserCurrency(ctx context.Context, userId int64) (finance_models.Currency, error)
	// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
	ChangeCurrency(ctx context.Context, userId int64, newCurrency finance_models.Currency) error
	// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period.
	GetSpendLimit(ctx context.Context, userId int64, period date.Month) (limit finance_models.Amount, exists bool, err error)
	// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
	// Возаращает размер установленного лимита в валюте пользователя.
	SetSpendLimit(ctx context.Context, userId int64, limit decimal.Decimal, period date.Month) (finance_models.Amount, error)
	// RequestSpendingReport запрашивает формирование отчёта по тратам пользователя userId
	// за указанный период времени period, сгруппированный по категориям.
	RequestSpendingReport(ctx context.Context, userId int64, period date.Period) error
	// AddUserSpending добавляет трату spending пользователя userId.
	// Возаращает размер траты в валюте пользователя.
	AddUserSpending(ctx context.Context, userId int64, spending finance_models.Spending) (finance_models.Amount, error)
}

// Модуль для работы с финансами
type financesModule struct {
	// База данных.
	db *sql.DB
	// Сервис кеширования.
	cache *caching.RedisCache
	// Подключение к сервису отправки отчётов.
	reportSenderConn *grpc.ClientConn
	// Отправитель сообщений в очередь.
	messageQueueProducer *message_queue.MessageQueueProducer
	// Получатель сообщений очереди.
	messageQueueConsumer *message_queue.MessageQueueConsumer
	// Пользовательские сценарии.
	useCases useCases
	// Валютные настройки пользователей.
	currencySettings *finance_services_currency_settings.CurrencySettingsService
	// Хранилище трат пользователей.
	spendingStorage *finance_storages_spendings.SpendingsDbStorage
	// Конвертер курсов валют.
	currencyConverter *finance_services_currency_convert.CurrencyConvertService
}

// NewModule создаёт модуль для работы с финансами.
func NewModule() *financesModule {
	return &financesModule{}
}

// GetName возвращает имя модуля.
func (m *financesModule) GetName() string {
	return "finances"
}

// Init инициализирует модуль финансов.
func (m *financesModule) Init(args modules.ModuleInitArgs) error {
	cfg, err := finance_config.New(args.ConfigType, args.ServiceName, args.ConfigFile)
	if err != nil {
		return errors.Wrap(err, "config init failed")
	}

	defaultCurrency, err := finance_models.ParseCurrency(cfg.DefaultCurrency())
	if err != nil {
		return errors.Wrap(err, "parse default currency failed")
	}

	dbConnectionString := args.Config.DbConnectionString()
	database, err := db.ConnectToDatabase(db.PostgresDriver, dbConnectionString)
	if err != nil {
		return errors.Wrap(err, "connect to database failed")
	}

	userCurrencies, err := finance_storages_user_currency_settings.NewDbStorageFor(database)
	if err != nil {
		return errors.Wrap(err, "Create UserCurrenciesDbStorage failed")
	}
	m.currencySettings, err = finance_services_currency_settings.NewSettingService(defaultCurrency, userCurrencies)
	if err != nil {
		return errors.Wrap(err, "Create CurrencySettingsService failed")
	}

	currencyRatesStorage, err := finance_storages_currency_rates.NewDbStorageFor(database)
	if err != nil {
		return errors.Wrap(err, "Create CurrencyRatesDbStorage failed")
	}

	currencyRates, err := finance_services_currency_rates.NewExchangeRateService(
		finance_clients.NewCurrencyRateClient(),
		currencyRatesStorage,
		finance_models.AllCurrencies[:],
		defaultCurrency,
		cfg.CurrencyRatesUpdatePeriod(),
		func(defaultCurrency finance_models.Currency, currency finance_models.Currency, loader finance_services_currency_rates.CurrencyRateLoader,
			cache finance_services_currency_rates.CurrencyRateCache, updatePeriod time.Duration) finance_services_currency_rates.CurrencyRateUpdater {
			updater, error := finance_services_currency_rates.NewExchangeRatesUpdater(defaultCurrency, currency, loader, cache, updatePeriod)
			if error != nil {
				logger.Error(args.Ctx, "Create CurrencyExchangeRatesUpdater failed", logger.Fields.Error(err))
				return nil
			}
			return updater
		})
	if err != nil {
		return errors.Wrap(err, "Create CurrencyRatesService failed")
	}

	m.currencyConverter, err = finance_services_currency_convert.NewConvertService(m.currencySettings, currencyRates)
	if err != nil {
		return errors.Wrap(err, "Create CurrencyConvertService failed")
	}

	spendLimitSettings, err := finance_storages_user_spend_limit_settings.NewDbStorageFor(database)
	if err != nil {
		return errors.Wrap(err, "Create UserSpendLimitDbStorage failed")
	}

	var spendingReportsCache finance_services_spendings.SpendingReportsCache
	cacheConnectionString := args.Config.CacheConnectionString()
	if cacheConnectionString != "" {
		m.cache, err = caching.NewRedisCache(args.Ctx, "SpendingReportsCache", cacheConnectionString, true)
		if err != nil {
			return errors.Wrap(err, "Create cache implementation for SpendingReportsCache failed")
		}
		spendingReportsCache, err = finance_caches_spending_reports.NewCache(m.cache, finance_commands_spendings_report.GetSpendingReportPeriods())
		if err != nil {
			return errors.Wrap(err, "Create SpendingReportsCache failed")
		}
	}

	m.spendingStorage, err = finance_storages_spendings.NewDbStorageFor(database)
	if err != nil {
		return errors.Wrap(err, "Create SpendingsDbStorage failed")
	}

	spendingService, err := finance_services_spendings.NewService(m.spendingStorage, spendingReportsCache)
	if err != nil {
		return errors.Wrap(err, "Create SpendingService failed")
	}

	m.messageQueueProducer, err = message_queue.NewProducer(args.Ctx, args.AppName, args.Config.MessageQueueBrokers())
	if err != nil {
		return errors.Wrap(err, "Create MessageQueueSender failed")
	}
	reportsClient, err := finance_clients.NewReportsClient(m.messageQueueProducer, serialization.NewJsonSerializer())
	if err != nil {
		return errors.Wrap(err, "Create ReportsClient failed")
	}

	transactionManager := db.NewTransactionManager(database)
	useCases, err := finances.NewUseCases(defaultCurrency, m.currencySettings, spendLimitSettings,
		spendingService, m.currencyConverter, reportsClient, transactionManager)
	if err != nil {
		return errors.Wrap(err, "Create UsesCases failed")
	}
	m.useCases = useCases

	currencyRates.StartMonitoringRates(args.Ctx)
	return nil
}

// InitCommands выполняет инициализацию команд модуля финансов.
func (m *financesModule) InitCommands(args modules.ModuleInitArgs) error {
	commands := args.Commands

	// Команда "Добавить трату"
	commands.RegisterCommand(finance_commands_add_spending.Metadata,
		func(command core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return finance_commands_add_spending.NewHandler(command, userId, m.useCases)
		})

	// Команда "Получить отчёт по тратам"
	commands.RegisterCommand(finance_commands_spendings_report.Metadata,
		func(command core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return finance_commands_spendings_report.NewHandler(command, userId, m.useCases)
		})

	// Команда "Сменить валюту"
	commands.RegisterCommand(finance_commands_change_currency.Metadata,
		func(command core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return finance_commands_change_currency.NewHandler(command, userId, m.useCases)
		})

	// Команда "Показать текущую валюту"
	commands.RegisterCommand(finance_commands_show_currency.Metadata,
		func(command core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return finance_commands_show_currency.NewHandler(command, userId, m.useCases)
		})

	// Команда "Установить лимит трат"
	commands.RegisterCommand(finance_commands_set_spend_limit.Metadata,
		func(command core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return finance_commands_set_spend_limit.NewHandler(command, userId, m.useCases)
		})

	// Команда "Показать лимит трат"
	commands.RegisterCommand(finance_commands_show_spend_limit.Metadata,
		func(command core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return finance_commands_show_spend_limit.NewHandler(command, userId, m.useCases)
		})

	return nil
}

// InitMessageQueueHandlers выполняет инициализацию обработчиков сообщений из очереди.
func (m *financesModule) InitMessageQueueHandlers(args modules.ModuleInitArgs) error {
	var err error
	m.messageQueueConsumer, err = message_queue.NewConsumer(args.Ctx, args.AppName, args.Config.MessageQueueBrokers(),
		args.ServiceName, message_queue.BalanceStrategyRoundRobin)
	if err != nil {
		return errors.Wrap(err, "Create MessageQueueConsumer failed")
	}

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(m.currencySettings, m.spendingStorage, m.currencyConverter)
	if err != nil {
		return errors.Wrap(err, "Create SpendingReportBuilder failed")
	}

	m.reportSenderConn, err = grpc.Dial(args.Config.ReportSenderAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "Establish connection to ReportSendService failed")
	}

	reportSendClient, err := core_transport_grpc.NewReportSendClient(m.reportSenderConn)
	if err != nil {
		return errors.Wrap(err, "Create ReportSendClient failed")
	}

	deserializer := serialization.NewJsonSerializer()
	spendingReportRequestHandler, err := finance_transport_mq_handlers.NewSpendingReportMessageHandler(reportBuilder, reportSendClient, deserializer)
	if err != nil {
		return errors.Wrap(err, "Create SpendingReportMessageHandler failed")
	}

	err = m.messageQueueConsumer.Subscribe(args.Ctx, finance_transport_mq.ReportsQueueName, spendingReportRequestHandler)
	if err != nil {
		return errors.Wrap(err, "Subscribe to SpendingReportRequestMessage message queue failed")
	}

	return nil
}

// Stop завершает работу модуля.
func (m *financesModule) Stop() error {
	var err error
	if m.db != nil {
		err = multi_error.Append(err, m.db.Close())
	}
	if m.cache != nil {
		err = multi_error.Append(err, m.cache.Close())
	}
	if m.messageQueueProducer != nil {
		err = multi_error.Append(err, m.messageQueueProducer.Close())
	}
	if m.messageQueueConsumer != nil {
		err = multi_error.Append(err, m.messageQueueConsumer.Close())
	}
	if m.reportSenderConn != nil {
		err = multi_error.Append(err, m.reportSenderConn.Close())
	}
	m.useCases = nil
	m.currencySettings = nil
	m.spendingStorage = nil
	m.currencyConverter = nil
	return err
}

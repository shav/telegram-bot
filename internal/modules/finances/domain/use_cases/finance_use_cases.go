//go:generate minimock -i spendingStorage -o ./mocks/ -s ".go"
//go:generate minimock -i currencySettings -o ./mocks/ -s ".go"
//go:generate minimock -i currencyConverter -o ./mocks/ -s ".go"
//go:generate minimock -i reportsBuilder -o ./mocks/ -s ".go"

package finances

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/multi_error"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyAmount = finance_models.Amount{}

// currencySettings хранит валютные настройки.
type currencySettings interface {
	// GetCurrency возвращает текущую валюту для пользователя userId.
	GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (finance_models.Currency, error)
	// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
	ChangeCurrency(ctx context.Context, ts tr.Transaction, userId int64, newCurrency finance_models.Currency) error
}

// spendLimitSettings хранит настройки бюджетных лимитов.
type spendLimitSettings interface {
	// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period.
	GetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, period date.Month) (limit decimal.Decimal, exists bool, err error)
	// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
	SetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, limit decimal.Decimal, period date.Month) error
}

// spendingStorage является хранилищем трат.
type spendingStorage interface {
	// AddSpending добавляет информацию о трате spending пользователя userId в хранилище.
	AddSpending(ctx context.Context, ts tr.Transaction, userId int64, spending finance_models.Spending) error
	// GetSpendingsAmount возвращает общий размер трат по всем категориям пользователя userId за указанный промежуток времени interval.
	GetSpendingsAmount(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (decimal.Decimal, error)
	// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
	// за указанный промежуток времени interval, сгруппированный по категориям.
	GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error)
}

// currencyConverter выполняет конвертацию валют.
type currencyConverter interface {
	// ConvertToDefaultCurrency конвертирует денежную сумму amount из валюты пользователя в основную валюту приложения.
	ConvertToDefaultCurrency(ctx context.Context, userId int64, amount decimal.Decimal) (decimal.Decimal, error)
	// ConvertToUserCurrency конвертирует денежную сумму amount из основной валюты приложения в валюту пользователя userId.
	ConvertToUserCurrency(ctx context.Context, userId int64, defaultAmount decimal.Decimal) (decimal.Decimal, error)
	// ConvertSpendingsTableToUserCurrency конвертирует таблицу расходов spendingsTable в валюту пользователя userId.
	ConvertSpendingsTableToUserCurrency(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) (finance_models.SpendingsByCategoryTable, error)
}

// transactionManager занимается управлением транзакциями.
type transactionManager interface {
	// BeginTransaction стартует новую транзакцию.
	BeginTransaction(ctx context.Context) (tr.Transaction, error)
}

// reportsBuilder предоставляет АПИ для построения отчётов.
type reportsBuilder interface {
	// RequestSpendingReport запрашивает формирование отчёта о тратах пользователя.
	RequestSpendingReport(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) error
}

// financeUseCases реализует логику пользовательских сценариев модуля финансов.
type financeUseCases struct {
	// Валюта по-умолчанию.
	defaultCurrency finance_models.Currency
	// Валютные настройки.
	currencySettings currencySettings
	// Настройки бюджентых лимитов.
	spendLimitSettings spendLimitSettings
	// Хранилище трат.
	spendingStorage spendingStorage
	// Конвертер валют.
	converter currencyConverter
	// АПИ для формирования отчётов.
	reports reportsBuilder
	// Менеджер транзакций.
	tm transactionManager
}

// NewUseCases создаёт пользовательские сценарии модуля финансов.
func NewUseCases(defaultCurrency finance_models.Currency, currencySettings currencySettings, spendLimitSettings spendLimitSettings,
	spendingStorage spendingStorage, converter currencyConverter, reports reportsBuilder, tm transactionManager) (*financeUseCases, error) {
	if currencySettings == nil {
		return nil, errors.New("New FinanceUseCases: currency settings is not assigned")
	}
	if spendLimitSettings == nil {
		return nil, errors.New("New FinanceUseCases: spend limit settings is not assigned")
	}
	if spendingStorage == nil {
		return nil, errors.New("New FinanceUseCases: spendings storage is not assigned")
	}
	if converter == nil {
		return nil, errors.New("New FinanceUseCases: currency converter is not assigned")
	}
	if reports == nil {
		return nil, errors.New("New FinanceUseCases: client build api is not assigned")
	}
	if tm == nil {
		return nil, errors.New("New FinanceUseCases: transactions manager is not assigned")
	}

	return &financeUseCases{
		defaultCurrency:    defaultCurrency,
		currencySettings:   currencySettings,
		spendLimitSettings: spendLimitSettings,
		spendingStorage:    spendingStorage,
		converter:          converter,
		reports:            reports,
		tm:                 tm,
	}, nil
}

// GetUserCurrency возвращает текущую валюту пользователя userId.
func (c *financeUseCases) GetUserCurrency(ctx context.Context, userId int64) (finance_models.Currency, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.GetUserCurrency")
	defer span.Finish()

	userCurrency, err := c.currencySettings.GetCurrency(ctx, nil, userId)
	if err != nil {
		tracing.SetError(span)
	}
	return userCurrency, err
}

// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
func (c *financeUseCases) ChangeCurrency(ctx context.Context, userId int64, newCurrency finance_models.Currency) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.ChangeUserCurrency")
	defer span.Finish()

	var transaction tr.Transaction
	var currencyError, cacheError error
	defer func() {
		if transaction != nil {
			if currencyError != nil || cacheError != nil {
				err = transaction.Rollback()
			} else {
				err = transaction.Commit()
			}
		}
		err = multi_error.Append(err, currencyError, cacheError)
		if err != nil {
			tracing.SetError(span)
			err = errors.Wrap(err, "change currency")
		}
	}()

	transaction, err = c.tm.BeginTransaction(ctx)
	if err != nil {
		return
	}

	oldCurrency, currencyError := c.currencySettings.GetCurrency(ctx, transaction, userId)
	if currencyError != nil {
		currencyError = errors.Wrap(currencyError, "Get old currency from user settings before change currency")
		return
	}
	if oldCurrency == newCurrency {
		return
	}

	currencyError = c.currencySettings.ChangeCurrency(ctx, transaction, userId, newCurrency)
	if currencyError != nil {
		currencyError = errors.Wrap(currencyError, "Change currency in user settings")
		return
	}

	return
}

// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period.
func (c *financeUseCases) GetSpendLimit(ctx context.Context, userId int64, period date.Month) (limit finance_models.Amount, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.GetSpendLimit")
	defer span.Finish()

	spendLimit, exists, err := c.spendLimitSettings.GetSpendLimit(ctx, nil, userId, period)
	if err != nil {
		tracing.SetError(span)
		return emptyAmount, exists, err
	}
	if !exists {
		return emptyAmount, exists, nil
	}

	limitInUserCurrency, err := c.converter.ConvertToUserCurrency(ctx, userId, spendLimit)
	if err != nil {
		tracing.SetError(span)
		return emptyAmount, false, finance_models.NewCurrencyConvertError(err)
	}

	userCurrency, err := c.currencySettings.GetCurrency(ctx, nil, userId)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Get user currency failed", logger.Fields.Error(err))
	}
	return finance_models.NewAmount(limitInUserCurrency, userCurrency), true, nil
}

// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
// Возаращает размер установленного лимита в валюте пользователя.
func (c *financeUseCases) SetSpendLimit(ctx context.Context, userId int64, limit decimal.Decimal, period date.Month) (finance_models.Amount, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.SetSpendLimit")
	defer span.Finish()

	limitInDefaultCurrency, err := c.converter.ConvertToDefaultCurrency(ctx, userId, limit)
	if err != nil {
		tracing.SetError(span)
		return emptyAmount, finance_models.NewCurrencyConvertError(err)
	}

	err = c.setSpendLimitInSettings(ctx, userId, limitInDefaultCurrency, period)
	if err != nil {
		tracing.SetError(span)
		return emptyAmount, err
	}

	userCurrency, err := c.currencySettings.GetCurrency(ctx, nil, userId)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Get user currency failed", logger.Fields.Error(err))
	}
	return finance_models.NewAmount(limit, userCurrency), nil
}

// setSpendLimitInSettings устанавливает в настройках пользователя userId бюджет limit на период времени period.
func (c *financeUseCases) setSpendLimitInSettings(ctx context.Context, userId int64, limitInDefaultCurrency decimal.Decimal, period date.Month) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.SetSpendLimitInSettings")
	defer span.Finish()

	var transaction tr.Transaction
	var limitErr error
	defer func() {
		if transaction != nil {
			if limitErr != nil {
				err = transaction.Rollback()
			} else {
				err = transaction.Commit()
			}
		}
		err = multi_error.Append(limitErr, err)
		if err != nil {
			tracing.SetError(span)
			err = errors.Wrap(err, "set spend limit")
		}
	}()

	transaction, err = c.tm.BeginTransaction(ctx)
	if err != nil {
		return
	}

	limitErr = c.spendLimitSettings.SetSpendLimit(ctx, transaction, userId, limitInDefaultCurrency, period)
	return
}

// RequestSpendingReport запрашивает формирование отчёта по тратам пользователя userId
// за указанный период времени period, сгруппированный по категориям.
func (c *financeUseCases) RequestSpendingReport(ctx context.Context, userId int64, period date.Period) error {
	dateInterval := date.NewIntervalFromPeriod(period)
	periodName := period.String(date.PeriodDisplayFormats.In)
	return c.reports.RequestSpendingReport(ctx, userId, periodName, dateInterval)
}

// AddUserSpending добавляет трату spending пользователя userId.
// Возаращает размер траты в валюте пользователя.
func (c *financeUseCases) AddUserSpending(ctx context.Context, userId int64, spending finance_models.Spending) (finance_models.Amount, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.AddUserSpending")
	defer span.Finish()

	userAmount := spending.Amount
	amountInDefaultCurrency, err := c.converter.ConvertToDefaultCurrency(ctx, userId, userAmount)
	if err != nil {
		tracing.SetError(span)
		return emptyAmount, finance_models.NewCurrencyConvertError(err)
	}
	spending.Amount = amountInDefaultCurrency

	err = c.addUserSpendingIfLimitNotExceed(ctx, userId, spending)
	if err != nil {
		tracing.SetError(span)
		return emptyAmount, err
	}

	userCurrency, err := c.currencySettings.GetCurrency(ctx, nil, userId)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Get user currency failed", logger.Fields.Error(err))
	}
	return finance_models.NewAmount(userAmount, userCurrency), nil
}

// addUserSpendingIfLimitNotExceed сохраняет в хранилище трату spending пользователя userId,
// если не превышен лимит трат за бюджетный период.
func (c *financeUseCases) addUserSpendingIfLimitNotExceed(ctx context.Context, userId int64, spending finance_models.Spending) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "FinanceUseCases.AddUserSpendingIfLimitNotExceed")
	defer span.Finish()

	var transaction tr.Transaction
	var spendErr, cacheErr error
	defer func() {
		if transaction != nil {
			if spendErr != nil || cacheErr != nil {
				err = transaction.Rollback()
			} else {
				err = transaction.Commit()
			}
		}
		err = multi_error.Append(err, spendErr, cacheErr)
		if err != nil {
			tracing.SetError(span)
			err = errors.Wrap(err, "add user spending")
		}
	}()

	transaction, err = c.tm.BeginTransaction(ctx)
	if err != nil {
		return
	}

	currentMonth := date.MonthOf(spending.Date)
	spendLimit, hasSpendLimit, spendErr := c.spendLimitSettings.GetSpendLimit(ctx, transaction, userId, currentMonth)
	if spendErr != nil {
		spendErr = errors.Wrap(spendErr, "get user spend limit before add spending to storage")
		return
	}

	if hasSpendLimit {
		currentSpendAmount, amountErr := c.spendingStorage.GetSpendingsAmount(ctx, transaction, userId, date.NewIntervalFromMonth(currentMonth))
		if amountErr != nil {
			spendErr = errors.Wrap(amountErr, "get user current spending amount before add spending to storage")
			return
		}

		newSpendAmount := decimal.Sum(currentSpendAmount, spending.Amount)
		if newSpendAmount.GreaterThan(spendLimit) {
			spendErr = finance_models.SpendLimitExceededError
			return
		}
	}

	spendErr = c.spendingStorage.AddSpending(ctx, transaction, userId, spending)
	if spendErr != nil {
		spendErr = errors.Wrap(spendErr, "when add spending to storage")
		return
	}

	return
}

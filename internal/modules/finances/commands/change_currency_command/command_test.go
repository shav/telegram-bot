package finance_commands_change_currency_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/ngrok/sqlmw"
	"github.com/stretchr/testify/assert"

	csql "github.com/shav/telegram-bot/internal/common/db"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/change_currency_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/change_currency_command/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases"
	ucmocks "github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/settings"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/user_settings/currencies"
	"github.com/shav/telegram-bot/internal/testing"
)

var userId int64 = 1
var userCurrency = finance_models.Currencies.Dollar
var defaultCurrency = finance_models.Currencies.Ruble

var changeCurrencyError = errors.New("change currency error")

var ctx = context.Background()

func Test_OnUserSelectedCurrency_ShouldChangeUserCurrencyInSettings(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)

	useCaseMock.ChangeCurrencyMock.Inspect(func(ctx context.Context, uid int64, currency finance_models.Currency) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userCurrency, currency)
	}).Return(nil)

	command, err := finance_commands_change_currency.NewHandler(finance_commands_change_currency.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, string(userCurrency.Code))
	assert.NoError(t, err)

	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Текущая валюта изменена на доллар", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.ChangeCurrencyAfterCounter())
}

func Test_OnUserSelectedCurrency_ShouldAnswerError_WhenChangeCurrencyFailed(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.ChangeCurrencyMock.Inspect(func(ctx context.Context, uid int64, currency finance_models.Currency) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userCurrency, currency)
	}).Return(changeCurrencyError)

	command, err := finance_commands_change_currency.NewHandler(finance_commands_change_currency.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, string(userCurrency.Code))

	assert.ErrorIs(t, err, changeCurrencyError)
	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Не удалось сменить валюту, произошла ошибка!!!", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.ChangeCurrencyAfterCounter())
}

//***********************************************************************************************
// Sql trace tests
//***********************************************************************************************

func newCurrencySettingsDbStorage(t *testing.T, tracer sqlmw.Interceptor) *finance_storages_user_currency_settings.UserCurrenciesDbStorage {
	const traceDriver = "postgres-tracer"
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	csql.RegisterDriverIfNotExists(traceDriver, sqlmw.Driver(pq.Driver{}, tracer))
	dbStorage, err := finance_storages_user_currency_settings.NewDbStorageWithDriver(traceDriver, dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func TestSql_OnUserSelectedCurrency(t *testing.T) {
	tracer := csql.NewSqlTracer()
	userCurrencies := newCurrencySettingsDbStorage(t, tracer)

	currencySettings, err := finance_services_currency.NewSettingService(defaultCurrency, userCurrencies)
	assert.NoError(t, err)

	spendLimitSettingsMock := ucmocks.NewSpendLimitSettingsMock(t)
	spendingStorageMock := ucmocks.NewSpendingStorageMock(t)

	currencyConverterMock := ucmocks.NewCurrencyConverterMock(t)
	reportsMock := ucmocks.NewReportsBuilderMock(t)

	ts := tmocks.NewTransactionMock(t)
	ts.CommitMock.Return(nil)
	transactionManagerMock := ucmocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(ts, nil)

	useCase, _ := finances.NewUseCases(defaultCurrency, currencySettings, spendLimitSettingsMock, spendingStorageMock,
		currencyConverterMock, reportsMock, transactionManagerMock)

	tracer.Reset()
	command, _ := finance_commands_change_currency.NewHandler(finance_commands_change_currency.Metadata, userId, useCase)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, string(userCurrency.Code))
	assert.NoError(t, err)

	expectedSqlTrace := "SELECT currency FROM user_currency_settings WHERE user_id = $1\n" +
		"INSERT INTO user_currency_settings (user_id,currency) VALUES ($1,$2) ON CONFLICT (user_id) DO UPDATE SET currency = $3"
	actualSqlTrace := tracer.GetTrace()
	assert.Equal(t, expectedSqlTrace, actualSqlTrace)
}

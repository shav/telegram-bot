package finance_commands_add_spending_test

import (
	"context"
	"testing"

	"github.com/lib/pq"
	"github.com/ngrok/sqlmw"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	csql "github.com/shav/telegram-bot/internal/common/db"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/add_spending_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/add_spending_command/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases"
	ucmocks "github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/spendings"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/user_settings/spend_limits"
	"github.com/shav/telegram-bot/internal/testing"
)

var userId int64 = 1
var user2, user3 int64 = 2, 3
var userCurrency = finance_models.Currencies.Dollar
var defaultCurrency = finance_models.Currencies.Ruble

var money10 = decimal.NewFromInt(10)
var money630 = decimal.NewFromInt(630)
var money1000 = decimal.NewFromInt(1000)

var userSpending = finance_models.Spending{
	Category: finance_models.Categories.Food,
	Amount:   money10,
	Date:     date.New(2022, 10, 13),
}

var thisMonth = date.NewMonth(2022, 10)

var emptyAmount = finance_models.Amount{}

var convertCurrencyError = finance_models.NewCurrencyConvertError(nil)

var ctx = context.Background()

func Test_OnAddSpending_ShouldAnswerSpendingAdded_WhenUserHasNotCustomCurrency(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.AddUserSpendingMock.Inspect(func(ctx context.Context, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending, spending)
	}).Return(finance_models.NewAmount(userSpending.Amount, defaultCurrency), nil)

	command, err := finance_commands_add_spending.NewHandler(finance_commands_add_spending.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Category.DisplayText)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Amount.String())
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, userSpending.Date.String())
	assert.NoError(t, err)

	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Трата из категории \"Еда\" за 13.10.2022 на сумму 10₽ успешно добавлена", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.AddUserSpendingAfterCounter())
}

func Test_OnAddSpending_ShouldAnswerSpendingAdded_WhenUserHasCustomCurrency(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.AddUserSpendingMock.Inspect(func(ctx context.Context, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending, spending)
	}).Return(finance_models.NewAmount(userSpending.Amount, userCurrency), nil)

	command, err := finance_commands_add_spending.NewHandler(finance_commands_add_spending.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Category.DisplayText)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Amount.String())
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, userSpending.Date.String())
	assert.NoError(t, err)

	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Трата из категории \"Еда\" за 13.10.2022 на сумму 10$ успешно добавлена", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.AddUserSpendingAfterCounter())
}

func Test_OnAddSpending_ShouldAnswerError_WhenConvertCurrencyFailed(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.AddUserSpendingMock.Inspect(func(ctx context.Context, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending, spending)
	}).Return(emptyAmount, convertCurrencyError)

	command, err := finance_commands_add_spending.NewHandler(finance_commands_add_spending.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Category.DisplayText)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Amount.String())
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, userSpending.Date.String())

	assert.ErrorIs(t, err, convertCurrencyError)
	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Не удалось добавить трату: \nНе удалось выполнить конвертацию валюты", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.AddUserSpendingAfterCounter())
}

func Test_OnAddSpending_ShouldAnswerError_WhenSpendLimitExceeded(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.AddUserSpendingMock.Inspect(func(ctx context.Context, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending, spending)
	}).Return(emptyAmount, finance_models.SpendLimitExceededError)

	command, _ := finance_commands_add_spending.NewHandler(finance_commands_add_spending.Metadata, userId, useCaseMock)

	_, _, err := command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Category.DisplayText)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Amount.String())
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, userSpending.Date.String())
	assert.NoError(t, err)

	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Не удалось добавить трату: \nПревышен лимит трат на текущий месяц", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.AddUserSpendingAfterCounter())
}

//***********************************************************************************************
// Sql trace tests
//***********************************************************************************************

const traceDriver = "postgres-tracer"

var tracer = csql.NewSqlTracer()

func newSpendingsDbStorage(t *testing.T, tracer sqlmw.Interceptor) *finance_storages_spendings.SpendingsDbStorage {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	csql.RegisterDriverIfNotExists(traceDriver, sqlmw.Driver(pq.Driver{}, tracer))
	dbStorage, err := finance_storages_spendings.NewDbStorageWithDriver(traceDriver, dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func newSpendLimitsDbStorage(t *testing.T, tracer sqlmw.Interceptor) *finance_storages_user_spend_limit_settings.UserSpendLimitDbStorage {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	csql.RegisterDriverIfNotExists(traceDriver, sqlmw.Driver(pq.Driver{}, tracer))
	dbStorage, err := finance_storages_user_spend_limit_settings.NewDbStorageWithDriver(traceDriver, dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func TestSql_OnAddSpending_WhenSpendLimitNotSet(t *testing.T) {
	spendingStorage := newSpendingsDbStorage(t, tracer)

	currencySettingsMock := ucmocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	spendLimitSettings := newSpendLimitsDbStorage(t, tracer)
	reportsMock := ucmocks.NewReportsBuilderMock(t)

	currencyConverterMock := ucmocks.NewCurrencyConverterMock(t)
	currencyConverterMock.ConvertToDefaultCurrencyMock.Return(money630, nil)

	ts := tmocks.NewTransactionMock(t)
	ts.CommitMock.Return(nil)
	transactionManagerMock := ucmocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(ts, nil)

	useCase, _ := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettings, spendingStorage,
		currencyConverterMock, reportsMock, transactionManagerMock)

	tracer.Reset()
	command, err := finance_commands_add_spending.NewHandler(finance_commands_add_spending.Metadata, user2, useCase)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Category.DisplayText)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Amount.String())
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Date.String())
	assert.NoError(t, err)

	expectedSqlTrace := "SELECT spend_limit FROM user_spend_limit_settings WHERE period_month = $1 AND user_id = $2\n" +
		"INSERT INTO user_spendings (user_id,category,amount,date) VALUES ($1,$2,$3,$4)"
	actualSqlTrace := tracer.GetTrace()
	assert.Equal(t, expectedSqlTrace, actualSqlTrace)
}

func TestSql_OnAddSpending_WhenSpendLimitSet(t *testing.T) {
	spendingStorage := newSpendingsDbStorage(t, tracer)
	reportsMock := ucmocks.NewReportsBuilderMock(t)

	currencySettingsMock := ucmocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	spendLimitSettings := newSpendLimitsDbStorage(t, tracer)
	err := spendLimitSettings.SetSpendLimit(ctx, nil, user3, money1000, thisMonth)
	assert.NoError(t, err)

	currencyConverterMock := ucmocks.NewCurrencyConverterMock(t)
	currencyConverterMock.ConvertToDefaultCurrencyMock.Return(money630, nil)

	ts := tmocks.NewTransactionMock(t)
	ts.CommitMock.Return(nil)
	transactionManagerMock := ucmocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(ts, nil)

	useCase, _ := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettings, spendingStorage,
		currencyConverterMock, reportsMock, transactionManagerMock)

	tracer.Reset()
	command, err := finance_commands_add_spending.NewHandler(finance_commands_add_spending.Metadata, user3, useCase)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Category.DisplayText)
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Amount.String())
	assert.NoError(t, err)
	_, _, err = command.HandleNextMessage(ctx, userSpending.Date.String())
	assert.NoError(t, err)

	expectedSqlTrace := "SELECT spend_limit FROM user_spend_limit_settings WHERE period_month = $1 AND user_id = $2\n" +
		"SELECT COALESCE(SUM(amount), 0.0) FROM user_spendings WHERE user_id = $1 AND date >= $2 AND date <= $3\n" +
		"INSERT INTO user_spendings (user_id,category,amount,date) VALUES ($1,$2,$3,$4)"
	actualSqlTrace := tracer.GetTrace()
	assert.Equal(t, expectedSqlTrace, actualSqlTrace)
}

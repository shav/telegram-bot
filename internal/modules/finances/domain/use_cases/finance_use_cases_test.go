package finances_test

import (
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases/mocks"
)

var defaultCurrency = finance_models.Currencies.Ruble
var emptyCurrency = finance_models.Currency{}

var userId int64 = 1
var userCurrency = finance_models.Currencies.Dollar

var thisMonth = date.NewMonth(2022, 10)

var emptyAmount = finance_models.Amount{}

var money10 = decimal.NewFromInt(10)
var money630 = decimal.NewFromInt(630)
var money500 = decimal.NewFromInt(500)
var money1000 = decimal.NewFromInt(1000)

var userSpending = finance_models.Spending{
	Category: finance_models.Categories.Food,
	Amount:   money10,
	Date:     date.New(2022, 10, 13),
}

var defaultSpending = finance_models.Spending{
	Category: finance_models.Categories.Food,
	Amount:   money630,
	Date:     date.New(2022, 10, 13),
}

var convertCurrencyError = errors.New("convert currency error")
var changeCurrencyError = errors.New("change currency error")
var getCurrencyError = errors.New("get currency error")
var changeSpendLimitError = errors.New("change spend limit error")
var getSpendLimitError = errors.New("get spend limit error")
var addSpendingToStorageError = errors.New("add spending to storage error")
var requestReportError = errors.New("request report error")

var ctx = context.Background()

// ***********************************************************************************
// Смена валюты
// ***********************************************************************************

func Test_OnChangeCurrency_ShouldChangeCurrencyInUserSettings(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.CommitMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.ChangeCurrencyMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, newCurrency finance_models.Currency) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, userCurrency, newCurrency)
	}).Return(nil)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	err = useCase.ChangeCurrency(ctx, userId, userCurrency)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), currencySettingsMock.ChangeCurrencyAfterCounter())
	assert.Equal(t, uint64(1), transaction.CommitAfterCounter())
}

func Test_OnChangeCurrency_ShouldReturnError_WhenChangeCurrencyInUserSettingsFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.RollbackMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.ChangeCurrencyMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, newCurrency finance_models.Currency) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, userCurrency, newCurrency)
	}).Return(changeCurrencyError)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	err = useCase.ChangeCurrency(ctx, userId, userCurrency)

	assert.ErrorIs(t, err, changeCurrencyError)
	assert.Equal(t, uint64(1), currencySettingsMock.ChangeCurrencyAfterCounter())
	assert.Equal(t, uint64(1), transaction.RollbackAfterCounter())
}

// ***********************************************************************************
// Получение текущей валюты
// ***********************************************************************************

func Test_OnGetCurrency_ShouldReturnCurrencyFromUserSettingsOrDefault(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64) {
		assert.Equal(t, userId, uid)
	}).Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	actualCurrency, err := useCase.GetUserCurrency(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, actualCurrency.Code, userCurrency.Code)
	assert.Equal(t, uint64(1), currencySettingsMock.GetCurrencyAfterCounter())
}

func Test_OnGetCurrency_ShouldReturnError_WhenGetCurrencyFromUserSettingsFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64) {
		assert.Equal(t, userId, uid)
	}).Return(emptyCurrency, getCurrencyError)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	actualCurrency, err := useCase.GetUserCurrency(ctx, userId)

	assert.ErrorIs(t, err, getCurrencyError)
	assert.Equal(t, actualCurrency, emptyCurrency)
	assert.Equal(t, uint64(1), currencySettingsMock.GetCurrencyAfterCounter())
}

// ***********************************************************************************
// Установка бюджетных лимитов
// ***********************************************************************************

func Test_OnChangeSpendLimit_ShouldChangeSpendLimitInUserSettings_AndReturnNewSpendLimitInUserCurrency(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.CommitMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.SetSpendLimitMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, limit decimal.Decimal, period date.Month) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, money630, limit)
		assert.Equal(t, thisMonth, period)
	}).Return(nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, money10, amount)
	}).Return(money630, nil)

	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	newSpendLimit, err := useCase.SetSpendLimit(ctx, userId, money10, thisMonth)

	assert.NoError(t, err)
	assert.Equal(t, userCurrency, newSpendLimit.Currency)
	assert.Equal(t, money10, newSpendLimit.Value)
	assert.Equal(t, uint64(1), spendLimitSettingsMock.SetSpendLimitAfterCounter())
	assert.Equal(t, uint64(1), transaction.CommitAfterCounter())
}

func Test_OnChangeSpendLimit_ShouldReturnError_WhenChangeSpendLimitInUserSettingsFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.RollbackMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.SetSpendLimitMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, limit decimal.Decimal, period date.Month) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, money630, limit)
		assert.Equal(t, thisMonth, period)
	}).Return(changeSpendLimitError)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, money10, amount)
	}).Return(money630, nil)

	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	newSpendLimit, err := useCase.SetSpendLimit(ctx, userId, money10, thisMonth)

	assert.ErrorIs(t, err, changeSpendLimitError)
	assert.Equal(t, emptyAmount, newSpendLimit)
	assert.Equal(t, uint64(1), spendLimitSettingsMock.SetSpendLimitAfterCounter())
	assert.Equal(t, uint64(1), transaction.RollbackAfterCounter())
}

// ***********************************************************************************
// Получение бюджетных лимитов
// ***********************************************************************************

func Test_OnGetSpendLimit_ShouldReturnSpendLimitFromUserSettingsInUserCurrency(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToUserCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, money630, amount)
	}).Return(money10, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, period date.Month) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisMonth, period)
	}).Return(money630, true, nil)

	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	actualSpendLimit, exists, err := useCase.GetSpendLimit(ctx, userId, thisMonth)

	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, money10, actualSpendLimit.Value)
	assert.Equal(t, userCurrency, actualSpendLimit.Currency)
	assert.Equal(t, uint64(1), spendLimitSettingsMock.GetSpendLimitAfterCounter())
}

func Test_OnGetSpendLimit_ShouldReturnNotExists_WhenSpendLimitNotSetInUserSettings(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, period date.Month) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisMonth, period)
	}).Return(decimal.Zero, false, nil)

	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	actualSpendLimit, exists, err := useCase.GetSpendLimit(ctx, userId, thisMonth)

	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, emptyAmount, actualSpendLimit)
	assert.Equal(t, uint64(1), spendLimitSettingsMock.GetSpendLimitAfterCounter())
}

func Test_OnGetSpendLimit_ShouldReturnError_WhenGetSpendLimitFromUserSettingsFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, period date.Month) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisMonth, period)
	}).Return(decimal.Zero, false, getSpendLimitError)

	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	actualSpendLimit, exists, err := useCase.GetSpendLimit(ctx, userId, thisMonth)

	assert.ErrorIs(t, err, getSpendLimitError)
	assert.False(t, exists)
	assert.Equal(t, emptyAmount, actualSpendLimit)
	assert.Equal(t, uint64(1), spendLimitSettingsMock.GetSpendLimitAfterCounter())
}

// ***********************************************************************************
// Добавление траты
// ***********************************************************************************

func Test_OnAddSpending_ShouldAddSpendingToStorage_WhenSpendLimitIsNotSet(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.CommitMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Return(decimal.Zero, false, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, transaction, ts)
		assert.Equal(t, defaultSpending, spending)
	}).Return(nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending.Amount, amount)
	}).Return(defaultSpending.Amount, nil)

	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	spendingAmount, err := useCase.AddUserSpending(ctx, userId, userSpending)

	assert.NoError(t, err)
	assert.Equal(t, userCurrency, spendingAmount.Currency)
	assert.Equal(t, userSpending.Amount, spendingAmount.Value)
	assert.Equal(t, uint64(1), converterMock.ConvertToDefaultCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.AddSpendingAfterCounter())
	assert.Equal(t, uint64(1), transaction.CommitAfterCounter())
}

func Test_OnAddSpending_ShouldAddSpendingToStorage_WhenSpendLimitIsSetAndNotExceeded(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.CommitMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Return(money1000, true, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, transaction, ts)
		assert.Equal(t, defaultSpending, spending)
	}).Return(nil)
	spendingStorageMock.GetSpendingsAmountMock.Return(decimal.Zero, nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending.Amount, amount)
	}).Return(defaultSpending.Amount, nil)

	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	spendingAmount, err := useCase.AddUserSpending(ctx, userId, userSpending)

	assert.NoError(t, err)
	assert.Equal(t, userCurrency, spendingAmount.Currency)
	assert.Equal(t, userSpending.Amount, spendingAmount.Value)
	assert.Equal(t, uint64(1), converterMock.ConvertToDefaultCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.AddSpendingAfterCounter())
	assert.Equal(t, uint64(1), transaction.CommitAfterCounter())
}

func Test_OnAddSpending_ShouldNotAddSpendingToStorage_WhenSpendLimitIsSetAndExceeded(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.RollbackMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Return(money500, true, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, transaction, ts)
		assert.Equal(t, defaultSpending, spending)
	}).Return(nil)
	spendingStorageMock.GetSpendingsAmountMock.Return(decimal.Zero, nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending.Amount, amount)
	}).Return(defaultSpending.Amount, nil)

	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	spendingAmount, err := useCase.AddUserSpending(ctx, userId, userSpending)

	assert.ErrorContains(t, err, finance_models.SpendLimitExceededError.Error())
	assert.Equal(t, emptyCurrency, spendingAmount.Currency)
	assert.Equal(t, decimal.Zero.String(), spendingAmount.Value.String())
	assert.Equal(t, uint64(1), converterMock.ConvertToDefaultCurrencyAfterCounter())
	assert.Equal(t, uint64(0), spendingStorageMock.AddSpendingAfterCounter())
	assert.Equal(t, uint64(1), transaction.RollbackAfterCounter())
}

func Test_OnAddSpending_ShouldReturnError_WhenAddSpendingToStorageFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)
	transaction.RollbackMock.Return(nil)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(transaction, nil)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Return(decimal.Zero, false, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, transaction, ts)
		assert.Equal(t, defaultSpending, spending)
	}).Return(addSpendingToStorageError)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending.Amount, amount)
	}).Return(defaultSpending.Amount, nil)

	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	spendingAmount, err := useCase.AddUserSpending(ctx, userId, userSpending)

	assert.ErrorIs(t, err, addSpendingToStorageError)
	assert.Equal(t, emptyAmount, spendingAmount)
	assert.Equal(t, uint64(1), converterMock.ConvertToDefaultCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.AddSpendingAfterCounter())
	assert.Equal(t, uint64(1), transaction.RollbackAfterCounter())
}

func Test_OnAddSpending_ShouldReturnError_WhenConvertToDefaultCurrencyFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	spendLimitSettingsMock.GetSpendLimitMock.Return(decimal.Zero, false, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, spending finance_models.Spending) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, defaultSpending, spending)
	}).Return(nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertToDefaultCurrencyMock.Inspect(func(ctx context.Context, uid int64, amount decimal.Decimal) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending.Amount, amount)
	}).Return(defaultSpending.Amount, convertCurrencyError)

	transactionManagerMock := mocks.NewTransactionManagerMock(t)
	reportsMock := mocks.NewReportsBuilderMock(t)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	spendingAmount, err := useCase.AddUserSpending(ctx, userId, userSpending)

	assert.ErrorContains(t, err, convertCurrencyError.Error())
	assert.Equal(t, emptyAmount, spendingAmount)
	assert.Equal(t, uint64(1), converterMock.ConvertToDefaultCurrencyAfterCounter())
	assert.Equal(t, uint64(0), spendingStorageMock.AddSpendingAfterCounter())
}

// ***********************************************************************************
// Запрос отчёта о тратах
// ***********************************************************************************

func Test_RequestSpendingReport_ShouldNotReturnError_WhenSpendingReportSuccessfullyRequested(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)

	reportsMock := mocks.NewReportsBuilderMock(t)
	reportsMock.RequestSpendingReportMock.Return(nil)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	err = useCase.RequestSpendingReport(ctx, userId, date.Periods.Today)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), reportsMock.RequestSpendingReportAfterCounter())
}

func Test_RequestSpendingReport_ShouldReturnError_WhenRequestSpendingReportFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	converterMock := mocks.NewCurrencyConverterMock(t)
	spendLimitSettingsMock := mocks.NewSpendLimitSettingsMock(t)
	transactionManagerMock := mocks.NewTransactionManagerMock(t)

	reportsMock := mocks.NewReportsBuilderMock(t)
	reportsMock.RequestSpendingReportMock.Return(requestReportError)

	useCase, err := finances.NewUseCases(defaultCurrency, currencySettingsMock, spendLimitSettingsMock,
		spendingStorageMock, converterMock, reportsMock, transactionManagerMock)
	assert.NoError(t, err)

	err = useCase.RequestSpendingReport(ctx, userId, date.Periods.Today)

	assert.ErrorIs(t, err, requestReportError)
	assert.Equal(t, uint64(1), reportsMock.RequestSpendingReportAfterCounter())
}

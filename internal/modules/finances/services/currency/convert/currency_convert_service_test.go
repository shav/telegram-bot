package finance_services_currency_test

import (
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/convert"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/convert/mocks"
)

var defaultCurrency = finance_models.Currencies.Ruble
var userCurrency = finance_models.Currencies.Dollar

var userId int64 = 1
var userCurrencyRate = finance_models.NewActualCurrencyRate(userCurrency, decimal.NewFromFloat(63))
var emptyCurrencyRate = finance_models.CurrencyRate{}

var loadRateError = errors.New("load rate error")

var ctx = context.Background()

func Test_ConvertDefaultToUserCurrency(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetDefaultCurrencyMock.Return(defaultCurrency)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	rateGetterMock := mocks.NewCurrencyRateGetterMock(t)
	rateGetterMock.GetRateMock.Return(userCurrencyRate, nil)

	converter, err := finance_services_currency.NewConvertService(currencySettingsMock, rateGetterMock)
	assert.NoError(t, err)

	defaultAmount, _ := decimal.NewFromString("116.55")
	userAmount, err := converter.ConvertToUserCurrency(ctx, userId, defaultAmount)
	expectedUserAmount, _ := decimal.NewFromString("1.85")

	assert.NoError(t, err)
	assert.Equal(t, expectedUserAmount.String(), userAmount.String())
}

func Test_ConvertUserToDefaultCurrency(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetDefaultCurrencyMock.Return(defaultCurrency)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	rateGetterMock := mocks.NewCurrencyRateGetterMock(t)
	rateGetterMock.GetRateMock.Return(userCurrencyRate, nil)

	converter, err := finance_services_currency.NewConvertService(currencySettingsMock, rateGetterMock)
	assert.NoError(t, err)

	userAmount, _ := decimal.NewFromString("123.45")
	defaultAmount, err := converter.ConvertToDefaultCurrency(ctx, userId, userAmount)
	expectedDefaultAmount, _ := decimal.NewFromString("7777.35")

	assert.NoError(t, err)
	assert.Equal(t, expectedDefaultAmount.String(), defaultAmount.String())
}

func Test_ConvertDefaultToUserCurrencyMany(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetDefaultCurrencyMock.Return(defaultCurrency)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	rateGetterMock := mocks.NewCurrencyRateGetterMock(t)
	rateGetterMock.GetRateMock.Return(userCurrencyRate, nil)

	converter, err := finance_services_currency.NewConvertService(currencySettingsMock, rateGetterMock)
	assert.NoError(t, err)

	defaultAmount := make([]decimal.Decimal, 3)
	defaultAmount[0], _ = decimal.NewFromString("63")
	defaultAmount[1], _ = decimal.NewFromString("94.5")
	defaultAmount[2], _ = decimal.NewFromString("147.42")
	userAmount, err := converter.ConvertToUserCurrencyMany(ctx, userId, defaultAmount...)

	expectedUserAmount := make([]decimal.Decimal, 3)
	expectedUserAmount[0], _ = decimal.NewFromString("1")
	expectedUserAmount[1], _ = decimal.NewFromString("1.5")
	expectedUserAmount[2], _ = decimal.NewFromString("2.34")

	assert.NoError(t, err)
	assert.Equal(t, expectedUserAmount[0].String(), userAmount[0].String())
	assert.Equal(t, expectedUserAmount[1].String(), userAmount[1].String())
	assert.Equal(t, expectedUserAmount[2].String(), userAmount[2].String())
}

func Test_ConvertUserToDefaultCurrencyMany(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetDefaultCurrencyMock.Return(defaultCurrency)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	rateGetterMock := mocks.NewCurrencyRateGetterMock(t)
	rateGetterMock.GetRateMock.Return(userCurrencyRate, nil)

	converter, err := finance_services_currency.NewConvertService(currencySettingsMock, rateGetterMock)
	assert.NoError(t, err)

	userAmount := make([]decimal.Decimal, 3)
	userAmount[0], _ = decimal.NewFromString("1")
	userAmount[1], _ = decimal.NewFromString("1.5")
	userAmount[2], _ = decimal.NewFromString("2.34")
	defaultAmount, err := converter.ConvertToDefaultCurrencyMany(ctx, userId, userAmount...)

	expectedDefaultAmount := make([]decimal.Decimal, 3)
	expectedDefaultAmount[0], _ = decimal.NewFromString("63")
	expectedDefaultAmount[1], _ = decimal.NewFromString("94.5")
	expectedDefaultAmount[2], _ = decimal.NewFromString("147.42")

	assert.NoError(t, err)
	assert.Equal(t, expectedDefaultAmount[0].String(), defaultAmount[0].String())
	assert.Equal(t, expectedDefaultAmount[1].String(), defaultAmount[1].String())
	assert.Equal(t, expectedDefaultAmount[2].String(), defaultAmount[2].String())
}

func Test_ConvertUserToDefaultCurrency_ShouldReturnError_WhenGetRateFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetDefaultCurrencyMock.Return(defaultCurrency)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	rateGetterMock := mocks.NewCurrencyRateGetterMock(t)
	rateGetterMock.GetRateMock.Return(emptyCurrencyRate, loadRateError)

	converter, err := finance_services_currency.NewConvertService(currencySettingsMock, rateGetterMock)
	assert.NoError(t, err)

	userAmount, _ := decimal.NewFromString("123.45")
	defaultAmount, err := converter.ConvertToDefaultCurrency(ctx, userId, userAmount)
	expectedDefaultAmount := decimal.Zero

	assert.ErrorIs(t, err, loadRateError)
	assert.Equal(t, expectedDefaultAmount.String(), defaultAmount.String())
}

func Test_ConvertSpendingsTableToUserCurrency(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetDefaultCurrencyMock.Return(defaultCurrency)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	rateGetterMock := mocks.NewCurrencyRateGetterMock(t)
	rateGetterMock.GetRateMock.Return(userCurrencyRate, nil)

	converter, err := finance_services_currency.NewConvertService(currencySettingsMock, rateGetterMock)
	assert.NoError(t, err)

	defaultSpendings := make(finance_models.SpendingsByCategoryTable, 3)
	defaultSpendings[finance_models.Categories.Food], _ = decimal.NewFromString("63")
	defaultSpendings[finance_models.Categories.Medicines], _ = decimal.NewFromString("94.5")
	defaultSpendings[finance_models.Categories.Clothes], _ = decimal.NewFromString("147.42")
	userSpendings, err := converter.ConvertSpendingsTableToUserCurrency(ctx, userId, defaultSpendings)

	expectedUserSpendings := make(finance_models.SpendingsByCategoryTable, 3)
	expectedUserSpendings[finance_models.Categories.Food], _ = decimal.NewFromString("1")
	expectedUserSpendings[finance_models.Categories.Medicines], _ = decimal.NewFromString("1.5")
	expectedUserSpendings[finance_models.Categories.Clothes], _ = decimal.NewFromString("2.34")

	assert.NoError(t, err)
	assert.Equal(t, expectedUserSpendings[finance_models.Categories.Food].String(), userSpendings[finance_models.Categories.Food].String())
	assert.Equal(t, expectedUserSpendings[finance_models.Categories.Medicines].String(), userSpendings[finance_models.Categories.Medicines].String())
	assert.Equal(t, expectedUserSpendings[finance_models.Categories.Clothes].String(), userSpendings[finance_models.Categories.Clothes].String())
}

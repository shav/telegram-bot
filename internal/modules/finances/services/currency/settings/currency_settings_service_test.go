package finance_services_currency_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/settings"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/settings/mocks"
)

var defaultCurrency = finance_models.Currencies.Ruble
var emptyCurrency = finance_models.Currency{}

var userId int64 = 1

var ctx = context.Background()

var userCurrencySettingsError = errors.New("user currency settings error")

func Test_OnGetUserCurrency_ShouldReturnDefaultCurrency_WhenUserHasNotCurrencySetting(t *testing.T) {
	userCurrenciesMock := mocks.NewUserCurrenciesMock(t)
	userCurrenciesMock.GetCurrencyMock.Return(emptyCurrency, false, nil)

	currencySettings, err := finance_services_currency.NewSettingService(defaultCurrency, userCurrenciesMock)
	assert.NoError(t, err)
	actualCurrency, err := currencySettings.GetCurrency(ctx, nil, userId)

	assert.NoError(t, err)
	assert.Equal(t, defaultCurrency, actualCurrency)
}

func Test_OnGetUserCurrency_ShouldReturnUserCurrencyFromSettings_WhenUserHasCurrencySetting(t *testing.T) {
	var userCurrency = finance_models.Currencies.Dollar
	userCurrenciesMock := mocks.NewUserCurrenciesMock(t)
	userCurrenciesMock.GetCurrencyMock.Return(userCurrency, true, nil)

	currencySettings, err := finance_services_currency.NewSettingService(defaultCurrency, userCurrenciesMock)
	assert.NoError(t, err)

	actualCurrency, err := currencySettings.GetCurrency(ctx, nil, userId)

	assert.NoError(t, err)
	assert.Equal(t, userCurrency, actualCurrency)
}

func Test_OnGetDefaultCurrency_ShouldReturnCorrectValue(t *testing.T) {
	userCurrenciesMock := mocks.NewUserCurrenciesMock(t)
	currencySettings, err := finance_services_currency.NewSettingService(defaultCurrency, userCurrenciesMock)
	assert.NoError(t, err)

	assert.Equal(t, defaultCurrency, currencySettings.GetDefaultCurrency())
}

func Test_OnGetUserCurrency_ShouldReturnError_WhenGettingCurrencyFromUserSettingFailed(t *testing.T) {
	var userCurrency = finance_models.Currencies.Dollar
	userCurrenciesMock := mocks.NewUserCurrenciesMock(t)
	userCurrenciesMock.GetCurrencyMock.Return(userCurrency, true, userCurrencySettingsError)

	currencySettings, err := finance_services_currency.NewSettingService(defaultCurrency, userCurrenciesMock)
	assert.NoError(t, err)

	actualCurrency, err := currencySettings.GetCurrency(ctx, nil, userId)

	assert.ErrorIs(t, err, userCurrencySettingsError)
	assert.Equal(t, emptyCurrency, actualCurrency)
}

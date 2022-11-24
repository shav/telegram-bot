package finance_services_currency_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/rates"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/rates/mocks"
)

var updatePeriod = 10 * time.Minute

var defaultCurrency = finance_models.Currencies.Ruble
var userCurrency = finance_models.Currencies.Dollar

var emptyCurrencyRate = finance_models.CurrencyRate{}
var userCurrencyRate = finance_models.NewActualCurrencyRate(userCurrency, decimal.NewFromFloat(63))

var loadRateError = errors.New("load rate error")

var ctx = context.Background()

func Test_OnUpdateCurrencyRateAsync_ShouldReturnNewCurrencyRate(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	ratesClientMock.LoadRateMock.Return(userCurrencyRate.Rate, nil)

	rateCacheMock := mocks.NewCurrencyRateCacheMock(t)
	rateCacheMock.UpdateMock.Return(nil)

	updater, err := finance_services_currency.NewExchangeRatesUpdater(defaultCurrency, userCurrency, ratesClientMock, rateCacheMock, updatePeriod)
	assert.NoError(t, err)

	rateChanel := updater.UpdateRateAsync(ctx)
	actualRate, ok := <-rateChanel

	assert.True(t, ok)
	assert.Equal(t, userCurrencyRate.Rate.String(), actualRate.Rate.String())
}

func Test_OnUpdateCurrencyRateAsync_ShouldNotReturnNewCurrencyRate_WhenLoadRateFailed(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	ratesClientMock.LoadRateMock.Return(emptyCurrencyRate.Rate, loadRateError)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)
	ratesCacheMock.UpdateMock.Return(nil)

	updater, err := finance_services_currency.NewExchangeRatesUpdater(defaultCurrency, userCurrency, ratesClientMock, ratesCacheMock, updatePeriod)
	assert.NoError(t, err)

	rateChanel := updater.UpdateRateAsync(ctx)
	actualRate, ok := <-rateChanel

	assert.False(t, ok)
	assert.Equal(t, emptyCurrencyRate, actualRate)
}

func Test_OnUpdateCurrencyRateAsync_ShouldNotLoadNewRates_WhenRatesAreAlreadyUpdatingOrReady(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	loadRateCall := ratesClientMock.LoadRateMock.Return(userCurrencyRate.Rate, nil)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)
	ratesCacheMock.UpdateMock.Return(nil)

	updater, err := finance_services_currency.NewExchangeRatesUpdater(defaultCurrency, userCurrency, ratesClientMock, ratesCacheMock, updatePeriod)
	assert.NoError(t, err)

	rateChanel1 := updater.UpdateRateAsync(ctx)
	for updater.IsUpdating() {
		time.Sleep(10 * time.Millisecond)
	}
	rateChanel2 := updater.UpdateRateAsync(ctx)
	assert.Equal(t, rateChanel1, rateChanel2)

	actualRate, ok := <-rateChanel2
	assert.True(t, ok)
	assert.Equal(t, userCurrencyRate.Rate.String(), actualRate.Rate.String())

	assert.Equal(t, uint64(1), loadRateCall.LoadRateAfterCounter())
}

func Test_OnStartRatesMonitoring_ShouldUpdateRates(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	loadRateCall := ratesClientMock.LoadRateMock.Return(userCurrencyRate.Rate, nil)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)
	ratesCacheMock.UpdateMock.Return(nil)

	updater, err := finance_services_currency.NewExchangeRatesUpdater(defaultCurrency, userCurrency, ratesClientMock, ratesCacheMock, updatePeriod)
	assert.NoError(t, err)

	updater.StartMonitoringRate(ctx)
	for updater.IsUpdating() {
		time.Sleep(10 * time.Millisecond)
	}

	assert.True(t, updater.IsMonitoring())
	assert.Equal(t, uint64(1), loadRateCall.LoadRateAfterCounter())
}

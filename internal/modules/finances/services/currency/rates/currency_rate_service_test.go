package finance_services_currency_test

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/rates"
	"github.com/shav/telegram-bot/internal/modules/finances/services/currency/rates/mocks"
)

var allCurrencies = []finance_models.Currency{finance_models.Currencies.Euro, finance_models.Currencies.Dollar}

var euroRate = finance_models.NewActualCurrencyRate(finance_models.Currencies.Euro, decimal.NewFromFloat(60))
var dollarRate = finance_models.NewActualCurrencyRate(finance_models.Currencies.Dollar, decimal.NewFromFloat(63))
var dollarRate2 = finance_models.NewActualCurrencyRate(finance_models.Currencies.Dollar, decimal.NewFromFloat(66))
var emptyRate = finance_models.CurrencyRate{}

func Test_OnStartMonitoringCurrencyRates_ShouldStartMonitoringAllCurrencies(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Euro).Then(euroRate.Rate, nil)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Dollar).Then(dollarRate.Rate, nil)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)

	updaterMocks := make([]*mocks.CurrencyRateUpdaterMock, 0)
	service, err := finance_services_currency.NewExchangeRateService(ratesClientMock, ratesCacheMock, allCurrencies, defaultCurrency, updatePeriod,
		getRateUpdaterFactory(func(defaultCurrency finance_models.Currency, currency finance_models.Currency) finance_services_currency.CurrencyRateUpdater {
			updaterMock := mocks.NewCurrencyRateUpdaterMock(t)
			updaterMock.StartMonitoringRateMock.Expect(ctx).Return()
			updaterMocks = append(updaterMocks, updaterMock)
			return updaterMock
		}))
	assert.NoError(t, err)

	service.StartMonitoringRates(ctx)

	assert.Equal(t, len(allCurrencies), len(updaterMocks))
	for _, updaterMock := range updaterMocks {
		assert.Equal(t, uint64(1), updaterMock.StartMonitoringRateAfterCounter())
	}
}

func Test_OnGetDefaultCurrencyRate_ShouldNotLoadRate(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Euro).Then(euroRate.Rate, nil)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Dollar).Then(dollarRate.Rate, nil)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)

	var defaultUpdaterMock *mocks.CurrencyRateUpdaterMock
	updaterMocks := make([]*mocks.CurrencyRateUpdaterMock, 0)
	service, err := finance_services_currency.NewExchangeRateService(ratesClientMock, ratesCacheMock, allCurrencies, defaultCurrency, updatePeriod,
		getRateUpdaterFactory(func(defaultCurrency finance_models.Currency, currency finance_models.Currency) finance_services_currency.CurrencyRateUpdater {
			updaterMock := mocks.NewCurrencyRateUpdaterMock(t)
			updaterMock.UpdateRateAsyncMock.Expect(ctx).Return(nil)
			if currency.Code == defaultCurrency.Code {
				defaultUpdaterMock = updaterMock
			}
			updaterMocks = append(updaterMocks, updaterMock)
			return updaterMock
		}))
	assert.NoError(t, err)

	defaultRate, err := service.GetRate(ctx, defaultCurrency)

	assert.NoError(t, err)
	assert.Nil(t, defaultUpdaterMock)
	assert.Equal(t, defaultCurrency.Code, defaultRate.Currency.Code)
	assert.Equal(t, finance_models.One, defaultRate.Rate)

	assert.Equal(t, len(allCurrencies), len(updaterMocks))
	for _, updaterMock := range updaterMocks {
		assert.Equal(t, uint64(0), updaterMock.UpdateRateAsyncAfterCounter())
	}
}

func Test_OnGetNotSupportedCurrencyRate_ShouldReturnError(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Euro).Then(euroRate.Rate, nil)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Dollar).Then(dollarRate.Rate, nil)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)

	updaterMocks := make([]*mocks.CurrencyRateUpdaterMock, 0)
	service, err := finance_services_currency.NewExchangeRateService(ratesClientMock, ratesCacheMock, allCurrencies, defaultCurrency, updatePeriod,
		getRateUpdaterFactory(func(defaultCurrency finance_models.Currency, currency finance_models.Currency) finance_services_currency.CurrencyRateUpdater {
			updaterMock := mocks.NewCurrencyRateUpdaterMock(t)
			updaterMock.UpdateRateAsyncMock.Expect(ctx).Return(nil)
			updaterMocks = append(updaterMocks, updaterMock)
			return updaterMock
		}))
	assert.NoError(t, err)

	defaultRate, err := service.GetRate(ctx, finance_models.Currencies.Yuan)

	assert.Errorf(t, err, "Conversion CNY->RUB is not supported")
	assert.Equal(t, emptyRate, defaultRate)

	assert.Equal(t, len(allCurrencies), len(updaterMocks))
	for _, updaterMock := range updaterMocks {
		assert.Equal(t, uint64(0), updaterMock.UpdateRateAsyncAfterCounter())
	}
}

func Test_OnGetCurrencyRate_ForTheFirstTime_ShouldUpdateRate_And_ForTheSecondTime_ShouldReturnDataFromCache(t *testing.T) {
	ratesClientMock := mocks.NewCurrencyRateLoaderMock(t)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Euro).Then(euroRate.Rate, nil)
	ratesClientMock.LoadRateMock.When(ctx, defaultCurrency, finance_models.Currencies.Dollar).Then(dollarRate.Rate, nil)

	ratesCacheMock := mocks.NewCurrencyRateCacheMock(t)
	ratesCacheMock.UpdateMock.Inspect(func(ctx context.Context, ts tr.Transaction, rate finance_models.CurrencyRate) {
		assert.Equal(t, dollarRate, rate)
	}).Return(nil)

	updaterMocks := make(map[finance_models.CurrencyCode]*mocks.CurrencyRateUpdaterMock, 0)
	service, err := finance_services_currency.NewExchangeRateService(ratesClientMock, ratesCacheMock, allCurrencies, defaultCurrency, updatePeriod,
		getRateUpdaterFactory(func(defaultCurrency finance_models.Currency, currency finance_models.Currency) finance_services_currency.CurrencyRateUpdater {
			updaterMock := mocks.NewCurrencyRateUpdaterMock(t)
			rateChan := make(chan finance_models.CurrencyRate, 1)
			if currency.Code == finance_models.Currencies.Dollar.Code {
				rateChan <- dollarRate
			} else if currency.Code == finance_models.Currencies.Euro.Code {
				rateChan <- euroRate
			}
			close(rateChan)
			updaterMock.UpdateRateAsyncMock.Return(rateChan)
			updaterMocks[currency.Code] = updaterMock
			return updaterMock
		}))
	assert.NoError(t, err)

	// Получаем курс самый первый раз
	actualRate, err := service.GetRate(ctx, finance_models.Currencies.Dollar)

	assert.NoError(t, err)
	assert.Equal(t, dollarRate, actualRate)
	assert.Equal(t, uint64(0), ratesCacheMock.GetActualRateAfterCounter())

	assert.Equal(t, len(allCurrencies), len(updaterMocks))
	for currency, updaterMock := range updaterMocks {
		if currency == finance_models.Currencies.Dollar.Code {
			// Хоть метод UpdateRateAsync() и был вызван,
			// в реальном приложении он должен вернуть результат фоновой загрузки курсов при старте приложения
			assert.Equal(t, uint64(1), updaterMock.UpdateRateAsyncAfterCounter())
		} else {
			assert.Equal(t, uint64(0), updaterMock.UpdateRateAsyncAfterCounter())
		}
	}

	// Повторно получаем курс валюты
	ratesCacheMock.GetActualRateMock.Inspect(func(ctx context.Context, ts tr.Transaction, currency finance_models.Currency) {
		assert.Equal(t, finance_models.Currencies.Dollar, currency)
	}).Return(dollarRate2, true, nil)
	actualRate, err = service.GetRate(ctx, finance_models.Currencies.Dollar)

	assert.NoError(t, err)
	assert.Equal(t, dollarRate2, actualRate)
	assert.Equal(t, uint64(1), ratesCacheMock.GetActualRateAfterCounter())

	assert.Equal(t, len(allCurrencies), len(updaterMocks))
	for currency, updaterMock := range updaterMocks {
		if currency == finance_models.Currencies.Dollar.Code {
			// Это вызов от первого раза. Заново метод не должен вызываться.
			assert.Equal(t, uint64(1), updaterMock.UpdateRateAsyncAfterCounter())
		} else {
			assert.Equal(t, uint64(0), updaterMock.UpdateRateAsyncAfterCounter())
		}
	}
}

type SimpleCurrencyRateUpdaterFactory func(defaultCurrency finance_models.Currency, currency finance_models.Currency) finance_services_currency.CurrencyRateUpdater

func getRateUpdaterFactory(factory SimpleCurrencyRateUpdaterFactory) finance_services_currency.CurrencyRateUpdaterFactory {
	return func(defaultCurrency finance_models.Currency, currency finance_models.Currency, loader finance_services_currency.CurrencyRateLoader,
		cache finance_services_currency.CurrencyRateCache, updatePeriod time.Duration) finance_services_currency.CurrencyRateUpdater {
		return factory(defaultCurrency, currency)
	}
}

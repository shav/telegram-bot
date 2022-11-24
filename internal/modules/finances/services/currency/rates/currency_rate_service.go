//go:generate minimock -i CurrencyRateLoader -o ./mocks/ -s ".go"
//go:generate minimock -i CurrencyRateCache -o ./mocks/ -s ".go"
//go:generate minimock -i CurrencyRateUpdater -o ./mocks/ -s ".go"

package finance_services_currency

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyCurrencyRate = finance_models.CurrencyRate{}

// CurrencyRateLoader предоставляет АПИ для загрузки курсов валют.
type CurrencyRateLoader interface {
	// LoadRate загружает курс обмена валюты baseCurrency на валюту targetCurrency.
	LoadRate(ctx context.Context, baseCurrency finance_models.Currency, targetCurrency finance_models.Currency) (decimal.Decimal, error)
}

// CurrencyRateCache используется для кеширования курсов валют.
type CurrencyRateCache interface {
	// Update обновляет курсы валют в кеше.
	Update(ctx context.Context, ts tr.Transaction, rate finance_models.CurrencyRate) error
	// GetActualRate возвращает актуальный курс обмена валюты currency, а также признак наличия информации о курсе в хранилище.
	GetActualRate(ctx context.Context, ts tr.Transaction, currency finance_models.Currency) (rate finance_models.CurrencyRate, exists bool, err error)
}

// CurrencyRateUpdater выполняет периодическое обновление курса валюты.
type CurrencyRateUpdater interface {
	// GetCurrency возвращает целевую валюту курса.
	GetCurrency() finance_models.Currency
	// UpdateRateAsync запускает асинхронное обновление курса валюты.
	UpdateRateAsync(ctx context.Context) <-chan finance_models.CurrencyRate
	// StartMonitoringRate запускает периодическое обновление курсов валют.
	StartMonitoringRate(ctx context.Context)
}

// Фабрика updater-ов курса валют.
type CurrencyRateUpdaterFactory func(
	defaultCurrency finance_models.Currency, currency finance_models.Currency, loader CurrencyRateLoader,
	cache CurrencyRateCache, updatePeriod time.Duration) CurrencyRateUpdater

// CurrencyRateService реализует бизнес-логику для управления курсами валют.
type CurrencyRateService struct {
	// Загрузчик курсов валют.
	loader CurrencyRateLoader
	// Кеш курсов валют.
	cache CurrencyRateCache
	// Список всех валют, за курсами которых нужно следить.
	allCurrencies []finance_models.Currency
	// Основная валюта приложения.
	defaultCurrency finance_models.Currency
	// Период автоматического обновления курсов валют.
	updatePeriod time.Duration
	// Время последнего обновления курсов валют.
	lastUpdateTime time.Time
	// Updater-ы курсов валют.
	updaters map[finance_models.CurrencyCode]CurrencyRateUpdater
	// Объекты синхронизации для считывания результата обновления курсов валют.
	readRateLocks map[finance_models.CurrencyCode]*sync.Mutex
}

// NewExchangeRateService создаёт новый экземпляр сервиса для управления курсами обмена валют.
func NewExchangeRateService(loader CurrencyRateLoader, cache CurrencyRateCache,
	allCurrencies []finance_models.Currency, defaultCurrency finance_models.Currency,
	updatePeriod time.Duration, updatersFactory CurrencyRateUpdaterFactory) (*CurrencyRateService, error) {
	if loader == nil {
		return nil, errors.New("New CurrencyRateService: currency rate loader is not assigned")
	}
	if cache == nil {
		return nil, errors.New("New CurrencyRateService: currency rate cache is not assigned")
	}
	if len(allCurrencies) == 0 {
		return nil, errors.New("New CurrencyRateService: supported currencies is not assigned or empty")
	}
	if updatersFactory == nil {
		return nil, errors.New("New CurrencyRateService: currency rate updaters factory is not assigned")
	}

	if updatePeriod.Nanoseconds() < minUpdatePeriod.Nanoseconds() {
		updatePeriod = minUpdatePeriod
	}

	updaters := make(map[finance_models.CurrencyCode]CurrencyRateUpdater, len(allCurrencies))
	readRateLocks := make(map[finance_models.CurrencyCode]*sync.Mutex, len(allCurrencies))
	for _, currency := range allCurrencies {
		if currency.Code != defaultCurrency.Code {
			updaters[currency.Code] = updatersFactory(defaultCurrency, currency, loader, cache, updatePeriod)
			readRateLocks[currency.Code] = &sync.Mutex{}
		}
	}

	service := &CurrencyRateService{
		loader:          loader,
		cache:           cache,
		allCurrencies:   allCurrencies,
		defaultCurrency: defaultCurrency,
		updatePeriod:    updatePeriod,
		updaters:        updaters,
		lastUpdateTime:  time.Time{},
		readRateLocks:   readRateLocks,
	}
	return service, nil
}

// StartMonitoringRates запускает периодическое обновление курсов валют.
func (s *CurrencyRateService) StartMonitoringRates(ctx context.Context) {
	for _, updater := range s.updaters {
		updater.StartMonitoringRate(ctx)
	}
}

// GetRate возвращает курс конвертации из основной валюты приложения в целевую валюту currency.
func (s *CurrencyRateService) GetRate(ctx context.Context, currency finance_models.Currency) (finance_models.CurrencyRate, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyRateService.GetRate")
	defer span.Finish()

	if s.defaultCurrency.Code == currency.Code {
		return finance_models.GetIdentityCurrencyRate(s.defaultCurrency), nil
	}

	if time.Since(s.lastUpdateTime).Nanoseconds() < s.updatePeriod.Nanoseconds() {
		rate, exists, err := s.cache.GetActualRate(ctx, nil, currency)
		if err != nil {
			tracing.SetError(span)
			return emptyCurrencyRate, err
		}
		if exists {
			return rate, nil
		}
	}

	updater, exists := s.updaters[currency.Code]
	if !exists {
		return finance_models.CurrencyRate{},
			errors.New(fmt.Sprintf("Conversion %s->%s is not supported", currency.Code, s.defaultCurrency.Code))
	}

	actualRate := updater.UpdateRateAsync(ctx)
	rate, ok, err := s.tryReadActualRate(ctx, currency, actualRate)
	if ok && err == nil {
		return rate, nil
	}

	// Если после обновления курса валют результат успел считать какой-то другой поток,
	// то он должен положить эти данные в кеш.
	// Поэтому дожидаемся освобождение доступа к кешу и получаем данные из него.
	rate, exists, err = s.cache.GetActualRate(ctx, nil, currency)
	if err != nil {
		tracing.SetError(span)
		return emptyCurrencyRate, err
	}
	if exists {
		return rate, nil
	}

	// TODO: По идее ситуация маловероятная. На всякий случайно можно попробовать прикрутить SpinWait считывания курса из кеша.
	tracing.SetError(span)
	return emptyCurrencyRate, errors.New("WTF: could not neither load rates nor get rates from cache")
}

// tryReadActualRate пытается прочитать из канала actualRate результат обновления курса валюты currency.
func (s *CurrencyRateService) tryReadActualRate(ctx context.Context, currency finance_models.Currency, actualRate <-chan finance_models.CurrencyRate) (rate finance_models.CurrencyRate, ok bool, err error) {
	readRateLock := s.readRateLocks[currency.Code]
	readRateLock.Lock()
	defer readRateLock.Unlock()

	select {
	case <-ctx.Done():
		break
	// Если несколько потоков запустили обновление курса валюты, то только одному из них повезёт считать результат из канала.
	// Все остальные потоки будут читать результат уже из кеша.
	case rate, ok = <-actualRate:
		if ok {
			if rate.Timestamp.IsZero() {
				rate.Timestamp = time.Now()
			}
			err = s.cache.Update(ctx, nil, rate)
			if err != nil {
				exchange := logger.Fields.String("exchange", fmt.Sprintf("%s->%s", rate.Currency.Code, s.defaultCurrency.Code))
				logger.Error(ctx, "Update currency rate for {exchange} in storage failed", exchange, logger.Fields.Error(err))
				return
			}
			s.lastUpdateTime = time.Now()
			return
		}
	}
	return emptyCurrencyRate, false, err
}

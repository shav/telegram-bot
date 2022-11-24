package finance_services_currency

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Минимальный период обновления курсов валют.
const minUpdatePeriod = time.Second

// CurrencyRateUpdateService выполняет периодическое обновление курсов валют.
type CurrencyRateUpdateService struct {
	// Основная валюта приложения.
	defaultCurrency finance_models.Currency
	// Целевая валюта курса.
	currency finance_models.Currency
	// Загрузчик курсов валют.
	loader CurrencyRateLoader
	// Кеш курсов валют.
	cache CurrencyRateCache
	// Период автоматического обновления курсов валют.
	updatePeriod time.Duration
	// Признак того, что запущено периодическое обновление курсов валют.
	isMonitoring bool
	// Объект синхронизации для запуска мониторингга курсов валют.
	monitoringLock *sync.Mutex
	// Признак того, что курсы валют обновляются прямо в данный момент времени.
	isUpdating bool
	// Объект синхронизации для одноразовой загрузки курсов валют.
	updatingLock *sync.Mutex
	// Результат последнего обновления курсов валют.
	lastRate chan finance_models.CurrencyRate
}

// NewExchangeRatesUpdater создаёт новый экземпляр updater-а курса валюты currency.
func NewExchangeRatesUpdater(defaultCurrency finance_models.Currency, currency finance_models.Currency,
	loader CurrencyRateLoader, cache CurrencyRateCache, updatePeriod time.Duration) (*CurrencyRateUpdateService, error) {
	if loader == nil {
		return nil, errors.New("New CurrencyRateUpdateService: currency rate loader is not assigned")
	}

	if updatePeriod.Nanoseconds() < minUpdatePeriod.Nanoseconds() {
		updatePeriod = minUpdatePeriod
	}

	return &CurrencyRateUpdateService{
		defaultCurrency: defaultCurrency,
		currency:        currency,
		loader:          loader,
		cache:           cache,
		updatePeriod:    updatePeriod,
		updatingLock:    &sync.Mutex{},
		monitoringLock:  &sync.Mutex{},
		lastRate:        nil,
	}, nil
}

// GetCurrency возвращает целевую валюту курса.
func (u *CurrencyRateUpdateService) GetCurrency() finance_models.Currency {
	return u.currency
}

// StartMonitoringRate запускает периодическое обновление курсов валют.
func (u *CurrencyRateUpdateService) StartMonitoringRate(ctx context.Context) {
	var isAlreadyMonitoring bool
	func() {
		u.monitoringLock.Lock()
		defer u.monitoringLock.Unlock()
		if u.isMonitoring {
			isAlreadyMonitoring = true
			return
		}
		u.isMonitoring = true
	}()
	if isAlreadyMonitoring {
		return
	}

	exchange := logger.Fields.String("exchange", fmt.Sprintf("%s->%s", u.currency.Code, u.defaultCurrency.Code))
	logger.Info(ctx, "Start monitoring currency rate for {exchange}", exchange)
	u.resetLastRate()
	u.UpdateRateAsync(ctx)

	go func() {
		timer := time.NewTicker(u.updatePeriod)
		for {
			select {
			case <-ctx.Done():
				timer.Stop()
				logger.Info(ctx, "Stopped monitoring currency rate for {exchange}", exchange)
				return
			case <-timer.C:
				u.resetLastRate()
				u.UpdateRateAsync(ctx)
			}
		}
	}()
}

// UpdateRateAsync запускает асинхронное обновление курса валюты.
func (u *CurrencyRateUpdateService) UpdateRateAsync(ctx context.Context) (actualRate <-chan finance_models.CurrencyRate) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyRateUpdater.UpdateRateAsync")
	defer func() {
		if err := recover(); err != nil {
			tracing.SetError(span)
			span.Finish()
			panic(err)
		}
	}()

	func() {
		u.updatingLock.Lock()
		defer u.updatingLock.Unlock()

		// Если данные уже были загружены ранее в фоновом режиме, то возвращаем их.
		// Загружать их заново смысла нет.
		if u.isUpdating || len(u.lastRate) > 0 {
			actualRate = u.lastRate
			return
		}
		u.isUpdating = true
		u.lastRate = make(chan finance_models.CurrencyRate, 1)
	}()
	if actualRate != nil {
		span.Finish()
		return
	}

	go func() {
		defer span.Finish()
		defer func() {
			u.updatingLock.Lock()
			close(u.lastRate)
			u.isUpdating = false
			u.updatingLock.Unlock()
		}()

		exchange := logger.Fields.String("exchange", fmt.Sprintf("%s->%s", u.currency.Code, u.defaultCurrency.Code))
		select {
		case <-ctx.Done():
			logger.Info(ctx, "Stopped monitoring currency rate for {exchange}", exchange)
			return
		default:
			rateValue, err := u.loader.LoadRate(ctx, u.defaultCurrency, u.currency)
			if err != nil {
				tracing.SetError(span)
				logger.Error(ctx, "Update currency rate for {exchange} failed", exchange, logger.Fields.Error(err))
				return
			}
			logger.Info(ctx, "Loaded currency rate for {exchange}: {rate}", exchange,
				logger.Fields.String("rate", finance_models.FormatMoney(rateValue)))

			rate := finance_models.NewCurrencyRate(u.currency, rateValue, time.Now())
			if u.cache != nil {
				err = u.cache.Update(ctx, nil, rate)
				if err != nil {
					tracing.SetError(span)
					logger.Error(ctx, "Update currency rate for {exchange} in storage failed", exchange, logger.Fields.Error(err))
				}
			}

			u.lastRate <- rate
		}
	}()
	return u.lastRate
}

// IsUpdating возвращает признак того, что курсы валют обновляются прямо в данный момент времени.
func (u *CurrencyRateUpdateService) IsUpdating() bool {
	return u.isUpdating
}

// IsMonitoring возвращает признак того, что запущено периодическое обновление курсов валют.
func (u *CurrencyRateUpdateService) IsMonitoring() bool {
	return u.isMonitoring
}

// resetLastRate инвалидирует результат последнего обновления курса валют.
func (u *CurrencyRateUpdateService) resetLastRate() {
	u.updatingLock.Lock()
	defer u.updatingLock.Unlock()

	u.lastRate = nil
}

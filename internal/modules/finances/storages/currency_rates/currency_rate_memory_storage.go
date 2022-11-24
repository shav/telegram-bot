package finance_storages_currency_rates

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/wangjia184/sortedset"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyCurrencyRate = finance_models.CurrencyRate{}

// CurrencyRateMemoryStorage хранит в памяти актуальные курсы валют приложения.
type CurrencyRateMemoryStorage struct {
	// Коллекция курсов валют (курсы для каждой валюты отсортированы по времени).
	rates map[finance_models.CurrencyCode]*sortedset.SortedSet
	// Объект синхронизации доступа к коллекции курсов.
	lock *sync.RWMutex
}

// NewMemoryStorage создаёт новое хранилище курсов валют в памяти.
func NewMemoryStorage() *CurrencyRateMemoryStorage {
	return &CurrencyRateMemoryStorage{
		rates: make(map[finance_models.CurrencyCode]*sortedset.SortedSet),
		lock:  &sync.RWMutex{},
	}
}

// Update обновляет курс валюты в хранилище.
func (s *CurrencyRateMemoryStorage) Update(ctx context.Context, ts tr.Transaction, rate finance_models.CurrencyRate) error {
	span, _ := tracing.StartSpanFromContext(ctx, "CurrencyRateMemoryStorage.Update")
	defer span.Finish()

	s.lock.Lock()
	defer s.lock.Unlock()

	currencyRates, exists := s.rates[rate.Currency.Code]
	if !exists {
		currencyRates = sortedset.New()
	}

	currencyRates.AddOrUpdate(geTimeKey(rate), geTimeScore(rate.Timestamp), rate)

	s.rates[rate.Currency.Code] = currencyRates
	return nil
}

// GetActualRate возвращает актуальный курс обмена валюты currency, а также признак наличия информации о курсе в хранилище.
func (s *CurrencyRateMemoryStorage) GetActualRate(ctx context.Context, ts tr.Transaction, currency finance_models.Currency) (rate finance_models.CurrencyRate, exists bool, err error) {
	span, _ := tracing.StartSpanFromContext(ctx, "CurrencyRateMemoryStorage.GetActualRate")
	defer span.Finish()

	s.lock.RLock()
	defer s.lock.RUnlock()

	currencyRates, exists := s.rates[currency.Code]
	if !exists || currencyRates.GetCount() == 0 {
		return emptyCurrencyRate, false, nil
	}

	actualRates := currencyRates.GetByScoreRange(geTimeScore(time.Now()), 0, &sortedset.GetByScoreRangeOptions{
		Limit: 1,
	})
	if len(actualRates) > 0 {
		value := actualRates[0].Value
		if value != nil {
			actualRate, ok := value.(finance_models.CurrencyRate)
			if ok && actualRate != emptyCurrencyRate {
				return actualRate, true, nil
			}
		}
	}
	return emptyCurrencyRate, false, nil
}

// geTimeKey возвращает ключ времени котировки валют.
func geTimeKey(rate finance_models.CurrencyRate) string {
	return strconv.FormatInt(rate.Timestamp.UnixNano(), 10)
}

// geTimeScore возвращает метку времени в качестве ключа сортировки.
func geTimeScore(timestamp time.Time) sortedset.SCORE {
	return sortedset.SCORE(timestamp.UnixNano())
}

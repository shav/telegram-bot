package finance_storages_spendings

import (
	"context"
	"sync"

	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// dateKey представляет из себя ключ для хранения даты в хэш-таблице.
type dateKey int

// SpendingsMemoryStorage хранит данные обо всех тратах в памяти (все суммы хранятся в основной валюте расчётов).
type SpendingsMemoryStorage struct {
	// Коллекция трат пользователей по дате совершения траты.
	values map[int64]map[dateKey][]finance_models.Spending
	// Объект синхронизации доступа к коллекции трат.
	lock *sync.RWMutex
}

// NewMemoryStorage создаёт новый экземпляр хранилища трат в памяти.
func NewMemoryStorage() *SpendingsMemoryStorage {
	return &SpendingsMemoryStorage{
		values: make(map[int64]map[dateKey][]finance_models.Spending),
		lock:   &sync.RWMutex{},
	}
}

// AddSpending добавляет информацию о трате spending пользователя userId в хранилище.
func (s *SpendingsMemoryStorage) AddSpending(ctx context.Context, ts tr.Transaction, userId int64, spending finance_models.Spending) error {
	span, _ := tracing.StartSpanFromContext(ctx, "SpendingsMemoryStorage.Add")
	defer span.Finish()

	s.lock.Lock()
	defer s.lock.Unlock()

	userSpendings, exists := s.values[userId]
	if !exists {
		userSpendings = make(map[dateKey][]finance_models.Spending)
		s.values[userId] = userSpendings
	}

	dateKey := dateKey(spending.Date.GetOrderHash())
	spendingsOfDate, exists := userSpendings[dateKey]
	if !exists {
		spendingsOfDate = make([]finance_models.Spending, 0)
	}
	userSpendings[dateKey] = append(spendingsOfDate, spending)
	return nil
}

// GetSpendingsAmount возвращает общий размер трат по всем категориям пользователя userId за указанный промежуток времени interval.
func (s *SpendingsMemoryStorage) GetSpendingsAmount(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (decimal.Decimal, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "SpendingsMemoryStorage.GetSpendingsAmount")
	defer span.Finish()

	s.lock.RLock()
	defer s.lock.RUnlock()

	userSpendings, exists := s.values[userId]
	if !exists {
		return decimal.Zero, nil
	}

	startDay := dateKey(interval.Start().GetOrderHash())
	endDay := dateKey(interval.End().GetOrderHash())
	hasStartBound := startDay != 0
	hasEndBound := endDay != 0
	totalAmount := decimal.Zero
	for day, spendingsOfDate := range userSpendings {
		if (!hasStartBound || day >= startDay) && (!hasEndBound || day <= endDay) {
			for _, spending := range spendingsOfDate {
				totalAmount = decimal.Sum(totalAmount, spending.Amount)
			}
		}
	}
	return totalAmount, nil
}

// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
// за указанный промежуток времени interval, сгруппированный по категориям.
func (s *SpendingsMemoryStorage) GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "SpendingsMemoryStorage.GetSpendingsByCategories")
	defer span.Finish()

	s.lock.RLock()
	defer s.lock.RUnlock()

	emptyTable := make(finance_models.SpendingsByCategoryTable)
	userSpendings, exists := s.values[userId]
	if !exists {
		return emptyTable, nil
	}

	startDay := dateKey(interval.Start().GetOrderHash())
	endDay := dateKey(interval.End().GetOrderHash())
	hasStartBound := startDay != 0
	hasEndBound := endDay != 0
	result := emptyTable
	for day, spendingsOfDate := range userSpendings {
		if (!hasStartBound || day >= startDay) && (!hasEndBound || day <= endDay) {
			for _, spending := range spendingsOfDate {
				totalAmount := result[spending.Category]
				result[spending.Category] = decimal.Sum(totalAmount, spending.Amount)
			}
		}
	}
	return result, nil
}

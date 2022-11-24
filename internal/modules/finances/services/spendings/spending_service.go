//go:generate minimock -i SpendingReportsCache -o ./mocks/ -s ".go"
//go:generate minimock -i spendingStorage -o ./mocks/ -s ".go"

package finance_services_spendings

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// spendingStorage является хранилищем трат.
type spendingStorage interface {
	// AddSpending добавляет информацию о трате spending пользователя userId в хранилище.
	AddSpending(ctx context.Context, ts tr.Transaction, userId int64, spending finance_models.Spending) error
	// GetSpendingsAmount возвращает общий размер трат по всем категориям пользователя userId за указанный промежуток времени interval.
	GetSpendingsAmount(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (decimal.Decimal, error)
	// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
	// за указанный промежуток времени interval, сгруппированный по категориям.
	GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error)
}

// SpendingReportsCache кеширует отчёты о тратах пользователей.
type SpendingReportsCache interface {
	// Add добавляет в кеш отчёт по тратам report пользователя userId за период времени dateInterval.
	Add(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) error
	// Get получает из кеша отчёт по тратам пользователя userId за период времени dateInterval.
	Get(ctx context.Context, userId int64, dateInterval date.Interval) (report finance_models.SpendingsByCategoryTable, exists bool, err error)
	// InvalidateForDate удаляет из кеша отчёты по тратам пользователя userId, в которые входит указанная дата invalidDate.
	InvalidateForDate(ctx context.Context, userId int64, invalidDate date.Date) error
}

// SpendingService реализует логику обрабоки трат пользователя.
type SpendingService struct {
	// Хранилище трат.
	spendingStorage spendingStorage
	// Кеш отчётов о тратах пользователей.
	spendingReportsCache SpendingReportsCache
}

// NewService создвёт новый сервис для обработки трат пользователей.
func NewService(spendingStorage spendingStorage, spendingReportsCache SpendingReportsCache) (*SpendingService, error) {
	if spendingStorage == nil {
		return nil, errors.New("New SpendingService: spendings storage is not assigned")
	}
	// Кеш является необязательным. Если кеш не задан, то просто напрямую обращаемся в БД без кеширования.

	return &SpendingService{
		spendingStorage:      spendingStorage,
		spendingReportsCache: spendingReportsCache,
	}, nil
}

// AddSpending добавляет информацию о трате spending пользователя userId в хранилище.
func (s *SpendingService) AddSpending(ctx context.Context, ts tr.Transaction, userId int64, spending finance_models.Spending) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingService.AddSpending")
	defer span.Finish()

	err := s.spendingStorage.AddSpending(ctx, ts, userId, spending)
	if err != nil {
		tracing.SetError(span)
		return err
	}

	if s.spendingReportsCache != nil {
		cacheErr := s.spendingReportsCache.InvalidateForDate(ctx, userId, spending.Date)
		if cacheErr != nil {
			tracing.SetError(span)
			return errors.Wrap(cacheErr, "Invalidate userId spending report in cache after new spending added")
		}
	}

	return nil
}

// GetSpendingsAmount возвращает общий размер трат по всем категориям пользователя userId за указанный промежуток времени interval.
func (s *SpendingService) GetSpendingsAmount(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (decimal.Decimal, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingService.GetSpendingsAmount")
	defer span.Finish()

	amount, err := s.spendingStorage.GetSpendingsAmount(ctx, ts, userId, interval)

	if err != nil {
		tracing.SetError(span)
	}
	return amount, err
}

// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
// за указанный промежуток времени interval, сгруппированный по категориям.
func (s *SpendingService) GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingService.GetSpendingsByCategories")
	defer span.Finish()

	var getFromCacheErr error
	var cachedReport finance_models.SpendingsByCategoryTable
	var existInCache bool
	if s.spendingReportsCache != nil {
		cachedReport, existInCache, getFromCacheErr = s.spendingReportsCache.Get(ctx, userId, interval)
		if getFromCacheErr != nil {
			tracing.SetError(span)
			logger.Error(ctx, "Get spending report from cache failed", logger.Fields.Error(getFromCacheErr))
		}
		if getFromCacheErr == nil && existInCache {
			return cachedReport, nil
		}
	}

	report, err := s.spendingStorage.GetSpendingsByCategories(ctx, ts, userId, interval)
	if err != nil {
		tracing.SetError(span)
		return nil, err
	}

	if s.spendingReportsCache != nil {
		addToCacheErr := s.spendingReportsCache.Add(ctx, report, userId, interval)
		if addToCacheErr != nil {
			tracing.SetError(span)
			if getFromCacheErr != nil {
				// Т.к. в случае ошибки предварительного получения отчёта из кеша мы достоверно не знаем, был ли отчёт в кеше,
				// то после получения отчёта из БД нам нужно обязательно обновить его в кеше, чтобы в кеше не остался неактуальный отчёт
				return nil, addToCacheErr
			}
			// Если же отчёта ранее не было в кеше, то ошибка при добавлении в кеш отчёта из БД некритична -
			// при следующем обращении просто снова получим отчёт из БД, а не из кеша.
			logger.Error(ctx, "Add spending report to cache failed", logger.Fields.Error(addToCacheErr))
		}
	}

	return report, err
}

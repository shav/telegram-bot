//go:generate minimock -i cache -o ./mocks/ -s ".go"

package finance_caches_spending_reports

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/maps"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Фейковый пустой отчёт, т.к. кеш не поддерживает пустые карты.
var emptyCachedReport = map[string]string{"__EMPTY__": ""}

// cache описывает АПИ универсального кеша данных.
type cache interface {
	// SetMap устанавливает для ключа в кеше key непустое значение value типа map со временем окончания жизни записи expireAt.
	SetMap(ctx context.Context, key string, value map[string]string, expireAt time.Time) error
	// GetMap получает из кеша по ключу key значение типа map.
	GetMap(ctx context.Context, key string) (value map[string]string, exists bool, err error)
	// Delete удаляет из кеша записи с ключами keys.
	Delete(ctx context.Context, keys ...string) error
}

// SpendingReportsCache занимается кешированием отчётов по тратам пользователей.
type SpendingReportsCache struct {
	// Универсальный кеш данных.
	cache cache
	// Интервалы времени, за которые составляются отчёты по тратам пользователей, актуальные на сегодняшний день.
	reportDateIntervals *date.IntervalsCollection
}

// NewCache создаёт кеш отчётов по тратам пользователей.
func NewCache(cache cache, reportPeriods []date.Period) (*SpendingReportsCache, error) {
	if cache == nil {
		return nil, errors.New("New SpendingReportsCache: cache implementation is not assigned")
	}
	return &SpendingReportsCache{
		cache:               cache,
		reportDateIntervals: date.NewIntervalsForPeriods(reportPeriods),
	}, nil
}

// Add добавляет в кеш отчёт по тратам пользователя userId отчёт report за период времени dateInterval.
func (c *SpendingReportsCache) Add(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingReportsCache.Add")
	defer span.Finish()

	var cachedReport map[string]string
	if len(report) == 0 {
		cachedReport = emptyCachedReport
	} else {
		cachedReport = make(map[string]string)
		for category, spendingAmount := range report {
			cachedReport[category.String()] = spendingAmount.String()
		}
	}
	err := c.cache.SetMap(ctx, getCacheKey(userId, dateInterval), cachedReport, getExpireAt(dateInterval))

	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// Get получает из кеша отчёт по тратам пользователя userId за период времени dateInterval.
func (c *SpendingReportsCache) Get(ctx context.Context, userId int64, dateInterval date.Interval) (report finance_models.SpendingsByCategoryTable, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingReportsCache.Get")
	defer span.Finish()

	cachedReport, exists, err := c.cache.GetMap(ctx, getCacheKey(userId, dateInterval))
	if err == nil && exists && cachedReport != nil {
		report = make(finance_models.SpendingsByCategoryTable)
		if maps.Equal(cachedReport, emptyCachedReport) {
			return report, true, nil
		}
		for rawCategory, rawSpendingAmount := range cachedReport {
			spendingAmount, parseErr := decimal.NewFromString(rawSpendingAmount)
			if parseErr != nil {
				return nil, true, parseErr
			}
			category := finance_models.ParseCategory(rawCategory)
			report[category] = spendingAmount
		}
	}

	if err != nil {
		tracing.SetError(span)
	}

	return report, exists, err
}

// InvalidateForDate удаляет из кеша отчёты по тратам пользователя userId, в которые входит указанная дата invalidDate.
func (c *SpendingReportsCache) InvalidateForDate(ctx context.Context, userId int64, invalidDate date.Date) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingReportsCache.InvalidateForDate")
	defer span.Finish()

	reportDateIntervals := c.reportDateIntervals.Get(date.Today())

	invalidDateIntervalKeys := make([]string, 0)
	for _, dateInterval := range reportDateIntervals {
		if dateInterval.Contains(invalidDate) {
			invalidDateIntervalKeys = append(invalidDateIntervalKeys, getCacheKey(userId, dateInterval))
		}
	}

	err := c.cache.Delete(ctx, invalidDateIntervalKeys...)

	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// getCacheKey возвращает ключ в кеше для отчёта по тратам пользователя userId за период времени dateInterval.
func getCacheKey(userId int64, dateInterval date.Interval) string {
	return fmt.Sprintf("%d:%s:%s", userId, dateInterval.Start().SystemString(), dateInterval.End().SystemString())
}

// getExpireAt возвращает время окончания жизни в кеше записи для отчёта по тратам за период времени dateInterval.
func getExpireAt(dateInterval date.Interval) time.Time {
	intervalLastDay := dateInterval.End()
	return intervalLastDay.EndOfDay()
}

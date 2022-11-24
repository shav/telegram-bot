package finance_caches_spending_reports_test

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/finances/caches"
	"github.com/shav/telegram-bot/internal/modules/finances/caches/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
)

var reportPeriods = []date.Period{
	date.Periods.Today, date.Periods.ThisWeek, date.Periods.ThisMonth, date.Periods.ThisYear,
}

var cachedSpendingsReport = map[string]string{
	"Еда": "100",
}

var spendingsReport = finance_models.SpendingsByCategoryTable{
	finance_models.Categories.Food: decimal.NewFromInt(100),
}

var emptyReport finance_models.SpendingsByCategoryTable

const userId = 777

var firstDayOfOfMonth = date.New(2021, 11, 1)
var lastDayOfMonth = date.New(2021, 11, 30)
var endOfMonthTime = lastDayOfMonth.EndOfDay()
var thisMonthDatesInterval = date.NewInterval(firstDayOfOfMonth, lastDayOfMonth)
var someDayOfMonth = date.New(2021, 11, 4)

var cacheError = errors.New("Cache error")

var ctx = context.Background()

// *******************************************************************************************************************
// Добавление отчётов в кеш
// *******************************************************************************************************************

func Test_Add_ShouldNotReturnError_WhenReportSuccessfullyAddedToCache(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.SetMapMock.Inspect(func(ctx context.Context, key string, value map[string]string, expireAt time.Time) {
		assert.Equal(t, "777:2021-11-01:2021-11-30", key)
		assert.Equal(t, cachedSpendingsReport, value)
		assert.Equal(t, endOfMonthTime, expireAt)
	}).Return(nil)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	err = spendingReportsCache.Add(ctx, spendingsReport, userId, thisMonthDatesInterval)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), cacheImplementorMock.SetMapAfterCounter())
}

func Test_Add_ShouldReturnError_WhenAddReportToCacheFailed(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.SetMapMock.Inspect(func(ctx context.Context, key string, value map[string]string, expireAt time.Time) {
		assert.Equal(t, "777:2021-11-01:2021-11-30", key)
		assert.Equal(t, cachedSpendingsReport, value)
		assert.Equal(t, endOfMonthTime, expireAt)
	}).Return(cacheError)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	err = spendingReportsCache.Add(ctx, spendingsReport, userId, thisMonthDatesInterval)

	assert.ErrorIs(t, err, cacheError)
	assert.Equal(t, uint64(1), cacheImplementorMock.SetMapAfterCounter())
}

// *******************************************************************************************************************
// Получение отчётов из кеша
// *******************************************************************************************************************

func Test_Get_ShouldReturnReport_WhenReportExistsInCache(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.GetMapMock.Inspect(func(ctx context.Context, key string) {
		assert.Equal(t, "777:2021-11-01:2021-11-30", key)
	}).Return(cachedSpendingsReport, true, nil)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	actualReport, exists, err := spendingReportsCache.Get(ctx, userId, thisMonthDatesInterval)

	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, spendingsReport, actualReport)
	assert.Equal(t, uint64(1), cacheImplementorMock.GetMapAfterCounter())
}

func Test_Get_ShouldReturnNotExistingFlag_WhenReportNotExistInCache(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.GetMapMock.Inspect(func(ctx context.Context, key string) {
		assert.Equal(t, "777:2021-11-01:2021-11-30", key)
	}).Return(nil, false, nil)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	actualReport, exists, err := spendingReportsCache.Get(ctx, userId, thisMonthDatesInterval)

	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, emptyReport, actualReport)
	assert.Equal(t, uint64(1), cacheImplementorMock.GetMapAfterCounter())
}

func Test_Get_ShouldReturnError_WhenGetReportFromCacheFailed(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.GetMapMock.Inspect(func(ctx context.Context, key string) {
		assert.Equal(t, "777:2021-11-01:2021-11-30", key)
	}).Return(nil, false, cacheError)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	actualReport, exists, err := spendingReportsCache.Get(ctx, userId, thisMonthDatesInterval)

	assert.ErrorIs(t, err, cacheError)
	assert.False(t, exists)
	assert.Equal(t, emptyReport, actualReport)
	assert.Equal(t, uint64(1), cacheImplementorMock.GetMapAfterCounter())
}

// *******************************************************************************************************************
// Инвалидация отчётов в кеше
// *******************************************************************************************************************

func Test_InvalidateForDate_ShouldReturnErrorAndNotInvalidateUserReports_WhenDeleteUserReportsFromCacheFailed(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.DeleteMock.Inspect(func(ctx context.Context, keys ...string) {
		assert.Empty(t, keys)
	}).Return(cacheError)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	err = spendingReportsCache.InvalidateForDate(ctx, userId, someDayOfMonth)

	assert.ErrorIs(t, err, cacheError)
	assert.Equal(t, uint64(1), cacheImplementorMock.DeleteAfterCounter())
}

func Test_InvalidateForDate_ShouldInvalidateUserReportsForDate_WhenUserReportsSuccessfullyDeletedFromCache(t *testing.T) {
	cacheImplementorMock := mocks.NewCacheMock(t)
	cacheImplementorMock.DeleteMock.Inspect(func(ctx context.Context, keys ...string) {
		assert.Empty(t, keys)
	}).Return(nil)

	spendingReportsCache, err := finance_caches_spending_reports.NewCache(cacheImplementorMock, reportPeriods)
	assert.NoError(t, err)

	err = spendingReportsCache.InvalidateForDate(ctx, userId, someDayOfMonth)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), cacheImplementorMock.DeleteAfterCounter())
}

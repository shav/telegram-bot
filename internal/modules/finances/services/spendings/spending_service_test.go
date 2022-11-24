package finance_services_spendings_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/services/spendings"
	"github.com/shav/telegram-bot/internal/modules/finances/services/spendings/mocks"
)

var userId int64 = 1

var money10 = decimal.NewFromInt(10)

var userSpending = finance_models.Spending{
	Category: finance_models.Categories.Food,
	Amount:   money10,
	Date:     date.New(2022, 11, 9),
}

var spendingsReport = finance_models.SpendingsByCategoryTable{
	finance_models.Categories.Food: money10,
}

var defaultReport finance_models.SpendingsByCategoryTable

var thisYear = date.NewInterval(date.New(2022, 1, 1), date.New(2022, 12, 31))

var storageError = errors.New("storage error")
var cacheError = errors.New("cache error")
var addToCacheError = errors.New("add cache error")

var ctx = context.Background()

// ********************************************************************************************************************
// Добавление трат
// ********************************************************************************************************************
func Test_AddSpending_ShouldNotReturnError_WhenSpendingAddedToStorageSuccessfully(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, spending finance_models.Spending) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending, spending)
	}).Return(nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.InvalidateForDateMock.Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	err = spendingService.AddSpending(ctx, transaction, userId, userSpending)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), spendingStorageMock.AddSpendingAfterCounter())
}

func Test_AddSpending_ShouldInvalidateUserSpendingReportsInCache_WhenSpendingAddedToStorageSuccessfully(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Return(nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.InvalidateForDateMock.Inspect(func(ctx context.Context, uid int64, invalidDate date.Date) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, userSpending.Date, invalidDate)
	}).Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	err = spendingService.AddSpending(ctx, transaction, userId, userSpending)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.InvalidateForDateAfterCounter())
}

func Test_AddSpending_ShouldReturnError_WhenAddSpendingToStorageFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Return(storageError)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.InvalidateForDateMock.Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	err = spendingService.AddSpending(ctx, transaction, userId, userSpending)

	assert.ErrorIs(t, err, storageError)
	assert.Equal(t, uint64(1), spendingStorageMock.AddSpendingAfterCounter())
}

func Test_AddSpending_ShouldNotInvalidateUserSpendingReportsInCache_WhenAddSpendingToStorageFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Return(storageError)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.InvalidateForDateMock.Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	err = spendingService.AddSpending(ctx, transaction, userId, userSpending)

	assert.ErrorIs(t, err, storageError)
	assert.Equal(t, uint64(0), spendingReportsCacheMock.InvalidateForDateAfterCounter())
}

func Test_AddSpending_ShouldReturnError_WhenInvalidateUserSpendingReportsInCacheFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.AddSpendingMock.Return(nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.InvalidateForDateMock.Return(cacheError)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	err = spendingService.AddSpending(ctx, transaction, userId, userSpending)

	assert.ErrorIs(t, err, cacheError)
	assert.Equal(t, uint64(1), spendingStorageMock.AddSpendingAfterCounter())
	assert.Equal(t, uint64(1), spendingReportsCacheMock.InvalidateForDateAfterCounter())
}

// ********************************************************************************************************************
// Получение отчёта о тратах
// ********************************************************************************************************************
func Test_GetSpendingsByCategories_ShouldReturnReportFromCache_WhenReportExistInCache(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.GetMock.Inspect(func(ctx context.Context, uid int64, dateInterval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisYear, dateInterval)
	}).Return(spendingsReport, true, nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	actualSpendingsReport, err := spendingService.GetSpendingsByCategories(ctx, transaction, userId, thisYear)

	assert.NoError(t, err)
	assert.Equal(t, spendingsReport, actualSpendingsReport)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.GetAfterCounter())
	assert.Equal(t, uint64(0), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
	assert.Equal(t, uint64(0), spendingReportsCacheMock.AddAfterCounter())
}

func Test_GetSpendingsByCategories_ShouldReturnReportFromStorage_AndAddReportToCache_WhenGetReportFromCacheFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, dateInterval date.Interval) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisYear, dateInterval)
	}).Return(spendingsReport, nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.GetMock.Return(nil, false, cacheError)
	spendingReportsCacheMock.AddMock.Inspect(func(ctx context.Context, report finance_models.SpendingsByCategoryTable, uid int64, dateInterval date.Interval) {
		assert.Equal(t, spendingsReport, report)
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisYear, dateInterval)
	}).Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	actualSpendingsReport, err := spendingService.GetSpendingsByCategories(ctx, transaction, userId, thisYear)

	assert.NoError(t, err)
	assert.Equal(t, spendingsReport, actualSpendingsReport)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.GetAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
	assert.Equal(t, uint64(1), spendingReportsCacheMock.AddAfterCounter())
}

func Test_GetSpendingsByCategories_ShouldReturnReportFromStorage_AndAddReportToCache_WhenReportNotExistedInCache(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, dateInterval date.Interval) {
		assert.Equal(t, transaction, ts)
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisYear, dateInterval)
	}).Return(spendingsReport, nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.GetMock.Inspect(func(ctx context.Context, uid int64, dateInterval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisYear, dateInterval)
	}).Return(nil, false, nil)
	spendingReportsCacheMock.AddMock.Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	actualSpendingsReport, err := spendingService.GetSpendingsByCategories(ctx, transaction, userId, thisYear)

	assert.NoError(t, err)
	assert.Equal(t, spendingsReport, actualSpendingsReport)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.GetAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
	assert.Equal(t, uint64(1), spendingReportsCacheMock.AddAfterCounter())
}

func Test_GetSpendingsByCategories_ShouldReturnError_AndNotAddReportToCache_WhenReportNotExistedInCache_AndGetReportFromStorageFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Return(nil, storageError)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.GetMock.Inspect(func(ctx context.Context, uid int64, dateInterval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, thisYear, dateInterval)
	}).Return(nil, false, nil)
	spendingReportsCacheMock.AddMock.Return(nil)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	actualSpendingsReport, err := spendingService.GetSpendingsByCategories(ctx, transaction, userId, thisYear)

	assert.ErrorIs(t, err, storageError)
	assert.Equal(t, defaultReport, actualSpendingsReport)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.GetAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
	assert.Equal(t, uint64(0), spendingReportsCacheMock.AddAfterCounter())
}

// Если отчёта ранее не было в кеше, то ошибка при добавлении в кеш отчёта из БД некритична -
// при следующем обращении просто снова получим отчёт из БД, а не из кеша.
func Test_GetSpendingsByCategories_ShouldNotReturnError_WhenAddReportToCacheFailed_AndReportNotExistedInCacheBefore(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Return(spendingsReport, nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.GetMock.Return(nil, false, nil)
	spendingReportsCacheMock.AddMock.Return(addToCacheError)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	actualSpendingsReport, err := spendingService.GetSpendingsByCategories(ctx, transaction, userId, thisYear)

	assert.NoError(t, err)
	assert.Equal(t, spendingsReport, actualSpendingsReport)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.GetAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
	assert.Equal(t, uint64(1), spendingReportsCacheMock.AddAfterCounter())
}

// Т.к. в случае ошибки предварительного получения отчёта из кеша мы достоверно не знаем, был ли отчёт в кеше,
// то после получения отчёта из БД нам нужно обязательно обновить его в кеше, чтобы в кеше не остался неактуальный отчёт
func Test_GetSpendingsByCategories_ShouldReturnError_WhenAddReportToCacheFailed_AndGetReportFromCacheBeforeGetFromStorageFailed(t *testing.T) {
	transaction := tmocks.NewTransactionMock(t)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Return(spendingsReport, nil)

	spendingReportsCacheMock := mocks.NewSpendingReportsCacheMock(t)
	spendingReportsCacheMock.GetMock.Return(nil, false, cacheError)
	spendingReportsCacheMock.AddMock.Return(addToCacheError)

	spendingService, err := finance_services_spendings.NewService(spendingStorageMock, spendingReportsCacheMock)
	assert.NoError(t, err)

	actualSpendingsReport, err := spendingService.GetSpendingsByCategories(ctx, transaction, userId, thisYear)

	assert.ErrorIs(t, err, addToCacheError)
	assert.Equal(t, defaultReport, actualSpendingsReport)
	assert.Equal(t, uint64(1), spendingReportsCacheMock.GetAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
	assert.Equal(t, uint64(1), spendingReportsCacheMock.AddAfterCounter())
}

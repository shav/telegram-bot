package finance_storages_spendings_test

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/spendings"
	"github.com/shav/telegram-bot/internal/testing"
)

const user1, user2, user3 = 1, 2, 3

var minDate = date.New(1, 1, 1)
var maxDate = date.New(9999, 12, 31)

var today = date.New(2022, 10, 18)
var dayInThisWeek = date.New(2022, 10, 17)
var dayInThisMonth = date.New(2022, 10, 1)
var dayInThisYear = date.New(2022, 6, 6)
var dayInPreviousYear = date.New(2021, 4, 4)

var todayInterval = date.NewInterval(today, today)
var thisWeek = date.NewInterval(today.StartOfWeek(), today.EndOfWeek())
var thisMonth = date.NewInterval(today.StartOfMonth(), today.EndOfMonth())
var thisYear = date.NewInterval(today.StartOfYear(), today.EndOfYear())
var wholeTime = date.NewInterval(minDate, maxDate)

var money10 = decimal.NewFromInt(10)
var money20 = decimal.NewFromInt(20)
var money30 = decimal.NewFromInt(30)
var money40 = decimal.NewFromInt(40)
var money50 = decimal.NewFromInt(50)
var money60 = decimal.NewFromInt(60)
var money70 = decimal.NewFromInt(70)
var money80 = decimal.NewFromInt(80)
var money90 = decimal.NewFromInt(90)
var money100 = decimal.NewFromInt(100)
var money150 = decimal.NewFromInt(150)
var money160 = decimal.NewFromInt(160)
var money270 = decimal.NewFromInt(270)
var money400 = decimal.NewFromInt(400)
var money550 = decimal.NewFromInt(550)

var ctx = context.Background()

// spendingsStorage хранит данные обо всех тратах.
type spendingsStorage interface {
	// AddSpending добавляет информацию о трате spending пользователя userId в хранилище.
	AddSpending(ctx context.Context, ts tr.Transaction, userId int64, spending finance_models.Spending) error
	// GetSpendings возвращает общий размер трат по всем категориям пользователя userId за указанный промежуток времени interval.
	GetSpendingsAmount(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (decimal.Decimal, error)
	// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
	// за указанный промежуток времени interval, сгруппированный по категориям.
	GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error)
}

func prepareDbStorage(t *testing.T) *finance_storages_spendings.SpendingsDbStorage {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	dbStorage, err := finance_storages_spendings.NewDbStorage(dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func Test_OnGetSpendingsAmount_FromEmptyStorage_ShouldReturnZero(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)
		spendingsAmount, err := storage.GetSpendingsAmount(ctx, ts, user1, wholeTime)

		assert.NoError(t, err)
		assert.Equal(t, decimal.Zero, spendingsAmount)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetSpendingsTable_FromEmptyStorage_ShouldReturnEmptyTable(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)
		spendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, wholeTime)

		assert.NoError(t, err)
		assert.Empty(t, spendingsTable)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetSpendingsAmount_ShouldReturnDifferentAmount_ForDifferentUsers(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)

		// Траты первого пользователя
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Medicines, money20, today))
		// Траты второго пользователя
		_ = storage.AddSpending(ctx, ts, user2, finance_models.NewSpending(finance_models.Categories.Food, money30, today))
		_ = storage.AddSpending(ctx, ts, user2, finance_models.NewSpending(finance_models.Categories.Medicines, money40, today))

		spendingsAmount1, err := storage.GetSpendingsAmount(ctx, ts, user1, wholeTime)
		assert.NoError(t, err)
		spendingsAmount2, err := storage.GetSpendingsAmount(ctx, ts, user2, wholeTime)
		assert.NoError(t, err)
		assert.NotEqual(t, spendingsAmount1.String(), spendingsAmount2.String())
		assert.NotEqual(t, spendingsAmount1.String(), spendingsAmount2.String())
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetSpendingsTable_ShouldReturnDifferentTables_ForDifferentUsers(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)

		// Траты первого пользователя
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Medicines, money20, today))
		// Траты второго пользователя
		_ = storage.AddSpending(ctx, ts, user2, finance_models.NewSpending(finance_models.Categories.Food, money30, today))
		_ = storage.AddSpending(ctx, ts, user2, finance_models.NewSpending(finance_models.Categories.Medicines, money40, today))

		spendingsTable1, err := storage.GetSpendingsByCategories(ctx, ts, user1, wholeTime)
		assert.NoError(t, err)
		spendingsTable2, err := storage.GetSpendingsByCategories(ctx, ts, user2, wholeTime)
		assert.NoError(t, err)
		assert.NotEqual(t, spendingsTable1[finance_models.Categories.Food].String(), spendingsTable2[finance_models.Categories.Food].String())
		assert.NotEqual(t, spendingsTable1[finance_models.Categories.Medicines].String(), spendingsTable2[finance_models.Categories.Medicines].String())
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetSpendingsTable_ShouldReturnDifferentAmount_ForDifferentCategories(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Medicines, money20, today))

		spendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, wholeTime)

		assert.NoError(t, err)
		assert.NotEqual(t, spendingsTable[finance_models.Categories.Food].String(), spendingsTable[finance_models.Categories.Medicines].String())
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetSpendingsAmount_ShouldReturnAggregatedAmount_ForDatesPeriod(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Medicines, money60, today))

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money20, dayInThisWeek))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Services, money70, dayInThisWeek))

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money30, dayInThisMonth))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Clothes, money80, dayInThisMonth))

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money40, dayInThisYear))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Transport, money90, dayInThisYear))

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money50, dayInPreviousYear))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Electronics, money100, dayInPreviousYear))

		todayAmount, _ := storage.GetSpendingsAmount(ctx, ts, user1, todayInterval)
		assert.Equal(t, money70, todayAmount)

		thisWeekAmount, _ := storage.GetSpendingsAmount(ctx, ts, user1, thisWeek)
		assert.Equal(t, money160, thisWeekAmount)

		thisMonthAmount, _ := storage.GetSpendingsAmount(ctx, ts, user1, thisMonth)
		assert.Equal(t, money270, thisMonthAmount)

		thisYearAmount, _ := storage.GetSpendingsAmount(ctx, ts, user1, thisYear)
		assert.Equal(t, money400, thisYearAmount)

		allSpendingsAmount, _ := storage.GetSpendingsAmount(ctx, ts, user1, wholeTime)
		assert.Equal(t, money550, allSpendingsAmount)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetSpendingsTable_ShouldReturnAggregatedAmount_ForDatesPeriod(t *testing.T) {
	runTest := func(t *testing.T, storage spendingsStorage) {
		ts := tmocks.NewTransactionMock(t)

		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money20, dayInThisWeek))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money30, dayInThisMonth))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money40, dayInThisYear))
		_ = storage.AddSpending(ctx, ts, user1, finance_models.NewSpending(finance_models.Categories.Food, money50, dayInPreviousYear))

		todaySpendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, todayInterval)
		assert.NoError(t, err)
		assert.Equal(t, money10, todaySpendingsTable[finance_models.Categories.Food])

		thisWeekSpendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, thisWeek)
		assert.NoError(t, err)
		assert.Equal(t, money30, thisWeekSpendingsTable[finance_models.Categories.Food])

		thisMonthSpendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, thisMonth)
		assert.NoError(t, err)
		assert.Equal(t, money60, thisMonthSpendingsTable[finance_models.Categories.Food])

		thisYearSpendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, thisYear)
		assert.NoError(t, err)
		assert.Equal(t, money100, thisYearSpendingsTable[finance_models.Categories.Food])

		allSpendingsTable, err := storage.GetSpendingsByCategories(ctx, ts, user1, wholeTime)
		assert.NoError(t, err)
		assert.Equal(t, money150, allSpendingsTable[finance_models.Categories.Food])
	}

	// Memory-хранилище
	memoryStorage := finance_storages_spendings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func TestDb_OnGetSpendingsAmount_AfterCommitAddSpending_ShouldReturnNewSpending(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.AddSpending(ctx, ts, user3, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
	assert.NoError(t, err)

	err = ts.Commit()
	assert.NoError(t, err)

	amount, err := dbStorage.GetSpendingsAmount(ctx, nil, user3, wholeTime)
	assert.NoError(t, err)
	assert.Equal(t, money10, amount)
}

func TestDb_OnGetSpendingsAmount_AfterRollbackAddSpending_ShouldNotReturnNewSpending(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.AddSpending(ctx, ts, user3, finance_models.NewSpending(finance_models.Categories.Food, money10, today))
	assert.NoError(t, err)

	err = ts.Rollback()
	assert.NoError(t, err)

	amount, err := dbStorage.GetSpendingsAmount(ctx, nil, user3, wholeTime)
	assert.NoError(t, err)
	assert.Equal(t, decimal.Zero, amount)
}

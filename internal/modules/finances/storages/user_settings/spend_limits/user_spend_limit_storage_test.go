package finance_storages_user_spend_limit_settings_test

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
	"github.com/shav/telegram-bot/internal/modules/finances/storages/user_settings/spend_limits"
	"github.com/shav/telegram-bot/internal/testing"
)

const user1, user2, user3 = 1, 2, 3

var thisMonth = date.NewMonth(2022, 10)
var nextMonth = date.NewMonth(2022, 11)

var money10 = decimal.NewFromInt(10)
var money20 = decimal.NewFromInt(20)

var ctx = context.Background()

// userSpendLimitStorage хранит пользовательские настройки бюджетов на траты.
type userSpendLimitStorage interface {
	// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
	SetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, limit decimal.Decimal, period date.Month) error
	// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period,
	// а также признак того, задан ли в настройках пользователя бюджет на указанный период.
	GetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, period date.Month) (limit decimal.Decimal, exists bool, err error)
}

func prepareDbStorage(t *testing.T) *finance_storages_user_spend_limit_settings.UserSpendLimitDbStorage {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	dbStorage, err := finance_storages_user_spend_limit_settings.NewDbStorage(dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func Test_OnGetUserSpendLimit_FromEmptyStorage_ShouldReturnNotExists(t *testing.T) {
	runTest := func(t *testing.T, storage userSpendLimitStorage) {
		ts := tmocks.NewTransactionMock(t)

		spendLimit, exists, err := storage.GetSpendLimit(ctx, ts, user1, thisMonth)

		assert.NoError(t, err)
		assert.False(t, exists)
		assert.Equal(t, decimal.Zero, spendLimit)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_user_spend_limit_settings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetUserSpendLimit_AfterSetUserSpendLimitForDifferentUsers_ShouldReturnDifferentSpendLimit(t *testing.T) {
	runTest := func(t *testing.T, storage userSpendLimitStorage) {
		ts := tmocks.NewTransactionMock(t)

		// Меняем бюджет только у одного пользователя
		err := storage.SetSpendLimit(ctx, ts, user1, money10, thisMonth)
		assert.NoError(t, err)
		userSpendLimit1, exists1, err := storage.GetSpendLimit(ctx, ts, user1, thisMonth)
		assert.NoError(t, err)
		assert.True(t, exists1)
		assert.Equal(t, money10, userSpendLimit1)
		// На другого пользователя это никак не влияет
		userSpendLimit2, exists2, err := storage.GetSpendLimit(ctx, ts, user2, thisMonth)
		assert.NoError(t, err)
		assert.False(t, exists2)
		assert.Equal(t, decimal.Zero, userSpendLimit2)

		// Меняем бюджет у второго пользователя
		err = storage.SetSpendLimit(ctx, ts, user2, money20, thisMonth)
		assert.NoError(t, err)
		userSpendLimit2, exists2, err = storage.GetSpendLimit(ctx, ts, user2, thisMonth)
		assert.NoError(t, err)
		assert.True(t, exists2)
		assert.Equal(t, money20, userSpendLimit2)
		// На другого пользователя это тоже никак не должно повлиять
		userSpendLimit1, exists1, err = storage.GetSpendLimit(ctx, ts, user1, thisMonth)
		assert.NoError(t, err)
		assert.True(t, exists1)
		assert.Equal(t, money10, userSpendLimit1)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_user_spend_limit_settings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetUserSpendLimit_AfterChangeUserSpendLimitForDifferentMonths_ShouldReturnDifferentSpendLimitsForMonths(t *testing.T) {
	runTest := func(t *testing.T, storage userSpendLimitStorage) {
		ts := tmocks.NewTransactionMock(t)

		err := storage.SetSpendLimit(ctx, ts, user1, money10, thisMonth)
		assert.NoError(t, err)

		err = storage.SetSpendLimit(ctx, ts, user1, money20, nextMonth)
		assert.NoError(t, err)

		thisMonthSpendLimit, thisExists, err := storage.GetSpendLimit(ctx, ts, user1, thisMonth)
		assert.NoError(t, err)
		assert.True(t, thisExists)
		assert.Equal(t, money10, thisMonthSpendLimit)

		nextMonthSpendLimit, nextExists, err := storage.GetSpendLimit(ctx, ts, user1, nextMonth)
		assert.NoError(t, err)
		assert.True(t, nextExists)
		assert.Equal(t, money20, nextMonthSpendLimit)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_user_spend_limit_settings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetUserSpendLimit_AfterChangeBeforeSetUserSpendLimit_ShouldReturnUserNewSpendLimit(t *testing.T) {
	runTest := func(t *testing.T, storage userSpendLimitStorage) {
		ts := tmocks.NewTransactionMock(t)

		err := storage.SetSpendLimit(ctx, ts, user1, money10, thisMonth)
		assert.NoError(t, err)
		userSpendLimit, exists, err := storage.GetSpendLimit(ctx, ts, user1, thisMonth)
		assert.NoError(t, err)
		assert.True(t, exists)
		assert.Equal(t, money10, userSpendLimit)

		err = storage.SetSpendLimit(ctx, ts, user1, money20, thisMonth)
		assert.NoError(t, err)
		userNewSpendLimit, newExists, err := storage.GetSpendLimit(ctx, ts, user1, thisMonth)
		assert.NoError(t, err)
		assert.True(t, newExists)
		assert.Equal(t, money20, userNewSpendLimit)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_user_spend_limit_settings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func TestDb_OnGetUserSpendLimit_AfterCommitSetUserSpendLimit_ShouldReturnNewSpendLimit(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.SetSpendLimit(ctx, ts, user3, money10, thisMonth)
	assert.NoError(t, err)

	err = ts.Commit()
	assert.NoError(t, err)

	userSpendLimit, exists, err := dbStorage.GetSpendLimit(ctx, nil, user3, thisMonth)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, money10, userSpendLimit)
}

func TestDb_OnGetUserSpendLimit_AfterRollbackSetUserSpendLimit_ShouldReturnOldSpendLimit(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.SetSpendLimit(ctx, ts, user3, money10, thisMonth)
	assert.NoError(t, err)

	err = ts.Rollback()
	assert.NoError(t, err)

	userSpendLimit, exists, err := dbStorage.GetSpendLimit(ctx, nil, user3, thisMonth)
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, decimal.Zero, userSpendLimit)
}

package finance_storages_user_currency_settings_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/user_settings/currencies"
	"github.com/shav/telegram-bot/internal/testing"
)

const user1, user2, user3 = 1, 2, 3

var emptyCurrency = finance_models.Currency{}

var ctx = context.Background()

// userCurrenciesStorage хранит пользовательские настройки валют.
type userCurrenciesStorage interface {
	// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
	ChangeCurrency(ctx context.Context, ts tr.Transaction, userId int64, newCurrency finance_models.Currency) error
	// GetCurrency возвращает текущую валюту для пользователя userId,
	// а также признак того, задана ли в настройках пользователя текущая валюта.
	GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (currency finance_models.Currency, exists bool, err error)
}

func prepareDbStorage(t *testing.T) *finance_storages_user_currency_settings.UserCurrenciesDbStorage {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	dbStorage, err := finance_storages_user_currency_settings.NewDbStorage(dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func Test_OnGetUserCurrency_FromEmptyStorage_ShouldReturnNotExists(t *testing.T) {
	runTest := func(t *testing.T, storage userCurrenciesStorage) {
		ts := tmocks.NewTransactionMock(t)

		userCurrency, exists, err := storage.GetCurrency(ctx, ts, user1)

		assert.NoError(t, err)
		assert.False(t, exists)
		assert.Equal(t, emptyCurrency, userCurrency)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_user_currency_settings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetUserCurrency_AfterChangeUserCurrencyForDifferentUsers_ShouldReturnDifferentCurrency(t *testing.T) {
	runTest := func(t *testing.T, storage userCurrenciesStorage) {
		ts := tmocks.NewTransactionMock(t)

		// Меняем валюту только у одного пользователя
		err := storage.ChangeCurrency(ctx, ts, user1, finance_models.Currencies.Dollar)
		assert.NoError(t, err)
		userCurrency1, exists1, err := storage.GetCurrency(ctx, ts, user1)
		assert.NoError(t, err)
		assert.True(t, exists1)
		assert.Equal(t, finance_models.Currencies.Dollar, userCurrency1)
		// На другого пользователя это никак не влияет
		userCurrency2, exists2, err := storage.GetCurrency(ctx, ts, user2)
		assert.NoError(t, err)
		assert.False(t, exists2)
		assert.Equal(t, emptyCurrency, userCurrency2)

		// Меняем валюту у второго пользователя
		err = storage.ChangeCurrency(ctx, ts, user2, finance_models.Currencies.Euro)
		userCurrency2, exists2, _ = storage.GetCurrency(ctx, ts, user2)
		assert.NoError(t, err)
		assert.True(t, exists2)
		assert.Equal(t, finance_models.Currencies.Euro, userCurrency2)
		// На другого пользователя это тоже никак не должно повлиять
		userCurrency1, exists1, _ = storage.GetCurrency(ctx, ts, user1)
		assert.True(t, exists1)
		assert.Equal(t, finance_models.Currencies.Dollar, userCurrency1)

		// Еще раз меняем валюту у первого пользователя (ранее у него была выбрана другая валюта)
		err = storage.ChangeCurrency(ctx, ts, user1, finance_models.Currencies.Ruble)
		userCurrency1, exists1, _ = storage.GetCurrency(ctx, ts, user1)
		assert.NoError(t, err)
		assert.True(t, exists1)
		assert.Equal(t, finance_models.Currencies.Ruble, userCurrency1)
		// На другого пользователя это опять никак не влияет
		userCurrency2, exists2, _ = storage.GetCurrency(ctx, ts, user2)
		assert.True(t, exists2)
		assert.Equal(t, finance_models.Currencies.Euro, userCurrency2)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_user_currency_settings.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func TestDb_OnGetUserCurrency_AfterCommitChangeUserCurrency_ShouldReturnNewUserCurrency(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.ChangeCurrency(ctx, ts, user3, finance_models.Currencies.Dollar)
	assert.NoError(t, err)

	err = ts.Commit()
	assert.NoError(t, err)

	userCurrency, exists, err := dbStorage.GetCurrency(ctx, nil, user3)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, finance_models.Currencies.Dollar, userCurrency)
}

func TestDb_OnGetUserCurrency_AfterRollbackChangeUserCurrency_ShouldReturnOldUserCurrency(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.ChangeCurrency(ctx, ts, user3, finance_models.Currencies.Dollar)
	assert.NoError(t, err)

	err = ts.Rollback()
	assert.NoError(t, err)

	userCurrency, exists, err := dbStorage.GetCurrency(ctx, nil, user3)
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, emptyCurrency, userCurrency)
}

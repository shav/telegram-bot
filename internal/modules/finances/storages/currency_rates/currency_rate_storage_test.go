package finance_storages_currency_rates_test

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/currency_rates"
	"github.com/shav/telegram-bot/internal/testing"
)

var emptyCurrencyRate = finance_models.CurrencyRate{}

var money50 = decimal.NewFromInt(50)
var money53 = decimal.NewFromInt(53)
var money55 = decimal.NewFromInt(55)
var money56 = decimal.NewFromInt(56)
var money60 = decimal.NewFromInt(60)
var money63 = decimal.NewFromInt(63)
var money65 = decimal.NewFromInt(65)

var ctx = context.Background()

// currencyRateStorage - хранилище курсов валют.
type currencyRateStorage interface {
	// Update обновляет курс валюты в хранилище.
	Update(ctx context.Context, ts tr.Transaction, rate finance_models.CurrencyRate) error
	// GetActualRate возвращает актуальный курс обмена валюты currency, а также признак наличия информации о курсе в хранилище.
	GetActualRate(ctx context.Context, ts tr.Transaction, currency finance_models.Currency) (rate finance_models.CurrencyRate, exists bool, err error)
}

func prepareDbStorage(t *testing.T) *finance_storages_currency_rates.CurrencyRateDbStorage {
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	dbStorage, err := finance_storages_currency_rates.NewDbStorage(dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func Test_OnGetCurrencyRate_FromEmptyStorage_ShouldReturnNotExists(t *testing.T) {
	runTest := func(t *testing.T, storage currencyRateStorage) {
		ts := tmocks.NewTransactionMock(t)
		dollarRate, dollarRateExists, err := storage.GetActualRate(ctx, ts, finance_models.Currencies.Dollar)
		assert.NoError(t, err)
		assert.False(t, dollarRateExists)
		assert.Equal(t, emptyCurrencyRate, dollarRate)

		euroRate, euroRateExists, err := storage.GetActualRate(ctx, ts, finance_models.Currencies.Euro)
		assert.NoError(t, err)
		assert.False(t, euroRateExists)
		assert.Equal(t, emptyCurrencyRate, euroRate)

		yuanRate, yuanRateExists, err := storage.GetActualRate(ctx, ts, finance_models.Currencies.Yuan)
		assert.NoError(t, err)
		assert.False(t, yuanRateExists)
		assert.Equal(t, emptyCurrencyRate, yuanRate)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_currency_rates.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetCurrencyRate_AfterUpdate_ShouldReturnNewCurrencyRate(t *testing.T) {
	runTest := func(t *testing.T, storage currencyRateStorage) {
		ts := tmocks.NewTransactionMock(t)
		// Обновляем курс одной валюты
		err := storage.Update(ctx, ts, finance_models.NewActualCurrencyRate(finance_models.Currencies.Dollar, money50))
		assert.NoError(t, err)
		dollarRate, dollarRateExists, err := storage.GetActualRate(ctx, ts, finance_models.Currencies.Dollar)
		assert.NoError(t, err)
		assert.True(t, dollarRateExists)
		assert.Equal(t, money50, dollarRate.Rate)
		// На курс другой валюты это никак не влияет
		euroRate, euroRateExists, err := storage.GetActualRate(ctx, ts, finance_models.Currencies.Euro)
		assert.NoError(t, err)
		assert.False(t, euroRateExists)
		assert.Equal(t, emptyCurrencyRate, euroRate)

		// Обновляем курс другой валюты
		err = storage.Update(ctx, ts, finance_models.NewActualCurrencyRate(finance_models.Currencies.Euro, money56))
		assert.NoError(t, err)
		euroRate, euroRateExists, err = storage.GetActualRate(ctx, ts, finance_models.Currencies.Euro)
		assert.NoError(t, err)
		assert.True(t, euroRateExists)
		assert.Equal(t, money56, euroRate.Rate)
		// На курс первой валюты это никак не влияет
		dollarRate, dollarRateExists, err = storage.GetActualRate(ctx, ts, finance_models.Currencies.Dollar)
		assert.NoError(t, err)
		assert.True(t, dollarRateExists)
		assert.Equal(t, money50, dollarRate.Rate)

		// Ещё раз обновляем курс первой валюты
		err = storage.Update(ctx, ts, finance_models.NewActualCurrencyRate(finance_models.Currencies.Dollar, money63))
		assert.NoError(t, err)
		dollarRate, dollarRateExists, err = storage.GetActualRate(ctx, ts, finance_models.Currencies.Dollar)
		assert.NoError(t, err)
		assert.True(t, dollarRateExists)
		assert.Equal(t, money63, dollarRate.Rate)
		// На курс другой валюты это снова никак не влияет
		euroRate, euroRateExists, err = storage.GetActualRate(ctx, ts, finance_models.Currencies.Euro)
		assert.NoError(t, err)
		assert.True(t, euroRateExists)
		assert.Equal(t, money56, euroRate.Rate)
	}

	// Memory-хранилище
	memoryStorage := finance_storages_currency_rates.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func Test_OnGetCurrencyRate_AfterManyUpdates_ShouldReturnActualCurrencyRate(t *testing.T) {
	runTest := func(t *testing.T, storage currencyRateStorage) {
		ts := tmocks.NewTransactionMock(t)

		now := time.Now()
		// Время, в порядке следования на временной шкале.
		time1 := now.Add(-12 * time.Hour)
		time2 := now.Add(-6 * time.Hour)
		time3 := now.Add(-4 * time.Hour)
		time4 := now.Add(-1 * time.Hour)
		time5 := now.Add(1 * time.Hour)
		time6 := now.Add(2 * time.Hour)

		_ = storage.Update(ctx, ts, finance_models.NewCurrencyRate(finance_models.Currencies.Dollar, money50, time2))
		_ = storage.Update(ctx, ts, finance_models.NewCurrencyRate(finance_models.Currencies.Dollar, money53, time1))
		_ = storage.Update(ctx, ts, finance_models.NewCurrencyRate(finance_models.Currencies.Dollar, money60, time4))
		_ = storage.Update(ctx, ts, finance_models.NewCurrencyRate(finance_models.Currencies.Dollar, money55, time3))
		_ = storage.Update(ctx, ts, finance_models.NewCurrencyRate(finance_models.Currencies.Dollar, money65, time6))
		_ = storage.Update(ctx, ts, finance_models.NewCurrencyRate(finance_models.Currencies.Dollar, money63, time5))

		dollarRate, dollarRateExists, err := storage.GetActualRate(ctx, ts, finance_models.Currencies.Dollar)
		assert.NoError(t, err)
		assert.True(t, dollarRateExists)
		assert.Equal(t, money60.String(), dollarRate.Rate.String())
	}

	// Memory-хранилище
	memoryStorage := finance_storages_currency_rates.NewMemoryStorage()
	runTest(t, memoryStorage)

	// Database-хранилище
	dbStorage := prepareDbStorage(t)
	runTest(t, dbStorage)
}

func TestDb_OnGetCurrencyRate_AfterCommitUpdateCurrencyRate_ShouldReturnNewCurrencyRate(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.Update(ctx, ts, finance_models.NewActualCurrencyRate(finance_models.Currencies.Dollar, money60))
	assert.NoError(t, err)

	err = ts.Commit()
	assert.NoError(t, err)

	dollarRate, dollarRateExists, err := dbStorage.GetActualRate(ctx, nil, finance_models.Currencies.Dollar)
	assert.NoError(t, err)
	assert.True(t, dollarRateExists)
	assert.Equal(t, money60, dollarRate.Rate)
}

func TestDb_OnGetCurrencyRate_AfterRollbackUpdateCurrencyRate_ShouldNotReturnOldCurrencyRate(t *testing.T) {
	dbStorage := prepareDbStorage(t)

	tm := db.NewTransactionManager(dbStorage.GetDatabase())
	ts, err := tm.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = dbStorage.Update(ctx, ts, finance_models.NewActualCurrencyRate(finance_models.Currencies.Dollar, money60))
	assert.NoError(t, err)

	err = ts.Rollback()
	assert.NoError(t, err)

	dollarRate, dollarRateExists, err := dbStorage.GetActualRate(ctx, nil, finance_models.Currencies.Dollar)
	assert.NoError(t, err)
	assert.False(t, dollarRateExists)
	assert.Equal(t, decimal.Decimal{}, dollarRate.Rate)
}

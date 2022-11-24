package finance_reports_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/ngrok/sqlmw"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	csql "github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	tmocks "github.com/shav/telegram-bot/internal/common/transactions/mocks"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/reports/spendings"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/reports/spendings/mocks"
	ucmocks "github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/storages/spendings"
	"github.com/shav/telegram-bot/internal/testing"
)

var defaultCurrency = finance_models.Currencies.Ruble

var userId int64 = 1
var userCurrency = finance_models.Currencies.Dollar

const reportingPeriodName = "За этот месяц"

var reportingDateInterval = date.NewInterval(date.New(2022, 11, 1), date.New(2022, 11, 30))

var convertCurrencyError = finance_models.NewCurrencyConvertError(nil)
var storageError = errors.New("storage error")

var money10 = decimal.NewFromInt(10)
var money20 = decimal.NewFromInt(20)
var money30 = decimal.NewFromInt(30)
var money60 = decimal.NewFromInt(40)

var spendingsTableInDefaultCurrency = finance_models.SpendingsByCategoryTable{
	finance_models.Categories.Food:      money10,
	finance_models.Categories.Medicines: money20,
}

var spendingsTableInUserCurrency = finance_models.SpendingsByCategoryTable{
	finance_models.Categories.Food:      money30,
	finance_models.Categories.Medicines: money60,
}

var emptySpendingsTable = make(finance_models.SpendingsByCategoryTable)

var ctx = context.Background()

func Test_OnGetSpendingReport_ShouldReturnEmptyReport_WhenNoSpendings(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, interval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingDateInterval, interval)
	}).Return(emptySpendingsTable, nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertSpendingsTableToUserCurrencyMock.Return(emptySpendingsTable, nil)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, converterMock)
	assert.NoError(t, err)

	spendingsReport, err := reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.NoError(t, err)
	assert.Equal(t, "Статистика по тратам за этот месяц:", spendingsReport.Title)
	assert.Equal(t, "Трат пока нет", spendingsReport.Content)
	assert.Equal(t, uint64(1), converterMock.ConvertSpendingsTableToUserCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
}

func Test_OnGetSpendingReport_ShouldReturnReportInUserCurrency_WhenUserSelectedCurrency(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, interval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingDateInterval, interval)
	}).Return(spendingsTableInDefaultCurrency, nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertSpendingsTableToUserCurrencyMock.Return(spendingsTableInUserCurrency, nil)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, converterMock)
	assert.NoError(t, err)

	spendingsReport, err := reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.NoError(t, err)
	assert.Equal(t, "Статистика по тратам за этот месяц:", spendingsReport.Title)
	assert.Equal(t, "Еда:  30$\nЛекарства:  40$", spendingsReport.Content)
	assert.Equal(t, uint64(1), converterMock.ConvertSpendingsTableToUserCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
}

func Test_OnGetSpendingReport_ShouldReturnReportInDefaultCurrency_WhenUserNotSelectedCurrency(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, interval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingDateInterval, interval)
	}).Return(spendingsTableInDefaultCurrency, nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertSpendingsTableToUserCurrencyMock.Return(spendingsTableInDefaultCurrency, nil)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, converterMock)
	assert.NoError(t, err)

	spendingsReport, err := reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.NoError(t, err)
	assert.Equal(t, "Статистика по тратам за этот месяц:", spendingsReport.Title)
	assert.Equal(t, "Еда:  10₽\nЛекарства:  20₽", spendingsReport.Content)
	assert.Equal(t, uint64(1), converterMock.ConvertSpendingsTableToUserCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
}

func Test_OnGetSpendingsReport_ShouldReturnErrorReport_WhenConvertCurrencyFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, interval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingDateInterval, interval)
	}).Return(spendingsTableInDefaultCurrency, nil)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertSpendingsTableToUserCurrencyMock.Return(nil, convertCurrencyError)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, converterMock)
	assert.NoError(t, err)

	spendingsReport, err := reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.ErrorContains(t, err, convertCurrencyError.Error())
	assert.Equal(t, "При формировании отчёта о тратах за этот месяц произошла ошибка:", spendingsReport.Title)
	assert.Equal(t, "Не удалось выполнить конвертацию валюты", spendingsReport.Content)
	assert.Equal(t, uint64(1), converterMock.ConvertSpendingsTableToUserCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
}

func Test_OnGetSpendingsReport_ShouldReturnErrorReport_WhenGetSpendingsFailed(t *testing.T) {
	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(userCurrency, nil)

	spendingStorageMock := mocks.NewSpendingStorageMock(t)
	spendingStorageMock.GetSpendingsByCategoriesMock.Inspect(func(ctx context.Context, ts tr.Transaction, uid int64, interval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingDateInterval, interval)
	}).Return(emptySpendingsTable, storageError)

	converterMock := mocks.NewCurrencyConverterMock(t)
	converterMock.ConvertSpendingsTableToUserCurrencyMock.Return(nil, convertCurrencyError)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, converterMock)
	assert.NoError(t, err)

	spendingsReport, err := reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.ErrorIs(t, err, storageError)
	assert.Equal(t, "При формировании отчёта о тратах за этот месяц произошла ошибка:", spendingsReport.Title)
	assert.Equal(t, "Не удалось получить траты", spendingsReport.Content)
	assert.Equal(t, uint64(0), converterMock.ConvertSpendingsTableToUserCurrencyAfterCounter())
	assert.Equal(t, uint64(1), spendingStorageMock.GetSpendingsByCategoriesAfterCounter())
}

//***********************************************************************************************
// Sql trace tests
//***********************************************************************************************

func newSpendingsDbStorage(t *testing.T, tracer sqlmw.Interceptor) *finance_storages_spendings.SpendingsDbStorage {
	const traceDriver = "postgres-tracer"
	config, err := config.NewEnvConfig(test_settings.ServiceName)
	assert.NoError(t, err)
	dbConnString := config.DbConnectionString()
	assert.NotEmpty(t, dbConnString)
	csql.RegisterDriverIfNotExists(traceDriver, sqlmw.Driver(pq.Driver{}, tracer))
	dbStorage, err := finance_storages_spendings.NewDbStorageWithDriver(traceDriver, dbConnString)
	assert.NoError(t, err)
	err = dbStorage.Clear(ctx)
	assert.NoError(t, err)
	return dbStorage
}

func TestSql_OnGetSpendingReport_WhenReportIsNotCached(t *testing.T) {
	tracer := csql.NewSqlTracer()
	spendingStorageMock := newSpendingsDbStorage(t, tracer)

	currencySettingsMock := mocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	currencyConverterMock := mocks.NewCurrencyConverterMock(t)
	currencyConverterMock.ConvertSpendingsTableToUserCurrencyMock.Return(make(finance_models.SpendingsByCategoryTable), nil)

	ts := tmocks.NewTransactionMock(t)
	ts.CommitMock.Return(nil)
	transactionManagerMock := ucmocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(ts, nil)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, currencyConverterMock)
	assert.NoError(t, err)

	tracer.Reset()
	_, err = reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	expectedSqlTrace := "SELECT category, SUM(amount) FROM user_spendings WHERE user_id = $1 AND date >= $2 AND date <= $3 GROUP BY category"
	actualSqlTrace := tracer.GetTrace()
	assert.NoError(t, err)
	assert.Equal(t, expectedSqlTrace, actualSqlTrace)
}

func TestSql_OnGetSpendingReport_WhenReportIsCached(t *testing.T) {
	tracer := csql.NewSqlTracer()
	spendingStorageMock := newSpendingsDbStorage(t, tracer)

	currencySettingsMock := ucmocks.NewCurrencySettingsMock(t)
	currencySettingsMock.GetCurrencyMock.Return(defaultCurrency, nil)

	currencyConverterMock := ucmocks.NewCurrencyConverterMock(t)
	currencyConverterMock.ConvertSpendingsTableToUserCurrencyMock.Return(make(finance_models.SpendingsByCategoryTable), nil)

	ts := tmocks.NewTransactionMock(t)
	ts.CommitMock.Return(nil)
	transactionManagerMock := ucmocks.NewTransactionManagerMock(t)
	transactionManagerMock.BeginTransactionMock.Return(ts, nil)

	reportBuilder, err := finance_reports.NewSpendingReportBuilder(currencySettingsMock, spendingStorageMock, currencyConverterMock)
	assert.NoError(t, err)

	tracer.Reset()
	_, err = reportBuilder.GetSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	expectedSqlTrace := ""
	actualSqlTrace := tracer.GetTrace()
	assert.NoError(t, err)
	assert.Equal(t, expectedSqlTrace, actualSqlTrace)
}

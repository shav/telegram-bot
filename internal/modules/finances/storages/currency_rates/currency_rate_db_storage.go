package finance_storages_currency_rates

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

const ratesTable = "currency_rates"
const (
	currencyColumn  = "currency"
	timestampColumn = "timestamp"
	rateColumn      = "rate"
)

// CurrencyRateDbStorage хранит в БД актуальные курсы валют приложения.
type CurrencyRateDbStorage struct {
	// База данных.
	db *sql.DB
	// АПИ для построения sql-запросов.
	sql sq.StatementBuilderType
}

// NewDbStorage создаёт новое хранилище курсов валют в БД.
func NewDbStorage(connectionString string) (*CurrencyRateDbStorage, error) {
	return NewDbStorageWithDriver(db.PostgresDriver, connectionString)
}

// NewDbStorageWithDriver создаёт новое хранилище курсов валют в БД.
func NewDbStorageWithDriver(driver string, connectionString string) (*CurrencyRateDbStorage, error) {
	database, err := db.ConnectToDatabase(driver, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "ConnectToDatabase in CurrencyRateDbStorage")
	}
	return NewDbStorageFor(database)
}

// NewDbStorageFor создаёт новое хранилище курсов валют в БД.
func NewDbStorageFor(database *sql.DB) (*CurrencyRateDbStorage, error) {
	return &CurrencyRateDbStorage{
		db:  database,
		sql: db.GetQueryBuilder(database),
	}, nil
}

// Update обновляет курс валюты в хранилище.
func (s *CurrencyRateDbStorage) Update(ctx context.Context, ts tr.Transaction, rate finance_models.CurrencyRate) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyRateDbStorage.Update")
	defer span.Finish()

	query := s.sql.Insert(ratesTable).
		Columns(currencyColumn, timestampColumn, rateColumn).
		Values(rate.Currency.Code, rate.Timestamp, rate.Rate).
		Suffix(fmt.Sprintf("ON CONFLICT (%s, %s) DO UPDATE SET %s = CAST($4 AS DECIMAL)", currencyColumn, timestampColumn, rateColumn), rate.Rate.String()).
		RunWith(s.db)

	err := db.Insert(ctx, ts, query)
	err = errors.Wrap(err, "Update in CurrencyRateDbStorage")
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// GetActualRate возвращает актуальный курс обмена валюты currency, а также признак наличия информации о курсе в хранилище.
func (s *CurrencyRateDbStorage) GetActualRate(ctx context.Context, ts tr.Transaction, currency finance_models.Currency) (rate finance_models.CurrencyRate, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyRateDbStorage.GetActualRate")
	defer span.Finish()

	now := time.Now()
	query := s.sql.Select(timestampColumn, rateColumn).
		From(ratesTable).
		Where(sq.Eq{currencyColumn: currency.Code}).
		Where(sq.LtOrEq{timestampColumn: now}).
		OrderBy(fmt.Sprintf("%s DESC", timestampColumn)).
		Limit(1)

	var timestamp time.Time
	var rateValue decimal.Decimal
	exists, err = db.SelectRow(ctx, ts, query, &timestamp, &rateValue)
	if err != nil {
		tracing.SetError(span)
		return emptyCurrencyRate, false, errors.Wrap(err, "GetActualRate from CurrencyRateDbStorage")
	}
	if !exists {
		return emptyCurrencyRate, false, nil
	}

	rate = finance_models.CurrencyRate{
		Currency:  currency,
		Timestamp: timestamp,
		Rate:      rateValue,
	}
	return rate, true, nil
}

// Clear полностью очищает хранилище курсов валют.
// WARNING: Использовать только в тестах!
func (s *CurrencyRateDbStorage) Clear(ctx context.Context) error {
	query := s.sql.Delete(ratesTable)
	_, err := query.ExecContext(ctx)
	return errors.Wrap(err, "Clear CurrencyRateDbStorage")
}

// GetDatabase возвращает подключение к БД.
// WARNING: Использовать только в тестах!
func (s *CurrencyRateDbStorage) GetDatabase() *sql.DB {
	return s.db
}

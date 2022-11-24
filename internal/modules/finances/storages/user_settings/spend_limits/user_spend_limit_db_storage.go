package finance_storages_user_spend_limit_settings

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

const userSpendLimitsTable = "user_spend_limit_settings"
const (
	userColumn       = "user_id"
	spendLimitColumn = "spend_limit"
	periodColumn     = "period_month"
)

const monthsPerYear = 12

// UserSpendLimitDbStorage хранит пользовательские настройки бюджетов на траты в БД.
type UserSpendLimitDbStorage struct {
	// База данных.
	db *sql.DB
	// АПИ для построения sql-запросов.
	sql sq.StatementBuilderType
}

// NewDbStorage создаёт новый экземпляр хранилища пользовательских настроек бюджетов в БД.
func NewDbStorage(connectionString string) (*UserSpendLimitDbStorage, error) {
	return NewDbStorageWithDriver(db.PostgresDriver, connectionString)
}

// NewDbStorageWithDriver создаёт новый экземпляр хранилища пользовательских настроек бюджетов в БД.
func NewDbStorageWithDriver(driver string, connectionString string) (*UserSpendLimitDbStorage, error) {
	database, err := db.ConnectToDatabase(driver, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "ConnectToDatabase in UserSpendLimitDbStorage")
	}
	return NewDbStorageFor(database)
}

// NewDbStorageFor создаёт новый экземпляр хранилища пользовательских настроек бюджетов в БД.
func NewDbStorageFor(database *sql.DB) (*UserSpendLimitDbStorage, error) {
	return &UserSpendLimitDbStorage{
		db:  database,
		sql: db.GetQueryBuilder(database),
	}, nil
}

// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
func (s *UserSpendLimitDbStorage) SetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, limit decimal.Decimal, period date.Month) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserSpendLimitDbStorage.SetSpendLimit")
	defer span.Finish()

	query := s.sql.Insert(userSpendLimitsTable).
		Columns(userColumn, periodColumn, spendLimitColumn).
		Values(userId, serializePeriod(period), limit).
		Suffix(fmt.Sprintf("ON CONFLICT (%s, %s) DO UPDATE SET %s = $4", userColumn, periodColumn, spendLimitColumn), limit).
		RunWith(s.db)

	err := db.Insert(ctx, ts, query)
	err = errors.Wrap(err, "SetSpendLimit in UserSpendLimitDbStorage")
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period,
// а также признак того, задан ли в настройках пользователя бюджет на указанный период.
func (s *UserSpendLimitDbStorage) GetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, period date.Month) (limit decimal.Decimal, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserSpendLimitDbStorage.GetSpendLimit")
	defer span.Finish()

	query := s.sql.Select(spendLimitColumn).
		From(userSpendLimitsTable).
		Where(sq.Eq{userColumn: userId, periodColumn: serializePeriod(period)})

	var spendLimit decimal.Decimal
	exists, err = db.SelectRow(ctx, ts, query, &spendLimit)
	if err != nil {
		tracing.SetError(span)
		return decimal.Zero, false, errors.Wrap(err, "GetSpendLimit from UserSpendLimitDbStorage")
	}
	if !exists {
		return decimal.Zero, false, nil
	}

	return spendLimit, true, nil
}

// Clear полностью очищает хранилище.
// WARNING: Использовать только в тестах!
func (s *UserSpendLimitDbStorage) Clear(ctx context.Context) error {
	query := s.sql.Delete(userSpendLimitsTable)
	_, err := query.ExecContext(ctx)
	return errors.Wrap(err, "Clear UserSpendLimitDbStorage")
}

// GetDatabase возвращает подключение к БД.
// WARNING: Использовать только в тестах!
func (s *UserSpendLimitDbStorage) GetDatabase() *sql.DB {
	return s.db
}

// serializePeriod возвращает сериализованное значение для периода времени.
func serializePeriod(period date.Month) int {
	return period.Year*monthsPerYear + period.Month
}

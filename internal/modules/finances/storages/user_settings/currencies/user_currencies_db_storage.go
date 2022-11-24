package finance_storages_user_currency_settings

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/db"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

const userCurrenciesTable = "user_currency_settings"
const (
	userColumn     = "user_id"
	currencyColumn = "currency"
)

// UserCurrenciesDbStorage хранит пользовательские настройки валют в памяти.
type UserCurrenciesDbStorage struct {
	// База данных.
	db *sql.DB
	// АПИ для построения sql-запросов.
	sql sq.StatementBuilderType
}

// NewDbStorage создаёт новый экземпляр хранилища пользовательских настроек валют в БД.
func NewDbStorage(connectionString string) (*UserCurrenciesDbStorage, error) {
	return NewDbStorageWithDriver(db.PostgresDriver, connectionString)
}

// NewDbStorageWithDriver создаёт новое хранилище пользовательских настроек валют в БД.
func NewDbStorageWithDriver(driver string, connectionString string) (*UserCurrenciesDbStorage, error) {
	database, err := db.ConnectToDatabase(driver, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "ConnectToDatabase in UserCurrenciesDbStorage")
	}
	return NewDbStorageFor(database)
}

// NewDbStorageFor создаёт новое хранилище пользовательских настроек валют в БД.
func NewDbStorageFor(database *sql.DB) (*UserCurrenciesDbStorage, error) {
	return &UserCurrenciesDbStorage{
		db:  database,
		sql: db.GetQueryBuilder(database),
	}, nil
}

// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
func (s *UserCurrenciesDbStorage) ChangeCurrency(ctx context.Context, ts tr.Transaction, userId int64, newCurrency finance_models.Currency) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserCurrenciesDbStorage.ChangeCurrency")
	defer span.Finish()

	query := s.sql.Insert(userCurrenciesTable).
		Columns(userColumn, currencyColumn).
		Values(userId, newCurrency.Code).
		Suffix(fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s = $3", userColumn, currencyColumn), newCurrency.Code).
		RunWith(s.db)

	err := db.Insert(ctx, ts, query)
	err = errors.Wrap(err, "ChangeCurrency in UserCurrenciesDbStorage")
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// GetCurrency возвращает текущую валюту для пользователя userId,
// а также признак того, задана ли в настройках пользователя текущая валюта.
func (s *UserCurrenciesDbStorage) GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (currency finance_models.Currency, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserCurrenciesDbStorage.GetCurrency")
	defer span.Finish()

	query := s.sql.Select(currencyColumn).
		From(userCurrenciesTable).
		Where(sq.Eq{userColumn: userId})

	var currencyCode string
	exists, err = db.SelectRow(ctx, ts, query, &currencyCode)
	if err != nil {
		tracing.SetError(span)
		return emptyCurrency, false, errors.Wrap(err, "GetCurrency from UserCurrenciesDbStorage")
	}
	if !exists {
		return emptyCurrency, false, nil
	}

	currency, err = finance_models.ParseCurrency(currencyCode)
	if err != nil {
		tracing.SetError(span)
	}
	return currency, true, errors.Wrap(err, "GetCurrency from UserCurrenciesDbStorage")
}

// Clear полностью очищает хранилище.
// WARNING: Использовать только в тестах!
func (s *UserCurrenciesDbStorage) Clear(ctx context.Context) error {
	query := s.sql.Delete(userCurrenciesTable)
	_, err := query.ExecContext(ctx)
	return errors.Wrap(err, "Clear UserCurrenciesDbStorage")
}

// GetDatabase возвращает подключение к БД.
// WARNING: Использовать только в тестах!
func (s *UserCurrenciesDbStorage) GetDatabase() *sql.DB {
	return s.db
}

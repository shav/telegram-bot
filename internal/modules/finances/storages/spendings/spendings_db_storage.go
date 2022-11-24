package finance_storages_spendings

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
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

const userSpendingsTable = "user_spendings"
const (
	idColumn       = "id"
	userColumn     = "user_id"
	categoryColumn = "category"
	amountColumn   = "amount"
	dateColumn     = "date"
)

var emptyDate = date.Date{}

// SpendingsDbStorage хранит данные обо всех тратах в БД (все суммы хранятся в основной валюте расчётов).
type SpendingsDbStorage struct {
	// База данных.
	db *sql.DB
	// АПИ для построения sql-запросов.
	sql sq.StatementBuilderType
}

// NewDbStorage создаёт новый экземпляр хранилища трат в БД.
func NewDbStorage(connectionString string) (*SpendingsDbStorage, error) {
	return NewDbStorageWithDriver(db.PostgresDriver, connectionString)
}

// NewDbStorageWithDriver создаёт новый экземпляр хранилища трат в БД.
func NewDbStorageWithDriver(driver string, connectionString string) (*SpendingsDbStorage, error) {
	database, err := db.ConnectToDatabase(driver, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "ConnectToDatabase in SpendingsStorage")
	}
	return NewDbStorageFor(database)
}

// NewDbStorageFor создаёт новый экземпляр хранилища трат в БД.
func NewDbStorageFor(database *sql.DB) (*SpendingsDbStorage, error) {
	return &SpendingsDbStorage{
		db:  database,
		sql: db.GetQueryBuilder(database),
	}, nil
}

// AddSpending добавляет информацию о трате spending пользователя userId в хранилище.
func (s *SpendingsDbStorage) AddSpending(ctx context.Context, ts tr.Transaction, userId int64, spending finance_models.Spending) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingsDbStorage.Add")
	defer span.Finish()

	query := s.sql.Insert(userSpendingsTable).
		Columns(userColumn, categoryColumn, amountColumn, dateColumn).
		Values(userId, spending.Category.String(), spending.Amount, spending.Date.SystemString()).
		RunWith(s.db)

	err := db.Insert(ctx, ts, query)
	err = errors.Wrap(err, "Add to SpendingDbStorage")
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// GetSpendingsAmount возвращает общий размер трат по всем категориям пользователя userId за указанный промежуток времени interval.
func (s *SpendingsDbStorage) GetSpendingsAmount(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (decimal.Decimal, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingsDbStorage.GetSpendingsAmount")
	defer span.Finish()

	query := s.sql.Select(fmt.Sprintf("COALESCE(SUM(%s), 0.0)", amountColumn)).
		From(userSpendingsTable).
		Where(sq.Eq{userColumn: userId})
	query = applyIntervalFilter(query, interval)

	var amount decimal.Decimal
	exists, err := db.SelectRow(ctx, ts, query, &amount)
	if err != nil {
		tracing.SetError(span)
		return decimal.Zero, errors.Wrap(err, "GetSpendingsAmount from SpendingDbStorage")
	}
	if !exists {
		return decimal.Zero, nil
	}

	if amount.IsZero() {
		return decimal.Zero, nil
	}
	return amount, nil
}

// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
// за указанный промежуток времени interval, сгруппированный по категориям.
func (s *SpendingsDbStorage) GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingsDbStorage.GetSpendingsByCategories")
	defer span.Finish()

	emptyTable := make(finance_models.SpendingsByCategoryTable)

	query := s.sql.Select(categoryColumn, fmt.Sprintf("SUM(%s)", amountColumn)).
		From(userSpendingsTable).
		Where(sq.Eq{userColumn: userId})
	query = applyIntervalFilter(query, interval)
	query = query.GroupBy(categoryColumn)

	rows, err := db.QueryRows(ctx, ts, query)
	if err != nil {
		tracing.SetError(span)
		return emptyTable, errors.Wrap(err, "GetSpendingsByCategories from SpendingDbStorage")
	}
	if rows == nil {
		return emptyTable, nil
	}

	result := emptyTable
	for rows.Next() {
		var categoryName string
		var amount decimal.Decimal
		err = rows.Scan(&categoryName, &amount)
		if err != nil {
			tracing.SetError(span)
			return emptyTable, errors.Wrap(err, "GetSpendingsByCategories from SpendingDbStorage")
		}
		category := finance_models.ParseCategory(categoryName)
		result[category] = amount
	}
	return result, nil
}

// Clear полностью очищает хранилище.
// WARNING: Использовать только в тестах!
func (s *SpendingsDbStorage) Clear(ctx context.Context) error {
	query := s.sql.Delete(userSpendingsTable)
	_, err := query.ExecContext(ctx)
	return errors.Wrap(err, "Clear SpendingsDbStorage")
}

// GetDatabase возвращает подключение к БД.
// WARNING: Использовать только в тестах!
func (s *SpendingsDbStorage) GetDatabase() *sql.DB {
	return s.db
}

// applyIntervalFilter применяет фильтр по периоду к запросу на получение трат пользователя.
func applyIntervalFilter(query sq.SelectBuilder, interval date.Interval) sq.SelectBuilder {
	if interval.Start() != emptyDate {
		query = query.Where(sq.GtOrEq{dateColumn: interval.Start().SystemString()})
	}
	if interval.End() != emptyDate {
		query = query.Where(sq.LtOrEq{dateColumn: interval.End().SystemString()})
	}
	return query
}

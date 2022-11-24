package db

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
)

const PostgresDriver = "postgres"

// Query - это модель запроса к БД.
type Query interface {
	// ToSql возвращает sql-запрос.
	ToSql() (string, []interface{}, error)
	// ExecContext выполняет запрос.
	ExecContext(ctx context.Context) (sql.Result, error)
	// QueryRowContext выполняет select-запрос одной строки.
	QueryRowContext(ctx context.Context) sq.RowScanner
	// QueryContext выполняет select-запрос нескольких строк.
	QueryContext(ctx context.Context) (*sql.Rows, error)
}

// ConnectToDatabase подключается к БД.
func ConnectToDatabase(driver string, connectionString string) (*sql.DB, error) {
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetQueryBuilder возвращает построитель запросов к БД.
func GetQueryBuilder(db *sql.DB) sq.StatementBuilderType {
	dbCache := sq.NewStmtCache(db)
	return sq.StatementBuilder.RunWith(dbCache).PlaceholderFormat(sq.Dollar)
}

// Insert выполняет запрос на вставку данных query в транзакции (если она задана).
func Insert(ctx context.Context, ts tr.Transaction, query Query) error {
	result, err := ExecQuery(ctx, ts, query)
	if err != nil {
		return err
	}
	if result != nil {
		insertedOrUpdatedRowsCount, err := result.RowsAffected()
		if err == nil && insertedOrUpdatedRowsCount == 0 {
			return NoDataWasInsertedOrUpdatedError
		}
	}
	return nil
}

// SelectRow выполняет select-запрос одной строки в транзакции (если она задана) и означивает поля строки в коллекцию result.
func SelectRow(ctx context.Context, ts tr.Transaction, query Query, result ...any) (exists bool, err error) {
	row, err := QueryRow(ctx, ts, query)
	if err != nil {
		return false, err
	}
	if row == nil {
		return false, nil
	}

	err = row.Scan(result...)
	if err != nil {
		if err.Error() == NoRowsInResultErrorMessage {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// QueryRow выполняет select-запрос одной строки в транзакции (если она задана) и возвращает читателя данных строки.
func QueryRow(ctx context.Context, ts tr.Transaction, query Query) (sq.RowScanner, error) {
	if sqltx, ok := ts.(*sql.Tx); ok {
		queryStr, args, qerr := query.ToSql()
		if qerr != nil {
			return nil, qerr
		}
		return sqltx.QueryRowContext(ctx, queryStr, args...), nil
	} else {
		return query.QueryRowContext(ctx), nil
	}
}

// QueryRows выполняет select-запрос нескольких строк в транзакции (если она задана) и возвращает читателя данных строк.
func QueryRows(ctx context.Context, ts tr.Transaction, query Query) (*sql.Rows, error) {
	if sqltx, ok := ts.(*sql.Tx); ok {
		queryStr, args, qerr := query.ToSql()
		if qerr != nil {
			return nil, qerr
		}
		return sqltx.QueryContext(ctx, queryStr, args...)
	} else {
		return query.QueryContext(ctx)
	}
}

// ExecQuery выполняет запрос query в транзакции (если она задана).
func ExecQuery(ctx context.Context, ts tr.Transaction, query Query) (sql.Result, error) {
	if sqltx, ok := ts.(*sql.Tx); ok {
		queryStr, args, qerr := query.ToSql()
		if qerr != nil {
			return nil, qerr
		}
		return sqltx.ExecContext(ctx, queryStr, args...)
	} else {
		return query.ExecContext(ctx)
	}
}

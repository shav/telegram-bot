package db

import (
	"context"
	"database/sql"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
)

// TODO: Прикрутить настройку уровня изоляции транзакций в конфиге.
// Уровень изоляции транзакций, используемый для запросов в БД.
const isolationLevel = sql.LevelReadCommitted

// TransactionManager занимается управлением транзакциями.
type TransactionManager struct {
	// База данных.
	db *sql.DB
}

// NewTransactionManager создаёт менеджер транзакций к базе данных db.
func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

// BeginTransaction стартует новую транзакцию.
func (t *TransactionManager) BeginTransaction(ctx context.Context) (tr.Transaction, error) {
	transaction, err := t.db.BeginTx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	return transaction, err
}

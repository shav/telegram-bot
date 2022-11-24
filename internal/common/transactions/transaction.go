//go:generate minimock -i Transaction -o ./mocks/ -s ".go"

package transactions

// Transaction предоставляет АПИ для транзакций.
type Transaction interface {
	// Commit подтверждает транзакцию.
	Commit() error
	// Rollback откатывает транзакцию.
	Rollback() error
}

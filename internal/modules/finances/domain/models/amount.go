package finance_models

import "github.com/shopspring/decimal"

// Amount хранит денежную сумму.
type Amount struct {
	// Значение денежной суммы.
	Value decimal.Decimal
	// Валюта.
	Currency Currency
}

// NewAmount создаёт новую денежную сумму.
func NewAmount(value decimal.Decimal, currency Currency) Amount {
	return Amount{
		Value:    value,
		Currency: currency,
	}
}

// String возвращает стровоке представление денежной суммы.
func (a Amount) String() string {
	return FormatMoneyWithCurrency(a.Value, a.Currency)
}

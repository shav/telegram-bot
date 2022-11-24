package finance_models

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// FormatMoney форматирует денежную сумму.
func FormatMoney(amount decimal.Decimal) string {
	if amount.IsInteger() {
		return amount.String()
	}
	return amount.StringFixedBank(2)
}

// FormatMoneyWithCurrency форматирует денежную сумму с указанием валюты.
func FormatMoneyWithCurrency(amount decimal.Decimal, currency Currency) string {
	return fmt.Sprintf("%s%s", FormatMoney(amount), currency.Symbol)
}

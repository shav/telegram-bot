package finance_models

import (
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
)

// Spending представляет из себя модель финансовой траты.
type Spending struct {
	// Категория трат
	Category Category
	// Сумма
	Amount decimal.Decimal
	// Дата совершения траты
	Date date.Date
}

// NewSpending создаёт новую трату.
func NewSpending(category Category, amount decimal.Decimal, date date.Date) Spending {
	return Spending{
		Category: category,
		Amount:   amount,
		Date:     date,
	}
}

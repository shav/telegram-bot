package finance_models

import (
	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shopspring/decimal"
)

// UserSpendLimitSettings хранит пользовательские настройки с ограничением на траты.
type UserSpendLimitSettings struct {
	// ИД пользователя.
	UserId int64
	// TODO: Прикрутить бюджеты на другие типы периодов (неделя, год)
	// Месяц, на который задаётся бюджет.
	Period date.Month
	// Бюджет на траты в указанный период времени.
	Limit decimal.Decimal
}

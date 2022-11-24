//go:generate easyjson -all spendings_report_request.go

package finance_transport_mq_models

import "github.com/shav/telegram-bot/internal/common/date"

// Название запроса на составление отчёта о тратах пользователя.
const SpendingReportRequestName = "SpendingReportRequest"

// SpendingReportRequestMessage представляет из себя запрос на составление отчёта о тратах пользователя.
type SpendingReportRequestMessage struct {
	// ИД пользователя.
	UserId int64
	// Название периода для составления отчёта.
	PeriodName string
	// Начало периода для составления отчёта.
	StartDate date.Date
	// Конец периода для составления отчёта.
	EndDate date.Date
}

// NewSpendingReportRequestMessage создаёт новый запрос на формирование отчёта
// о тратах пользователя userId за период со startDate по endDate.
func NewSpendingReportRequestMessage(userId int64, periodName string, startDate date.Date, endDate date.Date) SpendingReportRequestMessage {
	return SpendingReportRequestMessage{
		UserId:     userId,
		PeriodName: periodName,
		StartDate:  startDate,
		EndDate:    endDate,
	}
}

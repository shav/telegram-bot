package finance_commands_spendings_report

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Этапы ввода данных команды составления отчёта по тратам.
var inputStages = inputStagesEnum{
	// Этап ввода периода для составления отчёта.
	period: core_models.InputStage("Period"),
}

// Перечисление этапов ввода данных команды составления отчёта по тратам.
type inputStagesEnum struct {
	// Этап ввода периода для составления отчёта.
	period core_models.InputStage
}

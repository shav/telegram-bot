package finance_commands_set_spend_limit

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Этапы ввода данных команды установки бюджетных лимитов.
var inputStages = inputStagesEnum{
	// Этап ввода бюджетного лимита.
	spendLimit: core_models.InputStage("SpendLimit"),
}

// Перечисление этапов ввода данных команды установки бюджетных лимитов.
type inputStagesEnum struct {
	// Этап ввода бюджетного лимита.
	spendLimit core_models.InputStage
}

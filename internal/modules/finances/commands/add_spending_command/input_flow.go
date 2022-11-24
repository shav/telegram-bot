package finance_commands_add_spending

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Этапы ввода данных команды добавления траты.
var inputStages = inputStagesEnum{
	// Этап ввода категории трат.
	category: core_models.InputStage("Category"),
	// Этап ввода суммы.
	amount: core_models.InputStage("Amount"),
	// Этап ввода даты.
	date: core_models.InputStage("Date"),
}

// Перечисление этапов ввода данных команды добавления траты.
type inputStagesEnum struct {
	// Этап ввода категории трат.
	category core_models.InputStage
	// Этап ввода суммы.
	amount core_models.InputStage
	// Этап ввода даты.
	date core_models.InputStage
}

// Схема ввода данных команды добавления траты через диалог с чат-ботом.
var addSpendingInputFlow = core_models.NewInputFlowMetadata([]core_models.InputStage{
	inputStages.category, inputStages.amount, inputStages.date,
})

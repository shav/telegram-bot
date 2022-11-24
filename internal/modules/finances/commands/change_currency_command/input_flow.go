package finance_commands_change_currency

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Этапы ввода данных команды смены валюты.
var inputStages = inputStagesEnum{
	// Этап ввода валюты.
	currency: core_models.InputStage("Currency"),
}

// Перечисление этапов ввода данных команды смены валюты.
type inputStagesEnum struct {
	// Этап ввода валюты.
	currency core_models.InputStage
}

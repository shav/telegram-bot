package finance_commands_change_currency

import (
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// Опции команды смены валюты.
var options = map[core_models.InputStage][]core_models.Option{
	inputStages.currency: {
		{Value: string(finance_models.Currencies.Dollar.Code), Text: finance_models.Currencies.Dollar.Name},
		{Value: string(finance_models.Currencies.Euro.Code), Text: finance_models.Currencies.Euro.Name},
		{Value: string(finance_models.Currencies.Yuan.Code), Text: finance_models.Currencies.Yuan.Name},
		{Value: string(finance_models.Currencies.Ruble.Code), Text: finance_models.Currencies.Ruble.Name},
	},
}

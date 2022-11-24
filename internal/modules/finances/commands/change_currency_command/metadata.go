package finance_commands_change_currency

import (
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

var Metadata = core_models.CommandMetadata{
	Name:        "change_currency",
	Description: "Сменить мою текущую валюту",
	Answers:     answers,
	Options:     options,
}

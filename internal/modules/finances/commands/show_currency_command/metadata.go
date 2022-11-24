package finance_commands_show_currency

import (
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

var Metadata = core_models.CommandMetadata{
	Name:        "show_currency",
	Description: "Показать мою текущую валюту",
	Answers:     answers,
	Options:     nil,
}

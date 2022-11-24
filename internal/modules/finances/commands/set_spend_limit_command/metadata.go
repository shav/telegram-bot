package finance_commands_set_spend_limit

import (
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

var Metadata = core_models.CommandMetadata{
	Name:        "set_spend_limit",
	Description: "Установить бюджет трат на текущий месяц",
	Answers:     answers,
	Options:     nil,
}

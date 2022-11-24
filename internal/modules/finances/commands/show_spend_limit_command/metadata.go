package finance_commands_show_spend_limit

import (
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

var Metadata = core_models.CommandMetadata{
	Name:        "show_spend_limit",
	Description: "Показать лимит трат на текущий месяц",
	Answers:     answers,
	Options:     nil,
}

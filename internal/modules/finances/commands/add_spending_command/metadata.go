package finance_commands_add_spending

import (
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

var Metadata = core_models.CommandMetadata{
	Name:        "add_spending",
	Description: "Добавить трату",
	Answers:     answers,
	Options:     options,
}

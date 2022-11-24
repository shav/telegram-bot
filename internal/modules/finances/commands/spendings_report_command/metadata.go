package finance_commands_spendings_report

import (
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

var Metadata = core_models.CommandMetadata{
	Name:        "spendings_report",
	Description: "Получить отчёт по тратам",
	Answers:     answers,
	Options:     options,
}

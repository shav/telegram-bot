package core_commands_help

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

var Metadata = core_models.CommandMetadata{
	Name:        "help",
	Description: "Справка по всем доступным командам",
	Answers: map[core_models.AnswerKey]string{
		core_models.DefaultAnswer: "Можете управлять мной с помощью команд:",
	},
}

package core_commands_start

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

var Metadata = core_models.CommandMetadata{
	Name:        "start",
	Description: "Начать диалог",
	Answers: map[core_models.AnswerKey]string{
		core_models.DefaultAnswer: "Привет!",
	},
}

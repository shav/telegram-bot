package core_commands_stop

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

var Metadata = core_models.CommandMetadata{
	Name:        "stop",
	Description: "Завершить диалог",
	Answers: map[core_models.AnswerKey]string{
		core_models.DefaultAnswer: "Пока!",
	},
}

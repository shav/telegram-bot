package finance_commands_spendings_report

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Ответы для команды составления отчёта по тратам.
var answers = map[core_models.AnswerKey]string{
	a.inputPeriod:         "Выберите, пожалуйста, период, за который нужно составить отчёт по тратам",
	a.invalidPeriodError:  "Неверный период",
	a.waitForMakingReport: "Секундочку, готовим ваш отчёт...",
	a.cannotRequestReport: "Не удалось запросить отчёт о тратах: произошла ошибка",
}

var a = answersKeyEnum{
	inputPeriod:         core_models.AnswerKey("inputPeriod"),
	invalidPeriodError:  core_models.AnswerKey("invalidPeriodError"),
	waitForMakingReport: core_models.AnswerKey("waitForMakingReport"),
	cannotRequestReport: core_models.AnswerKey("cannotRequestReport"),
}

// answersKeyEnum перечисляет ответы для команды составления отчёта по тратам.
type answersKeyEnum struct {
	inputPeriod         core_models.AnswerKey
	invalidPeriodError  core_models.AnswerKey
	waitForMakingReport core_models.AnswerKey
	cannotRequestReport core_models.AnswerKey
}

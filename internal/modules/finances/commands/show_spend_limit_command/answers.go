package finance_commands_show_spend_limit

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Ответы для команды показа лимита трат.
var answers = map[core_models.AnswerKey]string{
	a.spendLimitIsNotSet:  "Лимит трат на текущий месяц не задан",
	a.spendLimitTemplate:  "Лимит трат на текущий месяц: %s",
	a.cannotGetSpendLimit: "Не удалось получить лимит трат на текущий месяц, произошла ошибка!!!",
}

var a = answersKeyEnum{
	spendLimitIsNotSet:  core_models.AnswerKey("spendLimitIsNotSet"),
	spendLimitTemplate:  core_models.AnswerKey("spendLimitTemplate"),
	cannotGetSpendLimit: core_models.AnswerKey("cannotGetSpendLimit"),
}

// answersKeyEnum перечисляет ответы для команды показа лимита трат.
type answersKeyEnum struct {
	spendLimitIsNotSet  core_models.AnswerKey
	spendLimitTemplate  core_models.AnswerKey
	cannotGetSpendLimit core_models.AnswerKey
}

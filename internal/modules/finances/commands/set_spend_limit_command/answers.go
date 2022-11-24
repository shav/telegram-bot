package finance_commands_set_spend_limit

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Ответы для команды установки бюджетных лимитов.
var answers = map[core_models.AnswerKey]string{
	a.inputSpendLimit:                  "Введите, пожалуйста, лимит трат на текущий месяц:",
	a.cannotParseNumber:                "Не могу распознать число",
	a.spendLimitShouldBeNotNegative:    "Размер лимита должен быть не меньше нуля",
	a.cannotSetSpendLimit:              "Не удалось установить лимит трат, произошла ошибка!!!",
	a.spendLimitHasBeenChangedTemplate: "Лимит трат на текущий месяц установлен в %s",
}

var a = answersKeyEnum{
	inputSpendLimit:                  core_models.AnswerKey("inputSpendLimit"),
	cannotParseNumber:                core_models.AnswerKey("cannotParseNumber"),
	spendLimitShouldBeNotNegative:    core_models.AnswerKey("spendLimitShouldBeNotNegative"),
	cannotSetSpendLimit:              core_models.AnswerKey("cannotSetSpendLimit"),
	spendLimitHasBeenChangedTemplate: core_models.AnswerKey("spendLimitHasBeenChangedTemplate"),
}

// answersKeyEnum перечисляет ответы для команды установки бюджетных лимитов.
type answersKeyEnum struct {
	inputSpendLimit                  core_models.AnswerKey
	cannotParseNumber                core_models.AnswerKey
	spendLimitShouldBeNotNegative    core_models.AnswerKey
	cannotSetSpendLimit              core_models.AnswerKey
	spendLimitHasBeenChangedTemplate core_models.AnswerKey
}

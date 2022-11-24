package finance_commands_show_currency

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Ответы для команды показа текущей валюты.
var answers = map[core_models.AnswerKey]string{
	a.activeCurrencyTemplate: "Ваша текущая валюта: %s",
	a.cannotGetUserCurrency:  "Не удалось получить текущую валюту, произошла ошибка!!!",
}

var a = answersKeyEnum{
	activeCurrencyTemplate: core_models.AnswerKey("activeCurrencyTemplate"),
	cannotGetUserCurrency:  core_models.AnswerKey("cannotGetUserCurrency"),
}

// answersKeyEnum перечисляет ответы для команды показа текущей валюты.
type answersKeyEnum struct {
	activeCurrencyTemplate core_models.AnswerKey
	cannotGetUserCurrency  core_models.AnswerKey
}

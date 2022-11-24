package finance_commands_change_currency

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Ответы для команды смены валюты.
var answers = map[core_models.AnswerKey]string{
	a.inputCurrency:                  "Выберите, пожалуйста, валюту из списка ниже:",
	a.currencyIsNotSupported:         "Выбранная вами валюта не поддерживается",
	a.cannotChangeCurrency:           "Не удалось сменить валюту, произошла ошибка!!!",
	a.currencyHasBeenChangedTemplate: "Текущая валюта изменена на %s",
}

var a = answersKeyEnum{
	inputCurrency:                  core_models.AnswerKey("inputCurrency"),
	currencyIsNotSupported:         core_models.AnswerKey("currencyIsNotSupported"),
	cannotChangeCurrency:           core_models.AnswerKey("cannotChangeCurrency"),
	currencyHasBeenChangedTemplate: core_models.AnswerKey("currencyHasBeenChangedTemplate"),
}

// answersKeyEnum перечисляет ответы для команды смены валюты.
type answersKeyEnum struct {
	inputCurrency                  core_models.AnswerKey
	currencyIsNotSupported         core_models.AnswerKey
	cannotChangeCurrency           core_models.AnswerKey
	currencyHasBeenChangedTemplate core_models.AnswerKey
}

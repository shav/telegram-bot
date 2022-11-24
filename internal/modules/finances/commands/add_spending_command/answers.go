package finance_commands_add_spending

import "github.com/shav/telegram-bot/internal/modules/core/domain/models"

// Ответы для команды добавления траты.
var answers = map[core_models.AnswerKey]string{
	a.cannotParseDate:              "Не могу распознать дату",
	a.cannotParseNumber:            "Не могу распознать число",
	a.amountShouldBePositive:       "Размер траты должен быть больше нуля",
	a.cannotAddSpending:            "Не удалось добавить трату",
	a.unknownError:                 "произошла ошибка",
	a.cannotConvertCurrency:        "Не удалось выполнить конвертацию валюты",
	a.spendLimitExceeded:           "Превышен лимит трат на текущий месяц",
	a.categoryInputRequest:         "К какой категории относится ваша трата?",
	a.amountInputRequest:           "Какую сумму вы потратили?",
	a.dateInputRequest:             "В какой день вы потратили деньги?\n(Выберите из списка ниже или введите дату вручную)",
	a.spendingHasBeenAddedTemplate: "Трата из категории \"%s\" за %s на сумму %s успешно добавлена",
}

var a = answersKeyEnum{
	cannotParseDate:              core_models.AnswerKey("cannotParseDate"),
	cannotParseNumber:            core_models.AnswerKey("cannotParseNumber"),
	amountShouldBePositive:       core_models.AnswerKey("amountShouldBePositive"),
	cannotAddSpending:            core_models.AnswerKey("cannotAddSpending"),
	unknownError:                 core_models.AnswerKey("unknownError"),
	cannotConvertCurrency:        core_models.AnswerKey("cannotConvertCurrency"),
	spendLimitExceeded:           core_models.AnswerKey("spendLimitExceeded"),
	categoryInputRequest:         core_models.AnswerKey("categoryInputRequest"),
	amountInputRequest:           core_models.AnswerKey("amountInputRequest"),
	dateInputRequest:             core_models.AnswerKey("dateInputRequest"),
	spendingHasBeenAddedTemplate: core_models.AnswerKey("spendingHasBeenAddedTemplate"),
}

// answersKeyEnum перечисляет ответы для команды добавления траты.
type answersKeyEnum struct {
	cannotParseDate              core_models.AnswerKey
	cannotParseNumber            core_models.AnswerKey
	amountShouldBePositive       core_models.AnswerKey
	cannotAddSpending            core_models.AnswerKey
	unknownError                 core_models.AnswerKey
	cannotConvertCurrency        core_models.AnswerKey
	spendLimitExceeded           core_models.AnswerKey
	categoryInputRequest         core_models.AnswerKey
	amountInputRequest           core_models.AnswerKey
	dateInputRequest             core_models.AnswerKey
	spendingHasBeenAddedTemplate core_models.AnswerKey
}

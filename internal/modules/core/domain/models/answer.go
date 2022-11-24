package core_models

// Answer является ответом на сообщение пользователя.
type Answer struct {
	// Текст ответа.
	Text string
	// Набор опций для выбора.
	Options []Option
}

// AnswerKey является ключо для строки ответа.
type AnswerKey string

// Ответ по-умолчанию.
var DefaultAnswer = AnswerKey("")

package core_services_messages

// Ответы по-умолчанию на пользовательские команды.
var defaultAnswers = defaultAnswersEnum{
	// Ответ на неизвестную команду.
	unknownCommandTemplate: "Не знаю эту команду.\nНаберите %s для получения списка доступных команд",
	commandError:           "При обработке команды произошла ошибка",
}

// defaultAnswers перечисляет ответы по-умолчанию на команды.
type defaultAnswersEnum struct {
	unknownCommandTemplate string
	commandError           string
}

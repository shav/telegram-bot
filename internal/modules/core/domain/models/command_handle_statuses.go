package core_models

// CommandStatuses содежит возможные состояния обработки команд.
var CommandStatuses = CommandHandleStatusEnum{
	// Обработка команды завершена.
	Completed: CommandHandleStatus("Completed"),
	// Ожидание сообщения от пользователя.
	WaitForNextMessage: CommandHandleStatus("WaitForNextMessage"),
	// Выполнение команды отменено.
	Canceled: CommandHandleStatus("Canceled"),
}

// CommandHandleStatus описывает состояние обработки команды.
type CommandHandleStatus string

// CommandHandleStatusEnum перечисляет состояния обработки команды.
type CommandHandleStatusEnum struct {
	Completed          CommandHandleStatus
	WaitForNextMessage CommandHandleStatus
	Canceled           CommandHandleStatus
}

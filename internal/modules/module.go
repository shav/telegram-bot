package modules

// Module представляет из себя модуль приложения.
type Module interface {
	// GetName возвращает имя модуля.
	GetName() string
	// Init выполняет первоначальную настройку модуля.
	Init(args ModuleInitArgs) error
	// InitCommands выполняет инициализацию команд модуля.
	InitCommands(args ModuleInitArgs) error
	// InitMessageQueueHandlers выполняет инициализацию обработчиков сообщений из очереди.
	InitMessageQueueHandlers(args ModuleInitArgs) error
	// Stop завершает работу модуля.
	Stop() error
}

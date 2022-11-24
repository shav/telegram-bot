package modules

import (
	"context"

	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// ModuleInitArgs хранит параметры, необходимые для первоначальной настройки модуля.
type ModuleInitArgs struct {
	// Название приложения.
	AppName string
	// Название текущего сервиса.
	ServiceName string
	// Тип конфига приложения.
	ConfigType config.Type
	// Путь к конфиг-файлу приложения.
	ConfigFile string
	// Конфиг приложения.
	Config *config.ConfigService
	// Контейнер для хранения команд.
	Commands commandsContainer
	// Контекст.
	Ctx context.Context
}

// commandsContainer используется для хранения команд чат-бота.
type commandsContainer interface {
	// RegisterCommand регистрирует команду command чат-бота и фабрику handlerFactory для создания обработчика команды.
	RegisterCommand(command core_models.CommandMetadata, handlerFactory core_models.CommandHandlerFactory)
	// GetAllCommandsDescription возвращает описание всех команд чат-бота.
	GetAllCommandsDescription() string
}

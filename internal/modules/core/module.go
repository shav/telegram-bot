package core

import (
	"github.com/shav/telegram-bot/internal/modules"
	"github.com/shav/telegram-bot/internal/modules/core/commands/help_command"
	"github.com/shav/telegram-bot/internal/modules/core/commands/start_command"
	"github.com/shav/telegram-bot/internal/modules/core/commands/stop_command"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
)

// Модуль, являющийся ядром приложения.
type coreModule struct {
}

// NewModule создаёт модуля ядра приложения.
func NewModule() *coreModule {
	return &coreModule{}
}

// GetName возвращает имя модуля.
func (m *coreModule) GetName() string {
	return "core"
}

// Init инициализирует модуль ядра приложения.
func (m *coreModule) Init(args modules.ModuleInitArgs) error {
	return nil
}

// InitCommands выполняет инициализацию команд модуля ядра приложения.
func (m *coreModule) InitCommands(args modules.ModuleInitArgs) error {
	commands := args.Commands

	// Команда "Начать диалог"
	helpInfo, err := core_commands_help.NewHandler(core_commands_help.Metadata, commands)
	if err != nil {
		logger.Error(args.Ctx, "HelpInfo init failed", logger.Fields.Error(err))
	}
	commands.RegisterCommand(core_commands_start.Metadata,
		func(c core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return core_commands_start.NewHandler(c, helpInfo)
		})

	// Команда "Завершить диалог"
	commands.RegisterCommand(core_commands_stop.Metadata,
		func(c core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return core_commands_stop.NewHandler(c)
		})

	// Команда "Показать справку"
	commands.RegisterCommand(core_commands_help.Metadata,
		func(c core_models.CommandMetadata, userId int64) (core_models.CommandHandler, error) {
			return core_commands_help.NewHandler(c, commands)
		})

	return nil
}

// InitMessageQueueHandlers выполняет инициализацию обработчиков сообщений из очереди.
func (m *coreModule) InitMessageQueueHandlers(args modules.ModuleInitArgs) error {
	return nil
}

// Stop завершает работу модуля.
func (m *coreModule) Stop() error {
	return nil
}

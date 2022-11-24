package core_services_commands

import (
	"fmt"
	"strings"
	"sync"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// Префикс команды.
const commandPrefix = "/"

// CommandContainer хранит в себе список доступных команд чат-бота.
type CommandContainer struct {
	// Список доступных команд чат-бота.
	allCommands []core_models.Command
	// Метаданные команд чат-бота.
	commandsMetadata map[core_models.Command]core_models.CommandMetadata
	// Обработчики команд чат-бота.
	// TODO: Пока считаем, что у каждой команды только один обработчик.
	// Но в будущем можно сделать и несколько обработчиков для одной команды.
	commandHandlers map[core_models.Command]core_models.CommandHandlerFactory
	// Объект синхронизации доступа к коллекции обработчиков команд.
	lock *sync.RWMutex
}

// NewContainer создаёт контейнер для хранения команд чат-бота.
func NewContainer() *CommandContainer {
	return &CommandContainer{
		allCommands:      make([]core_models.Command, 0),
		commandsMetadata: make(map[core_models.Command]core_models.CommandMetadata),
		commandHandlers:  make(map[core_models.Command]core_models.CommandHandlerFactory),
		lock:             &sync.RWMutex{},
	}
}

// RegisterCommand регистрирует команду command чат-бота и фабрику handlerFactory для создания обработчика команды.
func (c *CommandContainer) RegisterCommand(command core_models.CommandMetadata, handlerFactory core_models.CommandHandlerFactory) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.allCommands = append(c.allCommands, command.Name)
	c.commandsMetadata[command.Name] = command
	c.commandHandlers[command.Name] = handlerFactory
}

// GetCommandHandler возвращает обработчик команды command для пользователя userId
func (c *CommandContainer) GetCommandHandler(command core_models.Command, userId int64) (core_models.CommandHandler, error) {
	// Чтение из колллекции команд не синхронизируем, т.к. регистрация команд происходит только единожды при старте приложения.
	// TODO: Прикрутить пул обработчиков команд
	commandName := c.parseCommandName(command)
	if handlerFactory, ok := c.commandHandlers[commandName]; ok {
		commandMetadata := c.commandsMetadata[commandName]
		return handlerFactory(commandMetadata, userId)
	}
	return nil, nil
}

// GetAllCommandsDescription возвращает описание всех команд чат-бота.
func (c *CommandContainer) GetAllCommandsDescription() string {
	// Чтение из колллекции команд не синхронизируем, т.к. регистрация команд происходит только единожды при старте приложения.
	sb := strings.Builder{}
	for _, commandName := range c.allCommands {
		command := c.commandsMetadata[commandName]
		sb.WriteString(fmt.Sprintf("%s - %s\n", c.FormatCommandName(command.Name), command.Description))
	}
	return sb.String()
}

// IsLikeCommand проверяет, похожа ли строка на команду чат-бота (т.е. имеет формат команды).
func (c *CommandContainer) IsLikeCommand(text string) bool {
	return strings.HasPrefix(text, commandPrefix)
}

// FormatCommandName форматирует отображаемое имя команды (с добавлением префикса).
func (c *CommandContainer) FormatCommandName(command core_models.Command) string {
	var commandName = string(command)
	if strings.HasPrefix(commandName, commandPrefix) {
		return commandName
	}
	return commandPrefix + commandName
}

// parseCommandName парсит "голое" имя команды (без префикса).
func (c *CommandContainer) parseCommandName(command core_models.Command) core_models.Command {
	return core_models.Command(strings.TrimPrefix(string(command), commandPrefix))
}

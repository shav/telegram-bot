package core_commands_help

import (
	"context"
	"fmt"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// ICommandsInfo содержит информацию о доступных командах бота.
type ICommandsInfo interface {
	// GetAllCommandsDescription возвращает описание всех команд чат-бота.
	GetAllCommandsDescription() string
}

// HelpCommandHandler является обработчиком команды получения справки по всем доступным командам.
type HelpCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Информация о командах.
	info ICommandsInfo
}

// NewHandler создаёт новый обработчик команды получения справки.
func NewHandler(command core_models.CommandMetadata, info ICommandsInfo) (*HelpCommandHandler, error) {
	return &HelpCommandHandler{command: command, info: info}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *HelpCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return []core_models.Answer{{Text: h.GetHelpMessage()}}, core_models.CommandStatuses.Completed, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *HelpCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return nil, core_models.CommandStatuses.Completed, nil
}

// GetHelpMessage возвращает справочное сообщение по всем доступным командам бота.
func (h *HelpCommandHandler) GetHelpMessage() string {
	return fmt.Sprintf("%s\n%s", h.command.GetDefaultAnswer(), h.info.GetAllCommandsDescription())
}

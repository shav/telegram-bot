package core_commands_start

import (
	"context"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// ICommandsHelp позволяет получить справочную информацию о доступных командах бота.
type ICommandsHelp interface {
	// GetHelpMessage возвращает справочное сообщение по всем доступным командам бота.
	GetHelpMessage() string
}

// StartCommandHandler является обработчиком команды старта диалога.
type StartCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Справка по командам.
	help ICommandsHelp
}

// NewHandler создаёт новый обработчик команды старта диалога.
func NewHandler(command core_models.CommandMetadata, help ICommandsHelp) (*StartCommandHandler, error) {
	return &StartCommandHandler{command: command, help: help}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *StartCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	answers = []core_models.Answer{{Text: h.command.GetDefaultAnswer()}}
	if h.help != nil {
		answers = append(answers, core_models.Answer{Text: h.help.GetHelpMessage()})
	}
	return answers, core_models.CommandStatuses.Completed, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *StartCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return nil, core_models.CommandStatuses.Completed, nil
}

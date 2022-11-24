package core_commands_stop

import (
	"context"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// StopCommandHandler является обработчиком команды старта диалога.
type StopCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
}

// NewHandler создаёт новый обработчик команды завршения диалога.
func NewHandler(command core_models.CommandMetadata) (*StopCommandHandler, error) {
	return &StopCommandHandler{command: command}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *StopCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return []core_models.Answer{{Text: h.command.GetDefaultAnswer()}}, core_models.CommandStatuses.Completed, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *StopCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return nil, core_models.CommandStatuses.Completed, nil
}

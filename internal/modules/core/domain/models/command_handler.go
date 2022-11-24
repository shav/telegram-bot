package core_models

import (
	"context"
)

// CommandHandler является обработчиком конкретной команды чат-бота.
type CommandHandler interface {
	// StartHandleCommand начинает обработку команды.
	StartHandleCommand(ctx context.Context) (answers []Answer, status CommandHandleStatus, err error)
	// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
	HandleNextMessage(ctx context.Context, message string) (answers []Answer, status CommandHandleStatus, err error)
}

// CommandHandlerFactory является фабрикой по созданию обработчиков команд.
type CommandHandlerFactory = func(command CommandMetadata, userId int64) (CommandHandler, error)

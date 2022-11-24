package finance_commands_show_spend_limit

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// useCase описывает пользовательский сценарий получения лимита трат.
type useCase interface {
	// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period.
	GetSpendLimit(ctx context.Context, userId int64, period date.Month) (limit finance_models.Amount, exists bool, err error)
}

// ShowSpendLimitCommandHandler является обработчиком команды показа лимита трат.
type ShowSpendLimitCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Пользователь, который выполняет команду.
	userId int64
	// Пользовательский сценарий.
	useCase useCase
}

// NewHandler создаёт новый обработчик команды показа лимита трат.
func NewHandler(command core_models.CommandMetadata, userId int64, useCase useCase) (*ShowSpendLimitCommandHandler, error) {
	if useCase == nil {
		return nil, errors.New("New ShowSpendLimitCommandHandler: useCase is not assigned")
	}

	return &ShowSpendLimitCommandHandler{
		command: command,
		userId:  userId,
		useCase: useCase,
	}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *ShowSpendLimitCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	spendLimit, exists, err := h.getSpendLimit(ctx, h.userId, date.ThisMonth())
	var answer string
	if err != nil {
		answer = h.command.Answers[a.cannotGetSpendLimit]
	} else if !exists {
		answer = h.command.Answers[a.spendLimitIsNotSet]
	} else {
		answer = fmt.Sprintf(h.command.Answers[a.spendLimitTemplate], spendLimit)
	}
	return []core_models.Answer{{Text: answer}}, core_models.CommandStatuses.Completed, err
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *ShowSpendLimitCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return nil, core_models.CommandStatuses.Completed, nil
}

// getSpendLimit возвращает для пользователя userId бюджет на указанный период времени period.
func (h *ShowSpendLimitCommandHandler) getSpendLimit(ctx context.Context, userId int64, period date.Month) (limit finance_models.Amount, exists bool, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "ShowSpendLimitCommandHandler.GetSpendLimit")
	defer span.Finish()
	span.SetTag("user", userId)

	spendLimit, exists, err := h.useCase.GetSpendLimit(ctx, userId, period)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Get spend limit failed", logger.Fields.Error(err))
		return spendLimit, exists, err
	}
	return spendLimit, exists, nil
}

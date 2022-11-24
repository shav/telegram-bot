//go:generate minimock -i useCase -o ./mocks/ -s ".go"

package finance_commands_set_spend_limit

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/numbers"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// useCase описывает пользовательский сценарий для установки бюджетного лимита.
type useCase interface {
	// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
	// Возаращает размер установленного лимита в валюте пользователя.
	SetSpendLimit(ctx context.Context, userId int64, limit decimal.Decimal, period date.Month) (finance_models.Amount, error)
}

// SetSpendLimitCommandHandler является обработчиком команды установки бюджетного лимита.
type SetSpendLimitCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Пользователь, который выполняет команду.
	userId int64
	// Пользовательский сценарий.
	useCase useCase
}

// NewHandler создаёт новый экземпляр обработчика команды установки бюджетного лимита.
func NewHandler(command core_models.CommandMetadata, userId int64, useCase useCase) (*SetSpendLimitCommandHandler, error) {
	if useCase == nil {
		return nil, errors.New("New SetSpendLimitCommandHandler: useCase is not assigned")
	}

	return &SetSpendLimitCommandHandler{
		command: command,
		userId:  userId,
		useCase: useCase,
	}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *SetSpendLimitCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return []core_models.Answer{{
			Text: h.command.Answers[a.inputSpendLimit],
		}},
		core_models.CommandStatuses.WaitForNextMessage, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *SetSpendLimitCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	spendLimit, errorMessage := h.inputDataFrom(ctx, message)
	if errorMessage != "" {
		return []core_models.Answer{
			{Text: errorMessage},
			{Text: h.command.Answers[a.inputSpendLimit]},
		}, core_models.CommandStatuses.WaitForNextMessage, nil
	}

	newSpendLimit, err := h.setSpendLimit(ctx, h.userId, spendLimit, date.ThisMonth())
	if err != nil {
		return []core_models.Answer{{Text: h.command.Answers[a.cannotSetSpendLimit]}}, core_models.CommandStatuses.WaitForNextMessage, err
	}

	return []core_models.Answer{{
			Text: fmt.Sprintf(h.command.Answers[a.spendLimitHasBeenChangedTemplate], newSpendLimit),
		}},
		core_models.CommandStatuses.Completed, nil
}

// inputDataFrom распознаёт данные из пользовательского сообщения message.
func (h *SetSpendLimitCommandHandler) inputDataFrom(ctx context.Context, message string) (spendLimit decimal.Decimal, errMessage string) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SetSpendLimitCommandHandler.ParseInputData")
	defer span.Finish()

	message = strings.TrimSpace(message)
	spendLimit, err := numbers.ParseDecimal(message)

	if err != nil {
		tracing.SetError(span)
		logger.Info(ctx, "Cannot parse spend limit", logger.Fields.Error(err))
		return decimal.Decimal{}, h.command.Answers[a.cannotParseNumber]
	}

	if spendLimit.LessThan(decimal.Zero) {
		tracing.SetError(span)
		err = errors.New(fmt.Sprintf("Spend limit should be not negative (current value = %s)", finance_models.FormatMoney(spendLimit)))
		logger.Info(ctx, "Cannot input spend limit", logger.Fields.Error(err))
		return decimal.Decimal{}, h.command.Answers[a.spendLimitShouldBeNotNegative]
	}

	return spendLimit, ""
}

// setSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
func (h *SetSpendLimitCommandHandler) setSpendLimit(ctx context.Context, userId int64, limit decimal.Decimal, period date.Month) (finance_models.Amount, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SetSpendLimitCommandHandler.SetSpendLimit")
	defer span.Finish()

	newSpendLimit, err := h.useCase.SetSpendLimit(ctx, userId, limit, period)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Set spend limit failed", logger.Fields.Error(err))
		return newSpendLimit, err
	}
	return newSpendLimit, nil
}

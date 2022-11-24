//go:generate minimock -i useCase -o ./mocks/ -s ".go"

package finance_commands_show_currency

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyCurrency = finance_models.Currency{}

// useCase описывает пользовательский сценарий получения текущей валюты пользователя.
type useCase interface {
	// GetUserCurrency возвращает текущую валюту пользователя userId.
	GetUserCurrency(ctx context.Context, userId int64) (finance_models.Currency, error)
}

// ShowCurrencyCommandHandler является обработчиком команды показа текущей валюты.
type ShowCurrencyCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Пользователь, который выполняет команду.
	userId int64
	// Пользовательский сценарий получения отчёта по тратам.
	useCase useCase
}

// NewHandler создаёт новый обработчик команды показа текущей валюты.
func NewHandler(command core_models.CommandMetadata, userId int64, useCase useCase) (*ShowCurrencyCommandHandler, error) {
	if useCase == nil {
		return nil, errors.New("New ShowCurrencyCommandHandler: useCase is not assigned")
	}

	return &ShowCurrencyCommandHandler{
		command: command,
		userId:  userId,
		useCase: useCase,
	}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *ShowCurrencyCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	userCurrency, err := h.getUserCurrency(ctx, h.userId)
	var answer string
	if err != nil {
		answer = h.command.Answers[a.cannotGetUserCurrency]
	} else {
		answer = fmt.Sprintf(h.command.Answers[a.activeCurrencyTemplate], strings.ToLower(userCurrency.Name))
	}
	return []core_models.Answer{{Text: answer}}, core_models.CommandStatuses.Completed, err
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *ShowCurrencyCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return nil, core_models.CommandStatuses.Completed, nil
}

// getUserCurrency возвращает текущую валюту пользователя userId.
func (h *ShowCurrencyCommandHandler) getUserCurrency(ctx context.Context, userId int64) (finance_models.Currency, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "ShowCurrencyCommandHandler.GetUserCurrency")
	defer span.Finish()
	span.SetTag("user", userId)

	userCurrency, err := h.useCase.GetUserCurrency(ctx, userId)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Get user currency failed", logger.Fields.Error(err))
		return emptyCurrency, err
	}
	return userCurrency, nil
}

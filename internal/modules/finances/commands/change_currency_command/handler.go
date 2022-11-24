//go:generate minimock -i useCase -o ./mocks/ -s ".go"

package finance_commands_change_currency

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyCurrency = finance_models.Currency{}

// useCase описывает пользовательский сценарий для смены валюты.
type useCase interface {
	// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
	ChangeCurrency(ctx context.Context, userId int64, newCurrency finance_models.Currency) error
}

// ChangeCurrencyCommandHandler является обработчиком команды смены валюты для отображения расчётов.
type ChangeCurrencyCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Пользователь, который выполняет команду.
	userId int64
	// Пользовательский сценарий смены валюты.
	useCase useCase
}

// NewHandler создаёт новый экземпляр обработчика команды смены валюты.
func NewHandler(command core_models.CommandMetadata, userId int64, useCase useCase) (*ChangeCurrencyCommandHandler, error) {
	if useCase == nil {
		return nil, errors.New("New ChangeCurrencyCommandHandler: useCase is not assigned")
	}

	return &ChangeCurrencyCommandHandler{
		command: command,
		userId:  userId,
		useCase: useCase,
	}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *ChangeCurrencyCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return []core_models.Answer{{
			Text:    h.command.Answers[a.inputCurrency],
			Options: h.command.Options[inputStages.currency],
		}},
		core_models.CommandStatuses.WaitForNextMessage, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *ChangeCurrencyCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	newCurrency, errorMessage := h.inputDataFrom(ctx, message)
	if errorMessage != "" {
		return []core_models.Answer{{Text: errorMessage}}, core_models.CommandStatuses.WaitForNextMessage, nil
	}

	err = h.changeCurrency(ctx, h.userId, newCurrency)
	if err != nil {
		return []core_models.Answer{{Text: h.command.Answers[a.cannotChangeCurrency]}}, core_models.CommandStatuses.WaitForNextMessage, err
	}

	return []core_models.Answer{{
			Text: fmt.Sprintf(h.command.Answers[a.currencyHasBeenChangedTemplate], strings.ToLower(newCurrency.Name)),
		}},
		core_models.CommandStatuses.WaitForNextMessage, nil
}

// inputDataFrom распознаёт данные из пользовательского сообщения message.
// Возвращает введённые данные и сообщение для пользователя об ошибке распознавания ввода.
func (h *ChangeCurrencyCommandHandler) inputDataFrom(ctx context.Context, message string) (currency finance_models.Currency, errMessage string) {
	span, ctx := tracing.StartSpanFromContext(ctx, "ChangeCurrencyCommandHandler.ParseInputData")
	defer span.Finish()
	span.SetTag("inputStage", "currency")

	currency, err := finance_models.ParseCurrency(message)
	if err != nil {
		logger.Info(ctx, "Cannot parse currency", logger.Fields.Error(err))
		tracing.SetError(span)
		return emptyCurrency, h.command.Answers[a.currencyIsNotSupported]
	}

	return currency, ""
}

// changeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
func (h *ChangeCurrencyCommandHandler) changeCurrency(ctx context.Context, userId int64, newCurrency finance_models.Currency) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "ChangeCurrencyCommandHandler.ChangeCurrency")
	defer span.Finish()
	span.SetTag("user", userId)

	err := h.useCase.ChangeCurrency(ctx, userId, newCurrency)
	if err != nil {
		tracing.SetError(span)
		logger.Error(ctx, "Change user currency failed", logger.Fields.Error(err))
		return err
	}
	return nil
}

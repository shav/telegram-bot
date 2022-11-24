//go:generate minimock -i useCase -o ./mocks/ -s ".go"
//go:generate minimock -i currencySettings -o ./mocks/ -s ".go"

package finance_commands_add_spending

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/common/numbers"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/services/input"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// useCase описывает пользовательский сценарий добавления траты.
type useCase interface {
	// AddUserSpending добавляет трату spending пользователя userId.
	// Возаращает размер траты в валюте пользователя.
	AddUserSpending(ctx context.Context, userId int64, spending finance_models.Spending) (finance_models.Amount, error)
}

// AddSpendingCommandHandler является обработчиком команды добавления траты.
type AddSpendingCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Пользователь, который выполняет команду.
	userId int64
	// Пользовательский сценарий добавления траты.
	useCase useCase
	// Состояние процесса ввода данных.
	inputFlow *core_services_input.FlowProcess
	// Введённые пользователем данные по трате.
	spending finance_models.Spending
	// Объект синхронизации шагов выполнения команды.
	lock *sync.Mutex
}

// NewHandler создаёт новый экземпляр обработчика команды добавления траты.
func NewHandler(command core_models.CommandMetadata, userId int64, useCase useCase) (*AddSpendingCommandHandler, error) {
	if useCase == nil {
		return nil, errors.New("New AddSpendingCommandHandler: useCase is not assigned")
	}

	return &AddSpendingCommandHandler{
		command:   command,
		userId:    userId,
		useCase:   useCase,
		inputFlow: core_services_input.NewFlowProcess(addSpendingInputFlow),
		lock:      &sync.Mutex{},
	}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *AddSpendingCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	h.inputFlow.Start()
	answer, options := h.getInputRequestAnswer(h.inputFlow.GetCurrentStage())
	return []core_models.Answer{{Text: answer, Options: options}}, core_models.CommandStatuses.WaitForNextMessage, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *AddSpendingCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()

	errMessage := h.inputDataFrom(ctx, message)
	if errMessage != "" {
		answer, options := h.getInputRequestAnswer(h.inputFlow.GetCurrentStage())
		return []core_models.Answer{
				{Text: errMessage},
				{Text: answer, Options: options},
			},
			core_models.CommandStatuses.WaitForNextMessage, nil
	}

	answer, options, status, err := h.goToNextStage(ctx)
	return []core_models.Answer{{Text: answer, Options: options}}, status, err
}

// inputDataFrom распознаёт данные из пользовательского сообщения message, в зависимости от текущей фазы диалога.
// Возвращает сообщение для пользователя об ошибке распознавания ввода.
func (h *AddSpendingCommandHandler) inputDataFrom(ctx context.Context, message string) (errMessage string) {
	span, ctx := tracing.StartSpanFromContext(ctx, "AddSpendingCommandHandler.ParseInputData")
	defer span.Finish()

	message = strings.TrimSpace(message)
	currentInputStage := h.inputFlow.GetCurrentStage()
	span.SetTag("inputStage", string(currentInputStage))

	switch currentInputStage {
	case inputStages.category:
		h.spending.Category = finance_models.ParseCategory(message)
	case inputStages.date:
		date, err := date.ParseDate(message)
		if err != nil {
			logger.Info(ctx, "Cannot parse date", logger.Fields.Error(err))
			tracing.SetError(span)
			errMessage = h.command.Answers[a.cannotParseDate]
			break
		}
		h.spending.Date = date
	case inputStages.amount:
		amount, err := numbers.ParseDecimal(message)
		if err != nil {
			logger.Info(ctx, "Cannot parse amount", logger.Fields.Error(err))
			tracing.SetError(span)
			errMessage = h.command.Answers[a.cannotParseNumber]
			break
		}
		if amount.LessThanOrEqual(decimal.Zero) {
			err = errors.New(fmt.Sprintf("Amount of spending should be positive (current value = %s)", finance_models.FormatMoney(amount)))
			logger.Info(ctx, "Cannot input amount", logger.Fields.Error(err))
			tracing.SetError(span)
			errMessage = h.command.Answers[a.amountShouldBePositive]
			break
		}
		h.spending.Amount = amount
	}
	return errMessage
}

// goToNextStage выполняет переход состояния диалога к следующему этапу - либо ввод, либо сохранение данных.
// Возвращает ответ для пользователя и новое состояние диалога.
func (h *AddSpendingCommandHandler) goToNextStage(ctx context.Context) (answer string, options []core_models.Option, status core_models.CommandHandleStatus, err error) {
	h.inputFlow.GoToNextStage()
	if h.inputFlow.GetCurrentStage() == core_models.EmptyInputStage {
		status = core_models.CommandStatuses.Completed
		answer, err = h.addUserSpending(ctx, h.spending)
	} else {
		status = core_models.CommandStatuses.WaitForNextMessage
		answer, options = h.getInputRequestAnswer(h.inputFlow.GetCurrentStage())
	}
	return answer, options, status, err
}

// addUserSpending добавляет трату пользователя.
func (h *AddSpendingCommandHandler) addUserSpending(ctx context.Context, spending finance_models.Spending) (answer string, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "AddSpendingCommandHandler.AddSpending")
	defer span.Finish()
	span.SetTag("user", h.userId)

	spendingAmount, err := h.useCase.AddUserSpending(ctx, h.userId, h.spending)
	if err != nil {
		var cce *finance_models.CurrencyConvertError
		if errors.Is(err, finance_models.SpendLimitExceededError) {
			// Превышение лимита трат является бизнес-ошибкой, а не ошибкой в работе программы
			logger.Info(ctx, "Cannot add spending", logger.Fields.Error(err))
			tracing.SetError(span)
			answer = fmt.Sprintf("%s: \n%s", h.command.Answers[a.cannotAddSpending], h.command.Answers[a.spendLimitExceeded])
			return answer, nil
		} else if errors.As(err, &cce) {
			logger.Error(ctx, "Add spending failed", logger.Fields.Error(err))
			tracing.SetError(span)
			answer = fmt.Sprintf("%s: \n%s", h.command.Answers[a.cannotAddSpending], h.command.Answers[a.cannotConvertCurrency])
			return answer, err
		} else {
			logger.Error(ctx, "Add spending failed", logger.Fields.Error(err))
			tracing.SetError(span)
			answer = fmt.Sprintf("%s: %s", h.command.Answers[a.cannotAddSpending], h.command.Answers[a.unknownError])
			return answer, err
		}
	}

	return fmt.Sprintf(h.command.Answers[a.spendingHasBeenAddedTemplate], spending.Category, spending.Date, spendingAmount), nil
}

// getInputRequestAnswer возвращает сообщение, которое просит пользователя ввести конкретное поле данных,
// соответствующее этапу ввода stage.
func (h *AddSpendingCommandHandler) getInputRequestAnswer(stage core_models.InputStage) (message string, options []core_models.Option) {
	switch stage {
	case inputStages.category:
		return h.command.Answers[a.categoryInputRequest], h.command.Options[inputStages.category]
	case inputStages.date:
		return h.command.Answers[a.dateInputRequest], h.command.Options[inputStages.date]
	case inputStages.amount:
		return h.command.Answers[a.amountInputRequest], nil
	}
	return "", nil
}

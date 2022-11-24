//go:generate minimock -i useCase -o ./mocks/ -s ".go"

package finance_commands_spendings_report

import (
	"context"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyPeriod = date.Period("")

// useCase описывает пользовательский сценарий получения отчёта по тратам.
type useCase interface {
	// RequestSpendingReport запрашивает формирование отчёта по тратам пользователя userId
	// за указанный период времени period, сгруппированный по категориям.
	RequestSpendingReport(ctx context.Context, userId int64, period date.Period) error
}

// SpendingsReportCommandHandler является обработчиком команды получения отчёта по тратам.
type SpendingsReportCommandHandler struct {
	// Метаданные команды.
	command core_models.CommandMetadata
	// Пользователь, который выполняет команду.
	userId int64
	// Пользовательский сценарий получения отчёта по тратам.
	useCase useCase
}

// NewHandler создаёт новый экземпляр обработчика команды добавления траты.
func NewHandler(command core_models.CommandMetadata, userId int64, useCase useCase) (*SpendingsReportCommandHandler, error) {
	if useCase == nil {
		return nil, errors.New("New SpendingsReportCommandHandler: useCase is not assigned")
	}

	return &SpendingsReportCommandHandler{
		command: command,
		userId:  userId,
		useCase: useCase,
	}, nil
}

// StartHandleCommand начинает обработку команды.
func (h *SpendingsReportCommandHandler) StartHandleCommand(ctx context.Context) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	return []core_models.Answer{{
			Text:    h.command.Answers[a.inputPeriod],
			Options: h.command.Options[inputStages.period],
		}},
		core_models.CommandStatuses.WaitForNextMessage, nil
}

// HandleNextMessage обрабатывает следующее сообщение в рамках активного процесса обработки.
func (h *SpendingsReportCommandHandler) HandleNextMessage(ctx context.Context, message string) (answers []core_models.Answer, status core_models.CommandHandleStatus, err error) {
	period, errorMessage := h.inputDataFrom(ctx, message)
	if errorMessage != "" {
		return []core_models.Answer{{Text: errorMessage}}, core_models.CommandStatuses.WaitForNextMessage, nil
	}

	var answer string
	err = h.useCase.RequestSpendingReport(ctx, h.userId, period)
	if err != nil {
		answer = h.command.Answers[a.cannotRequestReport]
	} else {
		answer = h.command.Answers[a.waitForMakingReport]
	}
	return []core_models.Answer{{Text: answer}}, core_models.CommandStatuses.WaitForNextMessage, err
}

// inputDataFrom распознаёт данные из пользовательского сообщения message.
// Возвращает введённые данные и сообщение для пользователя об ошибке распознавания ввода.
func (h *SpendingsReportCommandHandler) inputDataFrom(ctx context.Context, message string) (period date.Period, errMessage string) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingsReportCommandHandler.ParseInputData")
	defer span.Finish()
	span.SetTag("inputStage", "period")

	period, err := date.ParsePeriod(message)
	if err != nil {
		logger.Info(ctx, "Cannot parse period", logger.Fields.Error(err))
		tracing.SetError(span)
		return emptyPeriod, h.command.Answers[a.invalidPeriodError]
	}
	return period, ""
}

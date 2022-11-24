//go:generate minimock -i spendingReportBuilder -o ./mocks/ -s ".go"
//go:generate minimock -i reportSender -o ./mocks/ -s ".go"
//go:generate minimock -i deserializer -o ./mocks/ -s ".go"

package finance_transport_mq_handlers

import (
	"context"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// spendingReportBuilder выполняет построение финансовых отчётов о тратах.
type spendingReportBuilder interface {
	// GetSpendingReport возвращает отчёт о тратах пользователя userId за указанный период времени dateInterval.
	GetSpendingReport(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) (core_models.Report, error)
}

// Отправитель отчётов клиенту.
type reportSender interface {
	// SendReport отправляет пользователю userId отчёт report.
	SendReport(ctx context.Context, userId int64, report core_models.Report) error
}

// deserializer выполняет десериализацию объектов из тела сообщения.
type deserializer interface {
	// Unmarshal выполняет десериализацию сериализованного объекта serializedObj.
	Unmarshal(serializedObj []byte, object any) error
}

// SpendingReportMessageHandler является обработчиком сообщений на формирование отчёта о тратах.
type SpendingReportMessageHandler struct {
	// Построитель отчётов.
	reportBuilder spendingReportBuilder
	// Отправитель отчёта клиенту.
	reportSender reportSender
	// Десериализатор сообщений.
	deserializer deserializer
}

// NewSpendingReportMessageHandler создаёт новый обработчик сообщений на формирование отчёта о тратах.
func NewSpendingReportMessageHandler(reportBuilder spendingReportBuilder, reportSender reportSender, deserializer deserializer) (*SpendingReportMessageHandler, error) {
	if reportBuilder == nil {
		return nil, errors.New("New SpendingReportMessageHandler: reportBuilder is not assigned")
	}
	if reportSender == nil {
		return nil, errors.New("New SpendingReportMessageHandler: reportSender is not assigned")
	}
	if deserializer == nil {
		return nil, errors.New("New SpendingReportMessageHandler: message deserializer is not assigned")
	}

	return &SpendingReportMessageHandler{
		reportBuilder: reportBuilder,
		reportSender:  reportSender,
		deserializer:  deserializer,
	}, nil
}

// HandleMessage выполняет обработку сообщения message с ключом key.
func (h *SpendingReportMessageHandler) HandleMessage(ctx context.Context, key string, messageData []byte) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingReportMessageHandler.HandleMessage")
	defer span.Finish()

	message := finance_transport_mq_models.SpendingReportRequestMessage{}
	err := h.deserializer.Unmarshal(messageData, &message)
	if err != nil {
		tracing.SetError(span)
		return errors.Wrap(err, "Message deserialization failed")
	}

	dateInterval := date.NewInterval(message.StartDate, message.EndDate)
	report, err := h.reportBuilder.GetSpendingReport(ctx, message.UserId, message.PeriodName, dateInterval)
	if err != nil {
		tracing.SetError(span)
		return errors.Wrap(err, "Build spending report failed")
	}

	err = h.reportSender.SendReport(ctx, message.UserId, report)
	if err != nil {
		// TODO: Если не удалось сразу отправить отчёт пользователю,
		// то нужно положить его в долговременное хранилище (БД или очередь) и попытаться отправить чуть позже.
		tracing.SetError(span)
		return errors.Wrap(err, "Send spending report to client failed")
	}

	return nil
}

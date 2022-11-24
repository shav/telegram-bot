//go:generate minimock -i messageQueueSender -o ./mocks/ -s ".go"
//go:generate minimock -i serializer -o ./mocks/ -s ".go"

package finance_clients

import (
	"context"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// messageQueueSender занимается отправкой сообщений в очередь.
type messageQueueSender interface {
	// SendMessageAsync асинхронно отправляет сообщение payload с ключом key в очередь queue.
	SendMessageAsync(ctx context.Context, queue string, key string, payload []byte)
}

// serializer выполняет сериализацию объектов для отправки в теле сообщения.
type serializer interface {
	// Marshal выполняет сериализацию object-а.
	Marshal(object any) ([]byte, error)
}

// ReportsClient является клиентом для сервиса отчётов.
type ReportsClient struct {
	// Отправитель сообщений в очередь.
	mq messageQueueSender
	// Сериализатор объектов.
	serializer serializer
}

// NewReportsClient создаёт новый клиент для сервиса отчётов.
func NewReportsClient(mq messageQueueSender, serializer serializer) (*ReportsClient, error) {
	if mq == nil {
		return nil, errors.New("New ReportsClient: message queue client is not provided")
	}
	if serializer == nil {
		return nil, errors.New("New ReportsClient: serializer is not provided")
	}

	return &ReportsClient{
		mq:         mq,
		serializer: serializer,
	}, nil
}

// RequestSpendingReport запрашивает формирование отчёта о тратах пользователя у сервиса отчётов.
func (c *ReportsClient) RequestSpendingReport(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "ReportsClient.RequestSpendingReport")
	defer span.Finish()

	spendingReportRequest := finance_transport_mq_models.NewSpendingReportRequestMessage(userId, periodName, dateInterval.Start(), dateInterval.End())
	requestBytes, err := c.serializer.Marshal(spendingReportRequest)
	if err != nil {
		tracing.SetError(span)
		return err
	}
	c.mq.SendMessageAsync(ctx, finance_transport_mq.ReportsQueueName, finance_transport_mq_models.SpendingReportRequestName, requestBytes)
	return nil
}

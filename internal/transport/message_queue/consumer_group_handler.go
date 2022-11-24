package message_queue

import (
	"context"

	"github.com/Shopify/sarama"

	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Кастомный обработчик сообщений.
type messageHandler interface {
	// HandleMessage выполняет обработку сообщения message с ключом key.
	HandleMessage(ctx context.Context, key string, message []byte) error
}

// Обработчик сообщений группы получателей.
type consumerGroupHandler struct {
	// Контекст.
	ctx context.Context
	// Кастомный обработчик сообщений.
	messageHandler messageHandler
}

// newConsumerGroupHandler создаёт новый обработчик сообщений группы получаталей.
func newConsumerGroupHandler(ctx context.Context, messageHandler messageHandler) *consumerGroupHandler {
	return &consumerGroupHandler{
		ctx:            ctx,
		messageHandler: messageHandler,
	}
}

// Setup выполняется перед началом обработки сообщений.
func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup выполняется после завершения всех обработчиков сообщений.
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim потребляет сообщения из очереди.
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.handleMessage(message)
		session.MarkMessage(message, "")
	}

	return nil
}

// handleMessage обрабатывает сообщение message.
func (h *consumerGroupHandler) handleMessage(message *sarama.ConsumerMessage) {
	span, ctx := tracing.StartSpanFromContext(h.ctx, "MessageQueueConsumer.HandleMessage")
	defer span.Finish()

	key := string(message.Key)
	logger.Info(ctx, "Received message with {key} from {queue} at {offset} of {partition}",
		logger.Fields.String("queue", message.Topic), logger.Fields.String("key", key),
		logger.Fields.Int32("partition", message.Partition), logger.Fields.Int64("offset", message.Offset))

	if h.messageHandler != nil {
		err := h.messageHandler.HandleMessage(ctx, key, message.Value)
		if err != nil {
			logger.Error(ctx, "Failed handling message with {key} from {queue} at {offset} of {partition}",
				logger.Fields.String("queue", message.Topic), logger.Fields.String("key", key),
				logger.Fields.Int32("partition", message.Partition), logger.Fields.Int64("offset", message.Offset))
		}
	}
}

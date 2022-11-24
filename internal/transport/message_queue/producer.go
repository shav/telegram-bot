package message_queue

import (
	"context"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// MessageQueueProducer отправляет сообщения в очередь сообщений.
type MessageQueueProducer struct {
	// Название текущего приложения.
	appName string
	// Асинхронный продьюсер сообщений в очередь.
	asyncProducer sarama.AsyncProducer
}

// NewProducer создаёт отправителя в очередь сообщений, который подключается к брокерам очередей сообщений brokers.
func NewProducer(ctx context.Context, appName string, brokers []string) (*MessageQueueProducer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Idempotent = true
	config.Producer.Return.Successes = true
	if config.Producer.Idempotent {
		config.Producer.Retry.Max = 1
		config.Net.MaxOpenRequests = 1
	}

	brokersList := strings.Join(brokers, "; ")
	logger.Info(ctx, "Connecting to message queue {brokers}...", logger.Fields.String("brokers", brokersList))
	asyncProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "creating message queue async producer")
	}
	logger.Info(ctx, "Connection to message queue brokers successfully established")

	producer := &MessageQueueProducer{
		appName:       strings.TrimSpace(appName),
		asyncProducer: asyncProducer,
	}

	go producer.handleMessageSendSuccess(ctx)
	go producer.handleMessageSendError(ctx)

	return producer, nil
}

// SendMessageAsync асинхронно отправляет сообщение payload с ключом key в очередь queue.
func (p *MessageQueueProducer) SendMessageAsync(ctx context.Context, queue string, key string, payload []byte) {
	span, _ := tracing.StartSpanFromContext(ctx, "MessageQueueProducer.SendMessageAsync")
	defer span.Finish()

	queue = getFullQueueName(p.appName, queue)
	message := sarama.ProducerMessage{Topic: queue}
	if key != "" {
		message.Key = sarama.StringEncoder(key)
	}
	if payload != nil {
		message.Value = sarama.ByteEncoder(payload)
	}

	p.asyncProducer.Input() <- &message
}

// Close завершает работу отправителя соообщений.
func (p *MessageQueueProducer) Close() error {
	if p.asyncProducer != nil {
		return p.asyncProducer.Close()
	}
	return nil
}

// handleMessageSendSuccess обрабатывает успешно отправленные сообщения.
func (p *MessageQueueProducer) handleMessageSendSuccess(ctx context.Context) {
	for message := range p.asyncProducer.Successes() {
		key, err := message.Key.Encode()
		if err != nil {
			logger.Warn(ctx, "Parse message key failed", logger.Fields.Error(err))
		}
		logger.Info(ctx, "Message with {key} has sent to message {queue} at {offset} of {partition}",
			logger.Fields.String("queue", message.Topic), logger.Fields.String("key", string(key)),
			logger.Fields.Int32("partition", message.Partition), logger.Fields.Int64("offset", message.Offset))
	}
}

// handleMessageSendSuccess обрабатывает ошибки отправки сообщений.
func (p *MessageQueueProducer) handleMessageSendError(ctx context.Context) {
	for err := range p.asyncProducer.Errors() {
		message := err.Msg
		key, err := message.Key.Encode()
		if err != nil {
			logger.Warn(ctx, "Parse message key failed", logger.Fields.Error(err))
		}
		logger.Error(ctx, "Failed to send message with {key} to message {queue}",
			logger.Fields.String("queue", message.Topic), logger.Fields.String("key", string(key)), logger.Fields.Error(err))
	}
}

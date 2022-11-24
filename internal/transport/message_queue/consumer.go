package message_queue

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/observability/logger"
)

// Стратегия балансировки получателей сообщений в рамках группы.
type BalanceStrategy string

const (
	BalanceStrategySticky     = BalanceStrategy("Sticky")
	BalanceStrategyRoundRobin = BalanceStrategy("RoundRobin")
	BalanceStrategyRange      = BalanceStrategy("Range ")
)

// MessageQueueConsumer получает сообщения из очереди сообщений.
type MessageQueueConsumer struct {
	// Название текущего приложения.
	appName string
	// Группа получателей сообщений из очереди.
	consumerGroup sarama.ConsumerGroup
}

// NewConsumer создаёт получателя сообщений из очередей, который подключается к брокерам очередей сообщений brokers.
func NewConsumer(ctx context.Context, appName string, brokers []string, consumerGroupName string, balanceStrategy BalanceStrategy) (*MessageQueueConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	bs, err := convertBalanceStrategy(balanceStrategy)
	if err != nil {
		return nil, err
	}
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{bs}

	brokersList := strings.Join(brokers, "; ")
	logger.Info(ctx, "Connecting to message queue {brokers}...", logger.Fields.String("brokers", brokersList))
	consumerGroup, err := sarama.NewConsumerGroup(brokers, consumerGroupName, config)
	if err != nil {
		return nil, errors.Wrap(err, "creating consumer group")
	}
	logger.Info(ctx, "Connection to message queue brokers successfully established")

	return &MessageQueueConsumer{
		appName:       appName,
		consumerGroup: consumerGroup,
	}, nil
}

// Subscribe подписывается на получение сообщений из очереди queue.
func (c *MessageQueueConsumer) Subscribe(ctx context.Context, queue string, handler messageHandler) error {
	queue = getFullQueueName(c.appName, queue)
	go func() {
		// В документации к методу consumerGroup.Consume() написано, что нужно потдключаться к очереди в цикле, на случай перебалансировки.
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := c.consumerGroup.Consume(ctx, []string{queue}, newConsumerGroupHandler(ctx, handler))
				if err != nil {
					logger.Error(ctx, "consuming via handler failed", logger.Fields.Error(err))
					// TODO: Пока до конца непонятно, что делать в случае ошибки подключения к очереди сообщений -
					// выходить или продолжать пытаться считывать дальше (возможно с экспоненциальными интервалами между переповторами)
					return
				}
			}
		}
	}()
	return nil
}

// Close завершает работу получателя соообщений.
func (c *MessageQueueConsumer) Close() error {
	if c.consumerGroup != nil {
		return c.consumerGroup.Close()
	}
	return nil
}

// convertBalanceStrategy преобразует стратегию балансировки в формат библиотеки sarama.
func convertBalanceStrategy(balanceStrategy BalanceStrategy) (sarama.BalanceStrategy, error) {
	switch balanceStrategy {
	case BalanceStrategySticky:
		return sarama.BalanceStrategySticky, nil
	case BalanceStrategyRoundRobin:
		return sarama.BalanceStrategyRoundRobin, nil
	case BalanceStrategyRange:
		return sarama.BalanceStrategyRange, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported consumer group balance strategy: %s", balanceStrategy))
	}
}

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Счётчик входящих сообщений.
var IncomingMessagesCount *incomingMessagesCount

// Тип счётчика входящих сообщений.
type incomingMessagesCount struct {
	counter *prometheus.CounterVec
}

// Создаёт счётчик входящих сообщений.
func newIncomingMessagesCount() *incomingMessagesCount {
	return &incomingMessagesCount{
		counter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "incoming_messages_total",
				Help: "Total count of incoming messages by command.",
			},
			[]string{commandLabel},
		),
	}
}

// Inc увеличивает на единицу счётчик входящих сообщений для команды command.
func (c *incomingMessagesCount) Inc(command string) {
	if c != nil { // В тестах метрики не собираются.
		c.counter.With(prometheus.Labels{commandLabel: command}).Inc()
	}
}

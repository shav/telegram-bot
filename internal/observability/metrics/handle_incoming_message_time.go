package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Таймер длительности обработки входящих сообщений.
var IncomingMessageResponseTime *incomingMessageResponseTime

// Тип таймера длительности обработки входящих сообщений.
type incomingMessageResponseTime struct {
	gauge     *prometheus.GaugeVec
	histogram *prometheus.HistogramVec
}

// Создаёт таймер длительности обработки входящих сообщений.
func newIncomingMessageResponseTime() *incomingMessageResponseTime {
	return &incomingMessageResponseTime{
		gauge: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "incoming_message_response_time_ms",
				Help: "Time of response to incoming message by command.",
			},
			[]string{commandLabel},
		),
		histogram: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "histogram_incoming_message_response_time_ms",
				Help: "Time histogram of response to incoming message by command.",
				Buckets: []float64{
					1, 5, 10, 50, 100, 150, 200, 250, 300, 350, 400, 450, 500, 550, 650, 700, 750, 800, 850, 900, 950, 1000,
					1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000,
				},
			},
			[]string{commandLabel},
		),
	}
}

// Set устанавливает длительность обработки входящего сообщения для команды command.
func (c *incomingMessageResponseTime) Set(command string, duration time.Duration) {
	if c != nil { // В тестах метрики не собираются.
		durationMs := float64(duration.Milliseconds())
		c.gauge.With(prometheus.Labels{commandLabel: command}).Set(durationMs)
		c.histogram.With(prometheus.Labels{commandLabel: command}).Observe(durationMs)
	}
}

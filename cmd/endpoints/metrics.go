package cmd_endpoints

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Конечная точка доступа к метрикам приложения.
type MetricsEndpoint struct {
}

// NewMetrics создаёт конечную точку доступа к метрикам приложения.
func NewMetrics() *MetricsEndpoint {
	return &MetricsEndpoint{}
}

// Listen слушает на порту port запросы к точке доступа с метриками приложения.
func (e *MetricsEndpoint) Listen(port int) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

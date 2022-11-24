package tracing

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// Обёртка для измерения длительности операций.
type Span = opentracing.Span

// traceCloser завершает работу трейсинга.
var traceCloser io.Closer

// Init выполняет инициализацию трейсинга на сервисе serviceName, с долей записи сообщений sampling (от 0 до 1.0).
func Init(serviceName string, sampling float64) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}

	cfg.Sampler = &config.SamplerConfig{
		Type:  jaeger.SamplerTypeConst,
		Param: sampling,
	}

	traceCloser, err = cfg.InitGlobalTracer(serviceName)
	return err
}

// Stop завершает работу трейсинга.
func Stop() {
	if traceCloser != nil {
		traceCloser.Close()
	}
}

// StartSpanFromContext начинает измерение длительности операции operationName в контексте ctx.
func StartSpanFromContext(ctx context.Context, operationName string) (Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName)
}

// SetError устанавливает в span признак ошибки.
func SetError(span Span) {
	ext.Error.Set(span, true)
}

package tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/shav/telegram-bot/internal/observability/logger"
)

// AddTraceIdToLog добавляет в запись лога поле с ИД трейса.
func AddTraceIdToLog(ctx context.Context, fields ...logger.Field) []logger.Field {
	if ctx == nil {
		return fields
	}

	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return fields
	}

	if spanContext, ok := span.Context().(jaeger.SpanContext); ok {
		traceId := spanContext.TraceID().String()
		return append(fields, logger.Fields.String("trace", traceId))
	}

	return fields
}

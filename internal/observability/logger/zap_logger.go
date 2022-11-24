package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// NewZapLogger создаёт движок логирования на основе библиотеки Zap.
func NewZapLogger(logMode LogMode) (*zap.Logger, error) {
	options := []zap.Option{
		zap.AddCallerSkip(1),
	}

	switch strings.ToLower(string(logMode)) {
	case strings.ToLower(string(DevelopmentLogMode)):
		return zap.NewDevelopment(options...)
	case strings.ToLower(string(ProductionLogMode)):
		return zap.NewProduction(options...)
	default:
		return nil, fmt.Errorf("unknown log mode \"%s\"", logMode)
	}
}

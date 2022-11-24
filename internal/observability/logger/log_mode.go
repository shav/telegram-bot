package logger

// Профиль логирования.
type LogMode string

const (
	DevelopmentLogMode = LogMode("dev")
	ProductionLogMode  = LogMode("prod")
)

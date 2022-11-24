package report_service

import "github.com/shav/telegram-bot/cmd"

// Настройки по-умолчанию сервиса отчётов.
var DefaultSettings = cmd.ServiceSettings{
	ServiceName:          "telegram_bot_report",
	ConfigFile:           "config/report/config.yaml",
	MetricsPort:          8081,
	ReportSenderGrpcPort: 9200,
	ReportSenderHttpPort: 9201,
}

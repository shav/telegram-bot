package bot_service

import "github.com/shav/telegram-bot/cmd"

// Настройки по-умолчанию сервиса чат-бота.
var DefaultSettings = cmd.ServiceSettings{
	ServiceName: "telegram_bot",
	ConfigFile:  "config/bot/config.yaml",
	MetricsPort: 8080,
}

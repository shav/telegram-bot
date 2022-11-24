package config

import "github.com/shav/telegram-bot/internal/observability/logger"

// Профиль логирования по умолчанию.
var defaultLogMode = logger.ProductionLogMode

// Доля записываемых сообщений трейсинга по-умолчанию.
var defaultTraceSampling = 1.0

// Config содержит настройки приложения.
type Config struct {
	// Название сервиса.
	ServiceName string `yaml:"service_name" envconfig:"SERVICE_NAME"`
	// Токен доступа к АПИ telegram.
	Token string `yaml:"token" envconfig:"TOKEN"`
	// Строка подключения к БД.
	DbConnectionString string `yaml:"db_connection_string" envconfig:"DB_CONNECTION_STRING"`
	// Строка подключения к сервису кеширования.
	CacheConnectionString string `yaml:"cache_connection_string" envconfig:"CACHE_CONNECTION_STRING"`
	// Профиль логирования.
	LogMode string `yaml:"log_mode" envconfig:"LOG_MODE"`
	// HACK: Чтобы отличить нулевое значение от случая, когда настройка вообще отсутствует в конфиге, используем указатель.
	// Доля записываемых сообщений трейсинга.
	TraceSampling *float64 `yaml:"trace_sampling" envconfig:"TRACE_SAMPLING"`
	// Порт, по которому приложение предоставляет метрики.
	MetricsPort *int `yaml:"metrics_port" envconfig:"METRICS_PORT"`
	// Адреса брокеров очередей сообщений.
	MessageQueueBrokers []string `yaml:"message_queue_brokers" envconfig:"MESSAGE_QUEUE_BROKERS"`
	// Grpc-порт, по которому приложение отправляет отчёты пользователям.
	ReportSenderGrpcPort *int `yaml:"report_sender_grpc_port" envconfig:"REPORT_SENDER_GRPC_PORT"`
	// Http-порт, по которому приложение отправляет отчёты пользователям.
	ReportSenderHttpPort *int `yaml:"report_sender_http_port" envconfig:"REPORT_SENDER_HTTP_PORT"`
	// Адрес сервиса для отправки отчётов пользователям.
	ReportSenderAddress string `yaml:"sender_address" envconfig:"SENDER_ADDRESS"`
}

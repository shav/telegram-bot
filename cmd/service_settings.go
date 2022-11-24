package cmd

// ServiceSettings хранит настройки сервисов.
type ServiceSettings struct {
	// Название текущего сервиса по-умолчанию.
	ServiceName string
	// Путь к config-файлу по-умолчанию.
	ConfigFile string
	// Порт по-умолчанию, по которому приложение предоставляет метрики.
	MetricsPort int
	// Grpc-порт по-умолчанию, по которому приложение отправляет отчёты пользователям.
	ReportSenderGrpcPort int
	// Http-порт по-умолчанию, по которому приложение отправляет отчёты пользователям.
	ReportSenderHttpPort int
}

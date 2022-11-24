package config

import (
	"strings"

	"github.com/shav/telegram-bot/internal/observability/logger"
)

// ConfigService для получения настроек приложения.
type ConfigService struct {
	// Настройки из config-а.
	config Config
	// Название приложения.
	serviceName string
}

// newConfigService создаёт сервис настроек приложения serviceName.
func newConfigService(serviceName string, config Config) *ConfigService {
	return &ConfigService{
		config:      config,
		serviceName: serviceName,
	}
}

// ServiceName возвращает из конфига имя сервиса.
func (s *ConfigService) ServiceName() string {
	serviceName := strings.TrimSpace(s.config.ServiceName)
	if serviceName == "" {
		return s.serviceName
	}
	return serviceName
}

// Token возвращает из конфига токен доступа.
func (s *ConfigService) Token() string {
	return strings.TrimSpace(s.config.Token)
}

// DbConnectionString возвращает из конфига строку подключения к БД.
func (s *ConfigService) DbConnectionString() string {
	return strings.TrimSpace(s.config.DbConnectionString)
}

// CacheConnectionString возвращает из конфига строку подключения к сервису кеширования.
func (s *ConfigService) CacheConnectionString() string {
	return strings.TrimSpace(s.config.CacheConnectionString)
}

// LogMode возвращает из конфига профиль логирования.
func (s *ConfigService) LogMode() logger.LogMode {
	rawLogMode := strings.TrimSpace(s.config.LogMode)
	if rawLogMode == "" {
		return defaultLogMode
	}
	return logger.LogMode(rawLogMode)
}

// TraceSampling возвращает из конфига долю записываемых сообщений трейсинга.
func (s *ConfigService) TraceSampling() float64 {
	traceSampling := s.config.TraceSampling
	if traceSampling == nil {
		return defaultTraceSampling
	}
	return *traceSampling
}

// MetricsPort возвращает из конфига порт, по которому приложение предоставляет метрики.
func (s *ConfigService) MetricsPort(defaultPort int) int {
	metricsPort := s.config.MetricsPort
	if metricsPort == nil {
		return defaultPort
	}
	return *metricsPort
}

// MessageQueueBrokers возвращает адреса брокеров очередей сообщений.
func (s *ConfigService) MessageQueueBrokers() []string {
	brokers := s.config.MessageQueueBrokers
	if brokers == nil {
		return make([]string, 0)
	}
	return brokers
}

// ReportSenderGrpcPort возвращает из конфига grpc-порт, по которому приложение отправляет отчёты пользователям.
func (s *ConfigService) ReportSenderGrpcPort(defaultPort int) int {
	reportSenderPort := s.config.ReportSenderGrpcPort
	if reportSenderPort == nil {
		return defaultPort
	}
	return *reportSenderPort
}

// ReportSenderHttpPort возвращает из конфига http-порт, по которому приложение отправляет отчёты пользователям.
func (s *ConfigService) ReportSenderHttpPort(defaultPort int) int {
	reportSenderPort := s.config.ReportSenderHttpPort
	if reportSenderPort == nil {
		return defaultPort
	}
	return *reportSenderPort
}

// ReportSenderAddress возвращает адрес сервиса для отправки отчётов пользователям.
func (s *ConfigService) ReportSenderAddress() string {
	return strings.TrimSpace(s.config.ReportSenderAddress)
}

package finance_config

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/config"
)

// New создаёт финансовые настройки приложения serviceName.
func New(configType config.Type, serviceName string, configFile string) (*configService, error) {
	switch configType {
	case config.File:
		return NewFileConfig(configFile)
	case config.Env:
		return NewEnvConfig(serviceName)
	default:
		return nil, fmt.Errorf("unsupported config source: %s", configType)
	}
}

// NewFileConfig создаёт финансовые настройки приложения serviceName на основе конфиг-файла configFile.
func NewFileConfig(configFile string) (*configService, error) {
	config, err := config.NewFileSource[Config](configFile)
	if err != nil {
		return nil, errors.Wrap(err, "Create file config source failed")
	}
	return newConfigService(*config), nil
}

// NewEnvConfig создаёт финансовые настройки приложения serviceName на основе переменных окружения.
func NewEnvConfig(serviceName string) (*configService, error) {
	config, err := config.NewEnvSource[Config](serviceName)
	if err != nil {
		return nil, errors.Wrap(err, "Create env config source failed")
	}
	return newConfigService(*config), nil
}

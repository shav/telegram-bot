package config

import (
	"fmt"

	"github.com/pkg/errors"
)

// New создаёт настройки приложения serviceName.
func New(configType Type, serviceName string, configFile string) (*ConfigService, error) {
	switch configType {
	case File:
		return NewFileConfig(serviceName, configFile)
	case Env:
		return NewEnvConfig(serviceName)
	default:
		return nil, fmt.Errorf("unsupported config source: %s", configType)
	}
}

// NewFileConfig создаёт настройки приложения serviceName на основе конфиг-файла configFile.
func NewFileConfig(serviceName string, configFile string) (*ConfigService, error) {
	config, err := NewFileSource[Config](configFile)
	if err != nil {
		return nil, errors.Wrap(err, "Create file config source failed")
	}
	return newConfigService(serviceName, *config), nil
}

// NewEnvConfig создаёт настройки приложения serviceName на основе переменных окружения.
func NewEnvConfig(serviceName string) (*ConfigService, error) {
	config, err := NewEnvSource[Config](serviceName)
	if err != nil {
		return nil, errors.Wrap(err, "Create env config source failed")
	}
	return newConfigService(serviceName, *config), nil
}

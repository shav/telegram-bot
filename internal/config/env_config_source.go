package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// NewEnvSource создаёт конфиг приложения serviceName на основе переменных окружения.
func NewEnvSource[TConfig any](serviceName string) (*TConfig, error) {
	var config TConfig
	err := envconfig.Process(serviceName, &config)
	if err != nil {
		return nil, errors.Wrap(err, "New EnvConfigSource")
	}
	return &config, nil
}

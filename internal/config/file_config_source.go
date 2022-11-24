package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// NewFileSource создаёт конфиг приложения на основе файла configFile.
func NewFileSource[TConfig any](configFile string) (*TConfig, error) {
	config, err := ReadFile[TConfig](configFile)
	if err != nil {
		return nil, errors.Wrap(err, "New FileConfigSource")
	}

	return &config, nil
}

// ReadFile считывает настройки приложения типа TConfig из файла конфига configFile.
func ReadFile[TConfig any](configFile string) (TConfig, error) {
	var config TConfig
	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return config, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &config)
	if err != nil {
		return config, errors.Wrap(err, "parsing yaml")
	}

	return config, nil
}

package finance_config

import (
	"time"
)

// Периодичность обновления курсов валют по-умолчанию.
const defaultCurrencyRatesUpdatePeriod = time.Duration(10 * time.Minute)

// Config содержит финансовые настройки приложения.
type Config struct {
	// Основная валюта для выполнения расчётов.
	DefaultCurrency string `yaml:"default_currency" envconfig:"DEFAULT_CURRENCY"`
	// Периодичность обновления курсов валют.
	CurrencyRatesUpdatePeriod time.Duration `yaml:"currency_rates_update_period" envconfig:"CURRENCY_RATES_UPDATE_PERIOD"`
}

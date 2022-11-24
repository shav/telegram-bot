package finance_config

import "time"

// configService для получения финансовых настроек приложения.
type configService struct {
	// Настройки из config-а.
	config Config
}

// newConfigService создаёт сервис финансовых настроек приложения.
func newConfigService(config Config) *configService {
	return &configService{config}
}

// DefaultCurrency возвращает из конфига основную валюту для расчётов.
func (s *configService) DefaultCurrency() string {
	return s.config.DefaultCurrency
}

// CurrencyRatesUpdatePeriod возвращает из конфига периодичность обновления курсов валют.
func (s *configService) CurrencyRatesUpdatePeriod() time.Duration {
	updatePeriod := s.config.CurrencyRatesUpdatePeriod
	if updatePeriod == time.Duration(0) {
		return defaultCurrencyRatesUpdatePeriod
	}
	return updatePeriod
}

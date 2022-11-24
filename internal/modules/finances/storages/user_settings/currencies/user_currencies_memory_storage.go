package finance_storages_user_currency_settings

import (
	"context"
	"sync"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyCurrency = finance_models.Currency{}

// UserCurrenciesMemoryStorage хранит пользовательские настройки валют в памяти.
type UserCurrenciesMemoryStorage struct {
	// Коллекция валютных настроек пользователей.
	settings map[int64]*finance_models.UserCurrencySettings
	// Объект синхронизации доступа к коллекции настроек.
	lock *sync.RWMutex
}

// NewMemoryStorage создаёт новый экземпляр хранилища пользовательских настроек валют в памяти.
func NewMemoryStorage() *UserCurrenciesMemoryStorage {
	return &UserCurrenciesMemoryStorage{
		settings: make(map[int64]*finance_models.UserCurrencySettings),
		lock:     &sync.RWMutex{},
	}
}

// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
func (s *UserCurrenciesMemoryStorage) ChangeCurrency(ctx context.Context, ts tr.Transaction, userId int64, newCurrency finance_models.Currency) error {
	span, _ := tracing.StartSpanFromContext(ctx, "UserCurrenciesMemoryStorage.ChangeCurrency")
	defer span.Finish()

	s.lock.Lock()
	defer s.lock.Unlock()

	userSettings, exists := s.settings[userId]
	if !exists {
		userSettings = &finance_models.UserCurrencySettings{UserId: userId}
		s.settings[userId] = userSettings
	}
	userSettings.Currency = newCurrency
	return nil
}

// GetCurrency возвращает текущую валюту для пользователя userId,
// а также признак того, задана ли в настройках пользователя текущая валюта.
func (s *UserCurrenciesMemoryStorage) GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (currency finance_models.Currency, exists bool, err error) {
	span, _ := tracing.StartSpanFromContext(ctx, "UserCurrenciesMemoryStorage.GetCurrency")
	defer span.Finish()

	s.lock.RLock()
	defer s.lock.RUnlock()

	userSettings, exists := s.settings[userId]
	if !exists {
		return emptyCurrency, false, nil
	}
	return userSettings.Currency, true, nil
}

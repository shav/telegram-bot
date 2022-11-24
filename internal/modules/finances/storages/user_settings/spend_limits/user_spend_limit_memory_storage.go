package finance_storages_user_spend_limit_settings

import (
	"context"
	"sync"

	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/date"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// UserSpendLimitMemoryStorage хранит пользовательские настройки бюджетов на траты в памяти.
type UserSpendLimitMemoryStorage struct {
	// Коллекция настроек бюджетов пользователей.
	settings map[int64][]*finance_models.UserSpendLimitSettings
	// Объект синхронизации доступа к коллекции настроек.
	lock *sync.RWMutex
}

// NewMemoryStorage создаёт новый экземпляр хранилища пользовательских настроек бюджетов в памяти.
func NewMemoryStorage() *UserSpendLimitMemoryStorage {
	return &UserSpendLimitMemoryStorage{
		settings: make(map[int64][]*finance_models.UserSpendLimitSettings),
		lock:     &sync.RWMutex{},
	}
}

// SetSpendLimit устанавливает в настройках пользователя userId бюджет limit на период времени period.
func (s *UserSpendLimitMemoryStorage) SetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, limit decimal.Decimal, period date.Month) error {
	span, _ := tracing.StartSpanFromContext(ctx, "UserCurrenciesMemoryStorage.SetSpendLimit")
	defer span.Finish()

	s.lock.Lock()
	defer s.lock.Unlock()

	userSettings, exists := s.settings[userId]
	if !exists {
		userSettings = make([]*finance_models.UserSpendLimitSettings, 0)
		s.settings[userId] = userSettings
	}

	spendLimit := searchSettingsForPeriod(userSettings, period)
	if spendLimit == nil {
		spendLimit = &finance_models.UserSpendLimitSettings{UserId: userId}
		userSettings = append(userSettings, spendLimit)
		s.settings[userId] = userSettings
	}
	spendLimit.Period = period
	spendLimit.Limit = limit
	return nil
}

// GetSpendLimit возвращает для пользователя userId бюджет на указанный период времени period,
// а также признак того, задан ли в настройках пользователя бюджет на указанный период.
func (s *UserSpendLimitMemoryStorage) GetSpendLimit(ctx context.Context, ts tr.Transaction, userId int64, period date.Month) (limit decimal.Decimal, exists bool, err error) {
	span, _ := tracing.StartSpanFromContext(ctx, "UserCurrenciesMemoryStorage.GetSpendLimit")
	defer span.Finish()

	s.lock.RLock()
	defer s.lock.RUnlock()

	userSettings, exists := s.settings[userId]
	if !exists {
		return decimal.Zero, false, nil
	}

	spendLimit := searchSettingsForPeriod(userSettings, period)
	if spendLimit != nil {
		return spendLimit.Limit, true, nil
	}

	return decimal.Zero, false, nil
}

// searchSettingsForPeriod ищет среди всех настроек settings ту, которая хранит бюджет на указанный период period.
func searchSettingsForPeriod(settings []*finance_models.UserSpendLimitSettings, period date.Month) *finance_models.UserSpendLimitSettings {
	for _, s := range settings {
		if s.Period == period {
			return s
		}
	}
	return nil
}

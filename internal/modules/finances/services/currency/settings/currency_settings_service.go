//go:generate minimock -i userCurrencies -o ./mocks/ -s ".go"

package finance_services_currency

import (
	"context"

	"github.com/pkg/errors"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

var emptyCurrency = finance_models.Currency{}

// userCurrencies хранит пользовательские настройки валют.
type userCurrencies interface {
	// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
	ChangeCurrency(ctx context.Context, ts tr.Transaction, userId int64, newCurrency finance_models.Currency) error
	// GetCurrency возвращает текущую валюту для пользователя userId.
	GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (currency finance_models.Currency, exists bool, err error)
}

// CurrencySettingsService реализует бизнес-логику управления настройками валют.
type CurrencySettingsService struct {
	// Валюта по-умолчанию.
	defaultCurrency finance_models.Currency
	// Пользовательские настройки валют.
	userCurrencies userCurrencies
}

// NewSettingService создаёт новый экземпляр сервиса управления настройками валют.
func NewSettingService(defaultCurrency finance_models.Currency, userCurrencies userCurrencies) (*CurrencySettingsService, error) {
	if userCurrencies == nil {
		return nil, errors.New("New CurrencySettingsService: user currency settings is not assigned")
	}

	return &CurrencySettingsService{
		defaultCurrency: defaultCurrency,
		userCurrencies:  userCurrencies,
	}, nil
}

// ChangeCurrency меняет в настройках пользователя userId текущую валюту на другую newCurrency.
func (s *CurrencySettingsService) ChangeCurrency(ctx context.Context, ts tr.Transaction, userId int64, newCurrency finance_models.Currency) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencySettingsService.ChangeCurrency")
	defer span.Finish()

	err := s.userCurrencies.ChangeCurrency(ctx, ts, userId, newCurrency)
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

// GetCurrency возвращает текущую валюту для пользователя userId.
func (s *CurrencySettingsService) GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (finance_models.Currency, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencySettingsService.GetCurrency")
	defer span.Finish()

	activeCurrency, exists, err := s.userCurrencies.GetCurrency(ctx, ts, userId)
	if err != nil {
		tracing.SetError(span)
		return emptyCurrency, err
	}
	if exists {
		return activeCurrency, nil
	}
	return s.defaultCurrency, nil
}

// GetDefaultCurrency возвращает основную валюту для расчётов.
func (s *CurrencySettingsService) GetDefaultCurrency() finance_models.Currency {
	return s.defaultCurrency
}

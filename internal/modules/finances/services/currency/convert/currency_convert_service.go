//go:generate minimock -i currencySettings -o ./mocks/ -s ".go"
//go:generate minimock -i currencyRateGetter -o ./mocks/ -s ".go"

package finance_services_currency

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// currencySettings хранит валютные настройки.
type currencySettings interface {
	// GetCurrency возвращает текущую валюту для пользователя userId.
	GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (finance_models.Currency, error)
	// GetDefaultCurrency возвращает основную валюту для расчётов.
	GetDefaultCurrency() finance_models.Currency
}

// currencyRateGetter получает курсы валют.
type currencyRateGetter interface {
	// GetRate возвращает курс конвертации из основной валюты приложения в целевую валюту currency.
	GetRate(ctx context.Context, currency finance_models.Currency) (finance_models.CurrencyRate, error)
}

type converter func(amount decimal.Decimal, rate decimal.Decimal) decimal.Decimal

// CurrencyConvertService выполняет конвертацию денежных сумм из одной валюты в другую.
type CurrencyConvertService struct {
	// Валютные настройки.
	currencySettings currencySettings
	// Курсы валют.
	currencyRates currencyRateGetter
}

// NewConvertService создаёт новый экземпляр сервиса конвертации валют.
func NewConvertService(currencySettings currencySettings, currencyRates currencyRateGetter) (*CurrencyConvertService, error) {
	if currencySettings == nil {
		return nil, errors.New("New CurrencyConvertService: currency settings is not assigned")
	}
	if currencyRates == nil {
		return nil, errors.New("New CurrencyConvertService: currency rates getter is not assigned")
	}

	return &CurrencyConvertService{
		currencySettings: currencySettings,
		currencyRates:    currencyRates,
	}, nil
}

// ConvertToUserCurrency конвертирует денежную сумму amount из основной валюты приложения в валюту пользователя userId.
func (s *CurrencyConvertService) ConvertToUserCurrency(ctx context.Context, userId int64, defaultAmount decimal.Decimal) (decimal.Decimal, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyConverter.ConvertToUserCurrency")
	defer span.Finish()

	rate, err := s.getDefaultToUserCurrencyRate(ctx, userId)
	if err != nil {
		tracing.SetError(span)
		return decimal.Zero, err
	}
	if rate == finance_models.One {
		return defaultAmount, nil
	}
	return convertToUserCurrency(defaultAmount, rate), nil
}

// ConvertToDefaultCurrency конвертирует денежную сумму amount из валюты пользователя в основную валюту приложения.
func (s *CurrencyConvertService) ConvertToDefaultCurrency(ctx context.Context, userId int64, userAmount decimal.Decimal) (decimal.Decimal, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyConverter.ConvertToDefaultCurrency")
	defer span.Finish()

	rate, err := s.getDefaultToUserCurrencyRate(ctx, userId)
	if err != nil {
		tracing.SetError(span)
		return decimal.Zero, err
	}
	if rate == finance_models.One {
		return userAmount, nil
	}
	return convertToDefaultCurrency(userAmount, rate), nil
}

// ConvertToUserCurrencyMany конвертирует денежные суммы amounts из основной валюты расчётов в валюту пользователя userId.
func (s *CurrencyConvertService) ConvertToUserCurrencyMany(ctx context.Context, userId int64, defaultAmounts ...decimal.Decimal) ([]decimal.Decimal, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyConverter.ConvertToUserCurrencyMany")
	defer span.Finish()

	rate, err := s.getDefaultToUserCurrencyRate(ctx, userId)
	if err != nil {
		tracing.SetError(span)
		return nil, err
	}
	if rate == finance_models.One {
		return defaultAmounts, nil
	}
	return s.convertMany(rate, defaultAmounts, convertToUserCurrency), nil
}

// ConvertToDefaultCurrencyMany конвертирует денежные суммы amounts из валюты пользователя в основную валюту приложения.
func (s *CurrencyConvertService) ConvertToDefaultCurrencyMany(ctx context.Context, userId int64, userAmounts ...decimal.Decimal) ([]decimal.Decimal, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyConverter.ConvertToDefaultCurrencyMany")
	defer span.Finish()

	rate, err := s.getDefaultToUserCurrencyRate(ctx, userId)
	if err != nil {
		tracing.SetError(span)
		return nil, err
	}
	if rate == finance_models.One {
		return userAmounts, nil
	}
	return s.convertMany(rate, userAmounts, convertToDefaultCurrency), nil
}

// ConvertSpendingsTableToUserCurrency конвертирует таблицу расходов spendingsTable в валюту пользователя userId.
func (s *CurrencyConvertService) ConvertSpendingsTableToUserCurrency(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) (finance_models.SpendingsByCategoryTable, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CurrencyConverter.ConvertSpendingsTableToUserCurrency")
	defer span.Finish()

	rate, err := s.getDefaultToUserCurrencyRate(ctx, userId)
	if err != nil {
		tracing.SetError(span)
		return nil, err
	}
	if rate == finance_models.One {
		return spendingsTable, nil
	}

	result := make(finance_models.SpendingsByCategoryTable, len(spendingsTable))
	for category, amount := range spendingsTable {
		result[category] = convertToUserCurrency(amount, rate)
	}
	return result, nil
}

// convertMany конвертирует денежные суммы amounts из одной валюты в другую по курсу rate.
func (s *CurrencyConvertService) convertMany(rate decimal.Decimal, amounts []decimal.Decimal, converter converter) []decimal.Decimal {
	result := make([]decimal.Decimal, len(amounts))
	for i := 0; i < len(amounts); i++ {
		result[i] = converter(amounts[i], rate)
	}
	return result
}

// convertToUserCurrency выполняет конвертацию денежной суммы defaultAmount из основной валюты в пользовательскую валюту по курсу rate.
func convertToUserCurrency(defaultAmount decimal.Decimal, rate decimal.Decimal) decimal.Decimal {
	return defaultAmount.Div(rate)
}

// convertToDefaultCurrency выполняет конвертацию денежной суммы userAmount из пользовательской валюты в основную валюту по курсу rate.
func convertToDefaultCurrency(userAmount decimal.Decimal, rate decimal.Decimal) decimal.Decimal {
	return userAmount.Mul(rate)
}

// getDefaultToUserCurrencyRate возвращает курс основной валюты приложения к валюте пользователя.
func (s *CurrencyConvertService) getDefaultToUserCurrencyRate(ctx context.Context, userId int64) (decimal.Decimal, error) {
	userCurrency, err := s.currencySettings.GetCurrency(ctx, nil, userId)
	if err != nil {
		return decimal.Zero, err
	}
	r, err := s.currencyRates.GetRate(ctx, userCurrency)
	return r.Rate, err
}

package finance_models

import (
	"github.com/shopspring/decimal"
	"time"
)

// Единица типа decimal.
var One = decimal.NewFromInt(1)

// CurrencyRate представляет из себя курс обмена валюты к основной валюте приложения
// по формуле Rate = DefaultCurrency / Currency.
type CurrencyRate struct {
	// Валюта.
	Currency Currency
	// Курс.
	Rate decimal.Decimal
	// Время, на которое актуален курс.
	Timestamp time.Time
}

// NewActualCurrencyRate создаёт актуальный на данный момент курс обмена валюты currency на основную валюту приложения.
func NewActualCurrencyRate(currency Currency, rate decimal.Decimal) CurrencyRate {
	return NewCurrencyRate(currency, rate, time.Now())
}

// NewCurrencyRate создаёт курс обмена валюты currency на основную валюту приложения, актуальный на момент времени time.
func NewCurrencyRate(currency Currency, rate decimal.Decimal, timestamp time.Time) CurrencyRate {
	return CurrencyRate{
		Currency:  currency,
		Rate:      rate,
		Timestamp: timestamp,
	}
}

// GetIdentityCurrencyRate возвращает курс 1:1 для конвертации валюты currency саму в себя.
func GetIdentityCurrencyRate(currency Currency) CurrencyRate {
	return CurrencyRate{
		Currency: currency,
		Rate:     One,
	}
}

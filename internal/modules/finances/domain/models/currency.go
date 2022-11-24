package finance_models

import (
	"github.com/pkg/errors"
	"strings"
)

var unknownCurrencyError = errors.New("currency is unknown or not supported")

// Currencies содержит список валют.
var Currencies = currencyEnum{
	Ruble:  Currency{CurrencyCode("RUB"), "Рубль", "₽"},
	Dollar: Currency{CurrencyCode("USD"), "Доллар", "$"},
	Euro:   Currency{CurrencyCode("EUR"), "Евро", "€"},
	Yuan:   Currency{CurrencyCode("CNY"), "Юань", "¥"},
}

// Список всех валют.
var AllCurrencies = [...]Currency{Currencies.Ruble, Currencies.Yuan, Currencies.Euro, Currencies.Dollar}

// CurrencyCode - это код валюты.
type CurrencyCode string

// Currency хранит информацию о валюте.
type Currency struct {
	// Код.
	Code CurrencyCode
	// Название.
	Name string
	// Символ.
	Symbol string
}

// currencyEnum является перчислением валют.
type currencyEnum struct {
	Ruble  Currency
	Dollar Currency
	Euro   Currency
	Yuan   Currency
}

// String возвращает строковое представление валюты.
func (c Currency) String() string {
	return c.Symbol
}

// ParseCurrency распознаёт валюту из строки text.
func ParseCurrency(text string) (Currency, error) {
	text = strings.TrimSpace(strings.ToLower(text))

	switch text {
	// По коду
	case strings.ToLower(string(Currencies.Ruble.Code)):
		return Currencies.Ruble, nil
	case strings.ToLower(string(Currencies.Dollar.Code)):
		return Currencies.Dollar, nil
	case strings.ToLower(string(Currencies.Euro.Code)):
		return Currencies.Euro, nil
	case strings.ToLower(string(Currencies.Yuan.Code)):
		return Currencies.Yuan, nil

	// По названию
	case strings.ToLower(Currencies.Ruble.Name):
		return Currencies.Ruble, nil
	case strings.ToLower(Currencies.Dollar.Name):
		return Currencies.Dollar, nil
	case strings.ToLower(Currencies.Euro.Name):
		return Currencies.Euro, nil
	case strings.ToLower(Currencies.Yuan.Name):
		return Currencies.Yuan, nil

	// По символу
	case Currencies.Ruble.Symbol:
		return Currencies.Ruble, nil
	case Currencies.Dollar.Symbol:
		return Currencies.Dollar, nil
	case Currencies.Euro.Symbol:
		return Currencies.Euro, nil
	case Currencies.Yuan.Symbol:
		return Currencies.Yuan, nil

	// Неизвестная или неподдерживаемая приложением валюта
	default:
		return Currency{}, errors.Wrap(unknownCurrencyError, text)
	}
}

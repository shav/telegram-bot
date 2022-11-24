package finance_models

import "errors"

// Ошибка "Превышен лимит трат".
var SpendLimitExceededError = errors.New("spend limit exceeded")

// CurrencyConvertError описывает ошибку конвертации валюты.
type CurrencyConvertError struct {
	// Оригинальная ошибка.
	innerError error
}

// NewCurrencyConvertError создёт ошибку конвертации валюты.
func NewCurrencyConvertError(innerError error) *CurrencyConvertError {
	return &CurrencyConvertError{
		innerError: innerError,
	}
}

// Error возвращает сообщение об ошибке.
func (e *CurrencyConvertError) Error() string {
	var innerErrorText string = ""
	if e.innerError != nil {
		innerErrorText = ": " + e.innerError.Error()
	}
	return "currency convert error" + innerErrorText
}

// Unwrap возвращает оригиналльную ошибку.
func (e *CurrencyConvertError) Unwrap() error {
	return e.innerError
}

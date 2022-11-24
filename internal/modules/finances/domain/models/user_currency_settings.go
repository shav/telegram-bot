package finance_models

// UserCurrencySettings хранит пользовательские настройки валют.
type UserCurrencySettings struct {
	// ИД пользователя.
	UserId int64
	// Ввалюта.
	Currency Currency
}

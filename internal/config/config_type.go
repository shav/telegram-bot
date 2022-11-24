package config

// Type задаёт тип конфига приложения.
type Type string

const (
	// Тип конфига на основе файла.
	File = Type("File")
	// Тип конфига на основе переменных окружения.
	Env = Type("Env")
)

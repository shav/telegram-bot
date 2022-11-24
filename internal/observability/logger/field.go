package logger

import "go.uber.org/zap"

// Field является полем данных, привязанных к записи лога.
type Field = zap.Field

// Конструктор полей данных.
var Fields *logFieldsFactory

// logFieldsFactory является конструктором полей данных.
type logFieldsFactory struct {
}

// newFieldsFactory создаёт конструктор полей данных.
func newFieldsFactory() *logFieldsFactory {
	return &logFieldsFactory{}
}

// Error создаёт поле с ошибкой.
func (f *logFieldsFactory) Error(err error) Field {
	return zap.Error(err)
}

// String создаёт строковое поле.
func (f *logFieldsFactory) String(key string, value string) Field {
	return zap.String(key, value)
}

// Int32 создаёт целочисленное поле типа Int32.
func (f *logFieldsFactory) Int32(key string, value int32) Field {
	return zap.Int32(key, value)
}

// Int64 создаёт целочисленное поле типа Int64.
func (f *logFieldsFactory) Int64(key string, value int64) Field {
	return zap.Int64(key, value)
}

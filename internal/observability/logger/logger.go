//go:generate minimock -i logger -o ./mocks/ -s ".go"

package logger

import (
	"context"
	"github.com/pkg/errors"
)

// logger описывает контракт движка логирования.
type logger interface {
	// Записывает в лог сообщение с уровнем Debug.
	Debug(msg string, fields ...Field)
	// Записывает в лог сообщение с уровнем Info.
	Info(msg string, fields ...Field)
	// Записывает в лог сообщение с уровнем Warn.
	Warn(msg string, fields ...Field)
	// Записывает в лог сообщение с уровнем Error.
	Error(msg string, fields ...Field)
	// Записывает в лог сообщение с уровнем Fatal.
	Fatal(msg string, fields ...Field)
	// Сбрасывает буферизованные логи в поток вывода.
	Sync() error
}

// Хук на запись полей в лог.
type WriteFieldsHook func(ctx context.Context, fields ...Field) []Field

// Логгер.
var loggerCore logger

// Хук перед записью полей в лог.
var beforeWriteHook WriteFieldsHook

// Init выполняет инициализацию логгера.
func Init(logger logger, beforeWrite WriteFieldsHook) error {
	if logger == nil {
		return errors.New("Cannot init logger: logger core is not assigned")
	}

	loggerCore = logger
	beforeWriteHook = beforeWrite
	Fields = newFieldsFactory()
	return nil
}

// Stop завершает работу логгера.
func Stop() {
	if loggerCore != nil {
		err := loggerCore.Sync()
		if err != nil {
			loggerCore.Error("Log flush failed", Fields.Error(err))
		}
	}
	loggerCore = nil
	beforeWriteHook = nil
}

// Записывает в лог сообщение с уровнем Debug.
func Debug(ctx context.Context, msg string, fields ...Field) {
	if loggerCore != nil {
		loggerCore.Debug(msg, applyBeforeWriteHook(ctx, fields...)...)
	}
}

// Записывает в лог сообщение с уровнем Info.
func Info(ctx context.Context, msg string, fields ...Field) {
	if loggerCore != nil {
		loggerCore.Info(msg, applyBeforeWriteHook(ctx, fields...)...)
	}
}

// Записывает в лог сообщение с уровнем Warn.
func Warn(ctx context.Context, msg string, fields ...Field) {
	if loggerCore != nil {
		loggerCore.Warn(msg, applyBeforeWriteHook(ctx, fields...)...)
	}
}

// Записывает в лог сообщение с уровнем Error.
func Error(ctx context.Context, msg string, fields ...Field) {
	if loggerCore != nil {
		loggerCore.Error(msg, applyBeforeWriteHook(ctx, fields...)...)
	}
}

// Записывает в лог сообщение с уровнем Fatal.
func Fatal(ctx context.Context, msg string, fields ...Field) {
	if loggerCore != nil {
		loggerCore.Fatal(msg, applyBeforeWriteHook(ctx, fields...)...)
	}
}

// applyBeforeWriteHook применяет хук перед записью полей в лог.
func applyBeforeWriteHook(ctx context.Context, fields ...Field) []Field {
	// TODO: Сейчас beforeWriteHook создаёт новый список полей на основе старого.
	// Нужно приудмать, как добавлять в записи лога новые поля без аллокации памяти.
	if beforeWriteHook != nil {
		return beforeWriteHook(ctx, fields...)
	}
	return fields
}

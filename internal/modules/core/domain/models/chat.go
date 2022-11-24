package core_models

import (
	"context"
	"sync"
)

// Chat представляет из себя чат с пользователем.
type Chat struct {
	// ИД пользователя.
	userId int64
	// TODO: Доступ к ActiveCommand и ActiveHandler-у тоже можно сделать потокобезопасным,
	// но в текущих условиях использования это пока не нужно, т.к. обращение к ним происходит внутри блокировки чата.
	// Активная в данный момент команда.
	ActiveCommand Command
	// Обработчик активной в данный момент команды.
	ActiveHandler CommandHandler
	// Текущее состояние обработки команды.
	Status CommandHandleStatus
	// Контекст.
	ctx context.Context
	// Отмена операциий в контексте.
	cancel context.CancelFunc
	// Объект синхронизации для обработки сообщений от пользователя.
	lock *sync.Mutex
}

// NewChat создаёт новый чат с пользователем.
func NewChat(userId int64) *Chat {
	return &Chat{
		userId: userId,
		lock:   &sync.Mutex{},
	}
}

// StartHandleMessage начинает обработку сообщения от пользователя.
func (c *Chat) StartHandleMessage(ctx context.Context) {
	c.lock.Lock()
	c.ctx, c.cancel = context.WithCancel(ctx)
}

// EndHandleMessage завершает обработку сообщения от пользователя.
func (c *Chat) EndHandleMessage() {
	c.lock.Unlock()
	if c.cancel != nil {
		c.cancel()
	}
	c.ctx = nil
}

// GetUserId возвращает ИД пользователя, с которым ведётся чат.
func (c *Chat) GetUserId() int64 {
	return c.userId
}

// IsActive возвращает признак того, что чат с пользователем активен.
func (c *Chat) IsActive() bool {
	return c.ActiveCommand != ""
}

package core_storages_chats

import (
	"context"
	"sync"

	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// ChatsMemoryStorage представляет из себя хранилище чатов с пользователями в памяти.
type ChatsMemoryStorage struct {
	// Чаты с пользователями.
	chats map[int64]*core_models.Chat
	// Объект синхронизации доступа к коллекции.
	lock *sync.Mutex
}

// NewMemoryStorage создаёт новое хранилище чатов в памяти.
func NewMemoryStorage() *ChatsMemoryStorage {
	return &ChatsMemoryStorage{
		chats: make(map[int64]*core_models.Chat),
		lock:  &sync.Mutex{},
	}
}

// GetOrAdd получает чат с пользователем userId, если он существует.
// Если чата с пользователем не существует, то добавляет его, используя конструктор chatFactory.
func (c *ChatsMemoryStorage) GetOrAdd(ctx context.Context, ts tr.Transaction, userId int64, chatFactory func(u int64) *core_models.Chat,
	prepare func(c *core_models.Chat)) (chat *core_models.Chat, existed bool, err error) {

	span, _ := tracing.StartSpanFromContext(ctx, "ChatsMemoryStorage.GetOrAdd")
	defer span.Finish()

	c.lock.Lock()
	defer c.lock.Unlock()

	chat, existed = c.chats[userId]
	if !existed {
		if chatFactory != nil {
			chat = chatFactory(userId)
			c.chats[userId] = chat
		}
	}
	if prepare != nil {
		prepare(chat)
	}
	return
}

// Update обновляет состояние чата с пользователем userId.
func (c *ChatsMemoryStorage) Update(ctx context.Context, ts tr.Transaction, chat *core_models.Chat) error {
	span, _ := tracing.StartSpanFromContext(ctx, "ChatsMemoryStorage.Update")
	defer span.Finish()

	c.lock.Lock()
	defer c.lock.Unlock()

	userId := chat.GetUserId()
	if !chat.IsActive() {
		// Если чат более не активен, то удаляем его из хранилища,
		// чтобы не утекала память в случае, если пользователь больше не будет общаться с ботом.
		delete(c.chats, userId)
	} else {
		c.chats[userId] = chat
	}
	return nil
}

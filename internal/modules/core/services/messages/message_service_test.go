package core_services_messages_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/modules"
	"github.com/shav/telegram-bot/internal/modules/core"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/services/commands"
	"github.com/shav/telegram-bot/internal/modules/core/services/messages"
	"github.com/shav/telegram-bot/internal/modules/core/services/messages/mocks"
)

var userId = int64(123)

var ctx = context.Background()

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	senderMock := mocks.NewMessageSenderMock(t)
	cmdContainer := core_services_commands.NewContainer()

	chat := core_models.NewChat(userId)
	chat.StartHandleMessage(ctx)
	chatStorageMock := mocks.NewChatStorageMock(t)
	chatStorageMock.GetOrAddMock.Return(chat, false, nil)
	chatStorageMock.UpdateMock.Return(nil)

	service, err := core_services_messages.NewService(senderMock, cmdContainer, chatStorageMock)
	assert.NoError(t, err)

	initArgs := modules.ModuleInitArgs{
		Commands: cmdContainer,
	}
	err = core.NewModule().InitCommands(initArgs)
	assert.NoError(t, err)

	helpMessage := "Можете управлять мной с помощью команд:\n/start - Начать диалог\n/stop - Завершить диалог\n/help - Справка по всем доступным командам"
	senderMock.SendMessageMock.Inspect(func(ctx context.Context, uid int64, text string) {
		assert.Equal(t, userId, uid)
		assert.Contains(t, [...]string{helpMessage, "Привет!"}, text)
	}).Return(nil)

	err = service.HandleIncomingMessage(ctx, core_models.Message{
		Text:   "/start",
		UserID: userId,
	})

	assert.NoError(t, err)
}

func Test_OnUnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	senderMock := mocks.NewMessageSenderMock(t)
	cmdContainer := core_services_commands.NewContainer()

	chat := core_models.NewChat(userId)
	chat.StartHandleMessage(ctx)
	chatStorageMock := mocks.NewChatStorageMock(t)
	chatStorageMock.GetOrAddMock.Return(chat, false, nil)
	chatStorageMock.UpdateMock.Return(nil)

	service, err := core_services_messages.NewService(senderMock, cmdContainer, chatStorageMock)
	assert.NoError(t, err)

	senderMock.SendMessageMock.Inspect(func(ctx context.Context, uid int64, text string) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, "Не знаю эту команду.\nНаберите /help для получения списка доступных команд", text)
	}).Return(nil)

	err = service.HandleIncomingMessage(ctx, core_models.Message{
		Text:   "some text",
		UserID: userId,
	})

	assert.NoError(t, err)
}

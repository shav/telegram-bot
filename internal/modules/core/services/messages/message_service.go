//go:generate minimock -i messageSender -o ./mocks/ -s ".go"
//go:generate minimock -i chatStorage -o ./mocks/ -s ".go"

package core_services_messages

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/multi_error"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/observability/metrics"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Максимальное количество попыток отправить сообщение пользователю.
const maxSendMessageRetryCount = 10

// chatStorage представляет из себя хранилище чатов с пользователями.
type chatStorage interface {
	// GetOrAdd получает чат с пользователем userId, если он существует.
	// Если чата с пользователем не существует, то добавляет его, используя конструктор chatFactory.
	GetOrAdd(ctx context.Context, ts tr.Transaction, userId int64, chatFactory func(u int64) *core_models.Chat,
		prepare func(c *core_models.Chat)) (chat *core_models.Chat, existed bool, err error)
	// Update обновляет состояние чата с пользователем userId.
	Update(ctx context.Context, ts tr.Transaction, chat *core_models.Chat) error
}

// messageSender позволяет отправлять сообщения в какой-либо мессенджер.
type messageSender interface {
	// SendMessage отправляет в мессенджер пользователю userId сообщение с текстом text.
	SendMessage(ctx context.Context, userId int64, text string) error
	// SendMessageWithOptions отправляет в мессенджер пользователю userID
	// сообщение с текстом text и набором опций options для выбора.
	SendMessageWithOptions(ctx context.Context, userID int64, text string, options []core_models.Option) error
}

// commandsContainer используется для хранения команд чат-бота.
type commandsContainer interface {
	// GetCommandHandler возвращает обработчик команды command для пользователя userId
	GetCommandHandler(command core_models.Command, userId int64) (core_models.CommandHandler, error)
	// IsLikeCommand проверяет, похожа ли строка на команду чат-бота (т.е. имеет формат команды).
	IsLikeCommand(text string) bool
	// FormatCommandName форматирует отображаемое имя команды.
	FormatCommandName(command core_models.Command) string
}

// MessageService реализует бизнес-логику обработки сообщений из мессенджера.
type MessageService struct {
	// Клиент для отправки сообщений в мессенджер.
	client messageSender
	// Контейнер с командами чат-бота.
	commands commandsContainer
	// Активные чаты с пользователями по обработке команд.
	activeChats chatStorage
}

// NewService создаёт сервис для обработки сообщений из мессенджера,
// используя client для отправки ответных сообщений в мессенджер
// и commands для получения команд.
func NewService(client messageSender, commands commandsContainer, chats chatStorage) (*MessageService, error) {
	if client == nil {
		return nil, errors.New("New MessageService: message sender is not assigned")
	}
	if commands == nil {
		return nil, errors.New("New MessageService: commands container is not assigned")
	}
	if chats == nil {
		return nil, errors.New("New MessageService: chats storage is not assigned")
	}

	return &MessageService{
		client:      client,
		commands:    commands,
		activeChats: chats,
	}, nil
}

// HandleIncomingMessage обрабатывает входящее сообщение message из мессенджера.
func (s *MessageService) HandleIncomingMessage(ctx context.Context, message core_models.Message) error {
	actualCommand := "unknown"
	startTime := time.Now()
	defer func() {
		metrics.IncomingMessagesCount.Inc(actualCommand)
		metrics.IncomingMessageResponseTime.Set(actualCommand, time.Since(startTime))
	}()

	message.Text = strings.TrimSpace(message.Text)

	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.HandleIncomingMessage")
	defer span.Finish()
	span.SetTag("user", strconv.FormatInt(message.UserID, 10))
	defer func() { span.SetTag("command", actualCommand) }()
	if s.commands.IsLikeCommand(message.Text) {
		span.SetTag("rawCommand", message.Text)
	}

	var answers []core_models.Answer
	var aggregateError error

	// TODO: Прикрутить сквозную транзакцию на весь метод обработки сообщения, когда появится хранилище чатов в БД.
	chat, err := s.getChatWithUser(ctx, message.UserID)
	if chat != nil {
		defer chat.EndHandleMessage()
	}
	if err != nil {
		tracing.SetError(span)
		aggregateError = multi_error.Append(aggregateError, errors.Wrap(err, "get chat with user"))
		answers = []core_models.Answer{{Text: defaultAnswers.commandError}}
		return multi_error.Append(aggregateError, s.sendAnswerToUser(ctx, message.UserID, answers))
	}

	answers, chat.Status, err = s.getAnswerToMessage(ctx, chat, message)
	if err != nil {
		tracing.SetError(span)
		aggregateError = multi_error.Append(aggregateError, errors.Wrap(err, "get answer to user message"))
	}

	if chat.ActiveCommand != "" {
		actualCommand = string(chat.ActiveCommand)
	}

	if chat.Status != core_models.CommandStatuses.WaitForNextMessage {
		chat.ActiveCommand = ""
	}

	err = s.updateChatInStorage(ctx, chat)
	if err != nil {
		tracing.SetError(span)
		aggregateError = multi_error.Append(aggregateError, errors.Wrap(err, "update chat in storage"))
		// TODO:  Для хранилища в БД нужно ещё откатывать транзакцию.
		answers = []core_models.Answer{{Text: defaultAnswers.commandError}}
		return multi_error.Append(aggregateError, s.sendAnswerToUser(ctx, message.UserID, answers))
	}

	return multi_error.Append(aggregateError, s.sendAnswerToUser(ctx, message.UserID, answers))
}

// getAnswerToMessage возвращает ответ на сообщение пользователя message в рамках chat.
func (s *MessageService) getAnswerToMessage(ctx context.Context, chat *core_models.Chat,
	message core_models.Message) ([]core_models.Answer, core_models.CommandHandleStatus, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.GetAnswerToMessage")
	defer span.Finish()

	var answers []core_models.Answer
	var status core_models.CommandHandleStatus

	// Если поступила новая команда, то обработку предыдущей команды отменяем
	if chat.Status == core_models.CommandStatuses.WaitForNextMessage && chat.IsActive() {
		if s.commands.IsLikeCommand(message.Text) {
			chat.ActiveCommand = ""
			chat.ActiveHandler = nil
		}
	}

	var command core_models.Command
	if chat.IsActive() {
		command = chat.ActiveCommand
	} else {
		command = core_models.Command(message.Text)
	}

	if !chat.IsActive() {
		// На самом деле строка может быть просто обычным текстом, а не командой чат-бота (команды начинаются с префикса /),
		// но это неважно, т.к. в этом случае для обычного текста просто не будет найден обработчик команды.
		var err error
		chat.ActiveHandler, err = s.getCommandHandler(ctx, command, message.UserID)
		if err != nil {
			tracing.SetError(span)
			err = errors.Wrap(err, "get command handler")
			answers = []core_models.Answer{{Text: defaultAnswers.commandError}}
			if chat.IsActive() {
				status = chat.Status
			} else {
				status = core_models.CommandStatuses.Canceled
			}
			return answers, status, err
		}
	}

	if chat.ActiveHandler == nil {
		chat.ActiveCommand = ""
		answers = []core_models.Answer{{
			Text: fmt.Sprintf(defaultAnswers.unknownCommandTemplate, s.commands.FormatCommandName(core_commands_help.Metadata.Name)),
		}}
		return answers, core_models.CommandStatuses.Canceled, nil
	}

	answers, status, err := s.handleMessageByCommand(ctx, chat, message, command)
	if err != nil {
		tracing.SetError(span)
	}

	if !chat.IsActive() {
		chat.ActiveCommand = command
	}

	return answers, status, err
}

// getCommandHandler возвращает обработчик команды command для пользователя userId
func (s *MessageService) getCommandHandler(ctx context.Context, command core_models.Command, userId int64) (core_models.CommandHandler, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "MessageService.GetCommandHandler")
	defer span.Finish()
	span.SetTag("command", string(command))

	commandHandler, err := s.commands.GetCommandHandler(command, userId)
	if err != nil {
		tracing.SetError(span)
	}
	return commandHandler, err
}

// handleMessageByCommand обрабатывает сообщение message с помощью команды command.
func (s *MessageService) handleMessageByCommand(ctx context.Context, chat *core_models.Chat, message core_models.Message,
	command core_models.Command) ([]core_models.Answer, core_models.CommandHandleStatus, error) {
	span, commandCtx := tracing.StartSpanFromContext(ctx, "MessageService.CommandHandleMessage")
	defer span.Finish()
	span.SetTag("command", string(command))

	var err error
	var answers []core_models.Answer
	var status core_models.CommandHandleStatus

	if !chat.IsActive() {
		answers, status, err = chat.ActiveHandler.StartHandleCommand(commandCtx)
	} else {
		answers, status, err = chat.ActiveHandler.HandleNextMessage(commandCtx, message.Text)
	}

	if err != nil {
		tracing.SetError(span)
	}

	return answers, status, err
}

// sendAnswerToUser отправляет answers в ответ пользователю userId.
func (s *MessageService) sendAnswerToUser(ctx context.Context, userId int64, answers []core_models.Answer) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.SendAnswerToUser")
	defer span.Finish()
	span.SetTag("user", userId)
	span.SetTag("answersCount", len(answers))

	if answers == nil {
		return nil
	}

	var aggregateError error
	for _, answer := range answers {
		answerText := answer.Text
		answerText = strings.TrimSpace(answerText)
		if answerText == "" {
			continue
		}
		err := s.sendMessageToUser(ctx, userId, answerText, answer.Options)
		aggregateError = multi_error.Append(aggregateError, err)
	}

	if aggregateError != nil {
		tracing.SetError(span)
	}

	return aggregateError
}

// sendMessageToUser отправляет пользователю userId сообщение message с опциями на выбор options.
func (s *MessageService) sendMessageToUser(ctx context.Context, userId int64, message string, options []core_models.Option) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.SendMessageToUser")
	span.SetTag("user", userId)
	defer span.Finish()

	var aggregateError error
	for attempt := 0; attempt < maxSendMessageRetryCount; attempt++ {
		delay := time.Duration(100*attempt) * time.Millisecond
		<-time.After(delay)
		err := s.trySendMessageToUser(ctx, userId, message, options, attempt, delay)
		if err == nil {
			return nil
		}
		aggregateError = multi_error.Append(aggregateError, errors.Wrap(err, "send answer to user"))
	}

	if aggregateError != nil {
		tracing.SetError(span)
	}

	return aggregateError
}

// TrySendMessageToUser выполняет одну попытку отправить пользователю userId сообщение message с опциями на выбор options.
func (s *MessageService) trySendMessageToUser(ctx context.Context, userId int64, message string,
	options []core_models.Option, attempt int, delayBefore time.Duration) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.TrySendMessageToUser")
	defer span.Finish()
	span.SetTag("user", userId)
	span.SetTag("attempt", attempt)
	span.SetTag("delayBefore", delayBefore)

	if len(options) > 0 {
		err = s.client.SendMessageWithOptions(ctx, userId, message, options)
	} else {
		err = s.client.SendMessage(ctx, userId, message)
	}

	if err != nil {
		tracing.SetError(span)
	}

	return err
}

// getChatWithUser возвращает чат с пользователем userId.
func (s *MessageService) getChatWithUser(ctx context.Context, userId int64) (chat *core_models.Chat, err error) {
	// TODO: Нужно прикрутить сохранение и восстановление состояния обработчиков команд:
	// * на случай перезапуска сервиса во время диалога
	// * в случае масштабируемой системы, когда следующее собщение от пользователя в рамках одного и того же диалога
	//   может прийти на другой инстанс сервиса

	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.GetChatWithUser")
	defer span.Finish()

	// TODO: Прикрутить транзакции, когда появится хранилище чатов в БД.
	chat, _, err = s.activeChats.GetOrAdd(ctx, nil, userId,
		func(userId int64) *core_models.Chat { return core_models.NewChat(userId) },
		func(c *core_models.Chat) { c.StartHandleMessage(ctx) })

	if err != nil {
		tracing.SetError(span)
	}

	return chat, err
}

// updateChatInStorage обновляет чат с пользователем в хранилище.
func (s *MessageService) updateChatInStorage(ctx context.Context, chat *core_models.Chat) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "MessageService.UpdateChatInStorage")
	defer span.Finish()

	// TODO: Прикрутить транзакции, когда появится хранилище чатов в БД.
	err := s.activeChats.Update(ctx, nil, chat)
	if err != nil {
		tracing.SetError(span)
	}
	return err
}

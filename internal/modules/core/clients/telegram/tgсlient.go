package core_clients

import (
	"context"
	"math"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Максимальное количество кнопок в строке с кнопками выбора опций.
const maxButtonsCountInOptionsRow = 3

// Смещение, относительно которого считываем новые сообщения из telegram при подключении к нему.
const updatesOffset = 0 // - считываем все непрочитанные сообщения

// Таймаут ожидания новых сообщений из telegram.
const waitUpdatesTimeout = 60 // секунд

// messageHandler реализует бизнес-логику обработки сообщений из мессенджера.
type messageHandler interface {
	// HandleIncomingMessage обрабатывает входящее сообщение message из мессенджера.
	HandleIncomingMessage(ctx context.Context, message core_models.Message) error
}

// Client представляет из себя клиент telegram.
type Client struct {
	// АПИ для управления чат-ботом telegram.
	bot *tgbot.BotAPI
}

// NewTelegramClient создаёт telegram-клиента,
// используя token для подключения к АПИ telegram.
func NewTelegramClient(token string) (*Client, error) {
	bot, err := tgbot.NewBotAPI(token)
	if err != nil {
		return nil, errors.Wrap(err, "tgbot.NewBotAPI")
	}
	return &Client{bot}, nil
}

// SendMessage отправляет в telegram пользователю userId сообщение с текстом text.
func (c *Client) SendMessage(ctx context.Context, userId int64, text string) error {
	span, _ := tracing.StartSpanFromContext(ctx, "TelegramClient.SendMessage")
	defer span.Finish()
	span.SetTag("user", userId)

	_, err := c.bot.Send(tgbot.NewMessage(userId, text))
	if err != nil {
		tracing.SetError(span)
		return errors.Wrap(err, "bot.Send.Message")
	}
	return nil
}

// SendMessageWithOptions отправляет в telegram пользователю userId
// сообщение с текстом text и набором опций options для выбора.
func (c *Client) SendMessageWithOptions(ctx context.Context, userId int64, text string, options []core_models.Option) error {
	span, _ := tracing.StartSpanFromContext(ctx, "TelegramClient.SendMessageWithOptions")
	defer span.Finish()
	span.SetTag("user", userId)

	msg := tgbot.NewMessage(userId, text)
	msg.ReplyMarkup = getOptionsKeyboard(options)
	_, err := c.bot.Send(msg)
	if err != nil {
		tracing.SetError(span)
		return errors.Wrap(err, "bot.Send.ReplyMarkup")
	}
	return nil
}

// ListenUpdates слушает и обрабатывает входящие сообщения из telegram.
func (c *Client) ListenUpdates(handler messageHandler, ctx context.Context) {
	u := tgbot.NewUpdate(updatesOffset)
	u.Timeout = waitUpdatesTimeout

	updates := c.bot.GetUpdatesChan(u)

	logger.Info(ctx, "Started listening for messages from telegram")

	for {
		select {
		case <-ctx.Done():
			logger.Info(ctx, "Stopped listening for messages from telegram")
			return
		case update := <-updates:
			go func() {
				isHandled := c.handleMessage(ctx, update, handler)
				if !isHandled {
					c.handleSelectOption(ctx, update, handler)
				}
			}()
		}
	}
}

// getOptionsKeyboard возвращает набор кнопок для выбора из списка опций options в чате telegram.
func getOptionsKeyboard(options []core_models.Option) tgbot.InlineKeyboardMarkup {
	buttons := make([][]tgbot.InlineKeyboardButton, int(math.Ceil(float64(len(options))/maxButtonsCountInOptionsRow)))
	for i, row, column := 0, 0, 0; i < len(options); i++ {
		option := options[i]
		if buttons[row] == nil {
			buttons[row] = make([]tgbot.InlineKeyboardButton, 0, maxButtonsCountInOptionsRow)
		}
		buttons[row] = append(buttons[row], tgbot.NewInlineKeyboardButtonData(option.Text, option.Value))

		column++
		if column == maxButtonsCountInOptionsRow {
			column = 0
			row++
		}
	}
	return tgbot.NewInlineKeyboardMarkup(buttons...)
}

// handleMessage обрабатывает входящее сообщение пользователя, если оно есть.
// Возвращает true, если в обновлении из telegram было сообщение от пользователя, иначе false.
func (c *Client) handleMessage(ctx context.Context, update tgbot.Update, handler messageHandler) bool {
	if update.Message == nil {
		return false
	}

	user := update.Message.From
	message := update.Message.Text

	span, ctx := tracing.StartSpanFromContext(ctx, "TelegramClient.HandleMessage")
	defer span.Finish()
	span.SetTag("user", strconv.FormatInt(user.ID, 10))

	messageField := logger.Fields.String("message", message)
	userField := logger.Fields.String("user", user.UserName)
	logger.Info(ctx, "Processing {message} from {user}", messageField, userField)

	err := handler.HandleIncomingMessage(ctx, core_models.Message{
		Text:   message,
		UserID: user.ID,
	})
	if err != nil {
		logger.Error(ctx, "Processing {message} from {user} failed", messageField, userField, logger.Fields.Error(err))
	} else {
		logger.Info(ctx, "Processed {message} from {user}", messageField, userField)
	}
	return true
}

// handleSelectOption обрабатывает пользовательский выбор опции.
// Возвращает true, если в обновлении из telegram был выбор опции, иначе false.
func (c *Client) handleSelectOption(ctx context.Context, update tgbot.Update, handler messageHandler) bool {
	if update.CallbackQuery == nil {
		return false
	}

	user := update.CallbackQuery.From
	option := update.CallbackQuery.Data

	span, ctx := tracing.StartSpanFromContext(ctx, "TelegramClient.HandleSelectOption")
	defer span.Finish()
	span.SetTag("user", strconv.FormatInt(user.ID, 10))

	optionField := logger.Fields.String("option", option)
	userField := logger.Fields.String("user", user.UserName)
	logger.Info(ctx, "Processing {option} from {user}", optionField, userField)

	// TODO: Пока считаем, что опция - это просто удобный способ ввода сообщения, без набора текста на клавиатуре
	// Возможно, в будущем понадобится отдельный метод для обработки выбора опций.
	err := handler.HandleIncomingMessage(ctx, core_models.Message{
		Text:   option,
		UserID: user.ID,
	})
	if err != nil {
		logger.Error(ctx, "Processing {option} from {user} failed", optionField, userField, logger.Fields.Error(err))
	} else {
		logger.Info(ctx, "Processed {option} from {user}", optionField, userField)
	}
	return true
}

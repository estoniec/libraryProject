package builder

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log/slog"
)

type Builder struct {
	bot *telego.Bot
}

func NewBuilder(bot *telego.Bot) *Builder {
	return &Builder{
		bot: bot,
	}
}

func (b *Builder) NewMessage(msg telego.Update, text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error) {
	var chatID int64
	if msg.CallbackQuery != nil {
		chatID = msg.CallbackQuery.Message.Chat.ID
	} else {
		chatID = msg.Message.Chat.ID
	}
	message := tu.Message(
		tu.ID(chatID),
		text,
	)
	if keyboard != nil {
		message = message.WithReplyMarkup(keyboard)
	}
	sentMessage, err := b.bot.SendMessage(
		message,
	)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return sentMessage, nil
}

func (b *Builder) DeleteMessage(msg telego.Update) error {
	var msgID int
	var chatID int64
	var username string
	if msg.CallbackQuery != nil {
		msgID = msg.CallbackQuery.Message.MessageID
		chatID = msg.CallbackQuery.Message.Chat.ID
		username = msg.CallbackQuery.From.Username
	} else {
		msgID = msg.Message.MessageID
		chatID = msg.Message.Chat.ID
		username = msg.Message.From.Username
	}
	params := &telego.DeleteMessageParams{
		MessageID: msgID,
		ChatID: telego.ChatID{
			ID:       chatID,
			Username: username,
		},
	}
	err := b.bot.DeleteMessage(params)
	if err != nil {
		return err
	}
	return nil
}

func (b *Builder) NewCallbackMessage(msg *telego.CallbackQuery, text string) error {
	message := tu.CallbackQuery(msg.ID).WithText(text)

	err := b.bot.AnswerCallbackQuery(
		message,
	)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return err
}

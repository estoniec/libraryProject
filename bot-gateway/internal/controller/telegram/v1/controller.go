package v1

import (
	"context"
	"gateway/pkg/adapters/handling"
	"github.com/mymmrac/telego"
	"log/slog"
	"time"
)

type Builder interface {
	NewMessage(msg telego.Update, text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error)
	NewMessageWithKeyboard(msg telego.Update, text string, keyboard *telego.ReplyKeyboardMarkup) (*telego.Message, error)
	DeleteMessage(msg telego.Update) error
	NewCallbackMessage(msg *telego.CallbackQuery, text string) error
	NewMessageByID(id int64, text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error)
	NewMessageAndDeleteKeyboard(msg telego.Update, text string, deleteKeyboard bool, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error)
}

type Router interface {
	AddGroup(group handling.HandlersGroup)
	Listen(ctx context.Context, countListeners uint) chan<- telego.Update
}

type Question interface {
	NewQuestion(message telego.Update, size int64) (chan *telego.Message, func())
	Middleware(ctx context.Context, message *telego.Message) bool
}

type CallbackQuestion interface {
	NewQuestion(message telego.Update, size int64) (chan telego.Update, func())
	Middleware(ctx context.Context, message telego.Update) bool
}

type Handler struct {
	builder          Builder
	question         Question
	router           Router
	callbackQuestion CallbackQuestion
	lastMsg          map[int64]time.Time
}

func NewHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion) *Handler {
	return &Handler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
		lastMsg:          make(map[int64]time.Time),
	}
}

func (h *Handler) HandleUpdates(ctx context.Context, updates <-chan telego.Update) {
	listen := h.router.Listen(ctx, 1)
	for update := range updates {
		if update.Message != nil {
			if last, ok := h.lastMsg[update.Message.From.ID]; ok {
				if time.Now().Unix()-last.Unix() < 1 {
					_, err := h.builder.NewMessage(update, "Вы можете отправлять боту не более 1 сообщения в секунду. Повторите поптыку позже.", nil)
					if err != nil {
						slog.Error(err.Error())
						continue
					}
					continue
				} else {
					h.lastMsg[update.Message.From.ID] = time.Now()
				}
			} else {
				h.lastMsg[update.Message.From.ID] = time.Now()
			}
			ok := h.question.Middleware(ctx, update.Message)
			if ok {
				listen <- update
			}
		}
		if update.CallbackQuery != nil {
			if last, ok := h.lastMsg[update.CallbackQuery.From.ID]; ok {
				if time.Now().Unix()-last.Unix() < 1 {
					_, err := h.builder.NewMessage(update, "Вы можете отправлять боту не более 1 сообщения в секунду. Повторите поптыку позже.", nil)
					if err != nil {
						slog.Error(err.Error())
						continue
					}
					continue
				} else {
					h.lastMsg[update.CallbackQuery.From.ID] = time.Now()
				}
			} else {
				h.lastMsg[update.CallbackQuery.From.ID] = time.Now()
			}
			ok := h.callbackQuestion.Middleware(ctx, update)
			if ok {
				listen <- update
			}
		}
	}
}

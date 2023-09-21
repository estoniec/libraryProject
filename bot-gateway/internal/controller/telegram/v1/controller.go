package v1

import (
	"context"
	"gateway/pkg/adapters/handling"
	"github.com/mymmrac/telego"
)

type Builder interface {
	NewMessage(msg telego.Update, text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error)
	DeleteMessage(msg telego.Update) error
	NewCallbackMessage(msg *telego.CallbackQuery, text string) error
}

type Router interface {
	AddGroup(group handling.HandlersGroup)
	Listen(ctx context.Context, countListeners uint) chan<- telego.Update
}

type Question interface {
	NewQuestion(message telego.Update) (chan *telego.Message, func())
	Middleware(ctx context.Context, message *telego.Message) bool
}

type CallbackQuestion interface {
	NewQuestion(message telego.Update) (chan telego.Update, func())
	Middleware(ctx context.Context, message telego.Update) bool
}

type Handler struct {
	builder          Builder
	question         Question
	router           Router
	callbackQuestion CallbackQuestion
}

func NewHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion) *Handler {
	return &Handler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
	}
}

func (h *Handler) HandleUpdates(ctx context.Context, updates <-chan telego.Update) {
	listen := h.router.Listen(ctx, 1)
	for update := range updates {
		if update.Message != nil {
			ok := h.question.Middleware(ctx, update.Message)
			if ok {
				listen <- update
			}
		}
		if update.CallbackQuery != nil {
			ok := h.callbackQuestion.Middleware(ctx, update)
			if ok {
				listen <- update
			}
		}
	}
}

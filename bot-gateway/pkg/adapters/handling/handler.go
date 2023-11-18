package handling

import (
	"context"
	"github.com/mymmrac/telego"
)

type Callback func(ctx context.Context, msg telego.Update)

type Handler struct {
	Aliases    []string
	Command    string
	Callback   Callback
	IsQuestion bool
}

type HandlersGroup struct {
	Handlers []*Handler
}

func newHandler(callback Callback, aliases ...string) *Handler {
	return &Handler{
		Aliases:  aliases,
		Callback: callback,
	}
}

func NewHandlersGroup(handlers ...*Handler) HandlersGroup {
	return HandlersGroup{
		Handlers: handlers,
	}
}

func (h *HandlersGroup) NewHandler(callback Callback, aliases ...string) *Handler {
	handler := newHandler(callback, aliases...)

	h.Add(handler)

	return handler
}

func (h *HandlersGroup) Add(handlers ...*Handler) {
	h.Handlers = append(h.Handlers, handlers...)
}

func (h *Handler) WithCommand(command string) {
	h.Command = command
}

func (h *Handler) Question() {
	h.IsQuestion = true
}

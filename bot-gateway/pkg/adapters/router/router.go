package router

import (
	"context"
	"gateway/pkg/adapters/handling"
	"github.com/buger/jsonparser"
	"github.com/mymmrac/telego"
	"github.com/sourcegraph/conc/pool"
	"log/slog"
	"strings"
)

type Middleware func(ctx context.Context) error

type Router struct {
	client        *telego.Bot
	handlersNames []string
	handlers      map[string]*handling.Handler
	wg            *pool.Pool
}

const (
	maxDistanceMessageLength = 50
)

func NewRouter(bot *telego.Bot) *Router {
	return &Router{
		client:   bot,
		wg:       pool.New(),
		handlers: make(map[string]*handling.Handler, 20),
	}
}

func (r *Router) AddGroup(group handling.HandlersGroup) {
	for _, handler := range group.Handlers {
		if handler.Command != "" {
			r.handlers[handler.Command] = handler
			r.handlersNames = append(r.handlersNames, handler.Command)
		}

		for _, alias := range handler.Aliases {
			lowerAlias := strings.ToLower(alias)

			r.handlersNames = append(r.handlersNames, lowerAlias)
			r.handlers[lowerAlias] = handler
		}
	}
}

func (r *Router) Listen(ctx context.Context, countListeners uint) chan<- telego.Update {
	messages := make(chan telego.Update)

	slog.Info("listeners starting...")

	for i := 0; uint(i) < countListeners; i++ {
		r.wg.Go(func() {
			defer func() {
				if rec := recover(); rec != nil {
					slog.Error("one of router listeners ended with panic %v", r)

					slog.Info("started new listener")

					r.wg.Go(func() {
						r.listener(ctx, messages)
					})
				}
			}()

			r.listener(ctx, messages)
		})
	}

	slog.Info("listeners started")

	return messages
}

func (r *Router) listener(ctx context.Context, messages <-chan telego.Update) {
	for msg := range messages {
		// TODO: Add send to default handler
		err := r.handleCommand(ctx, msg)
		if err != nil {
			continue
		}
	}
}

func (r *Router) handleCommand(ctx context.Context, msg telego.Update) error {
	if msg.CallbackQuery != nil {
		command, err := r.callback(msg)
		if err != nil {
			return err
		}

		if command == "" {
			return nil
		}

		if len(command) > maxDistanceMessageLength {
			return nil
		}

		handler, ok := r.handlers[command]
		if ok {
			handler.Callback(ctx, msg)
		}
		return nil
	}
	command, err := r.command(msg)
	if err != nil {
		return err
	}

	if command == "" {
		return nil
	}

	if len(command) > maxDistanceMessageLength {
		return nil
	}

	handler, ok := r.handlers[command]
	if ok {
		handler.Callback(ctx, msg)
	}
	return nil
}

func (r *Router) command(msg telego.Update) (string, error) {
	command, err := jsonparser.GetString([]byte(msg.Message.Text), "command")
	if err == nil {
		return strings.ToLower(command), err
	}

	return strings.ToLower(msg.Message.Text), nil
}

func (r *Router) callback(msg telego.Update) (string, error) {
	command, err := jsonparser.GetString([]byte(msg.CallbackQuery.Data), "command")
	if err == nil {
		return strings.ToLower(command), err
	}

	return strings.ToLower(msg.CallbackQuery.Data), nil
}

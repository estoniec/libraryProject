package app

import (
	"context"
	"gateway/internal/adapters/keyboard"
	"gateway/internal/config"
	v1 "gateway/internal/controller/telegram/v1"
	books_service "gateway/internal/domain/books/service"
	service2 "gateway/internal/domain/users/service"
	"gateway/pkg/adapters/builder"
	"gateway/pkg/adapters/question"
	"gateway/pkg/adapters/router"
	bookPb "github.com/estoniec/libraryProject/contracts/gen/go/books"
	regPb "github.com/estoniec/libraryProject/contracts/gen/go/registration"
	"github.com/mymmrac/telego"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
)

type App struct {
	config       *config.Config
	bot          *telego.Bot
	handler      *v1.Handler
	regHandler   *v1.RegHandler
	booksHandler *v1.BooksHandler
}

func NewApp(ctx context.Context, c *config.Config) *App {
	regCc, err := grpc.Dial(c.RegSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Could not connect:", err)
	}

	bookCc, err := grpc.Dial(c.BooksSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Could not connect:", err)
	}

	regClient := regPb.NewRegServiceClient(regCc)
	bookClient := bookPb.NewBooksServiceClient(bookCc)

	regService := service2.NewService(regClient)
	bookService := books_service.NewService(bookClient)

	bot, err := telego.NewBot(c.BotToken)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	keyboardManager := keyboard.NewKeyboardManager()
	builder := builder.NewBuilder(bot)
	questionManager := question.NewManager(ctx)
	callbackQuestionManager := question.NewCallbackManager(ctx)
	router := router.NewRouter(bot)
	handler := v1.NewHandler(builder, router, questionManager, callbackQuestionManager)
	regHandler := v1.NewRegHandler(builder, router, questionManager, callbackQuestionManager, regService, keyboardManager)
	booksHandler := v1.NewBooksHandler(builder, router, questionManager, callbackQuestionManager, bookService, keyboardManager)
	return &App{
		config:       c,
		bot:          bot,
		handler:      handler,
		regHandler:   regHandler,
		booksHandler: booksHandler,
	}
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.start(ctx)
	})

	return grp.Wait()
}

func (a *App) start(ctx context.Context) error {
	updates, err := a.bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return err
	}

	defer a.bot.StopLongPolling()

	a.regHandler.Register()
	a.booksHandler.Register()
	a.handler.HandleUpdates(ctx, updates)

	slog.Info("bot has started")
	return nil
}

package app

import (
	"context"
	"gateway/internal/adapters/keyboard"
	"gateway/internal/config"
	v1 "gateway/internal/controller/telegram/v1"
	books_service "gateway/internal/domain/books/service"
	books_storage "gateway/internal/domain/books/storage"
	"gateway/internal/domain/books/usecase"
	rentService "gateway/internal/domain/rent/usecase"
	"gateway/internal/domain/users/usecase"
	"gateway/pkg/adapters/builder"
	"gateway/pkg/adapters/question"
	"gateway/pkg/adapters/router"
	bookPb "github.com/estoniec/libraryProject/contracts/gen/go/books"
	rentPb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	regPb "github.com/estoniec/libraryProject/contracts/gen/go/users"
	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	"github.com/mymmrac/telego"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
	"time"
)

type App struct {
	config       *config.Config
	bot          *telego.Bot
	handler      *v1.Handler
	regHandler   *v1.RegHandler
	booksHandler *v1.BooksHandler
	adminHandler *v1.AdminHandler
	rentHandler  *v1.RentHandler
	scheduler    *gocron.Scheduler
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

	rentCc, err := grpc.Dial(c.RentSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Could not connect:", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	regClient := regPb.NewRegServiceClient(regCc)
	bookClient := bookPb.NewBooksServiceClient(bookCc)
	rentClient := rentPb.NewBooksUsersServiceClient(rentCc)

	bookRepository := books_storage.NewBooksStorage(client)

	regUsecase := usersService.NewUsecase(regClient)
	bookService := books_service.NewService(bookClient, bookRepository)
	rentUsecase := rentService.NewUsecase(rentClient)

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
	regHandler := v1.NewRegHandler(builder, router, questionManager, callbackQuestionManager, regUsecase, keyboardManager)
	booksUsecase := booksService.NewUsecase(bookClient, bookService)
	booksHandler := v1.NewBooksHandler(builder, router, questionManager, callbackQuestionManager, booksUsecase, keyboardManager)
	rentHandler := v1.NewRentHandler(builder, router, questionManager, callbackQuestionManager, rentUsecase, booksUsecase, keyboardManager)
	adminHandler := v1.NewAdminHandler(builder, router, questionManager, callbackQuestionManager, regUsecase, booksUsecase, rentUsecase, keyboardManager)
	s := gocron.NewScheduler(time.UTC)
	return &App{
		config:       c,
		bot:          bot,
		handler:      handler,
		regHandler:   regHandler,
		booksHandler: booksHandler,
		adminHandler: adminHandler,
		rentHandler:  rentHandler,
		scheduler:    s,
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
	a.adminHandler.Register()
	a.rentHandler.Register()
	a.scheduler.Every(1).Day().Do(a.rentHandler.GetDebt, ctx)
	a.scheduler.StartAsync()
	a.handler.HandleUpdates(ctx, updates)
	return nil
}

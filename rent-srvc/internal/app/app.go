package app

import (
	"context"
	"fmt"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"rent/internal/config"
	v1 "rent/internal/controller/grpc/v1"
	psql "rent/pkg/postgresql"
	"time"
)

type App struct {
	config *config.Config
	server *v1.Server
}

func NewApp(ctx context.Context, cfg *config.Config) App {
	pgDsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	// TODO to config
	pgClient, err := psql.NewClient(ctx, 5, 3*time.Second, pgDsn, false)
	if err != nil {
		return App{}
	}
	storage := dao.NewRegistrationStorage(pgClient)
	svc := service.NewService(storage)
	server := v1.NewServer(svc, pb.UnimplementedRegServiceServer{})
	return App{
		config: cfg,
		server: server,
	}
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.startGRPC()
	})

	return grp.Wait()
}

func (a *App) startGRPC() error {
	lis, err := net.Listen("tcp", a.config.Port)

	if err != nil {
		slog.Error("Failed to listing:", err)
	}

	slog.Info("Queue Svc on", a.config.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterRegServiceServer(grpcServer, a.server)

	return grpcServer.Serve(lis)
}

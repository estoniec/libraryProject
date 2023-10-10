package main

import (
	"context"
	"registration-svc/internal/app"
	"registration-svc/internal/config"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()
	app := app.NewApp(ctx, cfg)
	app.Run(ctx)
}

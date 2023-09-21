package main

import (
	"context"
	"project11/registration-svc/internal/app"
	"project11/registration-svc/internal/config"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()
	app := app.NewApp(ctx, cfg)
	app.Run(ctx)
}

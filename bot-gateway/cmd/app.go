package main

import (
	"context"
	app2 "gateway/internal/app"
	"gateway/internal/config"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()
	app := app2.NewApp(ctx, cfg)
	app.Run(ctx)
}

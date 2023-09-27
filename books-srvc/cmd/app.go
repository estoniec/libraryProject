package main

import (
	"books-srvc/internal/app"
	"books-srvc/internal/config"
	"context"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()
	app := app.NewApp(ctx, cfg)
	app.Run(ctx)
}

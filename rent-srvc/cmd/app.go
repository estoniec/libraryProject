package cmd

import (
	"context"
	"rent/internal/app"
	"rent/internal/config"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()
	app := app.NewApp(ctx, cfg)
	app.Run(ctx)
}

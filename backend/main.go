package main

import (
	"context"
	"os/signal"
	"pulse/internal/app"
	"syscall"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Parsing the config
	app := app.App{}
	app.Configure(ctx)
	app.Serve(ctx)
}

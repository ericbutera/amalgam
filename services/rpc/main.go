package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/services/rpc/internal/config"
	"github.com/ericbutera/amalgam/services/rpc/internal/server"
	"github.com/ericbutera/amalgam/services/rpc/internal/tasks"
	"github.com/samber/lo"
)

func main() {
	if err := run(); err != nil {
		slog.Error("run error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	config := lo.Must(env.New[config.Config]())

	db := lo.Must(db.NewFromEnv())

	tasks, err := tasks.NewTemporalFromEnv()
	if err != nil {
		return err
	}
	defer tasks.Close()

	opts := []server.Option{
		server.WithConfig(config),
		server.WithDbFromEnv(),
		server.WithService(service.NewGorm(db)),
		server.WithTasks(tasks),
	}

	if config.OtelEnable {
		opts = append(opts, server.WithOtel(ctx, config.IgnoredSpanNames))
	}

	srv, err := server.New(opts...)
	if err != nil {
		return err
	}
	return srv.Serve(ctx)
}

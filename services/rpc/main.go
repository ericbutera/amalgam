package main

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/services/rpc/internal/config"
	"github.com/ericbutera/amalgam/services/rpc/internal/server"
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

	opts := []server.Option{
		server.WithConfig(config),
		server.WithDbFromEnv(),
		server.WithService(service.NewGorm(db)),
	}

	if config.OtelEnable {
		opts = append(opts, server.WithOtel(ctx, strings.Split(config.IgnoredSpanNames, ",")))
	}

	srv, err := server.New(opts...)
	if err != nil {
		return err
	}
	return srv.Serve(ctx)
}

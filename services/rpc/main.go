package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/services/rpc/internal/config"
	"github.com/ericbutera/amalgam/services/rpc/internal/server"
	"github.com/samber/lo"
)

func main() {
	ctx := context.Background()
	config := lo.Must(env.New[config.Config]())

	opts := []server.Option{
		server.WithConfig(config),
		server.WithDbFromEnv(),
	}

	if config.OtelEnable {
		opts = append(opts, server.WithOtel(ctx, config.IgnoredSpanNames))
	}

	srv, err := server.New(opts...)
	if err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
	if err := srv.Serve(ctx); err != nil {
		slog.Error("serve error", "error", err)
		os.Exit(1)
	}
}

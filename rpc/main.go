package main

import (
	"context"
	"log/slog"
	"os"

	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/ericbutera/amalgam/rpc/internal/config"
	"github.com/ericbutera/amalgam/rpc/internal/server"
)

func main() {
	ctx := context.Background()

	cfg, err := cfg.NewFromEnv[config.Config]()
	if err != nil {
		slog.Error("config error: ", "error", err)
		os.Exit(1)
	}

	srv, err := server.New(
		server.WithOtel(ctx),
		server.WithDbFromEnv(),
		server.WithPort(cfg.Port),
		server.WithMetricAddress(cfg.MetricAddress),
	)
	if err != nil {
		slog.Error("server error: ", "error", err)
		os.Exit(1)
	}
	if err := srv.Serve(ctx); err != nil {
		slog.Error("server error: ", "error", err)
		os.Exit(1)
	}
}

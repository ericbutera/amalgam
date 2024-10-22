package main

import (
	"log/slog"
	"os"

	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/ericbutera/amalgam/rpc/internal/config"
	"github.com/ericbutera/amalgam/rpc/internal/server"
)

func main() {
	cfg, err := cfg.NewFromEnv[config.Config]()
	if err != nil {
		slog.Error("config error: ", "error", err)
		os.Exit(1)
	}

	srv, err := server.New(
		server.WithPort(cfg.Port),
		server.WithMetricAddress(cfg.MetricAddress),
	)
	if err != nil {
		slog.Error("rpc error: ", "error", err)
		os.Exit(1)
	}
	srv.Serve()
}

package main

import (
	"log/slog"
	"os"

	"github.com/ericbutera/amalgam/internal/db"
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
	d, err := db.NewFromEnv()
	if err != nil {
		slog.Error("database error: ", "error", err)
		os.Exit(1)
	}
	srv, err := server.New(
		server.WithDb(d),
		server.WithPort(cfg.Port),
		server.WithMetricAddress(cfg.MetricAddress),
	)
	if err != nil {
		slog.Error("server error: ", "error", err)
		os.Exit(1)
	}
	if err := srv.Serve(); err != nil {
		slog.Error("server error: ", "error", err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/rpc/internal/config"
	"github.com/ericbutera/amalgam/rpc/internal/server"
	"github.com/samber/lo"
)

func main() {
	ctx := context.Background()
	config := lo.Must(env.New[config.Config]())

	srv, err := server.New(
		server.WithConfig(config),
		server.WithOtel(ctx),
		server.WithDbFromEnv(),
	)
	if err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
	if err := srv.Serve(ctx); err != nil {
		slog.Error("serve error", "error", err)
		os.Exit(1)
	}
}

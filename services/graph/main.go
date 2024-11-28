package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ericbutera/amalgam/internal/logger"
	"github.com/ericbutera/amalgam/internal/tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/pkg/otel"
	"github.com/ericbutera/amalgam/services/graph/internal/config"
	"github.com/ericbutera/amalgam/services/graph/internal/server"
	rpc "github.com/ericbutera/amalgam/services/rpc/pkg/client"
	"github.com/samber/lo"
)

func main() {
	slog := logger.New()

	if err := run(slog); err != nil {
		slog.Error("graph server error", "error", err)
		os.Exit(1)
	}
}

func run(slog *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := lo.Must(env.New[config.Config]())

	if config.OtelEnable {
		shutdown := lo.Must(otel.Setup(ctx, config.IgnoredSpanNames))
		defer func() { lo.Must0(shutdown(ctx)) }()
	}

	client, closer := lo.Must2(rpc.New(config.RpcHost, config.RpcInsecure))
	defer func() { lo.Must0(closer()) }()

	tasks := lo.Must(tasks.NewTemporalFromEnv())
	defer tasks.Close()

	srv, err := server.New(config, client, tasks)
	if err != nil {
		return err
	}

	slog.Info("running graph", "port", config.Port)

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("server shutdown failed", "error", err)
		}
	}()

	return srv.ListenAndServe()
}

package cmd

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/ericbutera/amalgam/api/internal/config"
	"github.com/ericbutera/amalgam/api/internal/server"
	"github.com/ericbutera/amalgam/internal/db"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/ericbutera/amalgam/pkg/otel"
)

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Long:  "Run api server",
		Run:   runServer,
	}
	return cmd
}

func quit(ctx context.Context, err error) {
	slog.ErrorContext(ctx, err.Error())
	os.Exit(1)
}

func runServer(cmd *cobra.Command, args []string) {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(h))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	shutdown, err := otel.Setup(ctx)
	if err != nil {
		quit(ctx, err)
	}
	defer func() {
		err = errors.Join(err, shutdown(context.Background()))
	}()

	cfg, err := cfg.NewFromEnv[config.Config]()
	if err != nil {
		quit(ctx, err)
	}
	d, err := db.NewFromEnv()
	if err != nil {
		quit(ctx, err)
	}

	srv, err := server.New(
		server.WithConfig(cfg),
		server.WithDb(d),
	)
	if err != nil {
		quit(ctx, err)
	}

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.Run()
	}()

	select {
	case err = <-srvErr:
		quit(ctx, err)
	case <-ctx.Done():
		slog.Info("shutting down")
		stop()
	}
	quit(ctx, err)
}

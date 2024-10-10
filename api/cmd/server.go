package cmd

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/ericbutera/amalgam/api/internal/otel"
	"github.com/ericbutera/amalgam/api/internal/server"
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
		return
	}
	defer func() {
		err = errors.Join(err, shutdown(context.Background()))
	}()

	srv, err := server.New(server.WithDatabase())
	if err != nil {
		quit(ctx, err)
		return
	}

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.Run()
	}()

	select {
	case err = <-srvErr:
		slog.ErrorContext(ctx, "server error")
		return
	case <-ctx.Done():
		slog.Info("shutting down")
		stop()
	}
	quit(ctx, err)
}

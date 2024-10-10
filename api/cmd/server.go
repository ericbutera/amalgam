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
		Run:   RunServer,
	}
	return cmd
}

func RunServer(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	otelShutdown, err := otel.Setup(ctx)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- server.New().Run()
	}()

	select {
	case err = <-srvErr:
		slog.ErrorContext(ctx, "server error")
		return
	case <-ctx.Done():
		slog.Info("shutting down")
		stop()
	}
}

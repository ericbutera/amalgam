package main

import (
	"log/slog"
	"os"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/generate"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"

	rpc "github.com/ericbutera/amalgam/rpc/pkg/client"

	"go.temporal.io/sdk/worker"
)

func main() {
	config := lo.Must(config.NewFromEnv[generate.Config]())

	config.RpcHost = "localhost:50055"
	config.RpcInsecure = true

	// TODO: use graph
	rpc := lo.Must(rpc.NewClient(config.RpcHost, config.RpcInsecure))
	defer rpc.Conn.Close()
	a := generate.NewActivities(rpc.Client)

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	w := worker.New(client, config.TaskQueue, worker.Options{})
	w.RegisterWorkflow(generate.GenerateFeedsWorkflow)
	w.RegisterActivity(a)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}

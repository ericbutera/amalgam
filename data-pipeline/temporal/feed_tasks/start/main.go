package main

// TODO: https://docs.temporal.io/develop/go/failure-detection#workflow-timeouts

import (
	"context"
	"log/slog"
	"os"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

func main() {
	ctx := context.Background()

	config := lo.Must(env.New[feed_tasks.Config]())

	client := lo.Must(sdk.Dial(sdk.Options{
		HostPort: config.TemporalHost,
	}))
	defer client.Close()

	opts := sdk.StartWorkflowOptions{
		TaskQueue: config.TaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}

	we, err := client.ExecuteWorkflow(ctx, opts, feed_tasks.GenerateFeedsWorkflow, config.FakeHost, config.GenerateCount)
	if err != nil {
		slog.Error("unable to execute workflow", "error", err)
		os.Exit(1) //nolint: gocritic
	}

	slog.Info("started workflow", "workflow_id", we.GetID())
}

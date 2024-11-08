package main

// TODO: https://docs.temporal.io/develop/go/failure-detection#workflow-timeouts

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

var retryPolicy = &temporal.RetryPolicy{
	MaximumAttempts: 1,
}

func main() {
	if err := runWorker(); err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}

func runWorker() error {
	ctx := context.Background()

	config := lo.Must(cfg.NewFromEnv[config.Config]())

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	if config.UseSchedule {
		return runSchedule(ctx, config, client)
	} else {
		return runExecute(ctx, config, client)
	}
}

func runSchedule(ctx context.Context, config *config.Config, client sdk.Client) error {
	// docs: https://docs.temporal.io/develop/go/schedules
	handle := client.ScheduleClient().GetHandle(ctx, config.ScheduleID)
	if err := handle.Delete(ctx); err != nil {
		return fmt.Errorf("failed to delete schedule: %w", err)
	}
	schedule, err := client.ScheduleClient().Create(ctx, sdk.ScheduleOptions{
		ID: config.ScheduleID,
		Spec: sdk.ScheduleSpec{
			Intervals: []sdk.ScheduleIntervalSpec{
				{Every: 1 * time.Minute},
			},
		},
		Action: &sdk.ScheduleWorkflowAction{
			ID:          config.WorkflowID,
			Workflow:    app.FetchFeedsWorkflow,
			TaskQueue:   config.TaskQueue,
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to schedule workflow: %w", err)
	}
	slog.Info("started workflow", "schedule", schedule.GetID())
	return nil
}

func runExecute(ctx context.Context, config *config.Config, client sdk.Client) error {
	opts := sdk.StartWorkflowOptions{
		TaskQueue:   config.TaskQueue,
		RetryPolicy: retryPolicy,
	}
	we, err := client.ExecuteWorkflow(ctx, opts, app.FetchFeedsWorkflow)
	if err != nil {
		return fmt.Errorf("failed to execute workflow: %w", err)
	}
	slog.Info("started workflow", "workflow_id", we.GetID())
	return nil
}

package main

// TODO: https://docs.temporal.io/develop/go/failure-detection#workflow-timeouts

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_fetch"
	clientHelper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config/env"
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

type Config struct {
	UseSchedule bool   `env:"USE_SCHEDULE" envDefault:"false"`
	ScheduleID  string `env:"SCHEDULE_ID" envDefault:"feed-fetch-schedule-id"`
	WorkflowID  string `env:"WORKFLOW_ID" envDefault:"fetch-feeds-workflow-id"`
	TaskQueue   string `env:"TASK_QUEUE"`
}

func runWorker() error {
	ctx := context.Background()
	config := lo.Must(env.New[Config]())
	client := lo.Must(clientHelper.NewTemporalClientFromEnv())
	if config.UseSchedule {
		return runSchedule(ctx, config, client)
	}
	return runExecute(ctx, config, client)
}

func runSchedule(ctx context.Context, config *Config, client sdk.Client) error {
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

func runExecute(ctx context.Context, config *Config, client sdk.Client) error {
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

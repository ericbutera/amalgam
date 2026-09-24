package main

// TODO: https://docs.temporal.io/develop/go/failure-detection#workflow-timeouts

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	clientHelper "github.com/ericbutera/amalgam/internal/temporal/client"
	app "github.com/ericbutera/amalgam/internal/temporal/feed_fetch"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

var retryPolicy = &temporal.RetryPolicy{
	MaximumAttempts: 1,
}

func main() {
	err := runWorker()
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}

type Config struct {
	UseSchedule bool   `env:"USE_SCHEDULE" envDefault:"false"`
	ScheduleID  string `env:"SCHEDULE_ID"  envDefault:"feed-fetch-schedule-id"`
	WorkflowID  string `env:"WORKFLOW_ID"  envDefault:"fetch-feeds-workflow-id"`
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
	options := sdk.ScheduleOptions{
		ID: config.ScheduleID,
		Spec: sdk.ScheduleSpec{
			Intervals: []sdk.ScheduleIntervalSpec{
				{Every: 1 * time.Minute},
			},
		},
		Overlap:        enums.SCHEDULE_OVERLAP_POLICY_SKIP,
		PauseOnFailure: false,
		Action: &sdk.ScheduleWorkflowAction{
			ID:          config.WorkflowID,
			Workflow:    app.FetchFeedsWorkflow,
			TaskQueue:   config.TaskQueue,
			RetryPolicy: retryPolicy,
		},
	}

	if _, err := handle.Describe(ctx); err != nil {
		var notFound *serviceerror.NotFound
		if !errors.As(err, &notFound) {
			return fmt.Errorf("failed to inspect schedule: %w", err)
		}

		schedule, err := client.ScheduleClient().Create(ctx, options)
		if err != nil {
			return fmt.Errorf("failed to create schedule: %w", err)
		}

		slog.Info("created workflow schedule", "schedule", schedule.GetID())
		return nil
	}

	if err := handle.Update(ctx, sdk.ScheduleUpdateOptions{
		DoUpdate: func(input sdk.ScheduleUpdateInput) (*sdk.ScheduleUpdate, error) {
			return &sdk.ScheduleUpdate{
				Schedule: &sdk.Schedule{
					Action: options.Action,
					Spec:   &options.Spec,
					Policy: &sdk.SchedulePolicies{
						Overlap:        options.Overlap,
						CatchupWindow:  options.CatchupWindow,
						PauseOnFailure: options.PauseOnFailure,
					},
					State: input.Description.Schedule.State,
				},
			}, nil
		},
	}); err != nil {
		return fmt.Errorf("failed to update schedule: %w", err)
	}

	slog.Info("updated workflow schedule", "schedule", config.ScheduleID)

	return nil
}

func runExecute(ctx context.Context, config *Config, client sdk.Client) error {
	opts := sdk.StartWorkflowOptions{
		ID:          config.WorkflowID,
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

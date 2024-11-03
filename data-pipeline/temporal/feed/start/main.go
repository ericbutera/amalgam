package main

// TODO: https://docs.temporal.io/develop/go/failure-detection#workflow-timeouts

import (
	"context"
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

func main() {
	ctx := context.Background()

	config := lo.Must(cfg.NewFromEnv[config.Config]())

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	retryPolicy := &temporal.RetryPolicy{
		MaximumAttempts: 1,
	}

	if config.UseSchedule {
		// docs: https://docs.temporal.io/develop/go/schedules
		handle := client.ScheduleClient().GetHandle(ctx, config.ScheduleID)
		if err := handle.Delete(ctx); err != nil {
			slog.Error("unable to delete schedule", "error", err)
			os.Exit(1)
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
			slog.Error("unable to schedule workflow", "error", err)
			os.Exit(1)
		}
		slog.Info("started workflow", "schedule", schedule.GetID())
	} else {
		opts := sdk.StartWorkflowOptions{
			TaskQueue:   config.TaskQueue,
			RetryPolicy: retryPolicy,
		}
		we, err := client.ExecuteWorkflow(ctx, opts, app.FetchFeedsWorkflow)
		if err != nil {
			slog.Error("unable to execute workflow", "error", err)
			os.Exit(1)
		}
		slog.Info("started workflow", "workflow_id", we.GetID())
	}
}

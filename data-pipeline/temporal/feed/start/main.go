package main

// TODO: https://docs.temporal.io/develop/go/failure-detection#workflow-timeouts

import (
	"context"
	"log/slog"
	"os"
	"time"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"

	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"
	"go.temporal.io/sdk/client"
)

func main() {
	ctx := context.Background()

	config := lo.Must(cfg.NewFromEnv[config.Config]())

	c := lo.Must(app.NewTemporalClient(config.TemporalHost))
	defer c.Close()

	if config.UseSchedule {
		// docs: https://docs.temporal.io/develop/go/schedules
		handle := c.ScheduleClient().GetHandle(ctx, config.ScheduleID)
		if err := handle.Delete(ctx); err != nil {
			slog.Error("unable to delete schedule", "error", err)
		}
		schedule, err := c.ScheduleClient().Create(ctx, client.ScheduleOptions{
			ID: config.ScheduleID,
			Spec: client.ScheduleSpec{
				Intervals: []client.ScheduleIntervalSpec{
					{Every: 2 * time.Minute},
				},
			},
			Action: &client.ScheduleWorkflowAction{
				ID:        config.WorkflowID,
				Workflow:  app.FeedWorkflow,
				TaskQueue: config.TaskQueue,
			},
		})
		if err != nil {
			slog.Error("unable to schedule workflow", "error", err)
		}
		slog.Info("started workflow", "schedule", schedule.GetID())
	} else {
		opts := client.StartWorkflowOptions{
			TaskQueue: config.TaskQueue,
		}
		we, err := c.ExecuteWorkflow(ctx, opts, app.FeedWorkflow)
		if err != nil {
			slog.Error("unable to execute workflow", "error", err)
			os.Exit(1)
		}
		slog.Info("started workflow", "workflow_id", we.GetID())
	}
}

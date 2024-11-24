package feed_tasks

import (
	"context"
	"fmt"
	"time"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func GenerateFeedsWorkflow(ctx workflow.Context, host string, count int) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	err := workflow.ExecuteActivity(ctx, a.GenerateFeeds, host, count).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

// returns external workflow job id
func RefreshFeedsWorkflow(ctx workflow.Context) error {
	// start the FetchFeedsWorkflow workflow inside the external worker
	//
	// this might seem convoluted, but the point is that graph should be able to ask
	// "feed tasks" to perform an action without worrying how. feed tasks are designed
	// to be background jobs. graph shouldn't care how that happens.
	config := lo.Must(env.New[Config]())
	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	opts := sdk.StartWorkflowOptions{
		TaskQueue: config.FeedTaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	we, err := client.ExecuteWorkflow(context.Background(), opts, app.FetchFeedsWorkflow /*"FetchFeedsWorkflow"*/)
	if err != nil {
		return fmt.Errorf("failed to execute workflow: %w", err)
	}
	workflow.GetLogger(ctx).Info("started workflow", "ID", we.GetID())
	return nil
}

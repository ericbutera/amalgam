package feed_tasks

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Khan/genqlient/graphql"
	graph_client "github.com/ericbutera/amalgam/pkg/clients/graphql"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/google/uuid"
	"github.com/samber/lo"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

const UrlFormat = "http://%s/feed/%s"

type Activities struct {
	graphClient graphql.Client
	logger      *slog.Logger
	feedClient  sdk.Client // separate temporal instance for feed worker
}

func NewActivities(graphClient graphql.Client, feedClient sdk.Client) *Activities {
	// note: tClient only talks to another external worker which could be on a different host.
	// do not use it to enqueue items in feed_task.
	return &Activities{
		graphClient: graphClient,
		logger:      slog.Default(),
		feedClient:  feedClient,
	}
}

func (a *Activities) GenerateFeeds(ctx context.Context, host string, count int) error {
	for x := 0; x < count; x++ {
		url := fmt.Sprintf(UrlFormat, host, uuid.New().String())
		resp, err := graph_client.AddFeed(ctx, a.graphClient, url, fmt.Sprintf("generated-%d", x))
		if err != nil {
			return err
		}
		a.logger.Debug("created feed", "feed_id", resp.AddFeed.Id)
	}
	return nil
}

func (a *Activities) RefreshFeeds(ctx context.Context) error {
	// start the Refresh workflow inside the external worker
	//
	// this might seem convoluted, but the point is that graph should be able to ask
	// "feed tasks" to perform an action without worrying how. feed tasks are designed
	// to be background jobs. graph shouldn't care how that happens.

	// TODO: dependency injection of client
	config := lo.Must(env.New[Config]())

	opts := sdk.StartWorkflowOptions{
		TaskQueue: config.FeedTaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	args := []any{}
	// _, err := a.feedClient.ExecuteWorkflow(ctx, opts, app.FetchFeedsWorkflow, args...)
	_, err := a.feedClient.ExecuteWorkflow(ctx, opts, "FetchFeedsWorkflow", args...)
	if err != nil {
		return fmt.Errorf("failed to execute workflow: %w", err)
	}
	return nil
}

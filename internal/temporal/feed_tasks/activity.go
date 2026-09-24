package feed_tasks

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/internal/temporal/feed_add"
	"github.com/ericbutera/amalgam/internal/temporal/feed_fetch"
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
	RetryPolicy temporal.RetryPolicy
	Config      *Config
}

func NewActivities(graphClient graphql.Client, feedClient sdk.Client) *Activities {
	// note: tClient only talks to another external worker which could be on a different host.
	// do not use it to enqueue items in feed_task.
	return &Activities{
		graphClient: graphClient,
		logger:      slog.Default(),
		feedClient:  feedClient,
		Config:      lo.Must(env.New[Config]()),
		RetryPolicy: temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
}

func (a *Activities) GenerateFeeds(ctx context.Context, host string, count int /*, userID string*/) error {
	for x := 0; x < count; x++ {
		url := fmt.Sprintf(UrlFormat, host, uuid.New().String())

		resp, err := graph_client.AddFeed(ctx, a.graphClient, url, fmt.Sprintf("generated-%d", x) /*, userID*/)
		if err != nil {
			return err
		}

		a.logger.Debug("enqueued feed add", "job_id", resp.AddFeed.Id)
	}

	return nil
}

func (a *Activities) RefreshFeeds(ctx context.Context) error {
	// Compatibility path only. New enqueue requests start FetchFeedsWorkflow
	// directly through internal/tasks.
	opts := sdk.StartWorkflowOptions{
		TaskQueue: a.Config.FeedFetchQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	args := []any{}

	_, err := a.feedClient.ExecuteWorkflow(ctx, opts, feed_fetch.WorkflowName, args...)
	return err
}

func (a *Activities) AddFeed(ctx context.Context, url string, userID string) (string, error) {
	// Compatibility path only. New enqueue requests start AddFeedWorkflow
	// directly through internal/tasks.
	opts := sdk.StartWorkflowOptions{
		TaskQueue:   a.Config.FeedAddQueue,
		RetryPolicy: &a.RetryPolicy,
	}
	var feedID string
	run, err := a.feedClient.ExecuteWorkflow(ctx, opts, feed_add.LegacyWorkflowName, url, userID)
	if err != nil {
		return "", err
	}
	err = run.Get(ctx, &feedID)
	if err != nil {
		return feedID, err
	}
	return feedID, nil
}

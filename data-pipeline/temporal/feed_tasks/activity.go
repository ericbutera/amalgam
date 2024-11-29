package feed_tasks

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/db/models"
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
	config := lo.Must(env.New[Config]())

	opts := sdk.StartWorkflowOptions{
		TaskQueue: config.FeedTaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	args := []any{}
	_, err := a.feedClient.ExecuteWorkflow(ctx, opts, "FetchFeedsWorkflow", args...)
	if err != nil {
		return fmt.Errorf("failed to execute workflow: %w", err)
	}
	return nil
}

func (*Activities) AssociateFeeds(_ context.Context) error {
	// find feeds without user_feeds - SELECT f.id FROM feeds f LEFT JOIN user_feeds uf ON f.id = uf.feed_id WHERE uf.feed_id IS NULL
	// try to insert into user_feeds

	// TODO: move into rpc if worth keeping (no direct db access here)
	db := lo.Must(db.NewFromEnv())

	var feeds []string
	res := db.Raw(`SELECT f.id FROM feeds f LEFT JOIN user_feeds uf ON f.id = uf.feed_id WHERE uf.feed_id IS NULL`).Scan(&feeds)
	if res.Error != nil {
		return fmt.Errorf("failed to find feeds: %w", res.Error)
	}

	slog.Info("found feeds", "count", len(feeds))

	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	userId := "0e597e90-ece5-463e-8608-ff687bf286da" // TODO: resolve user
	for _, feedID := range feeds {
		res := db.Create(&models.UserFeed{
			UserID:        userId,
			FeedID:        feedID,
			UnreadStartAt: oneMonthAgo,
		})
		slog.Info("associated feed", "feed_id", feedID)
		if res.Error != nil {
			return fmt.Errorf("failed to associate feed: %w", res.Error)
		}
	}

	return nil
}

package generate

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Khan/genqlient/graphql"
	graph_client "github.com/ericbutera/amalgam/graph/pkg/client"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/google/uuid"
)

const UrlFormat = "http://%s/feed/%s"

type Activities struct {
	client      pb.FeedServiceClient
	graphClient graphql.Client
	logger      *slog.Logger
}

func NewActivities(graphClient graphql.Client) *Activities {
	return &Activities{
		graphClient: graphClient,
		logger:      slog.Default(),
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

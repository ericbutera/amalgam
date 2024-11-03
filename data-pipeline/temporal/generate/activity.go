package generate

import (
	"context"
	"fmt"
	"log/slog"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/google/uuid"
)

type Activities struct {
	client pb.FeedServiceClient
}

func NewActivities(client pb.FeedServiceClient) *Activities {
	return &Activities{
		client: client,
	}
}

func (a *Activities) GenerateFeeds(ctx context.Context) error {
	logger := slog.Default() //logger := workflow.GetLogger(ctx)
	fmt.Print("generate feeds!")

	base := "http://%s/feed/%s"
	for x := 0; x < 10; x++ {
		url := fmt.Sprintf(base, "faker:8080", uuid.New().String())
		resp, err := a.client.CreateFeed(context.Background(), &pb.CreateFeedRequest{
			Feed: &pb.CreateFeedRequest_Feed{
				Url: url,
			},
		})
		if err != nil {
			return err
		}
		logger.Debug("created feed", "feed_id", resp.Id)
	}
	return nil
}

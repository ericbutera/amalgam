package feeds

import (
	"context"
	"fmt"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	rss "github.com/ericbutera/amalgam/pkg/feed/parse"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	rpc "github.com/ericbutera/amalgam/rpc/pkg/client"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Feed struct {
	Url string
	ID  string
}

type FeedHelper struct {
	client  pb.FeedServiceClient
	conn    *grpc.ClientConn
	closers []func() error
}

func NewFeeds(config *config.Config) (*FeedHelper, error) {
	rpc, err := rpc.NewClient(config.RpcHost, config.RpcInsecure)
	if err != nil {
		return nil, err
	}
	return &FeedHelper{
		client: rpc.Client,
		conn:   rpc.Conn,
		closers: []func() error{
			rpc.Conn.Close,
		},
	}, nil
}

func (h *FeedHelper) Close() error {
	for _, closer := range h.closers {
		if err := closer(); err != nil {
			return err
		}
	}
	return nil
}

func (h *FeedHelper) GetFeeds() ([]Feed, error) {
	// TODO: use rpc ListFeeds
	// resp, err := h.client.ListFeeds(context.Background(), &pb.ListFeedsRequest{})
	// if err != nil {
	// 	return nil, err
	// }
	// feeds := []Feed{}
	// for _, feed := range resp.Feeds {
	// 	feeds = append(feeds, Feed{
	// 		ID:  feed.Id,
	// 		Url: feed.Url,
	// 	})
	// }
	// return feeds, nil
	// TODO: toggle between real & faker
	base := "http://%s/feed/%s"
	feeds := []Feed{}
	for x := 0; x < 10; x++ {
		id := uuid.New().String()
		feeds = append(feeds, Feed{
			ID:  id,
			Url: fmt.Sprintf(base, "faker:8080", id),
		})
	}
	return feeds, nil
}

// Returns the article ID on success
func (h *FeedHelper) SaveArticle(ctx context.Context, article rss.Article) (string, error) {
	res, err := h.client.SaveArticle(ctx, &pb.SaveArticleRequest{
		Article: &pb.Article{
			FeedId:      article.FeedId,
			Title:       article.Title,
			Url:         article.Url,
			Preview:     article.Preview,
			Content:     article.Content,
			ImageUrl:    article.ImageUrl,
			Guid:        article.GUID,
			AuthorName:  article.AuthorName,
			AuthorEmail: article.AuthorEmail,
		},
	})
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

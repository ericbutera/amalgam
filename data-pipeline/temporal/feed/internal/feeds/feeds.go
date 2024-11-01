package feeds

import (
	"context"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	rss "github.com/ericbutera/amalgam/pkg/feed/parse"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	rpc "github.com/ericbutera/amalgam/rpc/pkg/client"
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
	return []Feed{
		// TODO: this will expire in 30 days from 2024-10-29
		{ID: "0e597e90-ece5-463e-8608-ff687bf286da", Url: "https://run.mocky.io/v3/883f6eb9-81d3-4648-9adf-6395c4e1567c"},
	}, nil
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

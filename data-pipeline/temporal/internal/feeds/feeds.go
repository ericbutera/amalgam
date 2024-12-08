package feeds

import (
	"context"

	rss "github.com/ericbutera/amalgam/pkg/feed/parse"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	rpc "github.com/ericbutera/amalgam/services/rpc/pkg/client"
)

// helper to interact with feed rpc service
type Feeds interface {
	Close() error
	GetFeeds(ctx context.Context) ([]Feed, error)
	SaveArticle(ctx context.Context, article rss.Article) (string, error)
	UpdateStats(ctx context.Context, feedID string) error
}

type Feed struct {
	Url string
	ID  string
}

type FeedHelper struct {
	client  pb.FeedServiceClient
	closers []func() error
}

func NewFeeds(host string, insecure bool) (Feeds, error) {
	c, closer, err := rpc.New(host, insecure)
	if err != nil {
		return nil, err
	}
	return &FeedHelper{
		client: c,
		closers: []func() error{
			closer,
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

func (h *FeedHelper) GetFeeds(ctx context.Context) ([]Feed, error) {
	resp, err := h.client.ListFeeds(ctx, &pb.ListFeedsRequest{})
	if err != nil {
		return nil, err
	}
	feeds := []Feed{}
	for _, feed := range resp.GetFeeds() {
		feeds = append(feeds, Feed{
			ID:  feed.GetId(),
			Url: feed.GetUrl(),
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
			Description: article.Description,
			ImageUrl:    article.ImageUrl,
			Guid:        article.GUID,
			AuthorName:  article.AuthorName,
			AuthorEmail: article.AuthorEmail,
		},
	})
	if err != nil {
		return "", err
	}
	return res.GetId(), nil
}

func (h *FeedHelper) UpdateStats(ctx context.Context, feedID string) error {
	// TODO: move Update Stats (feed + article count) into a durable workflow
	_, err := h.client.UpdateStats(ctx, &pb.UpdateStatsRequest{
		FeedId: feedID,
		Stat:   pb.UpdateStatsRequest_STAT_FEED_ARTICLE_COUNT,
	})
	return err
}

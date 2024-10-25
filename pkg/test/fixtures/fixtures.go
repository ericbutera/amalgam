package fixtures

import (
	db_model "github.com/ericbutera/amalgam/internal/db/models"
)

// TODO: figure out how to auto-generate these. for now copilot makes this easy enough
// inspired from https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/

func NewDbFeed(opts ...DbFeedOption) *db_model.Feed {
	f := &db_model.Feed{
		Name: "test feed name",
		Url:  "https://example.com/test-feed",
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type DbFeedOption func(*db_model.Feed)

func WithDbFeedID(id string) DbFeedOption {
	return func(f *db_model.Feed) {
		f.Base.ID = id
	}
}

func WithDbFeedName(name string) DbFeedOption {
	return func(f *db_model.Feed) {
		f.Name = name
	}
}

func WithDbFeedUrl(url string) DbFeedOption {
	return func(f *db_model.Feed) {
		f.Url = url
	}
}

func NewDbArticle(opts ...DbArticleOption) *db_model.Article {
	a := &db_model.Article{
		FeedID:   "feedidaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		Url:      "https://example.com/test-article",
		Title:    "test article title",
		ImageUrl: "https://example.com/test-article-image.jpg",
		Preview:  "test article preview",
		Content:  "test article content",
		Guid:     "guidaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

type DbArticleOption func(*db_model.Article)

func WithFeed(f *db_model.Feed) DbArticleOption {
	return func(a *db_model.Article) {
		a.Feed = *f
	}
}

func WithDbArticleFeedID(feedID string) DbArticleOption {
	return func(a *db_model.Article) {
		a.FeedID = feedID
	}
}

func WithDbArticleUrl(url string) DbArticleOption {
	return func(a *db_model.Article) {
		a.Url = url
	}
}

func WithDbArticleTitle(title string) DbArticleOption {
	return func(a *db_model.Article) {
		a.Title = title
	}
}

func WithDbArticlePreview(preview string) DbArticleOption {
	return func(a *db_model.Article) {
		a.Preview = preview
	}
}

func WithDbArticleContent(content string) DbArticleOption {
	return func(a *db_model.Article) {
		a.Content = content
	}
}

func WithDbArticleGuid(guid string) DbArticleOption {
	return func(a *db_model.Article) {
		a.Guid = guid
	}
}

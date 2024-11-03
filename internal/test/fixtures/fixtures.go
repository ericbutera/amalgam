package fixtures

import (
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
)

// TODO: figure out how to auto-generate these. for now copilot makes this easy enough
// inspired from https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/
// TODO: go faker

func NewFeed(opts ...FeedOption) *svc_model.Feed {
	f := &svc_model.Feed{
		Name:     "test feed name",
		Url:      "https://example.com/test-feed",
		IsActive: true,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type FeedOption func(*svc_model.Feed)

func WithFeedID(id string) FeedOption {
	return func(f *svc_model.Feed) {
		f.ID = id
	}
}

func WithFeedName(name string) FeedOption {
	return func(f *svc_model.Feed) {
		f.Name = name
	}
}

func WithFeedUrl(url string) FeedOption {
	return func(f *svc_model.Feed) {
		f.Url = url
	}
}

func WithFeedIsActive(isActive bool) FeedOption {
	return func(f *svc_model.Feed) {
		f.IsActive = isActive
	}
}

func NewArticle(opts ...ArticleOption) *svc_model.Article {
	a := &svc_model.Article{
		ID:       "articlea-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
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

type ArticleOption func(*svc_model.Article)

func WithArticleID(id string) ArticleOption {
	return func(a *svc_model.Article) {
		a.ID = id
	}
}

func WithArticleFeedID(feedID string) ArticleOption {
	return func(a *svc_model.Article) {
		a.FeedID = feedID
	}
}

func WithArticleUrl(url string) ArticleOption {
	return func(a *svc_model.Article) {
		a.Url = url
	}
}

func WithArticleTitle(title string) ArticleOption {
	return func(a *svc_model.Article) {
		a.Title = title
	}
}

func WithArticlePreview(preview string) ArticleOption {
	return func(a *svc_model.Article) {
		a.Preview = preview
	}
}

func WithArticleContent(content string) ArticleOption {
	return func(a *svc_model.Article) {
		a.Content = content
	}
}

func WithArticleGuid(guid string) ArticleOption {
	return func(a *svc_model.Article) {
		a.Guid = guid
	}
}

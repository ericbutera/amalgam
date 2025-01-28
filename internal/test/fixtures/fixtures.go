package fixtures

import (
	"net/http"

	"github.com/ericbutera/amalgam/internal/service/models"
	"github.com/google/uuid"
)

// TODO: figure out how to auto-generate these. for now copilot makes this easy enough
// inspired from https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/
// TODO: go faker

func NewID() string {
	return uuid.New().String()
}

func NewFeed(opts ...FeedOption) *models.Feed {
	f := &models.Feed{
		Name:     "test feed name",
		URL:      "https://example.com/test-feed",
		IsActive: true,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type FeedOption func(*models.Feed)

func WithFeedID(id string) FeedOption {
	return func(f *models.Feed) {
		f.ID = id
	}
}

func WithFeedName(name string) FeedOption {
	return func(f *models.Feed) {
		f.Name = name
	}
}

func WithFeedURL(url string) FeedOption {
	return func(f *models.Feed) {
		f.URL = url
	}
}

func WithFeedIsActive(isActive bool) FeedOption {
	return func(f *models.Feed) {
		f.IsActive = isActive
	}
}

func NewArticle(opts ...ArticleOption) *models.Article {
	a := &models.Article{
		ID:          "articlea-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		FeedID:      "feedidaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		URL:         "https://example.com/test-article",
		Title:       "test article title",
		ImageURL:    "https://example.com/test-article-image.jpg",
		Preview:     "test article preview",
		Content:     "test article content",
		Description: "test article description",
		GUID:        "guidaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

type ArticleOption func(*models.Article)

func WithArticleID(id string) ArticleOption {
	return func(a *models.Article) {
		a.ID = id
	}
}

func WithArticleFeedID(feedID string) ArticleOption {
	return func(a *models.Article) {
		a.FeedID = feedID
	}
}

func WithArticleURL(url string) ArticleOption {
	return func(a *models.Article) {
		a.URL = url
	}
}

func WithArticleTitle(title string) ArticleOption {
	return func(a *models.Article) {
		a.Title = title
	}
}

func WithArticlePreview(preview string) ArticleOption {
	return func(a *models.Article) {
		a.Preview = preview
	}
}

func WithArticleContent(content string) ArticleOption {
	return func(a *models.Article) {
		a.Content = content
	}
}

func WithArticleDescription(description string) ArticleOption {
	return func(a *models.Article) {
		a.Description = description
	}
}

func WithArticleGUID(guid string) ArticleOption {
	return func(a *models.Article) {
		a.GUID = guid
	}
}

func NewFeedVerification() *models.FeedVerification {
	return &models.FeedVerification{
		ID:         1,
		URL:        "https://example.com/test-feed",
		UserID:     "test-user-id",
		WorkflowID: "test-workflow-id",
	}
}

func NewFetchHistory(opts ...FetchHistoryOpt) *models.FetchHistory {
	h := &models.FetchHistory{
		ID:                 0,
		FeedID:             "test-feed-id",
		FeedVerificationID: 1,
		ResponseCode:       http.StatusOK,
		ETag:               "test-etag",
		WorkflowID:         "test-workflow-id",
		Bucket:             "test-bucket",
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

type FetchHistoryOpt func(*models.FetchHistory)

func WithFetchHistoryID(id int64) FetchHistoryOpt {
	return func(h *models.FetchHistory) {
		h.ID = id
	}
}

func WithFetchHistoryFeedID(feedID string) FetchHistoryOpt {
	return func(h *models.FetchHistory) {
		h.FeedID = feedID
	}
}

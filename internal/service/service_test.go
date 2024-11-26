package service_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/service"
	svcModel "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/test"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/ericbutera/amalgam/internal/test/seed"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type TestHelper struct {
	svc service.Service
	db  *gorm.DB
}

func newTestHelper(t *testing.T) *TestHelper {
	db := test.NewDB(t)
	svc := service.NewGorm(db)
	return &TestHelper{svc: svc, db: db}
}

func TestFeeds(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)
	feed := data.Feed

	feeds, err := h.svc.Feeds(context.Background())
	require.NoError(t, err)

	assert.Len(t, feeds, 1)
	assert.Equal(t, feed.URL, feeds[0].URL)
	assert.Equal(t, feed.Name, feeds[0].Name)
}

func TestFeeds_IsActive(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	feed := fixtures.NewFeed(fixtures.WithFeedIsActive(false))
	require.NoError(t, h.db.Create(&feed).Error)

	feeds, err := h.svc.Feeds(context.Background())
	require.NoError(t, err)

	assert.Empty(t, feeds)
}

func TestCreateFeed(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	expected := fixtures.NewFeed(fixtures.WithFeedURL("https://example.com/moo"))
	result, err := h.svc.CreateFeed(context.Background(), expected)
	require.NoError(t, err)

	actual := &svcModel.Feed{}
	res := h.db.First(actual, "url=?", expected.URL)
	require.NoError(t, res.Error)

	assert.Empty(t, result.ValidationErrors)
	assert.Equal(t, expected.URL, actual.URL)
	helpers.Diff(t, *expected, *actual, "ID")
}

func TestUpdateFeed(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)
	feed := data.Feed

	expected := &svcModel.Feed{
		Name: feed.Name,
		URL:  feed.URL,
	}
	require.NoError(t, h.svc.UpdateFeed(context.Background(), feed.ID, expected))

	var actual svcModel.Feed
	require.NoError(t, h.db.First(&actual, "id=?", feed.ID).Error)

	assert.Equal(t, expected.URL, actual.URL)
	assert.Equal(t, expected.Name, actual.Name)
}

func TestGetFeed(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	feed := data.Feed
	actual, err := h.svc.GetFeed(context.Background(), feed.ID)
	require.NoError(t, err)

	assert.Equal(t, feed.URL, actual.URL)
	assert.Equal(t, feed.Name, actual.Name)
}

func TestGetArticlesByFeed(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 2)
	require.NoError(t, err)

	res, err := h.svc.GetArticlesByFeed(context.Background(), data.Feed.ID, service.ListOptions{
		Cursor: "",
		Limit:  0,
	})
	require.NoError(t, err)

	expected := data.Articles
	actual := res.Articles
	assert.Len(t, actual, 2)
	helpers.Diff(t, *expected[0], actual[0], "ID")
	helpers.Diff(t, *expected[1], actual[1], "ID")
}

func TestGetArticle(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)
	article := data.Articles[0]

	actual, err := h.svc.GetArticle(context.Background(), article.ID)
	require.NoError(t, err)

	assert.Equal(t, article.URL, actual.URL)
	assert.Equal(t, article.Title, actual.Title)
}

func TestSaveArticle(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	expected := fixtures.NewArticle(
		fixtures.WithArticleURL("https://example.com/moo"),
		fixtures.WithArticleFeedID("0e597e90-ece5-463e-8608-ff687bf286da"),
	)
	result, err := h.svc.SaveArticle(context.Background(), expected)
	require.NoError(t, err)

	actual := &svcModel.Article{}
	res := h.db.First(actual, "url=?", expected.URL)
	require.NoError(t, res.Error)

	assert.Empty(t, result.ValidationErrors)
	assert.Equal(t, expected.URL, actual.URL)
	helpers.Diff(t, *expected, *actual, "ID")
}

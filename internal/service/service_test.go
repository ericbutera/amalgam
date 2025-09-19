package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/ericbutera/amalgam/internal/db/pagination"
	"github.com/ericbutera/amalgam/internal/service"
	svcModel "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/test"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/ericbutera/amalgam/internal/test/seed"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type testHelper struct {
	svc   service.Service
	db    *gorm.DB
	start time.Time
}

func newTestHelper(t *testing.T) *testHelper {
	t.Helper()
	db := test.NewDB(t)
	svc := service.NewGorm(db)

	return &testHelper{
		svc:   svc,
		db:    db,
		start: time.Now().UTC(),
	}
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

func TestCreateFeed_Duplicate(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	feed := fixtures.NewFeed()
	require.NoError(t, h.db.Create(&feed).Error)

	_, err := h.svc.CreateFeed(context.Background(), feed)
	assert.ErrorIs(t, err, service.ErrDuplicate)
}

func TestCreateFeed_Validation(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	feed := fixtures.NewFeed(fixtures.WithFeedURL("invalid-url"))
	result, err := h.svc.CreateFeed(context.Background(), feed)
	require.ErrorIs(t, err, service.ErrValidation)
	/*
		validate.ValidationError {
			Field: "URL",
			Tag: "url",
			RawMessage: "Key: 'Feed.URL' Error:Field validation for 'URL' failed on the 'url' tag",
			FriendlyMessage: "The URL must be valid."
		}
	*/
	e := result.ValidationErrors[0]
	assert.NotEmpty(t, result.ValidationErrors)
	assert.Equal(t, "URL", e.Field)
	assert.Equal(t, "url", e.Tag)
	assert.Contains(t, e.RawMessage, "Feed.URL")
	assert.Contains(t, e.FriendlyMessage, "URL")
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

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	res, err := h.svc.GetArticlesByFeed(context.Background(), data.Feed.ID, pagination.ListOptions{})
	require.NoError(t, err)

	expected := data.Articles
	actual := res.Articles
	assert.Len(t, actual, 1)
	helpers.Diff(t, *expected[0], actual[0], "ID")
}

func TestGetArticlesByFeed_Pagination(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 15)
	require.NoError(t, err)

	page1, err := h.svc.GetArticlesByFeed(context.Background(), data.Feed.ID, pagination.ListOptions{
		Limit: 10,
	})
	require.NoError(t, err)
	assert.Len(t, page1.Articles, 10)

	page2, err := h.svc.GetArticlesByFeed(context.Background(), data.Feed.ID, pagination.ListOptions{
		Cursor: page1.Cursor,
		Limit:  10,
	})
	require.NoError(t, err)
	assert.Len(t, page2.Articles, 5)
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
	helpers.Diff(t, *expected, *actual, "ID", "UpdatedAt")
	assert.Greater(t, actual.UpdatedAt, h.start)
}

func TestUserFeeds(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	result, err := h.svc.GetUserFeeds(context.Background(), data.UserFeed.UserID)
	require.NoError(t, err)

	assert.Len(t, result.Feeds, 1)
	assert.Equal(t, data.UserFeed.FeedID, result.Feeds[0].FeedID)
}

func TestUserFeed(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	result, err := h.svc.GetUserFeed(context.Background(), data.UserFeed.UserID, data.UserFeed.FeedID)
	require.NoError(t, err)

	assert.Equal(t, data.UserFeed.FeedID, result.FeedID)
}

func TestSaveUserFeed(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	feed := fixtures.NewFeed()
	require.NoError(t, h.db.Create(&feed).Error)

	uf := svcModel.UserFeed{
		UserID: seed.UserID,
		FeedID: feed.ID,
	}
	err := h.svc.SaveUserFeed(context.Background(), &uf)
	require.NoError(t, err)

	actual := &svcModel.UserFeed{}
	res := h.db.First(actual, "user_id=? AND feed_id=?", seed.UserID, feed.ID)
	require.NoError(t, res.Error)

	helpers.Diff(t, uf, *actual, "CreatedAt", "ViewedAt", "UnreadStartAt")
}

func TestSaveUserArticle(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	ua := svcModel.UserArticle{
		UserID:    data.UserFeed.UserID,
		ArticleID: data.Articles[0].ID,
		ViewedAt:  lo.ToPtr(time.Now().UTC()),
	}
	err = h.svc.SaveUserArticle(context.Background(), &ua)
	require.NoError(t, err)

	actual := &svcModel.UserArticle{}
	res := h.db.First(actual, "user_id=? AND article_id=?", data.UserFeed.UserID, data.Articles[0].ID)
	require.NoError(t, res.Error)

	helpers.Diff(t, ua, *actual)
}

func TestUserFeedArticleCount(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	uf, err := h.svc.GetUserFeed(context.Background(), data.UserFeed.UserID, data.Feed.ID)
	require.NoError(t, err)

	assert.Equal(t, int32(1), uf.UnreadCount)
}

func TestUpdateFeedArticleCount(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	data, err := seed.FeedAndArticles(h.db, 1)
	require.NoError(t, err)

	// add another article which will should the unread count
	article := fixtures.NewArticle(
		fixtures.WithArticleFeedID(data.Feed.ID),
		fixtures.WithArticleID(uuid.New().String()),
	)
	require.NoError(t, h.db.Create(&article).Error)

	err = h.svc.UpdateFeedArticleCount(context.Background(), data.Feed.ID)
	require.NoError(t, err)

	actual, err := h.svc.GetUserFeed(context.Background(), data.UserFeed.UserID, data.UserFeed.FeedID)
	require.NoError(t, err)

	assert.Equal(t, int32(2), actual.UnreadCount)
}

func TestCreateFeedVerification(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	verification := fixtures.NewFeedVerification()
	result, err := h.svc.CreateFeedVerification(context.Background(), verification)
	require.NoError(t, err)

	actual := &svcModel.FeedVerification{}
	res := h.db.First(actual, "id=?", result.ID)
	require.NoError(t, res.Error)

	require.NoError(t, err)
	helpers.Diff(t, *verification, *actual, "CreatedAt")
}

// TODO: test url normalization
// TODO: test prevent duplicate URL

func TestCreateFetchHistory(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)

	expected := fixtures.NewFetchHistory()
	history, err := h.svc.CreateFetchHistory(context.Background(), expected)
	require.NoError(t, err)

	actual := &svcModel.FetchHistory{}
	res := h.db.First(actual, "id=?", history.ID)
	require.NoError(t, res.Error)

	helpers.Diff(t, *expected, *actual, "ID", "CreatedAt")
}

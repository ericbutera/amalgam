package feed_fetch_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_fetch"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	"github.com/ericbutera/amalgam/pkg/feed/parse"
	"github.com/ericbutera/amalgam/pkg/test"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const TestFeedID = "feed-id-123"

func newReader(s string) io.ReadCloser {
	return io.NopCloser(bytes.NewBufferString(s))
}

type activitySetup struct {
	fetcher    *fetch.MockFetch
	bucket     *bucket.MockBucket
	feeds      *feeds.MockFeeds
	activities *app.Activities
}

func setupActivities(t *testing.T) *activitySetup {
	fetcher := fetch.NewMockFetch(t)
	bucketClient := bucket.NewMockBucket(t)
	feeds := feeds.NewMockFeeds(t)

	return &activitySetup{
		fetcher:    fetcher,
		bucket:     bucketClient,
		feeds:      feeds,
		activities: app.NewActivities(fetcher, bucketClient, feeds),
	}
}

func TestDownloadActivity(t *testing.T) {
	t.Parallel()

	url := "http://localhost/feed.xml"
	rssFile := app.RssPath(TestFeedID)
	reader := newReader("test data")
	contentType := "application/xml"
	size := int64(9)

	s := setupActivities(t)
	matcher := mock.MatchedBy(func(fn fetch.Callback) bool {
		err := fn(fetch.CallbackParams{Reader: reader, Size: size, ContentType: contentType})
		return lo.Ternary(err == nil, true, false)
	})
	s.fetcher.EXPECT().
		Url(mock.Anything, url, matcher, mock.Anything). // TODO: use etag
		Return(nil)
	s.bucket.EXPECT().
		Create(mock.Anything, app.BucketName).
		Return(nil)
	s.bucket.EXPECT().
		WriteStream(mock.Anything, app.BucketName, rssFile, reader, contentType).
		Return(&bucket.UploadInfo{Key: rssFile, Bucket: app.BucketName}, nil)

	out, err := s.activities.DownloadActivity(context.Background(), TestFeedID, url)
	require.NoError(t, err)
	assert.Equal(t, "feeds/"+TestFeedID+"/raw.xml", out)
}

func dataToArticles(t *testing.T, data io.Reader) parse.Articles {
	decoder := json.NewDecoder(data)

	articles := parse.Articles{}
	for decoder.More() {
		var a parse.Article
		require.NoError(t, decoder.Decode(&a))
		articles = append(articles, &a)
	}

	return articles
}

func compareArticles(t *testing.T, expected parse.Articles, actual parse.Articles) {
	require.Len(t, actual, len(expected))

	for x := 0; x < len(expected); x++ {
		a := expected[x]
		b := actual[x]
		test.Diff(t, *a, *b, "FeedId", "DateAdded", "DateUpdated", "DatePublished")
	}
}

func getRssData(t *testing.T) io.ReadCloser {
	d, err := test.FileToReadCloser("feeds/minimal.xml") // reader is single use
	require.NoError(t, err)

	return d
}

func TestParseActivity(t *testing.T) {
	t.Parallel()

	rssFile := app.RssPath(TestFeedID)
	articlesFile := app.ArticlePath(TestFeedID)

	s := setupActivities(t)

	matcher := mock.MatchedBy(func(data any) bool {
		reader := getRssData(t)
		defer reader.Close()

		expected, err := parse.Parse(reader)
		require.NoError(t, err)

		r, ok := data.(io.Reader)
		require.True(t, ok)
		actual := dataToArticles(t, r)
		compareArticles(t, expected, actual)

		return true
	})

	s.bucket.EXPECT().
		Read(mock.Anything, app.BucketName, rssFile).
		Return(getRssData(t), nil)

	s.bucket.EXPECT().
		WriteStream(mock.Anything, app.BucketName, articlesFile, mock.AnythingOfType("*bytes.Buffer"), app.ArticleContentType).
		Return(&bucket.UploadInfo{Key: articlesFile, Bucket: app.BucketName}, nil)

	out, err := s.activities.ParseActivity(context.Background(), TestFeedID, rssFile)
	require.NoError(t, err)
	assert.Equal(t, articlesFile, out)
	s.bucket.AssertCalled(t, "WriteStream", mock.Anything, app.BucketName, articlesFile, matcher, app.ArticleContentType)
}

func TestSaveActivity(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)

	article := parse.Article{
		FeedId: TestFeedID,
		Title:  "Test Article",
		Url:    "http://example.com/test",
	}

	s.bucket.EXPECT().
		Read(mock.Anything, app.BucketName, app.ArticlePath(TestFeedID)).
		Return(newReader(`{"feed_id":"feed-id-123","title":"Test Article","url":"http://example.com/test"}`), nil)

	s.feeds.EXPECT().
		SaveArticle(mock.Anything, article).
		Return("id", nil)

	articlesFile := app.ArticlePath(TestFeedID)

	results, err := s.activities.SaveActivity(context.Background(), TestFeedID, articlesFile)
	require.NoError(t, err)
	assert.Equal(t, app.SaveResults{Succeeded: 1, Failed: 0}, results)
}

func TestStatsActivity(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)

	s.feeds.EXPECT().
		UpdateStats(mock.Anything, TestFeedID).
		Return(nil)

	err := s.activities.StatsActivity(context.Background(), TestFeedID)
	require.NoError(t, err)
}

func TestGetFeedsActivity(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)

	data := []feeds.Feed{
		{ID: "feed-id-1", Url: "http://localhost/feed1.xml"},
		{ID: "feed-id-2", Url: "http://localhost/feed2.xml"},
	}

	s.feeds.EXPECT().
		GetFeeds(mock.Anything).
		Return(data, nil)

	urls, err := s.activities.GetFeedsActivity(context.Background())
	require.NoError(t, err)
	assert.Equal(t, data, urls)
}

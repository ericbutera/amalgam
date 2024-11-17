package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/transforms"
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

func newReader(s string) io.ReadCloser {
	return io.NopCloser(bytes.NewBufferString(s))
}

type activitySetup struct {
	// transforms *transforms.MockTransforms // transforms.Transforms
	transforms transforms.Transforms
	fetcher    *fetch.MockFetch
	bucket     *bucket.MockBucket
	feeds      *feeds.MockFeeds
	activities *app.Activities
}

func setupActivities(t *testing.T) *activitySetup {
	// transforms := transforms.NewMockTransforms(t)
	transforms := transforms.New()
	fetcher := fetch.NewMockFetch(t)
	bucketClient := bucket.NewMockBucket(t)
	feeds := feeds.NewMockFeeds(t)

	return &activitySetup{
		transforms: transforms,
		fetcher:    fetcher,
		bucket:     bucketClient,
		feeds:      feeds,
		activities: app.NewActivities(transforms, fetcher, bucketClient, feeds),
	}
}

func TestDownloadActivity(t *testing.T) {
	feedId := "feed-id-123"
	url := "http://localhost/feed.xml"
	rssFile := app.RssPath(feedId)
	reader := newReader("test data")
	contentType := "application/xml"
	size := int64(9)

	s := setupActivities(t)
	matcher := mock.MatchedBy(func(fn fetch.Callback) bool {
		err := fn(fetch.CallbackParams{Reader: reader, Size: size, ContentType: contentType})
		return lo.Ternary(err == nil, true, false)
	})
	s.fetcher.EXPECT().
		Url(mock.Anything, url, matcher).
		Return(nil)

	s.bucket.EXPECT().
		WriteStream(mock.Anything, app.BucketName, rssFile, reader, contentType).
		Return(&bucket.UploadInfo{Key: rssFile, Bucket: app.BucketName}, nil)

	out, err := s.activities.DownloadActivity(context.Background(), feedId, url)
	require.NoError(t, err)
	assert.Equal(t, "feeds/feed-id-123/raw.xml", out)
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
	require.Equal(t, len(expected), len(actual))
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
	feedId := "feed-id-123"
	rssFile := app.RssPath(feedId)
	articlesFile := app.ArticlePath(feedId)

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

	out, err := s.activities.ParseActivity(context.Background(), feedId, rssFile)
	require.NoError(t, err)
	assert.Equal(t, articlesFile, out)
	s.bucket.AssertCalled(t, "WriteStream", mock.Anything, app.BucketName, articlesFile, matcher, app.ArticleContentType)
}

func TestSaveActivity(t *testing.T) {
	s := setupActivities(t)

	article := parse.Article{
		FeedId: "feed-id-123",
		Title:  "Test Article",
		Url:    "http://example.com/test",
	}

	s.bucket.EXPECT().
		Read(mock.Anything, app.BucketName, app.ArticlePath("feed-id-123")).
		Return(newReader(`{"feed_id":"feed-id-123","title":"Test Article","url":"http://example.com/test"}`), nil)

	s.feeds.EXPECT().
		SaveArticle(mock.Anything, article).
		Return("id", nil)

	articlesFile := app.ArticlePath("feed-id-123")

	results, err := s.activities.SaveActivity(context.Background(), "feed-id-123", articlesFile)
	require.NoError(t, err)
	assert.Equal(t, app.SaveResults{Succeeded: 1, Failed: 0}, results)
}

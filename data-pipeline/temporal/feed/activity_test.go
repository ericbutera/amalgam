package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	feedHelper *feeds.MockFeeds
	activities *app.Activities
}

func setupActivities(t *testing.T) *activitySetup {
	// transforms := transforms.NewMockTransforms(t)
	transforms := transforms.New()
	fetcher := fetch.NewMockFetch(t)
	bucketClient := bucket.NewMockBucket(t)
	feedHelper := feeds.NewMockFeeds(t)

	return &activitySetup{
		transforms: transforms,
		fetcher:    fetcher,
		bucket:     bucketClient,
		feedHelper: feedHelper,
		activities: app.NewActivities(transforms, fetcher, bucketClient, feedHelper),
	}
}

func TestDownloadActivity(t *testing.T) {
	feedId := "feed-id-123"
	url := "http://localhost/feed.xml"
	rssFile := fmt.Sprintf(app.RssPathFormat, feedId)
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
		WriteStream(mock.Anything, app.BucketName, rssFile, reader, contentType, size).
		Return(&bucket.UploadInfo{Key: rssFile, Bucket: app.BucketName}, nil)

	out, err := s.activities.DownloadActivity(context.Background(), feedId, url)
	require.NoError(t, err)
	assert.Equal(t, "/feeds/feed-id-123/raw.xml", out)
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
	rssFile := fmt.Sprintf(app.RssPathFormat, feedId)
	articlesFile := fmt.Sprintf(app.ArticlePathFormat, feedId)
	contentType := "application/json"
	size := int64(1)

	s := setupActivities(t)

	matcher := mock.MatchedBy(func(data any) bool {
		expected, err := parse.Parse(getRssData(t))
		require.NoError(t, err)

		actual := dataToArticles(t, data.(io.Reader))
		compareArticles(t, expected, actual)
		return true
	})

	s.bucket.EXPECT().
		Read(mock.Anything, app.BucketName, rssFile).
		Return(getRssData(t), nil)

	s.bucket.EXPECT().
		WriteStream(mock.Anything, app.BucketName, articlesFile, mock.AnythingOfType("*bytes.Buffer"), contentType, size).
		Return(&bucket.UploadInfo{Key: articlesFile, Bucket: app.BucketName}, nil)

	out, err := s.activities.ParseActivity(context.Background(), feedId, rssFile)
	require.NoError(t, err)
	assert.Equal(t, articlesFile, out)
	s.bucket.AssertCalled(t, "WriteStream", mock.Anything, app.BucketName, articlesFile, matcher, contentType, size)
}

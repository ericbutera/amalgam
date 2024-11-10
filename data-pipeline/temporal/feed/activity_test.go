package app_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newReader(s string) io.Reader {
	return io.NopCloser(bytes.NewBufferString(s))
}

type activitySetup struct {
	fetcher    *fetch.MockFetch
	bucket     *bucket.MockBucket
	feedHelper *feeds.MockFeeds
	activities *app.Activities
}

func setupActivities(t *testing.T) *activitySetup {
	fetcher := fetch.NewMockFetch(t)
	bucketClient := bucket.NewMockBucket(t)
	feedHelper := feeds.NewMockFeeds(t)

	return &activitySetup{
		fetcher:    fetcher,
		bucket:     bucketClient,
		feedHelper: feedHelper,
		activities: app.NewActivities(fetcher, bucketClient, feedHelper),
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

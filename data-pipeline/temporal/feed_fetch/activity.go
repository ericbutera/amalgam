package feed_fetch

// TODO: https://docs.temporal.io/develop/go/testing-suite

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/transforms"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	"github.com/ericbutera/amalgam/pkg/feed/parse"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	BucketName         = "feeds"
	RssPathFormat      = "feeds/%s/raw.xml" // TODO: add a time value to filename to keep historical data
	ArticlePathFormat  = "feeds/%s/articles.jsonl"
	ArticleContentType = "application/json"
)

var ErrContentNotChanged = errors.New("content not changed")

type Activities struct {
	transforms transforms.Transforms
	fetch      fetch.Fetch
	bucket     bucket.Bucket
	feeds      feeds.Feeds
	Closers    func()
}

func NewActivities(fetch fetch.Fetch, bucket bucket.Bucket, feeds feeds.Feeds) *Activities {
	return &Activities{
		transforms: transforms.New(),
		fetch:      fetch,
		bucket:     bucket,
		feeds:      feeds,
		Closers:    func() {},
	}
}

func NewActivitiesFromEnv() *Activities {
	feeds := lo.Must(feeds.NewFeedsFromEnv())
	a := NewActivities(
		lo.Must(fetch.New()),
		lo.Must(bucket.NewMinioFromEnv()),
		feeds,
	)
	a.Closers = func() { feeds.Close() }
	return a
}

func RssPath(feedId string) string {
	return fmt.Sprintf(RssPathFormat, feedId)
}

// Download RSS feeds.
func (a *Activities) DownloadActivity(ctx context.Context, feedId string, url string) (string, error) {
	// TODO: check if we have a newly verified feed that needs to be processed (but not downloaded)
	// TODO: check fetch_history to see last time this feed was fetched (cooldown)
	// TODO: research temporal metrics to see if there is a built in way to view how long "downloads" are taking
	// TODO: also research to get counts of how many downloads are happening
	rssFile := RssPath(feedId)
	entry := slog.Default().With(
		"feed_id", feedId,
		"file", rssFile,
		"url", url,
	)
	// TODO: ensure fetch.Url makes a fetch_history entry
	err := a.fetch.Url(ctx, url, func(params fetch.CallbackParams) error {
		if params.StatusCode == http.StatusNotModified {
			return ErrContentNotChanged
		}
		upload, err := a.bucket.WriteStream(ctx, BucketName, rssFile, params.Reader, params.ContentType)
		if err != nil {
			return err
		}
		entry.Debug("downloaded activity", "key", upload.Key, "bucket", upload.Bucket, "size", upload.Size)
		return nil
	}, &fetch.ExtraParams{
		// TODO: Etag: "select etag from fetch_history where feed_id = ? order by created_at desc limit 1",
	})
	if err != nil {
		return "", err
	}
	return rssFile, nil
}

func ArticlePath(feedId string) string {
	return fmt.Sprintf(ArticlePathFormat, feedId)
}

// Transform raw rss file into structured articles.
// TODO: actually support a streaming bucket read -> rss to articles -> bucket write
func (a *Activities) ParseActivity(ctx context.Context, feedId string, rssFile string) (string, error) { // TODO: ParseActivity -> ArticleActivity
	articlesFile := ArticlePath(feedId)
	entry := slog.Default().With(
		"feed_id", feedId,
		"article_file", articlesFile,
	)

	rssReader, err := a.bucket.Read(ctx, BucketName, rssFile)
	if err != nil {
		return articlesFile, err
	}
	defer rssReader.Close()

	articles, err := a.transforms.RssToArticles(rssReader)
	if err != nil {
		return articlesFile, err
	}

	jsonl, errs := a.transforms.ArticleToJsonl(feedId, articles)
	if len(errs) > 0 {
		for err := range errs {
			entry.Debug("article to jsonlines", "error", err)
		}
	}

	upload, err := a.bucket.WriteStream(ctx, BucketName, articlesFile, &jsonl, ArticleContentType)
	if err != nil {
		return articlesFile, err
	}
	entry.Debug("parse activity: upload info", "key", upload.Key, "bucket", upload.Bucket, "size", upload.Size)
	return articlesFile, nil
}

type SaveResults struct {
	Succeeded int
	Failed    int
}

// Load articles into database.
func (a *Activities) SaveActivity(ctx context.Context, feedId string, articlesPath string) (SaveResults, error) {
	entry := slog.Default().With(
		"feed_id", feedId,
	)
	results := SaveResults{}

	articleReader, err := a.bucket.Read(ctx, BucketName, articlesPath)
	if err != nil {
		return results, err
	}

	defer articleReader.Close()
	decoder := json.NewDecoder(articleReader)

	for {
		var article parse.Article
		if err := decoder.Decode(&article); err != nil {
			if err == io.EOF {
				break
			}

			results.Failed++
			entry.Error("save: decode error", "error", err)
			continue
		}

		id, err := a.feeds.SaveArticle(ctx, article)
		if err != nil {
			handleSaveError(err, article.Url, entry)
			results.Failed++
			continue
		}

		results.Succeeded++
		entry.Debug("saved article", "article_url", article.Url, "article_id", id)
	}

	entry.Info("save results", "succeeded", results.Succeeded, "failed", results.Failed)

	// TODO: increment failure counter for alerting (does temporal have built in metrics for this?)
	return results, nil
}

func (a *Activities) StatsActivity(ctx context.Context, feedId string) error {
	return a.feeds.UpdateStats(ctx, feedId)
}

func handleSaveError(err error, url string, entry *slog.Logger) {
	// TODO: research recording errors like this using grpc middleware
	// need to have to correlate errors with the feed & article that caused it
	code := status.Code(err)
	if code == codes.InvalidArgument {
		for _, detail := range status.Convert(err).Details() {
			entry.Error("save article", "article_url", url, "error", detail)
		}
		return
	}
	entry.Error("save article", "article_url", url, "error", err)
}

func (a *Activities) GetFeedsActivity(ctx context.Context) ([]feeds.Feed, error) {
	urls, err := a.feeds.GetFeeds(ctx)
	if err != nil {
		return nil, err
	}

	// TODO write urls to bucket, return file path
	return urls, nil
}

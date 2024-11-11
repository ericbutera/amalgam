package app

// TODO: https://docs.temporal.io/develop/go/testing-suite

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/transforms"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	"github.com/ericbutera/amalgam/pkg/feed/parse"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	BucketName         = "feeds"
	RssPathFormat      = "feeds/%s/raw.xml"
	ArticlePathFormat  = "feeds/%s/articles.jsonl"
	ArticleContentType = "application/json"
)

type Activities struct {
	transforms transforms.Transforms
	fetch      fetch.Fetch
	bucket     bucket.Bucket
	feeds      feeds.Feeds
}

func NewActivities(transforms transforms.Transforms, fetch fetch.Fetch, bucket bucket.Bucket, feeds feeds.Feeds) *Activities {
	return &Activities{
		transforms: transforms,
		fetch:      fetch,
		bucket:     bucket,
		feeds:      feeds,
	}
}

func RssPath(feedId string) string {
	return fmt.Sprintf(RssPathFormat, feedId)
}

// Download RSS feeds.
func (a *Activities) DownloadActivity(ctx context.Context, feedId string, url string) (string, error) {
	// TODO: research temporal metrics to see if there is a built in way to view how long "downloads" are taking
	// TODO: also research to get counts of how many downloads are happening
	rssFile := RssPath(feedId)
	entry := slog.Default().With(
		"feed_id", feedId,
		"rss_file", rssFile,
		"url", url,
	)
	err := a.fetch.Url(ctx, url, func(params fetch.CallbackParams) error {
		upload, err := a.bucket.WriteStream(ctx, BucketName, rssFile, params.Reader, params.ContentType)
		if err != nil {
			return err
		}
		entry.Debug("downloaded activity", "key", upload.Key, "bucket", upload.Bucket, "size", upload.Size)
		return nil
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

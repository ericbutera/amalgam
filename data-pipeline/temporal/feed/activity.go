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

// Download RSS feeds.
func (a *Activities) DownloadActivity(ctx context.Context, feedId string, url string) (string, error) {
	// TODO: research temporal metrics to see if there is a built in way to view how long "downloads" are taking
	// TODO: also research to get counts of how many downloads are happening
	rssFile := fmt.Sprintf(RssPathFormat, feedId)
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

// Transform raw rss file into structured articles.
// TODO: actually support a streaming bucket read -> rss to articles -> bucket write
func (a *Activities) ParseActivity(ctx context.Context, feedId string, rssFile string) (string, error) { // TODO: ParseActivity -> ArticleActivity
	articlesFile := fmt.Sprintf(ArticlePathFormat, feedId)
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
		entry.Error("parse activity: encode errors", "errors", len(errs))
		for err := range errs {
			entry.Debug("article to jsonlines", "error", err)
		}
	}

	upload, err := a.bucket.WriteStream(ctx, BucketName, articlesFile, &jsonl, ArticleContentType)
	if err != nil {
		entry.Error("parse activity: write error", "error", err)
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
func (a *Activities) SaveActivity(ctx context.Context, feedId string, articlesPath string) error {
	entry := slog.Default().With(
		"feed_id", feedId,
	)

	articleReader, err := a.bucket.Read(ctx, BucketName, articlesPath)
	if err != nil {
		entry.Error("save: read error", "file", articlesPath, "error", err)
		return err
	}

	results := SaveResults{}

	defer articleReader.Close()
	decoder := json.NewDecoder(articleReader)

	// TODO: refactor to use a batch upsert
	// TODO: reduce branch complexity
	for {
		var article parse.Article
		if err := decoder.Decode(&article); err != nil {
			if err == context.Canceled {
				entry.Error("save: context canceled")
				results.Failed++
				continue // return nil
			}
			if err == context.DeadlineExceeded {
				entry.Error("save: context deadline exceeded")
				continue // return nil
			}
			if err == io.EOF {
				break
			}

			results.Failed++
			entry.Error("save: decode error", "error", err)
			return err
		}

		id, err := a.feeds.SaveArticle(ctx, article) // TODO: ensure idempotent upsert
		if err != nil {
			code := status.Code(err)
			if code == codes.InvalidArgument {
				for _, detail := range status.Convert(err).Details() {
					entry.Error("save article", "article_url", article.Url, "error", detail)
				}
			}

			entry.Error("save article", "article_url", article.Url, "error", err)
			results.Failed++
			continue
		}

		results.Succeeded++
		entry.Debug("saved article", "article_url", article.Url, "article_id", id)
	}

	entry.Debug("save results", "succeeded", results.Succeeded, "failed", results.Failed)
	return nil
}

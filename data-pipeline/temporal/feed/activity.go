package app

// TODO: https://docs.temporal.io/develop/go/testing-suite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	parse "github.com/ericbutera/amalgam/pkg/feed/parse"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	BucketName        = "feeds"
	RssPathFormat     = "/feeds/%s/raw.xml"
	ArticlePathFormat = "/feeds/%s/articles.jsonlines"
)

type Activities struct {
	bucket *bucket.MinioBucket
	feeds  *feeds.FeedHelper
}

func NewActivities(bucket *bucket.MinioBucket, feeds *feeds.FeedHelper) *Activities {
	return &Activities{
		bucket: bucket,
		feeds:  feeds,
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
	entry.Info("download: start")
	err := fetch.FetchUrl(ctx, url, func(params fetch.FetchCallbackParams) error {
		upload, err := a.bucket.WriteStream(ctx, BucketName, rssFile, params.Reader, params.ContentType, params.Size)
		if err != nil {
			entry.Error("download error", "error", err)
			return err
		}
		entry.Info("downloaded", "key", upload.Key, "bucket", upload.Bucket)
		return nil
	})
	if err != nil {
		return "", err
	}
	return rssFile, nil
}

// Transform raw rss file into structured articles.
func (a *Activities) ParseActivity(ctx context.Context, feedId string, rssFile string) (string, error) { // TODO: ParseActivity -> ArticleActivity
	articlesFile := fmt.Sprintf(ArticlePathFormat, feedId)
	entry := slog.Default().With(
		"feed_id", feedId,
		"article_file", articlesFile,
	)

	rssReader, err := a.bucket.Read(ctx, BucketName, rssFile)
	if err != nil {
		entry.Error("bucket read error", "error", err)
		return articlesFile, err
	}
	defer rssReader.Close()

	articles, err := parse.Parse(rssReader)
	if err != nil {
		entry.Error("parse: parse error", "error", err)
		return articlesFile, err
	}

	// TODO: actually support a streaming write
	var reader bytes.Buffer
	encoder := json.NewEncoder(&reader)

	for _, article := range articles {
		article.FeedId = feedId
		if err := encoder.Encode(article); err != nil {
			entry.Error("parse: error writing jsonlines", "error", err)
			continue // next article
		}
	}

	size := int64(reader.Len())
	upload, err := a.bucket.WriteStream(ctx, BucketName, articlesFile, &reader, "application/json", size)
	if err != nil {
		return articlesFile, err
	}
	entry.Info("parse: upload info", "key", upload.Key, "bucket", upload.Bucket)
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

	entry.Info("save results", "succeeded", results.Succeeded, "failed", results.Failed)
	return nil
}

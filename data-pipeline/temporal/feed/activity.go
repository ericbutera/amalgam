package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	parse "github.com/ericbutera/amalgam/pkg/feed/parse"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
)

const BucketName = "feeds"
const RssPathFormat = "/feeds/%s/raw.xml"
const ArticlePathFormat = "/feeds/%s/articles.jsonlines"

// TODO: common log fields

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

// Download RSS feeds
func (a *Activities) DownloadActivity(ctx context.Context, feedId string, url string) (string, error) {
	rssFile := fmt.Sprintf(RssPathFormat, feedId)
	err := FetchUrl(url, func(params FetchCallbackParams) error {
		upload, err := a.bucket.WriteStream(ctx, BucketName, rssFile, params.Reader, params.ContentType, params.Size)
		if err != nil {
			slog.Error("fetch: upload error", "file", rssFile, "error", err)
			return err
		}
		slog.Info("fetch: upload info", "file", rssFile, "key", upload.Key, "bucket", upload.Bucket)
		return nil
	})
	if err != nil {
		return "", err
	}
	return rssFile, nil
}

// Transform raw rss file into structured articles
func (a *Activities) ParseActivity(ctx context.Context, rssFile string, feedId string) (string, error) { // TODO: ParseActivity -> ArticleActivity
	articlesFile := fmt.Sprintf(ArticlePathFormat, feedId)

	rssReader, err := a.bucket.Read(ctx, BucketName, rssFile)
	if err != nil {
		slog.Error("parse: read error", "file", rssFile, "error", err)
		return "", err
	}
	defer rssReader.Close()

	articles, err := parse.Parse(rssReader)
	if err != nil {
		slog.Error("parse: parse error", "file", rssFile, "error", err)
		return "", err
	}

	// TODO: actually support a streaming write
	var reader bytes.Buffer
	encoder := json.NewEncoder(&reader)

	for _, article := range articles {
		article.FeedId = feedId
		if err := encoder.Encode(article); err != nil {
			slog.Error("parse: error writing jsonlines", "article_url", article.Url, "feed_id", feedId, "error", err)
			return "", err
		}
	}

	size := int64(reader.Len())
	upload, err := a.bucket.WriteStream(ctx, BucketName, articlesFile, &reader, "application/json", size)
	if err != nil {
		return "", err
	}
	slog.Info("parse: upload info", "file", articlesFile, "key", upload.Key, "bucket", upload.Bucket)
	return articlesFile, nil
}

// Load articles into database
func (a *Activities) SaveActivity(ctx context.Context, articlesPath string, feedId string) error {
	articleReader, err := a.bucket.Read(ctx, BucketName, articlesPath)
	if err != nil {
		slog.Error("save: read error", "file", articlesPath, "error", err)
		return err
	}

	defer articleReader.Close()
	decoder := json.NewDecoder(articleReader)

	for {
		var article parse.Article
		if err := decoder.Decode(&article); err != nil {
			if err == context.Canceled {
				slog.Error("save: context canceled")
				return nil
			}
			if err == context.DeadlineExceeded {
				slog.Error("save: context deadline exceeded")
				return nil
			}
			if err == io.EOF {
				break
			}
			slog.Error("save: decode error", "error", err)
			return err
		}

		id, err := a.feeds.SaveArticle(ctx, article)
		if err != nil {
			slog.Error("save: error saving article", "article_url", article.Url, "feed_id", feedId, "error", err)
			continue
		}
		slog.Debug("save: saved article", "article_url", article.Url, "feed_id", feedId, "article_id", id)
	}

	return nil
}

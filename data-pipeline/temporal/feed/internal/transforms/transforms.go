package transforms

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/ericbutera/amalgam/pkg/feed/parse"
)

type Transforms interface {
	RssToArticles(rss io.ReadCloser) (parse.Articles, error)
	ArticleToJsonl(feedId string, articles parse.Articles) (io.Reader, []error)
}

type transforms struct{}

func New() Transforms {
	return &transforms{}
}

func (*transforms) RssToArticles(rss io.ReadCloser) (parse.Articles, error) {
	return parse.Parse(rss)
}

func (*transforms) ArticleToJsonl(feedId string, articles parse.Articles) (io.Reader, []error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	var errs []error

	for _, article := range articles {
		article.FeedId = feedId
		if err := encoder.Encode(article); err != nil {
			errs = append(errs, err)
			continue // next article
		}
	}

	// reader := bytes.NewReader(&buf)
	var reader io.Reader = &buf
	return reader, errs
}
package transforms

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/ericbutera/amalgam/pkg/feed/parse"
)

type Transforms interface {
	RssToArticles(rss io.ReadCloser) (parse.Articles, error)
	ArticleToJsonl(feedId string, articles parse.Articles) (bytes.Buffer, []error)
}

type transforms struct{}

func New() Transforms {
	return &transforms{}
}

func (*transforms) RssToArticles(rss io.ReadCloser) (parse.Articles, error) {
	// TODO prom err counter
	return parse.Parse(rss)
}

func (*transforms) ArticleToJsonl(feedId string, articles parse.Articles) (bytes.Buffer, []error) {
	// TODO: research returning a reader instead of a buffer. io.Pipe?
	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)

	var errs []error

	for _, article := range articles {
		article.FeedId = feedId

		err := encoder.Encode(article)
		if err != nil {
			errs = append(errs, err)
			continue // next article
		}
	}

	return buf, errs
}

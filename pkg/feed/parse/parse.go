package parse

// TODO: rename package to "rss"

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/ericbutera/amalgam/internal/sanitize"
	parser "github.com/mmcdole/gofeed"
	"github.com/samber/lo"
)

var ErrSanitizeFailure = errors.New("failed to sanitize article")

type Articles []*Article

type Article struct {
	FeedId        string    `json:"feed_id"`
	Title         string    `json:"title"                    san:"text"`
	Url           string    `json:"url"                      san:"trim,url"`
	Preview       string    `json:"preview"                  san:"html"`
	Content       string    `json:"content,omitempty"        san:"html"` // full text of the article or post (usually empty, use Description by default)
	Description   string    `json:"description,omitempty"    san:"html"` // a brief summary or excerpt of the content
	ImageUrl      string    `json:"image_url,omitempty"      san:"trim,url"`
	GUID          string    `json:"guid,omitempty"`
	AuthorName    string    `json:"author_name,omitempty"`
	AuthorEmail   string    `json:"author_email,omitempty"`
	DateAdded     time.Time `json:"date_added,omitempty"`
	DateUpdated   time.Time `json:"date_updated,omitempty"`
	DatePublished time.Time `json:"date_published,omitempty"`
}

// Parse an RSS feed from a file path
func Path(path string) (Articles, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return Parse(reader)
}

// Convert an RSS feed into articles
func Parse(reader io.Reader) (Articles, error) {
	parsed, err := parser.NewParser().Parse(reader)
	if err != nil {
		return nil, err
	}

	articles := make(Articles, 0, len(parsed.Items))
	for _, item := range parsed.Items {
		article, err := NewArticleFromItem(item)
		if err != nil {
			slog.Error("parse: failed to convert item to article", "error", err)
			continue
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func NewArticleFromItem(item *parser.Item) (*Article, error) {
	description := lo.CoalesceOrEmpty(item.Description, item.Content)
	content := lo.CoalesceOrEmpty(item.Content, item.Description)
	article := Article{
		Title:         item.Title,
		Url:           item.Link,
		Preview:       Preview(description),
		Content:       content,
		Description:   description,
		ImageUrl:      getImageUrl(item),
		GUID:          item.GUID,
		AuthorName:    getAuthorName(item),
		AuthorEmail:   getAuthorEmail(item),
		DateAdded:     time.Now().UTC(),
		DateUpdated:   getDateUpdated(item),
		DatePublished: getDatePublished(item),
	}

	// note: this does "duplicate" work from the service layer, but it's important to validate early and often
	// service should still perform these as it is the final gate before persisting to the database
	article, err := sanitize.Struct(article)
	if err != nil {
		slog.Error("parse: failed to sanitize article", "article_url", article.Url, "error", err)
		return nil, errors.Join(ErrSanitizeFailure, err)
	}

	return &article, nil
}

func getImageUrl(item *parser.Item) (s string) {
	if item.Image != nil && item.Image.URL != "" {
		s = item.Image.URL
	}
	return s
}

func getAuthorName(item *parser.Item) (s string) {
	if item.Author != nil && item.Author.Name != "" {
		s = item.Author.Name
	}
	return s
}

func getAuthorEmail(item *parser.Item) (s string) {
	if item.Author != nil && item.Author.Email != "" {
		s = item.Author.Email
	}
	return s
}

func getDateUpdated(item *parser.Item) time.Time {
	if item.UpdatedParsed != nil {
		return *item.UpdatedParsed
	}
	return time.Now().UTC()
}

func getDatePublished(item *parser.Item) time.Time {
	if item.PublishedParsed != nil {
		return *item.PublishedParsed
	}
	return time.Now().UTC()
}

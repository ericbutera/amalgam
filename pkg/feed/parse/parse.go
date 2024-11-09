package parse

// TODO: rename package to "rss"

import (
	"io"
	"os"
	"time"

	parser "github.com/mmcdole/gofeed"
)

type Articles []*Article

// TODO: possibly use service layer
type Article struct {
	FeedId        string    `json:"feed_id"`
	Title         string    `json:"title"`
	Url           string    `json:"url"`
	Preview       string    `json:"preview"`
	Content       string    `json:"content,omitempty"`
	ImageUrl      string    `json:"image_url,omitempty"`
	GUID          string    `json:"guid,omitempty"`
	AuthorName    string    `json:"author_name,omitempty"`
	AuthorEmail   string    `json:"author_email,omitempty"`
	DateAdded     time.Time `json:"date_added,omitempty"`
	DateUpdated   time.Time `json:"date_updated,omitempty"`
	DatePublished time.Time `json:"date_published,omitempty"`
}

// Parse an RSS feed from a file path
func ParseWithPath(path string) (Articles, error) {
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
		articles = append(articles, NewArticleFromItem(item))
	}
	return articles, nil
}

func NewArticleFromItem(item *parser.Item) *Article {
	return &Article{
		Title:         item.Title,
		Url:           item.Link,
		Preview:       item.Description,
		Content:       item.Content,
		ImageUrl:      getImageUrl(item),
		GUID:          item.GUID,
		AuthorName:    getAuthorName(item),
		AuthorEmail:   getAuthorEmail(item),
		DateAdded:     time.Now().UTC(),
		DateUpdated:   getDateUpdated(item),
		DatePublished: getDatePublished(item),
	}
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

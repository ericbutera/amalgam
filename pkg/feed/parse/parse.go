package parse

import (
	"os"
	"time"

	"github.com/microcosm-cc/bluemonday"
	parser "github.com/mmcdole/gofeed"
)

type Articles []*Article
type Article struct {
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
func Parse(reader *os.File) (articles Articles, err error) {
	parsed, err := parser.NewParser().Parse(reader)
	if err != nil {
		return
	}

	for _, item := range parsed.Items {
		articles = append(articles, newArticleFromItem(item))
	}
	return
}

func newArticleFromItem(item *parser.Item) *Article {
	return &Article{
		Title:         item.Title,
		Url:           item.Link,
		Preview:       sanitize(item.Description),
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

// attempt to sanitize HTML content
func sanitize(html string) string {
	// quoting https://github.com/microcosm-cc/bluemonday:
	// bluemonday takes untrusted user generated content as an input, and will return HTML that has been
	// sanitised against an allowlist of approved HTML elements and attributes so that you can safely
	// include the content in your web page.
	return bluemonday.UGCPolicy().Sanitize(html)
}

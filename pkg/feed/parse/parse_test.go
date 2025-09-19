package parse_test

import (
	"testing"
	"time"

	"github.com/ericbutera/amalgam/pkg/feed/parse"
	"github.com/ericbutera/amalgam/pkg/test"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newMinimal(t *testing.T) parse.Articles {
	return newFile(t, "feeds/minimal.xml")
}

func newFile(t *testing.T, path string) parse.Articles {
	require.NotEmpty(t, path)
	path, err := test.GetTestDataPath(path)
	require.NoError(t, err)
	articles, err := parse.Path(path)
	require.NoError(t, err)

	return articles
}

func Test_Parse_ReturnsArticles(t *testing.T) {
	t.Parallel()
	articles := newMinimal(t)
	assert.Len(t, articles, 2)
}

func Test_Parse_Article(t *testing.T) {
	t.Parallel()
	articles := newMinimal(t)
	article := articles[0]
	assert.Equal(t, "example article", article.Title)
}

func Test_ItemToArticle(t *testing.T) {
	t.Parallel()

	expected := &parse.Article{
		Title:         "title",
		Url:           "https://example.com",
		Preview:       "test description",
		Content:       "test content",
		Description:   "test description",
		GUID:          "8b5e2d99-45f2-44a1-8c29-f19d95dff832",
		AuthorName:    "author name",
		AuthorEmail:   "author@example.com",
		ImageUrl:      "https://example.com/image.jpg",
		DateAdded:     time.Now().UTC(),
		DatePublished: time.Date(2024, 1, 1, 8, 5, 0, 0, time.UTC),
		DateUpdated:   time.Date(2024, 10, 18, 8, 5, 0, 0, time.UTC),
	}

	item := &gofeed.Item{
		Title:           expected.Title,
		Link:            expected.Url,
		Description:     expected.Description,
		Content:         expected.Content,
		GUID:            expected.GUID,
		Image:           &gofeed.Image{URL: expected.ImageUrl},
		Author:          &gofeed.Person{Name: expected.AuthorName, Email: expected.AuthorEmail},
		Updated:         "Fri, 18 Oct 2024 10:10:10 +0000",
		UpdatedParsed:   &expected.DateUpdated,
		Published:       "Fri, 18 Oct 2024 10:10:10 +0000",
		PublishedParsed: &expected.DatePublished,
	}
	article, err := parse.NewArticleFromItem(item)
	require.NoError(t, err)
	test.Diff(t, *expected, *article, "DateAdded", "DateUpdated", "DatePublished")
}

func TestArticleSanitization(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                string
		input               string
		expectedPreview     string
		expectedContent     string
		expectedDescription string
		length              int
	}{
		{
			name:                "xss test",
			input:               "<script>alert('xss')</script>",
			expectedPreview:     "alert(&#39;xss&#39;)",
			expectedDescription: "",
			expectedContent:     "",
		},
		{
			name:                "xss attribute",
			input:               "<p onclick='alert(/ohno/)'>xss</p>",
			expectedPreview:     "xss",
			expectedDescription: "<p>xss</p>",
			expectedContent:     "<p>xss</p>",
		},
		{
			name:                "happy path",
			input:               "<h1>title</h1><p>content</p>",
			expectedPreview:     "title content",
			expectedDescription: "<h1>title</h1><p>content</p>",
			expectedContent:     "<h1>title</h1><p>content</p>",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			item := &gofeed.Item{
				Description: tc.input,
				Content:     tc.input,
			}
			article, err := parse.NewArticleFromItem(item)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedPreview, article.Preview, "preview mismatch")
			assert.Equal(t, tc.expectedDescription, article.Description, "description mismatch")
			assert.Equal(t, tc.expectedContent, article.Content, "content mismatch")
		})
	}
}

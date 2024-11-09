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

// func newAtom(t *testing.T) Articles {
// 	return newFile(t, "feeds/atom.xml")
// }

func newFile(t *testing.T, path string) parse.Articles {
	require.NotEmpty(t, path)
	path, err := test.GetTestDataPath(path)
	require.NoError(t, err)
	articles, err := parse.Path(path)
	require.NoError(t, err)
	return articles
}

func Test_Parse_ReturnsArticles(t *testing.T) {
	articles := newMinimal(t)
	assert.Len(t, articles, 2)
}

func Test_Parse_Article(t *testing.T) {
	articles := newMinimal(t)
	article := articles[0]
	assert.Equal(t, "example article", article.Title)
}

func Test_ItemToArticle(t *testing.T) {
	expectedDate := time.Date(2024, 10, 18, 8, 5, 0, 0, time.UTC)
	item := &gofeed.Item{
		Title:           "title",
		Link:            "link",
		Description:     "description",
		Updated:         "Fri, 18 Oct 2024 10:10:10 +0000",
		UpdatedParsed:   &expectedDate,
		Published:       "Fri, 18 Oct 2024 10:10:10 +0000",
		PublishedParsed: &expectedDate,
	}
	article := parse.NewArticleFromItem(item)
	// TODO: add coverage for all supported fields
	assert.Equal(t, "title", article.Title)
	assert.Equal(t, "link", article.Url)
	assert.Equal(t, "description", article.Preview)
	assert.Equal(t, expectedDate, article.DateUpdated)
	assert.Equal(t, expectedDate, article.DatePublished)
}

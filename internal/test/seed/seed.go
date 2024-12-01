package seed

import (
	"github.com/ericbutera/amalgam/internal/converters"
	dbModels "github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/test/faker/rss"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const UserID = "a3dce526-adf2-4f2d-bda9-e42dfd520ba5"

type Result struct {
	Feed     *models.Feed
	Articles []*models.Article
	UserFeed *dbModels.UserFeeds
}

// Generates a fake feed with articles. Returns the feed ID.
func FeedAndArticles(db *gorm.DB, articleCount int) (*Result, error) {
	res := &Result{}
	c := converters.New()

	uuid := uuid.New().String()
	r, err := rss.Generate(uuid, articleCount)
	if err != nil {
		return nil, err
	}

	feed := c.ServiceToDbFeed(fixtures.NewFeed(
		fixtures.WithFeedID(uuid),
		fixtures.WithFeedURL(r.Channel.Link),
		fixtures.WithFeedName(r.Channel.Title),
	))
	if err := db.Create(&feed).Error; err != nil {
		return nil, err
	}
	res.Feed = c.DbToServiceFeed(feed)

	for _, item := range r.Channel.Items {
		article := c.ServiceToDbArticle(fixtures.NewArticle(
			fixtures.WithArticleFeedID(feed.ID),
			fixtures.WithArticlePreview(item.Description),
			fixtures.WithArticleContent(item.Description),
			fixtures.WithArticleDescription(item.Description),
			fixtures.WithArticleTitle(item.Title),
			fixtures.WithArticleURL(item.Link),
		))
		if err := db.Create(&article).Error; err != nil {
			return nil, err
		}
		res.Articles = append(res.Articles, c.DbToServiceArticle(article))
	}

	uf := dbModels.UserFeeds{
		UserID:      UserID,
		FeedID:      feed.ID,
		UnreadCount: int32(articleCount), //nolint:gosec
	}
	if err := db.Create(&uf).Error; err != nil {
		return nil, err
	}
	res.UserFeed = &uf

	return res, nil
}

package service

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/pkg/test/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceSuite struct {
	suite.Suite
	svc *Service
	db  *gorm.DB
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	t := s.T()
	s.db = mustNewTestDb(t)
	s.svc = New(s.db)
}

func mustNewTestDb(t *testing.T) *gorm.DB {
	db, err := db.NewSqlite(
		"file::memory:?cache=shared",
		db.WithAutoMigrate(),
		// Debug: db.WithTraceAll(),
	)
	require.NoError(t, err)
	return db
}

func mustCreate[T models.AllowedModels](t *testing.T, db *gorm.DB, records ...*T) {
	require.NoError(t, models.Create(db, records...))
}

func (s *ServiceSuite) TestListFeeds() {
	t := s.T()

	feed := fixtures.NewDbFeed()
	mustCreate(t, s.db, &feed)

	feeds, err := s.svc.Feeds(context.Background())
	require.NoError(t, err)

	assert.Len(t, feeds, 1)
	assert.Equal(t, feed.Url, feeds[0].Url)
	assert.Equal(t, feed.Name, feeds[0].Name)
}

func (s *ServiceSuite) TestListArticles() {
	t := s.T()

	feed := fixtures.NewDbFeed()
	mustCreate(t, s.db, &feed)
	article1 := fixtures.NewDbArticle(fixtures.WithDbArticleFeedID(feed.Base.ID))
	mustCreate(t, s.db, &article1)
	article2 := fixtures.NewDbArticle(fixtures.WithDbArticleFeedID(feed.Base.ID))
	mustCreate(t, s.db, &article2)

	articles, err := s.svc.GetArticlesByFeed(context.Background(), feed.ID)
	require.NoError(t, err)

	assert.Len(t, articles, 2)
	assert.Equal(t, article1.FeedID, articles[0].FeedID)
	assert.Equal(t, article1.Url, articles[0].Url)
	assert.Equal(t, article1.Title, articles[0].Title)
	assert.Equal(t, article2.Url, articles[1].Url)
}

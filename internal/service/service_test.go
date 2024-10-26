package service

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/pkg/test/fixtures"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceSuite struct {
	suite.Suite
	svc Service
	db  *gorm.DB
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	t := s.T()
	s.db = mustNewTestDb(t)
	s.svc = NewGorm(s.db)
}

func mustNewTestDb(t *testing.T) *gorm.DB {
	db, err := db.NewSqlite(
		"file::memory:", // do not share (no cleanup)
		db.WithAutoMigrate(),
		db.WithTraceAll(),
	)
	require.NoError(t, err)
	return db
}

func mustCreate[T models.AllowedModels](t *testing.T, db *gorm.DB, records ...*T) {
	require.NoError(t, models.Create(db, records...))
}

func (s *ServiceSuite) TestFeeds() {
	t := s.T()

	feed := fixtures.NewDbFeed()
	mustCreate(t, s.db, &feed)

	feeds, err := s.svc.Feeds(context.Background())
	require.NoError(t, err)

	assert.Len(t, feeds, 1)
	assert.Equal(t, feed.Url, feeds[0].Url)
	assert.Equal(t, feed.Name, feeds[0].Name)
}

func (s *ServiceSuite) TestCreateFeed() {
	t := s.T()

	//feed := fixtures.NewDbFeed()
	err := s.svc.CreateFeed(context.Background(), &Feed{
		Name: "moo",
		Url:  "https://example.com/moo",
	})
	require.NoError(t, err)
	// TODO: verify feed was created in db
}

func (s *ServiceSuite) TestUpdateFeed() {
	t := s.T()

	feed := fixtures.NewDbFeed()
	mustCreate(t, s.db, &feed)

	expected := &Feed{
		Name: feed.Name,
		Url:  feed.Url,
	}
	err := s.svc.UpdateFeed(context.Background(), feed.ID, expected)
	require.NoError(t, err)

	var actual Feed //models.Feed
	require.NoError(t, s.db.First(&actual, "id=?", feed.ID).Error)

	assert.Equal(t, expected.Url, actual.Url)
	assert.Equal(t, expected.Name, actual.Name)
}

func (s *ServiceSuite) TestGetFeed() {
	t := s.T()

	feed := fixtures.NewDbFeed()
	mustCreate(t, s.db, &feed)

	actual, err := s.svc.GetFeed(context.Background(), feed.ID)
	require.NoError(t, err)

	assert.Equal(t, feed.Url, actual.Url)
	assert.Equal(t, feed.Name, actual.Name)
}

func (s *ServiceSuite) TestGetArticlesByFeed() {
	t := s.T()

	feed := fixtures.NewDbFeed()
	mustCreate(t, s.db, &feed)
	article0 := fixtures.NewDbArticle(fixtures.WithDbArticleFeedID(feed.Base.ID))
	mustCreate(t, s.db, &article0)
	article1 := fixtures.NewDbArticle(fixtures.WithDbArticleFeedID(feed.Base.ID))
	mustCreate(t, s.db, &article1)

	articles, err := s.svc.GetArticlesByFeed(context.Background(), feed.ID)
	require.NoError(t, err)

	assert.Len(t, articles, 2)
	assert.Equal(t, article0.FeedID, articles[0].FeedID)
	assert.Equal(t, article0.Url, articles[0].Url)
	assert.Equal(t, article0.Title, articles[0].Title)
	assert.Equal(t, article1.Url, articles[1].Url)

	ignored := cmpopts.IgnoreFields(Article{}, "ID")
	if diff := cmp.Diff(article0, articles[0], ignored); diff != "" {
		t.Errorf("article mismatch (-want +got):\n%s", diff)
	}
}

func (s *ServiceSuite) TestGetArticle() {
	t := s.T()

	article := fixtures.NewDbArticle()
	mustCreate(t, s.db, &article)

	actual, err := s.svc.GetArticle(context.Background(), article.ID)
	require.NoError(t, err)

	assert.Equal(t, article.Url, actual.Url)
	assert.Equal(t, article.Title, actual.Title)
}

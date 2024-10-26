package service_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/db"
	service "github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceSuite struct {
	suite.Suite
	svc service.Service
	db  *gorm.DB
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	t := s.T()
	s.db = mustNewTestDb(t)
	s.svc = service.NewGorm(s.db)
}

func mustNewTestDb(t *testing.T) *gorm.DB {
	db, err := db.NewSqlite(
		"file::memory:", // do not share (new db per test)
		db.WithAutoMigrate(),
		// db.WithTraceAll(),
	)
	require.NoError(t, err)
	return db
}

type AllowedModels interface {
	*service.Feed | *service.Article
}

func Create[T AllowedModels](db *gorm.DB, records ...*T) error {
	for _, record := range records {
		if err := db.Create(&record).Error; err != nil {
			return err
		}
	}
	return nil
}

func mustCreate[T AllowedModels](t *testing.T, db *gorm.DB, records ...*T) {
	require.NoError(t, Create(db, records...))
}

func (s *ServiceSuite) TestFeeds() {
	t := s.T()

	feed := fixtures.NewFeed()
	mustCreate(t, s.db, &feed)

	feeds, err := s.svc.Feeds(context.Background())
	require.NoError(t, err)

	assert.Len(t, feeds, 1)
	assert.Equal(t, feed.Url, feeds[0].Url)
	assert.Equal(t, feed.Name, feeds[0].Name)
}

func (s *ServiceSuite) TestCreateFeed() {
	t := s.T()

	expected := fixtures.NewFeed(fixtures.WithFeedUrl("https://example.com/moo"))
	err := s.svc.CreateFeed(context.Background(), expected)
	require.NoError(t, err)

	actual := &service.Feed{}
	res := s.db.First(actual, "url=?", expected.Url)
	require.NoError(t, res.Error)

	assert.Equal(t, expected.Url, actual.Url)
	helpers.Diff(t, *expected, *actual, "ID")
}

func (s *ServiceSuite) TestUpdateFeed() {
	t := s.T()

	feed := fixtures.NewFeed()
	mustCreate(t, s.db, &feed)

	expected := &service.Feed{
		Name: feed.Name,
		Url:  feed.Url,
	}
	err := s.svc.UpdateFeed(context.Background(), feed.ID, expected)
	require.NoError(t, err)

	var actual service.Feed
	require.NoError(t, s.db.First(&actual, "id=?", feed.ID).Error)

	assert.Equal(t, expected.Url, actual.Url)
	assert.Equal(t, expected.Name, actual.Name)
}

func (s *ServiceSuite) TestGetFeed() {
	t := s.T()

	feed := fixtures.NewFeed()
	mustCreate(t, s.db, &feed)

	actual, err := s.svc.GetFeed(context.Background(), feed.ID)
	require.NoError(t, err)

	assert.Equal(t, feed.Url, actual.Url)
	assert.Equal(t, feed.Name, actual.Name)
}

func (s *ServiceSuite) TestGetArticlesByFeed() {
	t := s.T()

	feed := fixtures.NewFeed()
	mustCreate(t, s.db, &feed)
	expected0 := fixtures.NewArticle(
		fixtures.WithArticleID("articlea-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
		fixtures.WithArticleFeedID(feed.ID),
		fixtures.WithArticleUrl("https://example.com/0"),
	)
	mustCreate(t, s.db, &expected0)
	expected1 := fixtures.NewArticle(
		fixtures.WithArticleID("articleb-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
		fixtures.WithArticleFeedID(feed.ID),
		fixtures.WithArticleUrl("https://example.com/0"),
	)
	mustCreate(t, s.db, &expected1)

	actual, err := s.svc.GetArticlesByFeed(context.Background(), feed.ID)
	require.NoError(t, err)

	assert.Len(t, actual, 2)
	assert.Equal(t, expected0.FeedID, actual[0].FeedID)
	assert.Equal(t, expected0.Url, actual[0].Url)
	assert.Equal(t, expected0.Title, actual[0].Title)
	assert.Equal(t, expected1.Url, actual[1].Url)

	helpers.Diff(t, *expected0, actual[0], "ID")
	helpers.Diff(t, *expected1, actual[1], "ID")
}

func (s *ServiceSuite) TestGetArticle() {
	t := s.T()

	article := fixtures.NewArticle()
	mustCreate(t, s.db, &article)

	actual, err := s.svc.GetArticle(context.Background(), article.ID)
	require.NoError(t, err)

	assert.Equal(t, article.Url, actual.Url)
	assert.Equal(t, article.Title, actual.Title)
}

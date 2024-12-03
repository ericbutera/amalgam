package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ericbutera/amalgam/internal/converters"
	db_model "github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/internal/db/pagination"
	"github.com/ericbutera/amalgam/internal/sanitize"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/validate"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const DefaultLimit = 100

var ErrQueryFailed = errors.New("query failed")

var (
	validateFeedCreate = validate.CustomMessages{
		"URL.required": "The URL is required.",
		"URL.url":      "The URL must be valid.",
	}
	validateArticleSave = validate.CustomMessages{
		"ID.required":       "The ID field is required.",
		"ID.uuid4":          "The ID must be a valid UUID.",
		"FeedID.required":   "The FeedID field is required.",
		"FeedID.uuid4":      "The FeedID must be a valid UUID.",
		"URL.required":      "The URL is required.",
		"URL.url":           "The URL must be valid.",
		"Title.required":    "The title is required.",
		"Title.min":         "The title must be at least 3 characters long.",
		"ImageURL.url":      "The Image URL must be a valid URL if provided.",
		"Preview.required":  "The preview is required.",
		"Preview.min":       "The preview must be at least 5 characters long.",
		"Content.required":  "The content is required.",
		"AuthorEmail.email": "The author email must be a valid email address if provided.",
	}
)

type Gorm struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) Service {
	return &Gorm{db: db}
}

// query returns a new query builder with the given context. required for otel
func (s *Gorm) query(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx)
}

func (s *Gorm) Feeds(ctx context.Context) ([]svc_model.Feed, error) {
	var feeds []svc_model.Feed
	result := s.query(ctx).
		Where("is_active=?", true).
		Order("name").       // TODO index name
		Limit(DefaultLimit). // TODO: pagination
		Find(&feeds)

	if result.Error != nil {
		return nil, ErrQueryFailed
	}
	return feeds, nil
}

func feedUrlExists(tx *gorm.DB, url string) error {
	res := tx.Find(&svc_model.Feed{}, "url=?", url)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return ErrDuplicate
	}
	return nil
}

type CreateFeedResult struct {
	ID               string
	ValidationErrors []validate.ValidationError
}

func (s *Gorm) CreateFeed(ctx context.Context, data *svc_model.Feed) (CreateFeedResult, error) {
	res := CreateFeedResult{}

	feed, err := sanitize.Struct(lo.FromPtr(data))
	if err != nil {
		return res, fmt.Errorf("unable to create feed: %w", err)
	}
	res.ValidationErrors = validate.Struct(feed, validateFeedCreate).Errors
	if len(res.ValidationErrors) > 0 {
		return res, ErrValidation
	}

	dbFeed := converters.New().ServiceToDbFeed(&feed)
	dbFeed.IsActive = true

	err = s.query(ctx).Transaction(func(tx *gorm.DB) error {
		if err := feedUrlExists(tx, feed.URL); err != nil {
			return err
		}
		if err := tx.Create(&dbFeed).Error; err != nil {
			return err
		}
		res.ID = dbFeed.ID
		return nil
	})

	return res, err
}

func (s *Gorm) UpdateFeed(ctx context.Context, id string, feed *svc_model.Feed) error {
	// note: no validation required for update
	dbFeed := converters.New().ServiceToDbFeed(feed)

	result := s.query(ctx).
		Model(&dbFeed).
		Where("id=?", id).
		Updates(map[string]any{
			"name": feed.Name,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *Gorm) GetFeed(ctx context.Context, id string) (*svc_model.Feed, error) {
	var feed svc_model.Feed
	result := s.query(ctx).First(&feed, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &feed, nil
}

func (s *Gorm) GetArticlesByFeed(ctx context.Context, feedId string, options pagination.ListOptions) (*ArticlesByFeedResult, error) {
	// TODO: revisit using `pagination.ListOptions`  - it should really be a service style
	query := s.query(ctx).
		Model(&svc_model.Article{}).
		Where("feed_id=?", feedId)

	res, err := pagination.Pager[svc_model.Article](query, options, []paginator.Rule{
		{
			Key:     "UpdatedAt",
			Order:   paginator.DESC,
			SQLRepr: "updated_at",
		},
	})
	if err != nil {
		return nil, err
	}
	return &ArticlesByFeedResult{
		Articles: res.Results,
		Cursor: pagination.Cursor{
			Next:     res.Cursor.Next,
			Previous: res.Cursor.Previous,
		},
	}, nil
}

func (s *Gorm) GetArticle(ctx context.Context, id string) (*svc_model.Article, error) {
	var article svc_model.Article
	result := s.query(ctx).First(&article, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &article, nil
}

type SaveArticleResult struct {
	ID               string
	ValidationErrors []validate.ValidationError
}

func (s *Gorm) SaveArticle(ctx context.Context, data *svc_model.Article) (SaveArticleResult, error) {
	res := SaveArticleResult{}

	article, err := sanitize.Struct(lo.FromPtr(data))
	if err != nil {
		return res, err
	}
	res.ValidationErrors = validate.Struct(article, validateArticleSave).Errors

	if len(res.ValidationErrors) > 0 {
		return res, ErrValidation
	}

	dbArticle := converters.New().ServiceToDbArticle(&article)

	result := s.query(ctx).
		Model(&db_model.Article{}).
		Omit(clause.Associations).
		Create(&dbArticle)

	if err := result.Error; err != nil {
		return res, err
	}

	res.ID = dbArticle.ID
	return res, err
}

func (s *Gorm) userFeeds(ctx context.Context, userID string) *gorm.DB {
	return s.query(ctx).
		Table("user_feeds uf").
		Select(
			"f.id feed_id", "f.name", "f.url",
			"uf.created_at", "uf.viewed_at", "uf.unread_start_at", "uf.unread_count",
		).
		Joins("JOIN feeds f ON uf.feed_id = f.id").
		Where("uf.user_id=?", userID).
		Order("f.name").
		Limit(DefaultLimit)
}

func (s *Gorm) GetUserFeed(ctx context.Context, userID string, feedID string) (*svc_model.UserFeed, error) {
	var feed svc_model.UserFeed
	result := s.userFeeds(ctx, userID).Where("uf.feed_id=?", feedID).First(&feed)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &feed, nil
}

func (s *Gorm) GetUserFeeds(ctx context.Context, userID string) (*GetUserFeedsResult, error) {
	result := &GetUserFeedsResult{}
	query := s.userFeeds(ctx, userID).Find(&result.Feeds)
	if query.Error != nil {
		return nil, query.Error
	}
	return result, nil
}

func (s *Gorm) SaveUserFeed(ctx context.Context, uf *svc_model.UserFeed) error {
	return s.query(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_id"},
				{Name: "feed_id"},
			},
			DoNothing: true,
		}).
		Create(&db_model.UserFeeds{ // TODO convets.SvcToDbUserFeed
			UserID:        uf.UserID,
			FeedID:        uf.FeedID,
			CreatedAt:     time.Now().UTC(),
			ViewedAt:      time.Now().UTC(),
			UnreadStartAt: time.Now().AddDate(0, -1, 0).UTC(), // -1 month
			UnreadCount:   0,                                  // TODO: show last N articles as new
		}).Error
}

func (s *Gorm) SaveUserArticle(ctx context.Context, ua *svc_model.UserArticle) error {
	return s.query(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_id"},
				{Name: "article_id"},
			},
			DoNothing: true,
		}).
		Create(&db_model.UserArticles{ // TODO converts.SvcToDbUserArticle
			UserID:    ua.UserID,
			ArticleID: ua.ArticleID,
			ViewedAt:  ua.ViewedAt, // TODO: ensure nil case (mark unread)
		}).Error
}

func (s *Gorm) findFeedUserIDs(ctx context.Context, feedID string) ([]string, error) {
	var userIds []string
	res := s.query(ctx).
		Model(&db_model.UserFeeds{}).
		Select("user_id").
		Find(&userIds, "feed_id=?", feedID)
	return userIds, res.Error
}

func (s *Gorm) updateArticleCount(ctx context.Context, userID string, feedID string) error {
	sql := `
	UPDATE user_feeds
	SET unread_count = (
		SELECT
			count(a.id)
		FROM articles a
		LEFT JOIN user_articles ua ON a.id = ua.article_id
		WHERE
			a.feed_id = ? AND
			ua.viewed_at IS NULL
	)
	WHERE user_id = ? AND feed_id = ?
	`
	return s.query(ctx).Exec(sql, feedID, userID, feedID).Error
}

func (s *Gorm) UpdateFeedArticleCount(ctx context.Context, feedID string) error {
	userIDs, err := s.findFeedUserIDs(ctx, feedID)
	if err != nil {
		return err
	}

	hasError := false
	for _, userID := range userIDs {
		if err := s.updateArticleCount(ctx, userID, feedID); err != nil {
			if !hasError {
				slog.Error("update user feed article count", "error", err) // TODO: rate limit
				hasError = true
			}
		}
	}
	if hasError {
		return ErrQueryFailed
	}
	return nil
}

func (s *Gorm) GetUserArticles(ctx context.Context, userID string, articleIDs []string) ([]*svc_model.UserArticle, error) {
	var userArticles []*svc_model.UserArticle
	err := s.query(ctx).
		Table("user_articles").
		Select("article_id", "viewed_at").
		Where("user_id=?", userID).
		Where("article_id IN (?)", articleIDs).
		Find(&userArticles).Error
	if err != nil {
		return nil, err
	}
	return userArticles, nil
}

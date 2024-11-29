package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feed struct {
	Base
	URL      string `json:"url"      gorm:"uniqueIndex" binding:"required,url" example:"https://example.com/"`
	Name     string `json:"name"     example:"Example"`
	IsActive bool   `json:"isActive" example:"true"`
}

type Article struct {
	Base
	FeedID      string `json:"feedId"            binding:"required"                                    example:"1"`
	Feed        Feed   `gorm:"foreignKey:FeedID"`
	URL         string `json:"url"               binding:"required"                                    example:"https://example.com/"`
	Title       string `json:"title"             example:"Example Article"`
	ImageURL    string `json:"imageUrl"          example:"https://example.com/image.jpg"`
	Preview     string `json:"preview"           example:"Article preview text. May contain HTML."`
	Content     string `json:"content"           example:"Article content text. May contain HTML."`
	Description string `json:"description"       example:"Description content text. May contain HTML."`
	GUID        string `json:"guid"              example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	AuthorName  string `json:"authorName"        example:"Eric Butera"`
	AuthorEmail string `json:"authorEmail"       example:"example@example.com"`
}

type User struct {
	Base
	UserUUID       string `json:"userUuid"       example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	ProviderUserID string `json:"providerUserId" example:"eric"`
	Username       string `json:"username"       binding:"required"                             example:"eric"`
	Name           string `json:"name"           example:"Eric Butera"`
	Email          string `json:"email"          example:"example@example.com"`
	PhotoURL       string `json:"photoUrl"       example:"https://example.com/image.jpg"`
}

type Base struct {
	ID        string         `json:"id"        gorm:"type:uuid;primary_key;"  binding:"required"`
	CreatedAt time.Time      `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type UserFeed struct {
	UserID        string    `gorm:"column:user_id;primaryKey"`
	FeedID        string    `gorm:"column:feed_id;primaryKey"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	ViewedAt      time.Time `gorm:"column:viewed_at;autoUpdateTime"`
	UnreadStartAt time.Time `gorm:"column:unread_start_at"`
}

type UserArticle struct {
	UserID    string    `gorm:"column:user_id;primaryKey"`
	ArticleID string    `gorm:"column:article_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	ViewedAt  time.Time `gorm:"column:viewed_at;autoUpdateTime"`
}

// Credit to https://medium.com/@amrilsyaifa_21001/how-to-use-uuid-in-gorm-golang-74be997d7087
// BeforeCreate will set a UUID rather than numeric ID.
func (b *Base) BeforeCreate(_ *gorm.DB) error {
	b.ID = uuid.New().String()
	return nil
}

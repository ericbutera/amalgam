package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO: decouple API models from database models

type Feed struct {
	Base
	Url  string `json:"url" gorm:"uniqueIndex" binding:"required,url" example:"https://example.com/"`
	Name string `json:"name" example:"Example"`
}

type Article struct {
	Base
	FeedID      uint   `json:"feedId" binding:"required" example:"1"`
	Feed        Feed   `gorm:"foreignKey:FeedID"`
	Url         string `json:"url" binding:"required" example:"https://example.com/"`
	Title       string `json:"title" example:"Example Article"`
	ImageUrl    string `json:"imageUrl" example:"https://example.com/image.jpg"`
	Preview     string `json:"preview" example:"Article preview text. May contain HTML."`
	Content     string `json:"content" example:"Article content text. May contain HTML."`
	Guid        string `json:"guid" example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	AuthorName  string `json:"authorName" example:"Eric Butera"`
	AuthorEmail string `json:"authorEmail" example:"example@example.com"`
}

type User struct {
	Base
	UserUUID       string `json:"userUuid" example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	ProviderUserID string `json:"providerUserId" example:"eric"`
	Username       string `json:"username" binding:"required" example:"eric"`
	Name           string `json:"name" example:"Eric Butera"`
	Email          string `json:"email" example:"example@example.com"`
	PhotoURL       string `json:"photoUrl" example:"https://example.com/image.jpg"`
}

type Base struct {
	ID        uint           `gorm:"primarykey" json:"id" binding:"required" example:"1"`
	CreatedAt time.Time      `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

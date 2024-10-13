package models

import (
	"time"

	"gorm.io/gorm"
)

type Feed struct {
	Base
	Url  string `example:"https://example.com/"`
	Name string `example:"Example"`
}

type Article struct {
	Base
	FeedID      uint   `example:"1"`
	Feed        Feed   `gorm:"foreignKey:FeedID"`
	Url         string `example:"https://example.com/"`
	Title       string `example:"Example Article"`
	ImageUrl    string `example:"https://example.com/image.jpg"`
	Preview     string `example:"Article preview text. May contain HTML."`
	Content     string `example:"Article content text. May contain HTML."`
	Guid        string `example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	AuthorName  string `example:"Eric Butera"`
	AuthorEmail string `example:"example@example.com"`
}

type User struct {
	Base
	UserUUID       string `example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	ProviderUserID string `example:"eric"`
	Username       string `example:"eric"`
	Name           string `example:"Eric Butera"`
	Email          string `example:"example@example.com"`
	PhotoURL       string `example:"https://example.com/image.jpg"`
}

type Base struct {
	gorm.Model
	ID        uint      `gorm:"primarykey" example:"1"`
	CreatedAt time.Time `example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `example:"2021-01-01T00:00:00Z"`
}

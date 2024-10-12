package models

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	Url  string
	Name string
}

type Article struct {
	gorm.Model
	FeedID      uint
	Feed        Feed `gorm:"foreignKey:FeedID"`
	Url         string
	Title       string
	ImageUrl    string
	Preview     string
	Content     string
	Guid        string
	AuthorName  string
	AuthorEmail string
}

type User struct {
	gorm.Model
	UserUUID       string
	ProviderUserID string
	Username       string
	Name           string
	Email          string
	PhotoURL       string
}

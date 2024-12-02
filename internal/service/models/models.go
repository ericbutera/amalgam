package models

import "time"

type Feed struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"      validate:"required,url"`
	IsActive bool   `json:"isActive"`
}

type Article struct {
	ID          string    `json:"id"`
	FeedID      string    `json:"feedId"      validate:"required,uuid4"`
	URL         string    `json:"url"         san:"trim,url"            validate:"required,url"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"imageUrl"    validate:"omitempty,url"`
	Preview     string    `json:"preview"     san:"html"`
	Content     string    `json:"content"     san:"html"`
	Description string    `json:"description" san:"html"`
	GUID        string    `json:"guid"`
	AuthorName  string    `json:"authorName"`
	AuthorEmail string    `json:"authorEmail"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserFeed struct {
	FeedID        string    `json:"feed_id"`
	UserID        string    `json:"user_id,omitempty"`
	Name          string    `json:"name"`
	URL           string    `json:"url"`
	CreatedAt     time.Time `json:"createdAt"`
	ViewedAt      time.Time `json:"viewedAt"`
	UnreadStartAt time.Time `json:"unreadStartAt"`
	UnreadCount   int32     `json:"unreadCount"`
}

type UserArticle struct {
	UserID    string     `json:"user_id"`
	ArticleID string     `json:"article_id"`
	ViewedAt  *time.Time `json:"viewed_at"`
}

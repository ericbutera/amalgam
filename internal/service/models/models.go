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
	ImageURL    string    `json:"imageUrl"    validate:"url"`
	Preview     string    `json:"preview"     san:"html"`
	Content     string    `json:"content"     san:"html"`
	Description string    `json:"description" san:"html"`
	GUID        string    `json:"guid"`
	AuthorName  string    `json:"authorName"`
	AuthorEmail string    `json:"authorEmail"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserFeed struct {
	FeedID        string    `json:"feedId"`
	UserID        string    `json:"userId,omitempty"`
	Name          string    `json:"name"`
	URL           string    `json:"url"`
	CreatedAt     time.Time `json:"createdAt"`
	ViewedAt      time.Time `json:"viewedAt"`
	UnreadStartAt time.Time `json:"unreadStartAt"`
	UnreadCount   int32     `json:"unreadCount"`
}

type UserArticle struct {
	UserID    string     `json:"userId"`
	ArticleID string     `json:"articleId"`
	ViewedAt  *time.Time `json:"viewedAt"`
}

type FeedVerification struct {
	ID         int64  `json:"id"`
	URL        string `json:"url" san:"trim,url" validate:"required,url"`
	UserID     string
	WorkflowID string
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type FetchHistory struct {
	ID                 int64  `json:"id"`
	FeedID             string `json:"feedId"`
	FeedVerificationID int64  `json:"feedVerificationId"`
	ResponseCode       int32  `json:"responseCode"`
	ETag               string `json:"eTag"`
	WorkflowID         string `json:"workflowId"`
	Bucket             string `json:"bucket"`
	Message            string `json:"message"`
	CreatedAt          time.Time
}

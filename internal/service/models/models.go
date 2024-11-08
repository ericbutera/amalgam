package models

type Feed struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"      validate:"required,url"`
	IsActive bool   `json:"isActive"`
}

type Article struct {
	ID          string `json:"id"`
	FeedID      string `json:"feedId"      validate:"required,uuid4"`
	Url         string `json:"url"         san:"trim,url"            validate:"required,url"`
	Title       string `json:"title"`
	ImageUrl    string `json:"imageUrl"    validate:"omitempty,url"`
	Preview     string `json:"preview"     san:"html"`
	Content     string `json:"content"     san:"html"`
	Guid        string `json:"guid"`
	AuthorName  string `json:"authorName"`
	AuthorEmail string `json:"authorEmail"`
}

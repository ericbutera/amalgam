package models

type Feed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Article struct {
	ID          string `json:"id"`
	FeedID      string `json:"feedId"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	ImageUrl    string `json:"imageUrl"`
	Preview     string `json:"preview"`
	Content     string `json:"content"`
	Guid        string `json:"guid"`
	AuthorName  string `json:"authorName"`
	AuthorEmail string `json:"authorEmail"`
}

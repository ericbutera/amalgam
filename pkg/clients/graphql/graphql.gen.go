// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package client

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// AddFeedAddFeedAddResponse includes the requested fields of the GraphQL type AddResponse.
type AddFeedAddFeedAddResponse struct {
	Id string `json:"id"`
}

// GetId returns AddFeedAddFeedAddResponse.Id, and is useful for accessing the field via an interface.
func (v *AddFeedAddFeedAddResponse) GetId() string { return v.Id }

// AddFeedResponse is returned by AddFeed on success.
type AddFeedResponse struct {
	AddFeed AddFeedAddFeedAddResponse `json:"addFeed"`
}

// GetAddFeed returns AddFeedResponse.AddFeed, and is useful for accessing the field via an interface.
func (v *AddFeedResponse) GetAddFeed() AddFeedAddFeedAddResponse { return v.AddFeed }

// GetArticleArticle includes the requested fields of the GraphQL type Article.
type GetArticleArticle struct {
	Id          string    `json:"id"`
	FeedId      string    `json:"feedId"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	ImageUrl    string    `json:"imageUrl"`
	Content     string    `json:"content"`
	Preview     string    `json:"preview"`
	Description string    `json:"description"`
	Guid        string    `json:"guid"`
	AuthorName  string    `json:"authorName"`
	AuthorEmail string    `json:"authorEmail"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GetId returns GetArticleArticle.Id, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetId() string { return v.Id }

// GetFeedId returns GetArticleArticle.FeedId, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetFeedId() string { return v.FeedId }

// GetUrl returns GetArticleArticle.Url, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetUrl() string { return v.Url }

// GetTitle returns GetArticleArticle.Title, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetTitle() string { return v.Title }

// GetImageUrl returns GetArticleArticle.ImageUrl, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetImageUrl() string { return v.ImageUrl }

// GetContent returns GetArticleArticle.Content, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetContent() string { return v.Content }

// GetPreview returns GetArticleArticle.Preview, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetPreview() string { return v.Preview }

// GetDescription returns GetArticleArticle.Description, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetDescription() string { return v.Description }

// GetGuid returns GetArticleArticle.Guid, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetGuid() string { return v.Guid }

// GetAuthorName returns GetArticleArticle.AuthorName, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetAuthorName() string { return v.AuthorName }

// GetAuthorEmail returns GetArticleArticle.AuthorEmail, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetAuthorEmail() string { return v.AuthorEmail }

// GetUpdatedAt returns GetArticleArticle.UpdatedAt, and is useful for accessing the field via an interface.
func (v *GetArticleArticle) GetUpdatedAt() time.Time { return v.UpdatedAt }

// GetArticleResponse is returned by GetArticle on success.
type GetArticleResponse struct {
	Article GetArticleArticle `json:"article"`
}

// GetArticle returns GetArticleResponse.Article, and is useful for accessing the field via an interface.
func (v *GetArticleResponse) GetArticle() GetArticleArticle { return v.Article }

// GetFeedFeed includes the requested fields of the GraphQL type Feed.
type GetFeedFeed struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

// GetId returns GetFeedFeed.Id, and is useful for accessing the field via an interface.
func (v *GetFeedFeed) GetId() string { return v.Id }

// GetUrl returns GetFeedFeed.Url, and is useful for accessing the field via an interface.
func (v *GetFeedFeed) GetUrl() string { return v.Url }

// GetName returns GetFeedFeed.Name, and is useful for accessing the field via an interface.
func (v *GetFeedFeed) GetName() string { return v.Name }

// GetFeedResponse is returned by GetFeed on success.
type GetFeedResponse struct {
	Feed GetFeedFeed `json:"feed"`
}

// GetFeed returns GetFeedResponse.Feed, and is useful for accessing the field via an interface.
func (v *GetFeedResponse) GetFeed() GetFeedFeed { return v.Feed }

// ListArticlesArticlesArticlesResponse includes the requested fields of the GraphQL type ArticlesResponse.
type ListArticlesArticlesArticlesResponse struct {
	Articles []ListArticlesArticlesArticlesResponseArticlesArticle `json:"articles"`
	Cursor   ListArticlesArticlesArticlesResponseCursor            `json:"cursor"`
}

// GetArticles returns ListArticlesArticlesArticlesResponse.Articles, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponse) GetArticles() []ListArticlesArticlesArticlesResponseArticlesArticle {
	return v.Articles
}

// GetCursor returns ListArticlesArticlesArticlesResponse.Cursor, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponse) GetCursor() ListArticlesArticlesArticlesResponseCursor {
	return v.Cursor
}

// ListArticlesArticlesArticlesResponseArticlesArticle includes the requested fields of the GraphQL type Article.
type ListArticlesArticlesArticlesResponseArticlesArticle struct {
	Id          string    `json:"id"`
	FeedId      string    `json:"feedId"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	ImageUrl    string    `json:"imageUrl"`
	Preview     string    `json:"preview"`
	AuthorName  string    `json:"authorName"`
	AuthorEmail string    `json:"authorEmail"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GetId returns ListArticlesArticlesArticlesResponseArticlesArticle.Id, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetId() string { return v.Id }

// GetFeedId returns ListArticlesArticlesArticlesResponseArticlesArticle.FeedId, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetFeedId() string { return v.FeedId }

// GetUrl returns ListArticlesArticlesArticlesResponseArticlesArticle.Url, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetUrl() string { return v.Url }

// GetTitle returns ListArticlesArticlesArticlesResponseArticlesArticle.Title, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetTitle() string { return v.Title }

// GetImageUrl returns ListArticlesArticlesArticlesResponseArticlesArticle.ImageUrl, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetImageUrl() string { return v.ImageUrl }

// GetPreview returns ListArticlesArticlesArticlesResponseArticlesArticle.Preview, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetPreview() string { return v.Preview }

// GetAuthorName returns ListArticlesArticlesArticlesResponseArticlesArticle.AuthorName, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetAuthorName() string {
	return v.AuthorName
}

// GetAuthorEmail returns ListArticlesArticlesArticlesResponseArticlesArticle.AuthorEmail, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetAuthorEmail() string {
	return v.AuthorEmail
}

// GetUpdatedAt returns ListArticlesArticlesArticlesResponseArticlesArticle.UpdatedAt, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseArticlesArticle) GetUpdatedAt() time.Time {
	return v.UpdatedAt
}

// ListArticlesArticlesArticlesResponseCursor includes the requested fields of the GraphQL type ResponseCursor.
type ListArticlesArticlesArticlesResponseCursor struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// GetNext returns ListArticlesArticlesArticlesResponseCursor.Next, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseCursor) GetNext() string { return v.Next }

// GetPrevious returns ListArticlesArticlesArticlesResponseCursor.Previous, and is useful for accessing the field via an interface.
func (v *ListArticlesArticlesArticlesResponseCursor) GetPrevious() string { return v.Previous }

// ListArticlesResponse is returned by ListArticles on success.
type ListArticlesResponse struct {
	Articles ListArticlesArticlesArticlesResponse `json:"articles"`
}

// GetArticles returns ListArticlesResponse.Articles, and is useful for accessing the field via an interface.
func (v *ListArticlesResponse) GetArticles() ListArticlesArticlesArticlesResponse { return v.Articles }

// ListFeedsFeedsFeedResponse includes the requested fields of the GraphQL type FeedResponse.
type ListFeedsFeedsFeedResponse struct {
	Feeds []ListFeedsFeedsFeedResponseFeedsFeed `json:"feeds"`
}

// GetFeeds returns ListFeedsFeedsFeedResponse.Feeds, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponse) GetFeeds() []ListFeedsFeedsFeedResponseFeedsFeed { return v.Feeds }

// ListFeedsFeedsFeedResponseFeedsFeed includes the requested fields of the GraphQL type Feed.
type ListFeedsFeedsFeedResponseFeedsFeed struct {
	Id            string    `json:"id"`
	Url           string    `json:"url"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	ViewedAt      time.Time `json:"viewedAt"`
	UnreadStartAt time.Time `json:"unreadStartAt"`
	UnreadCount   int       `json:"unreadCount"`
}

// GetId returns ListFeedsFeedsFeedResponseFeedsFeed.Id, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetId() string { return v.Id }

// GetUrl returns ListFeedsFeedsFeedResponseFeedsFeed.Url, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetUrl() string { return v.Url }

// GetName returns ListFeedsFeedsFeedResponseFeedsFeed.Name, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetName() string { return v.Name }

// GetCreatedAt returns ListFeedsFeedsFeedResponseFeedsFeed.CreatedAt, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetCreatedAt() time.Time { return v.CreatedAt }

// GetViewedAt returns ListFeedsFeedsFeedResponseFeedsFeed.ViewedAt, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetViewedAt() time.Time { return v.ViewedAt }

// GetUnreadStartAt returns ListFeedsFeedsFeedResponseFeedsFeed.UnreadStartAt, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetUnreadStartAt() time.Time { return v.UnreadStartAt }

// GetUnreadCount returns ListFeedsFeedsFeedResponseFeedsFeed.UnreadCount, and is useful for accessing the field via an interface.
func (v *ListFeedsFeedsFeedResponseFeedsFeed) GetUnreadCount() int { return v.UnreadCount }

// ListFeedsResponse is returned by ListFeeds on success.
type ListFeedsResponse struct {
	Feeds ListFeedsFeedsFeedResponse `json:"feeds"`
}

// GetFeeds returns ListFeedsResponse.Feeds, and is useful for accessing the field via an interface.
func (v *ListFeedsResponse) GetFeeds() ListFeedsFeedsFeedResponse { return v.Feeds }

// UpdateFeedResponse is returned by UpdateFeed on success.
type UpdateFeedResponse struct {
	UpdateFeed UpdateFeedUpdateFeedUpdateResponse `json:"updateFeed"`
}

// GetUpdateFeed returns UpdateFeedResponse.UpdateFeed, and is useful for accessing the field via an interface.
func (v *UpdateFeedResponse) GetUpdateFeed() UpdateFeedUpdateFeedUpdateResponse { return v.UpdateFeed }

// UpdateFeedUpdateFeedUpdateResponse includes the requested fields of the GraphQL type UpdateResponse.
type UpdateFeedUpdateFeedUpdateResponse struct {
	Id string `json:"id"`
}

// GetId returns UpdateFeedUpdateFeedUpdateResponse.Id, and is useful for accessing the field via an interface.
func (v *UpdateFeedUpdateFeedUpdateResponse) GetId() string { return v.Id }

// __AddFeedInput is used internally by genqlient
type __AddFeedInput struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

// GetUrl returns __AddFeedInput.Url, and is useful for accessing the field via an interface.
func (v *__AddFeedInput) GetUrl() string { return v.Url }

// GetName returns __AddFeedInput.Name, and is useful for accessing the field via an interface.
func (v *__AddFeedInput) GetName() string { return v.Name }

// __GetArticleInput is used internally by genqlient
type __GetArticleInput struct {
	Id string `json:"id"`
}

// GetId returns __GetArticleInput.Id, and is useful for accessing the field via an interface.
func (v *__GetArticleInput) GetId() string { return v.Id }

// __GetFeedInput is used internally by genqlient
type __GetFeedInput struct {
	Id string `json:"id"`
}

// GetId returns __GetFeedInput.Id, and is useful for accessing the field via an interface.
func (v *__GetFeedInput) GetId() string { return v.Id }

// __ListArticlesInput is used internally by genqlient
type __ListArticlesInput struct {
	FeedId string `json:"feedId"`
}

// GetFeedId returns __ListArticlesInput.FeedId, and is useful for accessing the field via an interface.
func (v *__ListArticlesInput) GetFeedId() string { return v.FeedId }

// __UpdateFeedInput is used internally by genqlient
type __UpdateFeedInput struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

// GetId returns __UpdateFeedInput.Id, and is useful for accessing the field via an interface.
func (v *__UpdateFeedInput) GetId() string { return v.Id }

// GetUrl returns __UpdateFeedInput.Url, and is useful for accessing the field via an interface.
func (v *__UpdateFeedInput) GetUrl() string { return v.Url }

// GetName returns __UpdateFeedInput.Name, and is useful for accessing the field via an interface.
func (v *__UpdateFeedInput) GetName() string { return v.Name }

// The query or mutation executed by AddFeed.
const AddFeed_Operation = `
mutation AddFeed ($url: String!, $name: String!) {
	addFeed(url: $url, name: $name) {
		id
	}
}
`

func AddFeed(
	ctx_ context.Context,
	client_ graphql.Client,
	url string,
	name string,
) (*AddFeedResponse, error) {
	req_ := &graphql.Request{
		OpName: "AddFeed",
		Query:  AddFeed_Operation,
		Variables: &__AddFeedInput{
			Url:  url,
			Name: name,
		},
	}
	var err_ error

	var data_ AddFeedResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetArticle.
const GetArticle_Operation = `
query GetArticle ($id: ID!) {
	article(id: $id) {
		id
		feedId
		url
		title
		imageUrl
		content
		preview
		description
		guid
		authorName
		authorEmail
		updatedAt
	}
}
`

func GetArticle(
	ctx_ context.Context,
	client_ graphql.Client,
	id string,
) (*GetArticleResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetArticle",
		Query:  GetArticle_Operation,
		Variables: &__GetArticleInput{
			Id: id,
		},
	}
	var err_ error

	var data_ GetArticleResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetFeed.
const GetFeed_Operation = `
query GetFeed ($id: ID!) {
	feed(id: $id) {
		id
		url
		name
	}
}
`

func GetFeed(
	ctx_ context.Context,
	client_ graphql.Client,
	id string,
) (*GetFeedResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetFeed",
		Query:  GetFeed_Operation,
		Variables: &__GetFeedInput{
			Id: id,
		},
	}
	var err_ error

	var data_ GetFeedResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by ListArticles.
const ListArticles_Operation = `
query ListArticles ($feedId: ID!) {
	articles(feedId: $feedId) {
		articles {
			id
			feedId
			url
			title
			imageUrl
			preview
			authorName
			authorEmail
			updatedAt
		}
		cursor {
			next
			previous
		}
	}
}
`

func ListArticles(
	ctx_ context.Context,
	client_ graphql.Client,
	feedId string,
) (*ListArticlesResponse, error) {
	req_ := &graphql.Request{
		OpName: "ListArticles",
		Query:  ListArticles_Operation,
		Variables: &__ListArticlesInput{
			FeedId: feedId,
		},
	}
	var err_ error

	var data_ ListArticlesResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by ListFeeds.
const ListFeeds_Operation = `
query ListFeeds {
	feeds {
		feeds {
			id
			url
			name
			createdAt
			viewedAt
			unreadStartAt
			unreadCount
		}
	}
}
`

func ListFeeds(
	ctx_ context.Context,
	client_ graphql.Client,
) (*ListFeedsResponse, error) {
	req_ := &graphql.Request{
		OpName: "ListFeeds",
		Query:  ListFeeds_Operation,
	}
	var err_ error

	var data_ ListFeedsResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by UpdateFeed.
const UpdateFeed_Operation = `
mutation UpdateFeed ($id: ID!, $url: String, $name: String) {
	updateFeed(id: $id, url: $url, name: $name) {
		id
	}
}
`

func UpdateFeed(
	ctx_ context.Context,
	client_ graphql.Client,
	id string,
	url string,
	name string,
) (*UpdateFeedResponse, error) {
	req_ := &graphql.Request{
		OpName: "UpdateFeed",
		Query:  UpdateFeed_Operation,
		Variables: &__UpdateFeedInput{
			Id:   id,
			Url:  url,
			Name: name,
		},
	}
	var err_ error

	var data_ UpdateFeedResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

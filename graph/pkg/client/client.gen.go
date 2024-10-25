// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package client

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// FeedsFeedsFeed includes the requested fields of the GraphQL type Feed.
type FeedsFeedsFeed struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

// GetId returns FeedsFeedsFeed.Id, and is useful for accessing the field via an interface.
func (v *FeedsFeedsFeed) GetId() string { return v.Id }

// GetUrl returns FeedsFeedsFeed.Url, and is useful for accessing the field via an interface.
func (v *FeedsFeedsFeed) GetUrl() string { return v.Url }

// GetName returns FeedsFeedsFeed.Name, and is useful for accessing the field via an interface.
func (v *FeedsFeedsFeed) GetName() string { return v.Name }

// FeedsResponse is returned by Feeds on success.
type FeedsResponse struct {
	Feeds []FeedsFeedsFeed `json:"feeds"`
}

// GetFeeds returns FeedsResponse.Feeds, and is useful for accessing the field via an interface.
func (v *FeedsResponse) GetFeeds() []FeedsFeedsFeed { return v.Feeds }

// The query or mutation executed by Feeds.
const Feeds_Operation = `
query Feeds {
	feeds {
		id
		url
		name
	}
}
`

func Feeds(
	ctx_ context.Context,
	client_ graphql.Client,
) (*FeedsResponse, error) {
	req_ := &graphql.Request{
		OpName: "Feeds",
		Query:  Feeds_Operation,
	}
	var err_ error

	var data_ FeedsResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
// contains rest api related resolvers
package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ericbutera/amalgam/graph/graph/model"
	"github.com/samber/lo"
)

// Fetch feeds from the REST API
func (r *queryResolver) Feeds(ctx context.Context) ([]*model.Feed, error) {
	var feeds []*model.Feed

	apiClient := NewApiClient(r.config.ApiScheme, r.config.ApiHost)
	res, _, err := apiClient.DefaultAPI.FeedsGet(ctx).Execute()
	if err != nil {
		return nil, err
	}

	for _, feed := range res.Feeds {
		// TODO: mapstructure!
		feeds = append(feeds, &model.Feed{
			ID:   fmt.Sprintf("%d", feed.Id),
			URL:  feed.Url,
			Name: lo.FromPtr(feed.Name),
		})
	}

	return feeds, nil
}

// Fetch a single feed by ID
func (r *queryResolver) Feed(ctx context.Context, id string) (*model.Feed, error) {
	resp, err := http.Get(r.config.ApiBaseUrl + "/feeds/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var feed model.Feed
	if err := json.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Fetch articles for a specific feed
func (r *queryResolver) Articles(ctx context.Context, feedId string) ([]*model.Article, error) {
	resp, err := http.Get(r.config.ApiBaseUrl + "/feeds/" + feedId + "/articles")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var articles []*model.Article
	if err := json.NewDecoder(resp.Body).Decode(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	resp, err := http.Get(r.config.ApiBaseUrl + "/articles/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var article model.Article
	if err := json.NewDecoder(resp.Body).Decode(&article); err != nil {
		return nil, err
	}

	return &article, nil
}

// Add a new feed by sending a POST request to the REST API
func (r *mutationResolver) AddFeed(ctx context.Context, url string, name string) (*model.Feed, error) {
	feed := &model.Feed{URL: url, Name: name}
	payload, err := json.Marshal(map[string]interface{}{
		"feed": feed,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(r.config.ApiBaseUrl+"/feeds", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, err
	}

	return feed, nil
}

// Update a feed by sending a PUT request
func (r *mutationResolver) UpdateFeed(ctx context.Context, id string, url *string, name *string) (*model.Feed, error) {
	update := map[string]interface{}{
		"url":  url,
		"name": name,
	}

	payload, err := json.Marshal(map[string]interface{}{
		"feed": update,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.config.ApiBaseUrl+"/feeds/"+id, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedFeed model.Feed
	if err := json.NewDecoder(resp.Body).Decode(&updatedFeed); err != nil {
		return nil, err
	}

	return &updatedFeed, nil
}

// Delete a feed by sending a DELETE request
func (r *mutationResolver) DeleteFeed(ctx context.Context, id string) (*model.Feed, error) {
	req, err := http.NewRequest(http.MethodDelete, r.config.ApiBaseUrl+"/feeds/"+id, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deletedFeed model.Feed
	if err := json.NewDecoder(resp.Body).Decode(&deletedFeed); err != nil {
		return nil, err
	}

	return &deletedFeed, nil
}

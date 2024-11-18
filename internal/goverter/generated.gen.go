// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package goverter

import (
	model "github.com/ericbutera/amalgam/graph/graph/model"
	models "github.com/ericbutera/amalgam/internal/db/models"
	models1 "github.com/ericbutera/amalgam/internal/service/models"
	graphql "github.com/ericbutera/amalgam/pkg/clients/graphql"
	v1 "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertBase(source struct {
	ID string
}) *models.Base {
	var modelsBase models.Base
	modelsBase.ID = source.ID
	return &modelsBase
}
func (c *ConverterImpl) DbToServiceArticle(source *models.Article) *models1.Article {
	var pModelsArticle *models1.Article
	if source != nil {
		var modelsArticle models1.Article
		modelsArticle.ID = (*source).Base.ID
		modelsArticle.FeedID = (*source).FeedID
		modelsArticle.URL = (*source).URL
		modelsArticle.Title = (*source).Title
		modelsArticle.ImageURL = (*source).ImageURL
		modelsArticle.Preview = (*source).Preview
		modelsArticle.Content = (*source).Content
		modelsArticle.GUID = (*source).GUID
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) DbToServiceFeed(source *models.Feed) *models1.Feed {
	var pModelsFeed *models1.Feed
	if source != nil {
		var modelsFeed models1.Feed
		modelsFeed.ID = (*source).Base.ID
		modelsFeed.Name = (*source).Name
		modelsFeed.URL = (*source).URL
		modelsFeed.IsActive = (*source).IsActive
		pModelsFeed = &modelsFeed
	}
	return pModelsFeed
}
func (c *ConverterImpl) GraphClientToApiArticle(source *graphql.GetArticleArticle) *models1.Article {
	var pModelsArticle *models1.Article
	if source != nil {
		var modelsArticle models1.Article
		modelsArticle.ID = (*source).Id
		modelsArticle.FeedID = (*source).FeedId
		modelsArticle.URL = (*source).Url
		modelsArticle.Title = (*source).Title
		modelsArticle.ImageURL = (*source).ImageUrl
		modelsArticle.Preview = (*source).Preview
		modelsArticle.Content = (*source).Content
		modelsArticle.GUID = (*source).Guid
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) GraphClientToApiArticleList(source *graphql.ListArticlesArticlesArticle) *models1.Article {
	var pModelsArticle *models1.Article
	if source != nil {
		var modelsArticle models1.Article
		modelsArticle.ID = (*source).Id
		modelsArticle.FeedID = (*source).FeedId
		modelsArticle.URL = (*source).Url
		modelsArticle.Title = (*source).Title
		modelsArticle.ImageURL = (*source).ImageUrl
		modelsArticle.Preview = (*source).Preview
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) GraphClientToApiFeedGet(source *graphql.GetFeedFeed) *models1.Feed {
	var pModelsFeed *models1.Feed
	if source != nil {
		var modelsFeed models1.Feed
		modelsFeed.ID = (*source).Id
		modelsFeed.Name = (*source).Name
		modelsFeed.URL = (*source).Url
		pModelsFeed = &modelsFeed
	}
	return pModelsFeed
}
func (c *ConverterImpl) GraphToServiceArticle(source *model.Article) *models1.Article {
	var pModelsArticle *models1.Article
	if source != nil {
		var modelsArticle models1.Article
		modelsArticle.ID = (*source).ID
		modelsArticle.FeedID = (*source).FeedID
		modelsArticle.URL = (*source).URL
		modelsArticle.Title = (*source).Title
		if (*source).ImageURL != nil {
			modelsArticle.ImageURL = *(*source).ImageURL
		}
		modelsArticle.Preview = (*source).Preview
		modelsArticle.Content = (*source).Content
		if (*source).GUID != nil {
			modelsArticle.GUID = *(*source).GUID
		}
		if (*source).AuthorName != nil {
			modelsArticle.AuthorName = *(*source).AuthorName
		}
		if (*source).AuthorEmail != nil {
			modelsArticle.AuthorEmail = *(*source).AuthorEmail
		}
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) GraphToServiceFeed(source *model.Feed) *models1.Feed {
	var pModelsFeed *models1.Feed
	if source != nil {
		var modelsFeed models1.Feed
		modelsFeed.ID = (*source).ID
		modelsFeed.Name = (*source).Name
		modelsFeed.URL = (*source).URL
		pModelsFeed = &modelsFeed
	}
	return pModelsFeed
}
func (c *ConverterImpl) ProtoCreateFeedToService(source *v1.CreateFeedRequest_Feed) *models1.Feed {
	var pModelsFeed *models1.Feed
	if source != nil {
		var modelsFeed models1.Feed
		modelsFeed.Name = (*source).Name
		modelsFeed.URL = (*source).Url
		pModelsFeed = &modelsFeed
	}
	return pModelsFeed
}
func (c *ConverterImpl) ProtoToServiceArticle(source *v1.Article) *models1.Article {
	var pModelsArticle *models1.Article
	if source != nil {
		var modelsArticle models1.Article
		modelsArticle.ID = (*source).Id
		modelsArticle.FeedID = (*source).FeedId
		modelsArticle.URL = (*source).Url
		modelsArticle.Title = (*source).Title
		modelsArticle.ImageURL = (*source).ImageUrl
		modelsArticle.Preview = (*source).Preview
		modelsArticle.Content = (*source).Content
		modelsArticle.GUID = (*source).Guid
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) ProtoUpdateFeedToService(source *v1.UpdateFeedRequest_Feed) *models1.Feed {
	var pModelsFeed *models1.Feed
	if source != nil {
		var modelsFeed models1.Feed
		modelsFeed.ID = (*source).Id
		modelsFeed.Name = (*source).Name
		modelsFeed.URL = (*source).Url
		pModelsFeed = &modelsFeed
	}
	return pModelsFeed
}
func (c *ConverterImpl) ServiceToDbArticle(source *models1.Article) *models.Article {
	var pModelsArticle *models.Article
	if source != nil {
		var modelsArticle models.Article
		modelsArticle.FeedID = (*source).FeedID
		modelsArticle.URL = (*source).URL
		modelsArticle.Title = (*source).Title
		modelsArticle.ImageURL = (*source).ImageURL
		modelsArticle.Preview = (*source).Preview
		modelsArticle.Content = (*source).Content
		modelsArticle.GUID = (*source).GUID
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) ServiceToDbFeed(source *models1.Feed) *models.Feed {
	var pModelsFeed *models.Feed
	if source != nil {
		var modelsFeed models.Feed
		modelsFeed.URL = (*source).URL
		modelsFeed.Name = (*source).Name
		modelsFeed.IsActive = (*source).IsActive
		pModelsFeed = &modelsFeed
	}
	return pModelsFeed
}
func (c *ConverterImpl) ServiceToGraphArticle(source *models1.Article) *model.Article {
	var pModelArticle *model.Article
	if source != nil {
		var modelArticle model.Article
		modelArticle.ID = (*source).ID
		modelArticle.FeedID = (*source).FeedID
		modelArticle.URL = (*source).URL
		modelArticle.Title = (*source).Title
		pString := (*source).ImageURL
		modelArticle.ImageURL = &pString
		modelArticle.Content = (*source).Content
		modelArticle.Preview = (*source).Preview
		pString2 := (*source).GUID
		modelArticle.GUID = &pString2
		pString3 := (*source).AuthorName
		modelArticle.AuthorName = &pString3
		pString4 := (*source).AuthorEmail
		modelArticle.AuthorEmail = &pString4
		pModelArticle = &modelArticle
	}
	return pModelArticle
}
func (c *ConverterImpl) ServiceToGraphFeed(source *models1.Feed) *model.Feed {
	var pModelFeed *model.Feed
	if source != nil {
		var modelFeed model.Feed
		modelFeed.ID = (*source).ID
		modelFeed.URL = (*source).URL
		modelFeed.Name = (*source).Name
		pModelFeed = &modelFeed
	}
	return pModelFeed
}
func (c *ConverterImpl) ServiceToProtoArticle(source *models1.Article) *v1.Article {
	var pFeedsArticle *v1.Article
	if source != nil {
		var feedsArticle v1.Article
		feedsArticle.Id = (*source).ID
		feedsArticle.Title = (*source).Title
		feedsArticle.Content = (*source).Content
		feedsArticle.FeedId = (*source).FeedID
		feedsArticle.Url = (*source).URL
		feedsArticle.ImageUrl = (*source).ImageURL
		feedsArticle.Preview = (*source).Preview
		feedsArticle.Guid = (*source).GUID
		feedsArticle.AuthorName = (*source).AuthorName
		feedsArticle.AuthorEmail = (*source).AuthorEmail
		pFeedsArticle = &feedsArticle
	}
	return pFeedsArticle
}
func (c *ConverterImpl) ServiceToProtoFeed(source *models1.Feed) *v1.Feed {
	var pFeedsFeed *v1.Feed
	if source != nil {
		var feedsFeed v1.Feed
		feedsFeed.Id = (*source).ID
		feedsFeed.Url = (*source).URL
		feedsFeed.Name = (*source).Name
		pFeedsFeed = &feedsFeed
	}
	return pFeedsFeed
}

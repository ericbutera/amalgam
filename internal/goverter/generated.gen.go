// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package goverter

import (
	models "github.com/ericbutera/amalgam/internal/db/models"
	models1 "github.com/ericbutera/amalgam/internal/service/models"
	graphql "github.com/ericbutera/amalgam/pkg/clients/graphql"
	v1 "github.com/ericbutera/amalgam/pkg/feeds/v1"
	model "github.com/ericbutera/amalgam/services/graph/graph/model"
	"time"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertBase(source struct {
	ID        string
	UpdatedAt time.Time
}) *models.Base {
	var modelsBase models.Base
	modelsBase.ID = source.ID
	modelsBase.UpdatedAt = Time(source.UpdatedAt)
	return &modelsBase
}
func (c *ConverterImpl) ConvertListCursor(source *model.ListCursor) *v1.Cursor {
	var pFeedsCursor *v1.Cursor
	if source != nil {
		var feedsCursor v1.Cursor
		if (*source).Previous != nil {
			feedsCursor.Previous = *(*source).Previous
		}
		if (*source).Next != nil {
			feedsCursor.Next = *(*source).Next
		}
		pFeedsCursor = &feedsCursor
	}
	return pFeedsCursor
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
		modelsArticle.Description = (*source).Description
		modelsArticle.GUID = (*source).GUID
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		modelsArticle.UpdatedAt = Time((*source).Base.UpdatedAt)
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
		modelsArticle.Description = (*source).Description
		modelsArticle.GUID = (*source).Guid
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		modelsArticle.UpdatedAt = Time((*source).UpdatedAt)
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) GraphClientToApiArticleList(source *graphql.ListArticlesArticlesArticlesResponseArticlesArticle) *models1.Article {
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
		modelsArticle.UpdatedAt = Time((*source).UpdatedAt)
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
func (c *ConverterImpl) GraphToProtoListOptions(source *model.ListOptions) *v1.ListOptions {
	var pFeedsListOptions *v1.ListOptions
	if source != nil {
		var feedsListOptions v1.ListOptions
		feedsListOptions.Cursor = c.ConvertListCursor((*source).Cursor)
		feedsListOptions.Limit = IntPtrToInt32((*source).Limit)
		pFeedsListOptions = &feedsListOptions
	}
	return pFeedsListOptions
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
func (c *ConverterImpl) ProtoToGraphArticle(source *v1.Article) *model.Article {
	var pModelArticle *model.Article
	if source != nil {
		var modelArticle model.Article
		modelArticle.ID = (*source).Id
		modelArticle.FeedID = (*source).FeedId
		modelArticle.URL = (*source).Url
		modelArticle.Title = (*source).Title
		pString := (*source).ImageUrl
		modelArticle.ImageURL = &pString
		modelArticle.Content = (*source).Content
		modelArticle.Description = (*source).Description
		modelArticle.Preview = (*source).Preview
		pString2 := (*source).Guid
		modelArticle.GUID = &pString2
		pString3 := (*source).AuthorName
		modelArticle.AuthorName = &pString3
		pString4 := (*source).AuthorEmail
		modelArticle.AuthorEmail = &pString4
		modelArticle.UpdatedAt = ProtoTimestampToTime((*source).UpdatedAt)
		pModelArticle = &modelArticle
	}
	return pModelArticle
}
func (c *ConverterImpl) ProtoToGraphCursor(source *v1.Cursor) *model.ResponseCursor {
	var pModelResponseCursor *model.ResponseCursor
	if source != nil {
		var modelResponseCursor model.ResponseCursor
		modelResponseCursor.Previous = (*source).Previous
		modelResponseCursor.Next = (*source).Next
		pModelResponseCursor = &modelResponseCursor
	}
	return pModelResponseCursor
}
func (c *ConverterImpl) ProtoToGraphUserArticle(source *v1.GetUserArticlesResponse_UserArticle) *model.UserArticle {
	var pModelUserArticle *model.UserArticle
	if source != nil {
		var modelUserArticle model.UserArticle
		modelUserArticle.ViewedAt = ProtoTimestampToTime((*source).ViewedAt)
		pModelUserArticle = &modelUserArticle
	}
	return pModelUserArticle
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
		modelsArticle.Description = (*source).Description
		modelsArticle.GUID = (*source).Guid
		modelsArticle.AuthorName = (*source).AuthorName
		modelsArticle.AuthorEmail = (*source).AuthorEmail
		modelsArticle.UpdatedAt = ProtoTimestampToTime((*source).UpdatedAt)
		pModelsArticle = &modelsArticle
	}
	return pModelsArticle
}
func (c *ConverterImpl) ProtoToServiceFeed(source *v1.Feed) *models1.Feed {
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
func (c *ConverterImpl) ProtoToServiceFeedVerification(source *v1.FeedVerification) *models1.FeedVerification {
	var pModelsFeedVerification *models1.FeedVerification
	if source != nil {
		var modelsFeedVerification models1.FeedVerification
		modelsFeedVerification.ID = (*source).Id
		modelsFeedVerification.URL = (*source).Url
		modelsFeedVerification.UserID = (*source).UserId
		modelsFeedVerification.WorkflowID = (*source).WorkflowId
		modelsFeedVerification.CreatedAt = ProtoTimestampToTime((*source).CreatedAt)
		modelsFeedVerification.UpdatedAt = ProtoTimestampToTime((*source).UpdatedAt)
		pModelsFeedVerification = &modelsFeedVerification
	}
	return pModelsFeedVerification
}
func (c *ConverterImpl) ProtoToServiceFetchHistory(source *v1.FetchHistory) *models1.FetchHistory {
	var pModelsFetchHistory *models1.FetchHistory
	if source != nil {
		var modelsFetchHistory models1.FetchHistory
		modelsFetchHistory.ID = (*source).Id
		modelsFetchHistory.FeedID = (*source).FeedId
		modelsFetchHistory.FeedVerificationID = (*source).FeedVerificationId
		modelsFetchHistory.ResponseCode = (*source).ResponseCode
		modelsFetchHistory.ETag = (*source).Etag
		modelsFetchHistory.WorkflowID = (*source).WorkflowId
		modelsFetchHistory.Bucket = (*source).Bucket
		modelsFetchHistory.Message = (*source).Message
		modelsFetchHistory.CreatedAt = ProtoTimestampToTime((*source).CreatedAt)
		pModelsFetchHistory = &modelsFetchHistory
	}
	return pModelsFetchHistory
}
func (c *ConverterImpl) ProtoToServiceUserFeed(source *v1.UserFeed) *models1.UserFeed {
	var pModelsUserFeed *models1.UserFeed
	if source != nil {
		var modelsUserFeed models1.UserFeed
		modelsUserFeed.FeedID = (*source).FeedId
		modelsUserFeed.Name = (*source).Name
		modelsUserFeed.URL = (*source).Url
		modelsUserFeed.CreatedAt = ProtoTimestampToTime((*source).CreatedAt)
		modelsUserFeed.ViewedAt = ProtoTimestampToTime((*source).ViewedAt)
		modelsUserFeed.UnreadStartAt = ProtoTimestampToTime((*source).UnreadStartAt)
		modelsUserFeed.UnreadCount = (*source).UnreadCount
		pModelsUserFeed = &modelsUserFeed
	}
	return pModelsUserFeed
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
func (c *ConverterImpl) ProtoUserFeedToGraphUserFeed(source *v1.UserFeed) *model.Feed {
	var pModelFeed *model.Feed
	if source != nil {
		var modelFeed model.Feed
		modelFeed.ID = (*source).FeedId
		modelFeed.URL = (*source).Url
		modelFeed.Name = (*source).Name
		modelFeed.CreatedAt = ProtoTimestampToTime((*source).CreatedAt)
		modelFeed.ViewedAt = ProtoTimestampToTime((*source).ViewedAt)
		modelFeed.UnreadStartAt = ProtoTimestampToTime((*source).UnreadStartAt)
		modelFeed.UnreadCount = Int32ToInt((*source).UnreadCount)
		pModelFeed = &modelFeed
	}
	return pModelFeed
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
		modelsArticle.Description = (*source).Description
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
		modelArticle.Description = (*source).Description
		modelArticle.Preview = (*source).Preview
		pString2 := (*source).GUID
		modelArticle.GUID = &pString2
		pString3 := (*source).AuthorName
		modelArticle.AuthorName = &pString3
		pString4 := (*source).AuthorEmail
		modelArticle.AuthorEmail = &pString4
		modelArticle.UpdatedAt = Time((*source).UpdatedAt)
		pModelArticle = &modelArticle
	}
	return pModelArticle
}
func (c *ConverterImpl) ServiceToGraphFeed(source *models1.UserFeed) *model.Feed {
	var pModelFeed *model.Feed
	if source != nil {
		var modelFeed model.Feed
		modelFeed.ID = (*source).FeedID
		modelFeed.URL = (*source).URL
		modelFeed.Name = (*source).Name
		modelFeed.CreatedAt = Time((*source).CreatedAt)
		modelFeed.ViewedAt = Time((*source).ViewedAt)
		modelFeed.UnreadStartAt = Time((*source).UnreadStartAt)
		modelFeed.UnreadCount = Int32ToInt((*source).UnreadCount)
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
		feedsArticle.Description = (*source).Description
		feedsArticle.UpdatedAt = TimeToProtoTimestamp((*source).UpdatedAt)
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
func (c *ConverterImpl) ServiceToProtoFeedVerification(source *models1.FeedVerification) *v1.FeedVerification {
	var pFeedsFeedVerification *v1.FeedVerification
	if source != nil {
		var feedsFeedVerification v1.FeedVerification
		feedsFeedVerification.Id = (*source).ID
		feedsFeedVerification.Url = (*source).URL
		feedsFeedVerification.UserId = (*source).UserID
		feedsFeedVerification.WorkflowId = (*source).WorkflowID
		feedsFeedVerification.CreatedAt = TimeToProtoTimestamp((*source).CreatedAt)
		feedsFeedVerification.UpdatedAt = TimeToProtoTimestamp((*source).UpdatedAt)
		pFeedsFeedVerification = &feedsFeedVerification
	}
	return pFeedsFeedVerification
}
func (c *ConverterImpl) ServiceToProtoFetchHistory(source *models1.FetchHistory) *v1.FetchHistory {
	var pFeedsFetchHistory *v1.FetchHistory
	if source != nil {
		var feedsFetchHistory v1.FetchHistory
		feedsFetchHistory.Id = (*source).ID
		feedsFetchHistory.FeedId = (*source).FeedID
		feedsFetchHistory.FeedVerificationId = (*source).FeedVerificationID
		feedsFetchHistory.ResponseCode = (*source).ResponseCode
		feedsFetchHistory.Etag = (*source).ETag
		feedsFetchHistory.WorkflowId = (*source).WorkflowID
		feedsFetchHistory.Bucket = (*source).Bucket
		feedsFetchHistory.Message = (*source).Message
		feedsFetchHistory.CreatedAt = TimeToProtoTimestamp((*source).CreatedAt)
		pFeedsFetchHistory = &feedsFetchHistory
	}
	return pFeedsFetchHistory
}
func (c *ConverterImpl) ServiceToProtoUserArticle(source *models1.UserArticle) *v1.GetUserArticlesResponse_UserArticle {
	var pFeedsGetUserArticlesResponse_UserArticle *v1.GetUserArticlesResponse_UserArticle
	if source != nil {
		var feedsGetUserArticlesResponse_UserArticle v1.GetUserArticlesResponse_UserArticle
		feedsGetUserArticlesResponse_UserArticle.ViewedAt = NillableTimeToProtoTimestamp((*source).ViewedAt)
		pFeedsGetUserArticlesResponse_UserArticle = &feedsGetUserArticlesResponse_UserArticle
	}
	return pFeedsGetUserArticlesResponse_UserArticle
}
func (c *ConverterImpl) ServiceToProtoUserFeed(source *models1.UserFeed) *v1.UserFeed {
	var pFeedsUserFeed *v1.UserFeed
	if source != nil {
		var feedsUserFeed v1.UserFeed
		feedsUserFeed.FeedId = (*source).FeedID
		feedsUserFeed.Url = (*source).URL
		feedsUserFeed.Name = (*source).Name
		feedsUserFeed.UnreadCount = (*source).UnreadCount
		feedsUserFeed.CreatedAt = TimeToProtoTimestamp((*source).CreatedAt)
		feedsUserFeed.ViewedAt = TimeToProtoTimestamp((*source).ViewedAt)
		feedsUserFeed.UnreadStartAt = TimeToProtoTimestamp((*source).UnreadStartAt)
		pFeedsUserFeed = &feedsUserFeed
	}
	return pFeedsUserFeed
}

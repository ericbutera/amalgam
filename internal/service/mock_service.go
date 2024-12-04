// Code generated by mockery v2.47.0. DO NOT EDIT.

package service

import (
	context "context"

	models "github.com/ericbutera/amalgam/internal/service/models"
	mock "github.com/stretchr/testify/mock"

	pagination "github.com/ericbutera/amalgam/internal/db/pagination"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// CreateFeed provides a mock function with given fields: ctx, feed
func (_m *MockService) CreateFeed(ctx context.Context, feed *models.Feed) (CreateFeedResult, error) {
	ret := _m.Called(ctx, feed)

	if len(ret) == 0 {
		panic("no return value specified for CreateFeed")
	}

	var r0 CreateFeedResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Feed) (CreateFeedResult, error)); ok {
		return rf(ctx, feed)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Feed) CreateFeedResult); ok {
		r0 = rf(ctx, feed)
	} else {
		r0 = ret.Get(0).(CreateFeedResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Feed) error); ok {
		r1 = rf(ctx, feed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_CreateFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFeed'
type MockService_CreateFeed_Call struct {
	*mock.Call
}

// CreateFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - feed *models.Feed
func (_e *MockService_Expecter) CreateFeed(ctx interface{}, feed interface{}) *MockService_CreateFeed_Call {
	return &MockService_CreateFeed_Call{Call: _e.mock.On("CreateFeed", ctx, feed)}
}

func (_c *MockService_CreateFeed_Call) Run(run func(ctx context.Context, feed *models.Feed)) *MockService_CreateFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Feed))
	})
	return _c
}

func (_c *MockService_CreateFeed_Call) Return(_a0 CreateFeedResult, _a1 error) *MockService_CreateFeed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_CreateFeed_Call) RunAndReturn(run func(context.Context, *models.Feed) (CreateFeedResult, error)) *MockService_CreateFeed_Call {
	_c.Call.Return(run)
	return _c
}

// Feeds provides a mock function with given fields: ctx
func (_m *MockService) Feeds(ctx context.Context) ([]models.Feed, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Feeds")
	}

	var r0 []models.Feed
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Feed, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Feed); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Feed)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_Feeds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Feeds'
type MockService_Feeds_Call struct {
	*mock.Call
}

// Feeds is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) Feeds(ctx interface{}) *MockService_Feeds_Call {
	return &MockService_Feeds_Call{Call: _e.mock.On("Feeds", ctx)}
}

func (_c *MockService_Feeds_Call) Run(run func(ctx context.Context)) *MockService_Feeds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_Feeds_Call) Return(_a0 []models.Feed, _a1 error) *MockService_Feeds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_Feeds_Call) RunAndReturn(run func(context.Context) ([]models.Feed, error)) *MockService_Feeds_Call {
	_c.Call.Return(run)
	return _c
}

// GetArticle provides a mock function with given fields: ctx, id
func (_m *MockService) GetArticle(ctx context.Context, id string) (*models.Article, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetArticle")
	}

	var r0 *models.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Article, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Article); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArticle'
type MockService_GetArticle_Call struct {
	*mock.Call
}

// GetArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockService_Expecter) GetArticle(ctx interface{}, id interface{}) *MockService_GetArticle_Call {
	return &MockService_GetArticle_Call{Call: _e.mock.On("GetArticle", ctx, id)}
}

func (_c *MockService_GetArticle_Call) Run(run func(ctx context.Context, id string)) *MockService_GetArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_GetArticle_Call) Return(_a0 *models.Article, _a1 error) *MockService_GetArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetArticle_Call) RunAndReturn(run func(context.Context, string) (*models.Article, error)) *MockService_GetArticle_Call {
	_c.Call.Return(run)
	return _c
}

// GetArticlesByFeed provides a mock function with given fields: ctx, feedId, options
func (_m *MockService) GetArticlesByFeed(ctx context.Context, feedId string, options pagination.ListOptions) (*ArticlesByFeedResult, error) {
	ret := _m.Called(ctx, feedId, options)

	if len(ret) == 0 {
		panic("no return value specified for GetArticlesByFeed")
	}

	var r0 *ArticlesByFeedResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, pagination.ListOptions) (*ArticlesByFeedResult, error)); ok {
		return rf(ctx, feedId, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, pagination.ListOptions) *ArticlesByFeedResult); ok {
		r0 = rf(ctx, feedId, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ArticlesByFeedResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, pagination.ListOptions) error); ok {
		r1 = rf(ctx, feedId, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetArticlesByFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArticlesByFeed'
type MockService_GetArticlesByFeed_Call struct {
	*mock.Call
}

// GetArticlesByFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - feedId string
//   - options pagination.ListOptions
func (_e *MockService_Expecter) GetArticlesByFeed(ctx interface{}, feedId interface{}, options interface{}) *MockService_GetArticlesByFeed_Call {
	return &MockService_GetArticlesByFeed_Call{Call: _e.mock.On("GetArticlesByFeed", ctx, feedId, options)}
}

func (_c *MockService_GetArticlesByFeed_Call) Run(run func(ctx context.Context, feedId string, options pagination.ListOptions)) *MockService_GetArticlesByFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(pagination.ListOptions))
	})
	return _c
}

func (_c *MockService_GetArticlesByFeed_Call) Return(_a0 *ArticlesByFeedResult, _a1 error) *MockService_GetArticlesByFeed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetArticlesByFeed_Call) RunAndReturn(run func(context.Context, string, pagination.ListOptions) (*ArticlesByFeedResult, error)) *MockService_GetArticlesByFeed_Call {
	_c.Call.Return(run)
	return _c
}

// GetFeed provides a mock function with given fields: ctx, id
func (_m *MockService) GetFeed(ctx context.Context, id string) (*models.Feed, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFeed")
	}

	var r0 *models.Feed
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Feed, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Feed); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Feed)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFeed'
type MockService_GetFeed_Call struct {
	*mock.Call
}

// GetFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockService_Expecter) GetFeed(ctx interface{}, id interface{}) *MockService_GetFeed_Call {
	return &MockService_GetFeed_Call{Call: _e.mock.On("GetFeed", ctx, id)}
}

func (_c *MockService_GetFeed_Call) Run(run func(ctx context.Context, id string)) *MockService_GetFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_GetFeed_Call) Return(_a0 *models.Feed, _a1 error) *MockService_GetFeed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetFeed_Call) RunAndReturn(run func(context.Context, string) (*models.Feed, error)) *MockService_GetFeed_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserArticles provides a mock function with given fields: ctx, userID, articleIDs
func (_m *MockService) GetUserArticles(ctx context.Context, userID string, articleIDs []string) ([]*models.UserArticle, error) {
	ret := _m.Called(ctx, userID, articleIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetUserArticles")
	}

	var r0 []*models.UserArticle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]*models.UserArticle, error)); ok {
		return rf(ctx, userID, articleIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []*models.UserArticle); ok {
		r0 = rf(ctx, userID, articleIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserArticle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, userID, articleIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetUserArticles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserArticles'
type MockService_GetUserArticles_Call struct {
	*mock.Call
}

// GetUserArticles is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - articleIDs []string
func (_e *MockService_Expecter) GetUserArticles(ctx interface{}, userID interface{}, articleIDs interface{}) *MockService_GetUserArticles_Call {
	return &MockService_GetUserArticles_Call{Call: _e.mock.On("GetUserArticles", ctx, userID, articleIDs)}
}

func (_c *MockService_GetUserArticles_Call) Run(run func(ctx context.Context, userID string, articleIDs []string)) *MockService_GetUserArticles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockService_GetUserArticles_Call) Return(_a0 []*models.UserArticle, _a1 error) *MockService_GetUserArticles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetUserArticles_Call) RunAndReturn(run func(context.Context, string, []string) ([]*models.UserArticle, error)) *MockService_GetUserArticles_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserFeed provides a mock function with given fields: ctx, userID, feedID
func (_m *MockService) GetUserFeed(ctx context.Context, userID string, feedID string) (*models.UserFeed, error) {
	ret := _m.Called(ctx, userID, feedID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserFeed")
	}

	var r0 *models.UserFeed
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.UserFeed, error)); ok {
		return rf(ctx, userID, feedID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.UserFeed); ok {
		r0 = rf(ctx, userID, feedID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserFeed)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, feedID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetUserFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserFeed'
type MockService_GetUserFeed_Call struct {
	*mock.Call
}

// GetUserFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - feedID string
func (_e *MockService_Expecter) GetUserFeed(ctx interface{}, userID interface{}, feedID interface{}) *MockService_GetUserFeed_Call {
	return &MockService_GetUserFeed_Call{Call: _e.mock.On("GetUserFeed", ctx, userID, feedID)}
}

func (_c *MockService_GetUserFeed_Call) Run(run func(ctx context.Context, userID string, feedID string)) *MockService_GetUserFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockService_GetUserFeed_Call) Return(_a0 *models.UserFeed, _a1 error) *MockService_GetUserFeed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetUserFeed_Call) RunAndReturn(run func(context.Context, string, string) (*models.UserFeed, error)) *MockService_GetUserFeed_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserFeeds provides a mock function with given fields: ctx, userID
func (_m *MockService) GetUserFeeds(ctx context.Context, userID string) (*GetUserFeedsResult, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserFeeds")
	}

	var r0 *GetUserFeedsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*GetUserFeedsResult, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *GetUserFeedsResult); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetUserFeedsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetUserFeeds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserFeeds'
type MockService_GetUserFeeds_Call struct {
	*mock.Call
}

// GetUserFeeds is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *MockService_Expecter) GetUserFeeds(ctx interface{}, userID interface{}) *MockService_GetUserFeeds_Call {
	return &MockService_GetUserFeeds_Call{Call: _e.mock.On("GetUserFeeds", ctx, userID)}
}

func (_c *MockService_GetUserFeeds_Call) Run(run func(ctx context.Context, userID string)) *MockService_GetUserFeeds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_GetUserFeeds_Call) Return(_a0 *GetUserFeedsResult, _a1 error) *MockService_GetUserFeeds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetUserFeeds_Call) RunAndReturn(run func(context.Context, string) (*GetUserFeedsResult, error)) *MockService_GetUserFeeds_Call {
	_c.Call.Return(run)
	return _c
}

// SaveArticle provides a mock function with given fields: ctx, article
func (_m *MockService) SaveArticle(ctx context.Context, article *models.Article) (SaveArticleResult, error) {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for SaveArticle")
	}

	var r0 SaveArticleResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Article) (SaveArticleResult, error)); ok {
		return rf(ctx, article)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Article) SaveArticleResult); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Get(0).(SaveArticleResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Article) error); ok {
		r1 = rf(ctx, article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_SaveArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveArticle'
type MockService_SaveArticle_Call struct {
	*mock.Call
}

// SaveArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - article *models.Article
func (_e *MockService_Expecter) SaveArticle(ctx interface{}, article interface{}) *MockService_SaveArticle_Call {
	return &MockService_SaveArticle_Call{Call: _e.mock.On("SaveArticle", ctx, article)}
}

func (_c *MockService_SaveArticle_Call) Run(run func(ctx context.Context, article *models.Article)) *MockService_SaveArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Article))
	})
	return _c
}

func (_c *MockService_SaveArticle_Call) Return(_a0 SaveArticleResult, _a1 error) *MockService_SaveArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_SaveArticle_Call) RunAndReturn(run func(context.Context, *models.Article) (SaveArticleResult, error)) *MockService_SaveArticle_Call {
	_c.Call.Return(run)
	return _c
}

// SaveUserArticle provides a mock function with given fields: ctx, userArticle
func (_m *MockService) SaveUserArticle(ctx context.Context, userArticle *models.UserArticle) error {
	ret := _m.Called(ctx, userArticle)

	if len(ret) == 0 {
		panic("no return value specified for SaveUserArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.UserArticle) error); ok {
		r0 = rf(ctx, userArticle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_SaveUserArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveUserArticle'
type MockService_SaveUserArticle_Call struct {
	*mock.Call
}

// SaveUserArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - userArticle *models.UserArticle
func (_e *MockService_Expecter) SaveUserArticle(ctx interface{}, userArticle interface{}) *MockService_SaveUserArticle_Call {
	return &MockService_SaveUserArticle_Call{Call: _e.mock.On("SaveUserArticle", ctx, userArticle)}
}

func (_c *MockService_SaveUserArticle_Call) Run(run func(ctx context.Context, userArticle *models.UserArticle)) *MockService_SaveUserArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.UserArticle))
	})
	return _c
}

func (_c *MockService_SaveUserArticle_Call) Return(_a0 error) *MockService_SaveUserArticle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_SaveUserArticle_Call) RunAndReturn(run func(context.Context, *models.UserArticle) error) *MockService_SaveUserArticle_Call {
	_c.Call.Return(run)
	return _c
}

// SaveUserFeed provides a mock function with given fields: ctx, userFeed
func (_m *MockService) SaveUserFeed(ctx context.Context, userFeed *models.UserFeed) error {
	ret := _m.Called(ctx, userFeed)

	if len(ret) == 0 {
		panic("no return value specified for SaveUserFeed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.UserFeed) error); ok {
		r0 = rf(ctx, userFeed)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_SaveUserFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveUserFeed'
type MockService_SaveUserFeed_Call struct {
	*mock.Call
}

// SaveUserFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - userFeed *models.UserFeed
func (_e *MockService_Expecter) SaveUserFeed(ctx interface{}, userFeed interface{}) *MockService_SaveUserFeed_Call {
	return &MockService_SaveUserFeed_Call{Call: _e.mock.On("SaveUserFeed", ctx, userFeed)}
}

func (_c *MockService_SaveUserFeed_Call) Run(run func(ctx context.Context, userFeed *models.UserFeed)) *MockService_SaveUserFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.UserFeed))
	})
	return _c
}

func (_c *MockService_SaveUserFeed_Call) Return(_a0 error) *MockService_SaveUserFeed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_SaveUserFeed_Call) RunAndReturn(run func(context.Context, *models.UserFeed) error) *MockService_SaveUserFeed_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateFeed provides a mock function with given fields: ctx, id, feed
func (_m *MockService) UpdateFeed(ctx context.Context, id string, feed *models.Feed) error {
	ret := _m.Called(ctx, id, feed)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFeed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.Feed) error); ok {
		r0 = rf(ctx, id, feed)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_UpdateFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateFeed'
type MockService_UpdateFeed_Call struct {
	*mock.Call
}

// UpdateFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - feed *models.Feed
func (_e *MockService_Expecter) UpdateFeed(ctx interface{}, id interface{}, feed interface{}) *MockService_UpdateFeed_Call {
	return &MockService_UpdateFeed_Call{Call: _e.mock.On("UpdateFeed", ctx, id, feed)}
}

func (_c *MockService_UpdateFeed_Call) Run(run func(ctx context.Context, id string, feed *models.Feed)) *MockService_UpdateFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*models.Feed))
	})
	return _c
}

func (_c *MockService_UpdateFeed_Call) Return(_a0 error) *MockService_UpdateFeed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_UpdateFeed_Call) RunAndReturn(run func(context.Context, string, *models.Feed) error) *MockService_UpdateFeed_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateFeedArticleCount provides a mock function with given fields: ctx, feedID
func (_m *MockService) UpdateFeedArticleCount(ctx context.Context, feedID string) error {
	ret := _m.Called(ctx, feedID)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFeedArticleCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, feedID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_UpdateFeedArticleCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateFeedArticleCount'
type MockService_UpdateFeedArticleCount_Call struct {
	*mock.Call
}

// UpdateFeedArticleCount is a helper method to define mock.On call
//   - ctx context.Context
//   - feedID string
func (_e *MockService_Expecter) UpdateFeedArticleCount(ctx interface{}, feedID interface{}) *MockService_UpdateFeedArticleCount_Call {
	return &MockService_UpdateFeedArticleCount_Call{Call: _e.mock.On("UpdateFeedArticleCount", ctx, feedID)}
}

func (_c *MockService_UpdateFeedArticleCount_Call) Run(run func(ctx context.Context, feedID string)) *MockService_UpdateFeedArticleCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_UpdateFeedArticleCount_Call) Return(_a0 error) *MockService_UpdateFeedArticleCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_UpdateFeedArticleCount_Call) RunAndReturn(run func(context.Context, string) error) *MockService_UpdateFeedArticleCount_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

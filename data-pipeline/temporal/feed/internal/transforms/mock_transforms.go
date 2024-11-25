// Code generated by mockery v2.47.0. DO NOT EDIT.

package transforms

import (
	bytes "bytes"
	io "io"

	mock "github.com/stretchr/testify/mock"

	parse "github.com/ericbutera/amalgam/pkg/feed/parse"
)

// MockTransforms is an autogenerated mock type for the Transforms type
type MockTransforms struct {
	mock.Mock
}

type MockTransforms_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransforms) EXPECT() *MockTransforms_Expecter {
	return &MockTransforms_Expecter{mock: &_m.Mock}
}

// ArticleToJsonl provides a mock function with given fields: feedId, articles
func (_m *MockTransforms) ArticleToJsonl(feedId string, articles parse.Articles) (bytes.Buffer, []error) {
	ret := _m.Called(feedId, articles)

	if len(ret) == 0 {
		panic("no return value specified for ArticleToJsonl")
	}

	var r0 bytes.Buffer
	var r1 []error
	if rf, ok := ret.Get(0).(func(string, parse.Articles) (bytes.Buffer, []error)); ok {
		return rf(feedId, articles)
	}
	if rf, ok := ret.Get(0).(func(string, parse.Articles) bytes.Buffer); ok {
		r0 = rf(feedId, articles)
	} else {
		r0 = ret.Get(0).(bytes.Buffer)
	}

	if rf, ok := ret.Get(1).(func(string, parse.Articles) []error); ok {
		r1 = rf(feedId, articles)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// MockTransforms_ArticleToJsonl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ArticleToJsonl'
type MockTransforms_ArticleToJsonl_Call struct {
	*mock.Call
}

// ArticleToJsonl is a helper method to define mock.On call
//   - feedId string
//   - articles parse.Articles
func (_e *MockTransforms_Expecter) ArticleToJsonl(feedId interface{}, articles interface{}) *MockTransforms_ArticleToJsonl_Call {
	return &MockTransforms_ArticleToJsonl_Call{Call: _e.mock.On("ArticleToJsonl", feedId, articles)}
}

func (_c *MockTransforms_ArticleToJsonl_Call) Run(run func(feedId string, articles parse.Articles)) *MockTransforms_ArticleToJsonl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(parse.Articles))
	})
	return _c
}

func (_c *MockTransforms_ArticleToJsonl_Call) Return(_a0 bytes.Buffer, _a1 []error) *MockTransforms_ArticleToJsonl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransforms_ArticleToJsonl_Call) RunAndReturn(run func(string, parse.Articles) (bytes.Buffer, []error)) *MockTransforms_ArticleToJsonl_Call {
	_c.Call.Return(run)
	return _c
}

// RssToArticles provides a mock function with given fields: rss
func (_m *MockTransforms) RssToArticles(rss io.ReadCloser) (parse.Articles, error) {
	ret := _m.Called(rss)

	if len(ret) == 0 {
		panic("no return value specified for RssToArticles")
	}

	var r0 parse.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(io.ReadCloser) (parse.Articles, error)); ok {
		return rf(rss)
	}
	if rf, ok := ret.Get(0).(func(io.ReadCloser) parse.Articles); ok {
		r0 = rf(rss)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(parse.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(io.ReadCloser) error); ok {
		r1 = rf(rss)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransforms_RssToArticles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RssToArticles'
type MockTransforms_RssToArticles_Call struct {
	*mock.Call
}

// RssToArticles is a helper method to define mock.On call
//   - rss io.ReadCloser
func (_e *MockTransforms_Expecter) RssToArticles(rss interface{}) *MockTransforms_RssToArticles_Call {
	return &MockTransforms_RssToArticles_Call{Call: _e.mock.On("RssToArticles", rss)}
}

func (_c *MockTransforms_RssToArticles_Call) Run(run func(rss io.ReadCloser)) *MockTransforms_RssToArticles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.ReadCloser))
	})
	return _c
}

func (_c *MockTransforms_RssToArticles_Call) Return(_a0 parse.Articles, _a1 error) *MockTransforms_RssToArticles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransforms_RssToArticles_Call) RunAndReturn(run func(io.ReadCloser) (parse.Articles, error)) *MockTransforms_RssToArticles_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransforms creates a new instance of MockTransforms. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransforms(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransforms {
	mock := &MockTransforms{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
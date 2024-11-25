// Code generated by mockery v2.47.0. DO NOT EDIT.

package fetch

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockFetch is an autogenerated mock type for the Fetch type
type MockFetch struct {
	mock.Mock
}

type MockFetch_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFetch) EXPECT() *MockFetch_Expecter {
	return &MockFetch_Expecter{mock: &_m.Mock}
}

// Url provides a mock function with given fields: ctx, url, fetchCb
func (_m *MockFetch) Url(ctx context.Context, url string, fetchCb Callback) error {
	ret := _m.Called(ctx, url, fetchCb)

	if len(ret) == 0 {
		panic("no return value specified for Url")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, Callback) error); ok {
		r0 = rf(ctx, url, fetchCb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFetch_Url_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Url'
type MockFetch_Url_Call struct {
	*mock.Call
}

// Url is a helper method to define mock.On call
//   - ctx context.Context
//   - url string
//   - fetchCb Callback
func (_e *MockFetch_Expecter) Url(ctx interface{}, url interface{}, fetchCb interface{}) *MockFetch_Url_Call {
	return &MockFetch_Url_Call{Call: _e.mock.On("Url", ctx, url, fetchCb)}
}

func (_c *MockFetch_Url_Call) Run(run func(ctx context.Context, url string, fetchCb Callback)) *MockFetch_Url_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(Callback))
	})
	return _c
}

func (_c *MockFetch_Url_Call) Return(_a0 error) *MockFetch_Url_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFetch_Url_Call) RunAndReturn(run func(context.Context, string, Callback) error) *MockFetch_Url_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFetch creates a new instance of MockFetch. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFetch(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFetch {
	mock := &MockFetch{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
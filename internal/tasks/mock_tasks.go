// Code generated by mockery v2.47.0. DO NOT EDIT.

package tasks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTasks is an autogenerated mock type for the Tasks type
type MockTasks struct {
	mock.Mock
}

type MockTasks_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTasks) EXPECT() *MockTasks_Expecter {
	return &MockTasks_Expecter{mock: &_m.Mock}
}

// Workflow provides a mock function with given fields: ctx, task, args
func (_m *MockTasks) Workflow(ctx context.Context, task TaskType, args []any) (*TaskResult, error) {
	ret := _m.Called(ctx, task, args)

	if len(ret) == 0 {
		panic("no return value specified for Workflow")
	}

	var r0 *TaskResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, TaskType, []any) (*TaskResult, error)); ok {
		return rf(ctx, task, args)
	}
	if rf, ok := ret.Get(0).(func(context.Context, TaskType, []any) *TaskResult); ok {
		r0 = rf(ctx, task, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TaskResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, TaskType, []any) error); ok {
		r1 = rf(ctx, task, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTasks_Workflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Workflow'
type MockTasks_Workflow_Call struct {
	*mock.Call
}

// Workflow is a helper method to define mock.On call
//   - ctx context.Context
//   - task TaskType
//   - args []any
func (_e *MockTasks_Expecter) Workflow(ctx interface{}, task interface{}, args interface{}) *MockTasks_Workflow_Call {
	return &MockTasks_Workflow_Call{Call: _e.mock.On("Workflow", ctx, task, args)}
}

func (_c *MockTasks_Workflow_Call) Run(run func(ctx context.Context, task TaskType, args []any)) *MockTasks_Workflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(TaskType), args[2].([]any))
	})
	return _c
}

func (_c *MockTasks_Workflow_Call) Return(_a0 *TaskResult, _a1 error) *MockTasks_Workflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTasks_Workflow_Call) RunAndReturn(run func(context.Context, TaskType, []any) (*TaskResult, error)) *MockTasks_Workflow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTasks creates a new instance of MockTasks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTasks(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTasks {
	mock := &MockTasks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.46.2. DO NOT EDIT.

package store

import (
	context "context"

	enum "github.com/den4ik117/examly/internal/enum"
	mock "github.com/stretchr/testify/mock"
)

// MockTestSessionsStore is an autogenerated mock type for the TestSessionsStore type
type MockTestSessionsStore struct {
	mock.Mock
}

type MockTestSessionsStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTestSessionsStore) EXPECT() *MockTestSessionsStore_Expecter {
	return &MockTestSessionsStore_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, test
func (_m *MockTestSessionsStore) Create(ctx context.Context, test *TestSession) error {
	ret := _m.Called(ctx, test)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *TestSession) error); ok {
		r0 = rf(ctx, test)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestSessionsStore_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTestSessionsStore_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - test *TestSession
func (_e *MockTestSessionsStore_Expecter) Create(ctx interface{}, test interface{}) *MockTestSessionsStore_Create_Call {
	return &MockTestSessionsStore_Create_Call{Call: _e.mock.On("Create", ctx, test)}
}

func (_c *MockTestSessionsStore_Create_Call) Run(run func(ctx context.Context, test *TestSession)) *MockTestSessionsStore_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*TestSession))
	})
	return _c
}

func (_c *MockTestSessionsStore_Create_Call) Return(_a0 error) *MockTestSessionsStore_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestSessionsStore_Create_Call) RunAndReturn(run func(context.Context, *TestSession) error) *MockTestSessionsStore_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockTestSessionsStore) GetByID(ctx context.Context, id int) (TestSession, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 TestSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (TestSession, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) TestSession); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(TestSession)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestSessionsStore_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockTestSessionsStore_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockTestSessionsStore_Expecter) GetByID(ctx interface{}, id interface{}) *MockTestSessionsStore_GetByID_Call {
	return &MockTestSessionsStore_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *MockTestSessionsStore_GetByID_Call) Run(run func(ctx context.Context, id int)) *MockTestSessionsStore_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockTestSessionsStore_GetByID_Call) Return(_a0 TestSession, _a1 error) *MockTestSessionsStore_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestSessionsStore_GetByID_Call) RunAndReturn(run func(context.Context, int) (TestSession, error)) *MockTestSessionsStore_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUUID provides a mock function with given fields: ctx, uuid
func (_m *MockTestSessionsStore) GetByUUID(ctx context.Context, uuid string) (TestSession, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for GetByUUID")
	}

	var r0 TestSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (TestSession, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) TestSession); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(TestSession)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestSessionsStore_GetByUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUUID'
type MockTestSessionsStore_GetByUUID_Call struct {
	*mock.Call
}

// GetByUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - uuid string
func (_e *MockTestSessionsStore_Expecter) GetByUUID(ctx interface{}, uuid interface{}) *MockTestSessionsStore_GetByUUID_Call {
	return &MockTestSessionsStore_GetByUUID_Call{Call: _e.mock.On("GetByUUID", ctx, uuid)}
}

func (_c *MockTestSessionsStore_GetByUUID_Call) Run(run func(ctx context.Context, uuid string)) *MockTestSessionsStore_GetByUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTestSessionsStore_GetByUUID_Call) Return(_a0 TestSession, _a1 error) *MockTestSessionsStore_GetByUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestSessionsStore_GetByUUID_Call) RunAndReturn(run func(context.Context, string) (TestSession, error)) *MockTestSessionsStore_GetByUUID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUserIDAndType provides a mock function with given fields: ctx, id, t
func (_m *MockTestSessionsStore) GetByUserIDAndType(ctx context.Context, id int, t enum.TestSessionType) (TestSession, error) {
	ret := _m.Called(ctx, id, t)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserIDAndType")
	}

	var r0 TestSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, enum.TestSessionType) (TestSession, error)); ok {
		return rf(ctx, id, t)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, enum.TestSessionType) TestSession); ok {
		r0 = rf(ctx, id, t)
	} else {
		r0 = ret.Get(0).(TestSession)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, enum.TestSessionType) error); ok {
		r1 = rf(ctx, id, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestSessionsStore_GetByUserIDAndType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUserIDAndType'
type MockTestSessionsStore_GetByUserIDAndType_Call struct {
	*mock.Call
}

// GetByUserIDAndType is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
//   - t enum.TestSessionType
func (_e *MockTestSessionsStore_Expecter) GetByUserIDAndType(ctx interface{}, id interface{}, t interface{}) *MockTestSessionsStore_GetByUserIDAndType_Call {
	return &MockTestSessionsStore_GetByUserIDAndType_Call{Call: _e.mock.On("GetByUserIDAndType", ctx, id, t)}
}

func (_c *MockTestSessionsStore_GetByUserIDAndType_Call) Run(run func(ctx context.Context, id int, t enum.TestSessionType)) *MockTestSessionsStore_GetByUserIDAndType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(enum.TestSessionType))
	})
	return _c
}

func (_c *MockTestSessionsStore_GetByUserIDAndType_Call) Return(_a0 TestSession, _a1 error) *MockTestSessionsStore_GetByUserIDAndType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestSessionsStore_GetByUserIDAndType_Call) RunAndReturn(run func(context.Context, int, enum.TestSessionType) (TestSession, error)) *MockTestSessionsStore_GetByUserIDAndType_Call {
	_c.Call.Return(run)
	return _c
}

// GetStats provides a mock function with given fields: ctx, userID
func (_m *MockTestSessionsStore) GetStats(ctx context.Context, userID int) ([]TestSessionStats, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetStats")
	}

	var r0 []TestSessionStats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]TestSessionStats, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []TestSessionStats); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TestSessionStats)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestSessionsStore_GetStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStats'
type MockTestSessionsStore_GetStats_Call struct {
	*mock.Call
}

// GetStats is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
func (_e *MockTestSessionsStore_Expecter) GetStats(ctx interface{}, userID interface{}) *MockTestSessionsStore_GetStats_Call {
	return &MockTestSessionsStore_GetStats_Call{Call: _e.mock.On("GetStats", ctx, userID)}
}

func (_c *MockTestSessionsStore_GetStats_Call) Run(run func(ctx context.Context, userID int)) *MockTestSessionsStore_GetStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockTestSessionsStore_GetStats_Call) Return(_a0 []TestSessionStats, _a1 error) *MockTestSessionsStore_GetStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestSessionsStore_GetStats_Call) RunAndReturn(run func(context.Context, int) ([]TestSessionStats, error)) *MockTestSessionsStore_GetStats_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, test
func (_m *MockTestSessionsStore) Update(ctx context.Context, test *TestSession) error {
	ret := _m.Called(ctx, test)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *TestSession) error); ok {
		r0 = rf(ctx, test)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestSessionsStore_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockTestSessionsStore_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - test *TestSession
func (_e *MockTestSessionsStore_Expecter) Update(ctx interface{}, test interface{}) *MockTestSessionsStore_Update_Call {
	return &MockTestSessionsStore_Update_Call{Call: _e.mock.On("Update", ctx, test)}
}

func (_c *MockTestSessionsStore_Update_Call) Run(run func(ctx context.Context, test *TestSession)) *MockTestSessionsStore_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*TestSession))
	})
	return _c
}

func (_c *MockTestSessionsStore_Update_Call) Return(_a0 error) *MockTestSessionsStore_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestSessionsStore_Update_Call) RunAndReturn(run func(context.Context, *TestSession) error) *MockTestSessionsStore_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTestSessionsStore creates a new instance of MockTestSessionsStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTestSessionsStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTestSessionsStore {
	mock := &MockTestSessionsStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

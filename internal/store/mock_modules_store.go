// Code generated by mockery v2.46.2. DO NOT EDIT.

package store

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockModulesStore is an autogenerated mock type for the ModulesStore type
type MockModulesStore struct {
	mock.Mock
}

type MockModulesStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockModulesStore) EXPECT() *MockModulesStore_Expecter {
	return &MockModulesStore_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, module
func (_m *MockModulesStore) Create(ctx context.Context, module *Module) error {
	ret := _m.Called(ctx, module)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Module) error); ok {
		r0 = rf(ctx, module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockModulesStore_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockModulesStore_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - module *Module
func (_e *MockModulesStore_Expecter) Create(ctx interface{}, module interface{}) *MockModulesStore_Create_Call {
	return &MockModulesStore_Create_Call{Call: _e.mock.On("Create", ctx, module)}
}

func (_c *MockModulesStore_Create_Call) Run(run func(ctx context.Context, module *Module)) *MockModulesStore_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Module))
	})
	return _c
}

func (_c *MockModulesStore_Create_Call) Return(_a0 error) *MockModulesStore_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockModulesStore_Create_Call) RunAndReturn(run func(context.Context, *Module) error) *MockModulesStore_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, module
func (_m *MockModulesStore) Delete(ctx context.Context, module *Module) error {
	ret := _m.Called(ctx, module)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Module) error); ok {
		r0 = rf(ctx, module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockModulesStore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockModulesStore_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - module *Module
func (_e *MockModulesStore_Expecter) Delete(ctx interface{}, module interface{}) *MockModulesStore_Delete_Call {
	return &MockModulesStore_Delete_Call{Call: _e.mock.On("Delete", ctx, module)}
}

func (_c *MockModulesStore_Delete_Call) Run(run func(ctx context.Context, module *Module)) *MockModulesStore_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Module))
	})
	return _c
}

func (_c *MockModulesStore_Delete_Call) Return(_a0 error) *MockModulesStore_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockModulesStore_Delete_Call) RunAndReturn(run func(context.Context, *Module) error) *MockModulesStore_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx
func (_m *MockModulesStore) Get(ctx context.Context) ([]Module, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Module, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Module); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Module)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockModulesStore_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockModulesStore_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockModulesStore_Expecter) Get(ctx interface{}) *MockModulesStore_Get_Call {
	return &MockModulesStore_Get_Call{Call: _e.mock.On("Get", ctx)}
}

func (_c *MockModulesStore_Get_Call) Run(run func(ctx context.Context)) *MockModulesStore_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockModulesStore_Get_Call) Return(_a0 []Module, _a1 error) *MockModulesStore_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockModulesStore_Get_Call) RunAndReturn(run func(context.Context) ([]Module, error)) *MockModulesStore_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUUID provides a mock function with given fields: ctx, uuid
func (_m *MockModulesStore) GetByUUID(ctx context.Context, uuid string) (Module, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for GetByUUID")
	}

	var r0 Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (Module, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) Module); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(Module)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockModulesStore_GetByUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUUID'
type MockModulesStore_GetByUUID_Call struct {
	*mock.Call
}

// GetByUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - uuid string
func (_e *MockModulesStore_Expecter) GetByUUID(ctx interface{}, uuid interface{}) *MockModulesStore_GetByUUID_Call {
	return &MockModulesStore_GetByUUID_Call{Call: _e.mock.On("GetByUUID", ctx, uuid)}
}

func (_c *MockModulesStore_GetByUUID_Call) Run(run func(ctx context.Context, uuid string)) *MockModulesStore_GetByUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockModulesStore_GetByUUID_Call) Return(_a0 Module, _a1 error) *MockModulesStore_GetByUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockModulesStore_GetByUUID_Call) RunAndReturn(run func(context.Context, string) (Module, error)) *MockModulesStore_GetByUUID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, module
func (_m *MockModulesStore) Update(ctx context.Context, module *Module) error {
	ret := _m.Called(ctx, module)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Module) error); ok {
		r0 = rf(ctx, module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockModulesStore_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockModulesStore_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - module *Module
func (_e *MockModulesStore_Expecter) Update(ctx interface{}, module interface{}) *MockModulesStore_Update_Call {
	return &MockModulesStore_Update_Call{Call: _e.mock.On("Update", ctx, module)}
}

func (_c *MockModulesStore_Update_Call) Run(run func(ctx context.Context, module *Module)) *MockModulesStore_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Module))
	})
	return _c
}

func (_c *MockModulesStore_Update_Call) Return(_a0 error) *MockModulesStore_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockModulesStore_Update_Call) RunAndReturn(run func(context.Context, *Module) error) *MockModulesStore_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, module
func (_m *MockModulesStore) UpdateStatus(ctx context.Context, module *Module) error {
	ret := _m.Called(ctx, module)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Module) error); ok {
		r0 = rf(ctx, module)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockModulesStore_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type MockModulesStore_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - module *Module
func (_e *MockModulesStore_Expecter) UpdateStatus(ctx interface{}, module interface{}) *MockModulesStore_UpdateStatus_Call {
	return &MockModulesStore_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, module)}
}

func (_c *MockModulesStore_UpdateStatus_Call) Run(run func(ctx context.Context, module *Module)) *MockModulesStore_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Module))
	})
	return _c
}

func (_c *MockModulesStore_UpdateStatus_Call) Return(_a0 error) *MockModulesStore_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockModulesStore_UpdateStatus_Call) RunAndReturn(run func(context.Context, *Module) error) *MockModulesStore_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockModulesStore creates a new instance of MockModulesStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockModulesStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockModulesStore {
	mock := &MockModulesStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

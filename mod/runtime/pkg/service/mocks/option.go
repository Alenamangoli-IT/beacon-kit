// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Option is an autogenerated mock type for the Option type
type Option[T interface{}] struct {
	mock.Mock
}

type Option_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *Option[T]) EXPECT() *Option_Expecter[T] {
	return &Option_Expecter[T]{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *Option[T]) Execute(_a0 *T) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*T) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Option_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type Option_Execute_Call[T interface{}] struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *T
func (_e *Option_Expecter[T]) Execute(_a0 interface{}) *Option_Execute_Call[T] {
	return &Option_Execute_Call[T]{Call: _e.mock.On("Execute", _a0)}
}

func (_c *Option_Execute_Call[T]) Run(run func(_a0 *T)) *Option_Execute_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*T))
	})
	return _c
}

func (_c *Option_Execute_Call[T]) Return(_a0 error) *Option_Execute_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Option_Execute_Call[T]) RunAndReturn(run func(*T) error) *Option_Execute_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewOption creates a new instance of Option. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOption[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Option[T] {
	mock := &Option[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

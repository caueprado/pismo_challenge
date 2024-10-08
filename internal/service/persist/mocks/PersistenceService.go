// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	domain "pismo/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// PersistenceService is an autogenerated mock type for the PersistenceService type
type PersistenceService struct {
	mock.Mock
}

type PersistenceService_Expecter struct {
	mock *mock.Mock
}

func (_m *PersistenceService) EXPECT() *PersistenceService_Expecter {
	return &PersistenceService_Expecter{mock: &_m.Mock}
}

// Insert provides a mock function with given fields: event
func (_m *PersistenceService) Insert(event domain.EventMessage) error {
	ret := _m.Called(event)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.EventMessage) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistenceService_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type PersistenceService_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - event domain.EventMessage
func (_e *PersistenceService_Expecter) Insert(event interface{}) *PersistenceService_Insert_Call {
	return &PersistenceService_Insert_Call{Call: _e.mock.On("Insert", event)}
}

func (_c *PersistenceService_Insert_Call) Run(run func(event domain.EventMessage)) *PersistenceService_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.EventMessage))
	})
	return _c
}

func (_c *PersistenceService_Insert_Call) Return(_a0 error) *PersistenceService_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersistenceService_Insert_Call) RunAndReturn(run func(domain.EventMessage) error) *PersistenceService_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// NewPersistenceService creates a new instance of PersistenceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPersistenceService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PersistenceService {
	mock := &PersistenceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

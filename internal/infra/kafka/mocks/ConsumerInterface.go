// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerInterface is an autogenerated mock type for the ConsumerInterface type
type ConsumerInterface struct {
	mock.Mock
}

type ConsumerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ConsumerInterface) EXPECT() *ConsumerInterface_Expecter {
	return &ConsumerInterface_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *ConsumerInterface) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsumerInterface_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type ConsumerInterface_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *ConsumerInterface_Expecter) Close() *ConsumerInterface_Close_Call {
	return &ConsumerInterface_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *ConsumerInterface_Close_Call) Run(run func()) *ConsumerInterface_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConsumerInterface_Close_Call) Return(_a0 error) *ConsumerInterface_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsumerInterface_Close_Call) RunAndReturn(run func() error) *ConsumerInterface_Close_Call {
	_c.Call.Return(run)
	return _c
}

// ReadMessage provides a mock function with given fields:
func (_m *ConsumerInterface) ReadMessage() (*kafka.Message, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadMessage")
	}

	var r0 *kafka.Message
	var r1 error
	if rf, ok := ret.Get(0).(func() (*kafka.Message, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *kafka.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafka.Message)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsumerInterface_ReadMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadMessage'
type ConsumerInterface_ReadMessage_Call struct {
	*mock.Call
}

// ReadMessage is a helper method to define mock.On call
func (_e *ConsumerInterface_Expecter) ReadMessage() *ConsumerInterface_ReadMessage_Call {
	return &ConsumerInterface_ReadMessage_Call{Call: _e.mock.On("ReadMessage")}
}

func (_c *ConsumerInterface_ReadMessage_Call) Run(run func()) *ConsumerInterface_ReadMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConsumerInterface_ReadMessage_Call) Return(_a0 *kafka.Message, _a1 error) *ConsumerInterface_ReadMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsumerInterface_ReadMessage_Call) RunAndReturn(run func() (*kafka.Message, error)) *ConsumerInterface_ReadMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewConsumerInterface creates a new instance of ConsumerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsumerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsumerInterface {
	mock := &ConsumerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

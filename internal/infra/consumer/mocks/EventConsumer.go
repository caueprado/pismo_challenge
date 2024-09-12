// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	kafka "pismo/internal/infra/kafka"

	mock "github.com/stretchr/testify/mock"
)

// EventConsumer is an autogenerated mock type for the EventConsumer type
type EventConsumer struct {
	mock.Mock
}

type EventConsumer_Expecter struct {
	mock *mock.Mock
}

func (_m *EventConsumer) EXPECT() *EventConsumer_Expecter {
	return &EventConsumer_Expecter{mock: &_m.Mock}
}

// Consume provides a mock function with given fields: doneChan
func (_m *EventConsumer) Consume(doneChan chan bool) error {
	ret := _m.Called(doneChan)

	if len(ret) == 0 {
		panic("no return value specified for Consume")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(chan bool) error); ok {
		r0 = rf(doneChan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventConsumer_Consume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Consume'
type EventConsumer_Consume_Call struct {
	*mock.Call
}

// Consume is a helper method to define mock.On call
//   - doneChan chan bool
func (_e *EventConsumer_Expecter) Consume(doneChan interface{}) *EventConsumer_Consume_Call {
	return &EventConsumer_Consume_Call{Call: _e.mock.On("Consume", doneChan)}
}

func (_c *EventConsumer_Consume_Call) Run(run func(doneChan chan bool)) *EventConsumer_Consume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chan bool))
	})
	return _c
}

func (_c *EventConsumer_Consume_Call) Return(_a0 error) *EventConsumer_Consume_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventConsumer_Consume_Call) RunAndReturn(run func(chan bool) error) *EventConsumer_Consume_Call {
	_c.Call.Return(run)
	return _c
}

// consumeMessages provides a mock function with given fields: _a0, doneChan
func (_m *EventConsumer) consumeMessages(_a0 kafka.ConsumerInterface, doneChan chan bool) {
	_m.Called(_a0, doneChan)
}

// EventConsumer_consumeMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'consumeMessages'
type EventConsumer_consumeMessages_Call struct {
	*mock.Call
}

// consumeMessages is a helper method to define mock.On call
//   - _a0 kafka.ConsumerInterface
//   - doneChan chan bool
func (_e *EventConsumer_Expecter) consumeMessages(_a0 interface{}, doneChan interface{}) *EventConsumer_consumeMessages_Call {
	return &EventConsumer_consumeMessages_Call{Call: _e.mock.On("consumeMessages", _a0, doneChan)}
}

func (_c *EventConsumer_consumeMessages_Call) Run(run func(_a0 kafka.ConsumerInterface, doneChan chan bool)) *EventConsumer_consumeMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(kafka.ConsumerInterface), args[1].(chan bool))
	})
	return _c
}

func (_c *EventConsumer_consumeMessages_Call) Return() *EventConsumer_consumeMessages_Call {
	_c.Call.Return()
	return _c
}

func (_c *EventConsumer_consumeMessages_Call) RunAndReturn(run func(kafka.ConsumerInterface, chan bool)) *EventConsumer_consumeMessages_Call {
	_c.Call.Return(run)
	return _c
}

// NewEventConsumer creates a new instance of EventConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventConsumer {
	mock := &EventConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

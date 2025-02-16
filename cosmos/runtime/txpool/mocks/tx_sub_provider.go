// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	core "github.com/ethereum/go-ethereum/core"
	event "github.com/ethereum/go-ethereum/event"

	mock "github.com/stretchr/testify/mock"
)

// TxSubProvider is an autogenerated mock type for the TxSubProvider type
type TxSubProvider struct {
	mock.Mock
}

type TxSubProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *TxSubProvider) EXPECT() *TxSubProvider_Expecter {
	return &TxSubProvider_Expecter{mock: &_m.Mock}
}

// SubscribeTransactions provides a mock function with given fields: ch, reorgs
func (_m *TxSubProvider) SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription {
	ret := _m.Called(ch, reorgs)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeTransactions")
	}

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(chan<- core.NewTxsEvent, bool) event.Subscription); ok {
		r0 = rf(ch, reorgs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	return r0
}

// TxSubProvider_SubscribeTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeTransactions'
type TxSubProvider_SubscribeTransactions_Call struct {
	*mock.Call
}

// SubscribeTransactions is a helper method to define mock.On call
//   - ch chan<- core.NewTxsEvent
//   - reorgs bool
func (_e *TxSubProvider_Expecter) SubscribeTransactions(ch interface{}, reorgs interface{}) *TxSubProvider_SubscribeTransactions_Call {
	return &TxSubProvider_SubscribeTransactions_Call{Call: _e.mock.On("SubscribeTransactions", ch, reorgs)}
}

func (_c *TxSubProvider_SubscribeTransactions_Call) Run(run func(ch chan<- core.NewTxsEvent, reorgs bool)) *TxSubProvider_SubscribeTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chan<- core.NewTxsEvent), args[1].(bool))
	})
	return _c
}

func (_c *TxSubProvider_SubscribeTransactions_Call) Return(_a0 event.Subscription) *TxSubProvider_SubscribeTransactions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxSubProvider_SubscribeTransactions_Call) RunAndReturn(run func(chan<- core.NewTxsEvent, bool) event.Subscription) *TxSubProvider_SubscribeTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxSubProvider creates a new instance of TxSubProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxSubProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxSubProvider {
	mock := &TxSubProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

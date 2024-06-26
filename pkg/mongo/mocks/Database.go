// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	mongo "github.com/quachhoang2002/Music-Library/pkg/mongo"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

type Database_Expecter struct {
	mock *mock.Mock
}

func (_m *Database) EXPECT() *Database_Expecter {
	return &Database_Expecter{mock: &_m.Mock}
}

// Client provides a mock function with given fields:
func (_m *Database) Client() mongo.Client {
	ret := _m.Called()

	var r0 mongo.Client
	if rf, ok := ret.Get(0).(func() mongo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Client)
		}
	}

	return r0
}

// Database_Client_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Client'
type Database_Client_Call struct {
	*mock.Call
}

// Client is a helper method to define mock.On call
func (_e *Database_Expecter) Client() *Database_Client_Call {
	return &Database_Client_Call{Call: _e.mock.On("Client")}
}

func (_c *Database_Client_Call) Run(run func()) *Database_Client_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Client_Call) Return(_a0 mongo.Client) *Database_Client_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Client_Call) RunAndReturn(run func() mongo.Client) *Database_Client_Call {
	_c.Call.Return(run)
	return _c
}

// Collection provides a mock function with given fields: _a0
func (_m *Database) Collection(_a0 string) mongo.Collection {
	ret := _m.Called(_a0)

	var r0 mongo.Collection
	if rf, ok := ret.Get(0).(func(string) mongo.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Collection)
		}
	}

	return r0
}

// Database_Collection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Collection'
type Database_Collection_Call struct {
	*mock.Call
}

// Collection is a helper method to define mock.On call
//   - _a0 string
func (_e *Database_Expecter) Collection(_a0 interface{}) *Database_Collection_Call {
	return &Database_Collection_Call{Call: _e.mock.On("Collection", _a0)}
}

func (_c *Database_Collection_Call) Run(run func(_a0 string)) *Database_Collection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Database_Collection_Call) Return(_a0 mongo.Collection) *Database_Collection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Collection_Call) RunAndReturn(run func(string) mongo.Collection) *Database_Collection_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	database "github.com/KarnerTh/mermerd/database"
	mock "github.com/stretchr/testify/mock"
)

// ConnectorFactory is an autogenerated mock type for the ConnectorFactory type
type ConnectorFactory struct {
	mock.Mock
}

// NewConnector provides a mock function with given fields: connectionString
func (_m *ConnectorFactory) NewConnector(connectionString string) (database.Connector, error) {
	ret := _m.Called(connectionString)

	var r0 database.Connector
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (database.Connector, error)); ok {
		return rf(connectionString)
	}
	if rf, ok := ret.Get(0).(func(string) database.Connector); ok {
		r0 = rf(connectionString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Connector)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(connectionString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewConnectorFactory creates a new instance of ConnectorFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectorFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectorFactory {
	mock := &ConnectorFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

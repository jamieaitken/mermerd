// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	config "github.com/KarnerTh/mermerd/config"

	mock "github.com/stretchr/testify/mock"
)

// RelationshipLabelMap is an autogenerated mock type for the RelationshipLabelMap type
type RelationshipLabelMap struct {
	mock.Mock
}

// AddRelationshipLabel provides a mock function with given fields: label
func (_m *RelationshipLabelMap) AddRelationshipLabel(label config.RelationshipLabel) {
	_m.Called(label)
}

// LookupRelationshipLabel provides a mock function with given fields: pkName, fkName
func (_m *RelationshipLabelMap) LookupRelationshipLabel(pkName string, fkName string) (config.RelationshipLabel, bool) {
	ret := _m.Called(pkName, fkName)

	var r0 config.RelationshipLabel
	var r1 bool
	if rf, ok := ret.Get(0).(func(string, string) (config.RelationshipLabel, bool)); ok {
		return rf(pkName, fkName)
	}
	if rf, ok := ret.Get(0).(func(string, string) config.RelationshipLabel); ok {
		r0 = rf(pkName, fkName)
	} else {
		r0 = ret.Get(0).(config.RelationshipLabel)
	}

	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(pkName, fkName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewRelationshipLabelMap creates a new instance of RelationshipLabelMap. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRelationshipLabelMap(t interface {
	mock.TestingT
	Cleanup(func())
}) *RelationshipLabelMap {
	mock := &RelationshipLabelMap{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

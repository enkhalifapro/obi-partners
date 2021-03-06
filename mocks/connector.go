// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Connector is an autogenerated mock type for the Connector type
type Connector struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: collectionName, document
func (_m *Connector) CreateOne(collectionName string, document interface{}) (string, error) {
	ret := _m.Called(collectionName, document)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, interface{}) string); ok {
		r0 = rf(collectionName, document)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(collectionName, document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: collectionName, filter, res
func (_m *Connector) FindOne(collectionName string, filter interface{}, res interface{}) error {
	ret := _m.Called(collectionName, filter, res)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, interface{}) error); ok {
		r0 = rf(collectionName, filter, res)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

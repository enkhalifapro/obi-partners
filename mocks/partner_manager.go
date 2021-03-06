// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	types "github.com/enkhalifapro/obi-partners/internal/types"
	mock "github.com/stretchr/testify/mock"
)

// PartnerManager is an autogenerated mock type for the PartnerManager type
type PartnerManager struct {
	mock.Mock
}

// Add provides a mock function with given fields: partner
func (_m *PartnerManager) Add(partner *types.Partner) error {
	ret := _m.Called(partner)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Partner) error); ok {
		r0 = rf(partner)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AppOptions is an autogenerated mock type for the AppOptions type
type AppOptions struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *AppOptions) Get(_a0 string) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewAppOptions interface {
	mock.TestingT
	Cleanup(func())
}

// NewAppOptions creates a new instance of AppOptions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAppOptions(t mockConstructorTestingTNewAppOptions) *AppOptions {
	mock := &AppOptions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

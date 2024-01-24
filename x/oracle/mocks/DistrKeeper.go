// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// DistrKeeper is an autogenerated mock type for the DistrKeeper type
type DistrKeeper struct {
	mock.Mock
}

// AllocateTokensToValidator provides a mock function with given fields: ctx, val, tokens
func (_m *DistrKeeper) AllocateTokensToValidator(ctx context.Context, val types.ValidatorI, tokens cosmos_sdktypes.DecCoins) error {
	ret := _m.Called(ctx, val, tokens)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ValidatorI, cosmos_sdktypes.DecCoins) error); ok {
		r0 = rf(ctx, val, tokens)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDistrKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewDistrKeeper creates a new instance of DistrKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDistrKeeper(t mockConstructorTestingTNewDistrKeeper) *DistrKeeper {
	mock := &DistrKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

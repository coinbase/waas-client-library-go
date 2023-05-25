// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	mock "github.com/stretchr/testify/mock"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// MPCWalletServiceServer is an autogenerated mock type for the MPCWalletServiceServer type
type MPCWalletServiceServer struct {
	mock.Mock
}

// CreateMPCWallet provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) CreateMPCWallet(_a0 context.Context, _a1 *v1.CreateMPCWalletRequest) (*longrunningpb.Operation, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *longrunningpb.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCWalletRequest) *longrunningpb.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateMPCWalletRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAddress provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) GenerateAddress(_a0 context.Context, _a1 *v1.GenerateAddressRequest) (*v1.Address, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Address
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GenerateAddressRequest) *v1.Address); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GenerateAddressRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddress provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) GetAddress(_a0 context.Context, _a1 *v1.GetAddressRequest) (*v1.Address, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Address
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetAddressRequest) *v1.Address); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetAddressRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMPCWallet provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) GetMPCWallet(_a0 context.Context, _a1 *v1.GetMPCWalletRequest) (*v1.MPCWallet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.MPCWallet
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCWalletRequest) *v1.MPCWallet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCWallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetMPCWalletRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAddresses provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) ListAddresses(_a0 context.Context, _a1 *v1.ListAddressesRequest) (*v1.ListAddressesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ListAddressesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListAddressesRequest) *v1.ListAddressesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListAddressesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListAddressesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBalanceDetails provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) ListBalanceDetails(_a0 context.Context, _a1 *v1.ListBalanceDetailsRequest) (*v1.ListBalanceDetailsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ListBalanceDetailsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListBalanceDetailsRequest) *v1.ListBalanceDetailsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListBalanceDetailsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListBalanceDetailsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBalances provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) ListBalances(_a0 context.Context, _a1 *v1.ListBalancesRequest) (*v1.ListBalancesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ListBalancesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListBalancesRequest) *v1.ListBalancesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListBalancesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListBalancesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMPCWallets provides a mock function with given fields: _a0, _a1
func (_m *MPCWalletServiceServer) ListMPCWallets(_a0 context.Context, _a1 *v1.ListMPCWalletsRequest) (*v1.ListMPCWalletsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ListMPCWalletsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCWalletsRequest) *v1.ListMPCWalletsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListMPCWalletsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListMPCWalletsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedMPCWalletServiceServer provides a mock function with given fields:
func (_m *MPCWalletServiceServer) mustEmbedUnimplementedMPCWalletServiceServer() {
	_m.Called()
}

type NewMPCWalletServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMPCWalletServiceServer creates a new instance of MPCWalletServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMPCWalletServiceServer(t NewMPCWalletServiceServerT) *MPCWalletServiceServer {
	mock := &MPCWalletServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

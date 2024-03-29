// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// MPCWalletServiceClient is an autogenerated mock type for the MPCWalletServiceClient type
type MPCWalletServiceClient struct {
	mock.Mock
}

// CreateMPCWallet provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) CreateMPCWallet(ctx context.Context, in *v1.CreateMPCWalletRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *longrunningpb.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCWalletRequest, ...grpc.CallOption) (*longrunningpb.Operation, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCWalletRequest, ...grpc.CallOption) *longrunningpb.Operation); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateMPCWalletRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAddress provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) GenerateAddress(ctx context.Context, in *v1.GenerateAddressRequest, opts ...grpc.CallOption) (*v1.Address, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GenerateAddressRequest, ...grpc.CallOption) (*v1.Address, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GenerateAddressRequest, ...grpc.CallOption) *v1.Address); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.GenerateAddressRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddress provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) GetAddress(ctx context.Context, in *v1.GetAddressRequest, opts ...grpc.CallOption) (*v1.Address, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetAddressRequest, ...grpc.CallOption) (*v1.Address, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetAddressRequest, ...grpc.CallOption) *v1.Address); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetAddressRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMPCWallet provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) GetMPCWallet(ctx context.Context, in *v1.GetMPCWalletRequest, opts ...grpc.CallOption) (*v1.MPCWallet, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.MPCWallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCWalletRequest, ...grpc.CallOption) (*v1.MPCWallet, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCWalletRequest, ...grpc.CallOption) *v1.MPCWallet); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCWallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetMPCWalletRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAddresses provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) ListAddresses(ctx context.Context, in *v1.ListAddressesRequest, opts ...grpc.CallOption) (*v1.ListAddressesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ListAddressesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListAddressesRequest, ...grpc.CallOption) (*v1.ListAddressesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListAddressesRequest, ...grpc.CallOption) *v1.ListAddressesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListAddressesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListAddressesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBalanceDetails provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) ListBalanceDetails(ctx context.Context, in *v1.ListBalanceDetailsRequest, opts ...grpc.CallOption) (*v1.ListBalanceDetailsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ListBalanceDetailsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListBalanceDetailsRequest, ...grpc.CallOption) (*v1.ListBalanceDetailsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListBalanceDetailsRequest, ...grpc.CallOption) *v1.ListBalanceDetailsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListBalanceDetailsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListBalanceDetailsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBalances provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) ListBalances(ctx context.Context, in *v1.ListBalancesRequest, opts ...grpc.CallOption) (*v1.ListBalancesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ListBalancesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListBalancesRequest, ...grpc.CallOption) (*v1.ListBalancesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListBalancesRequest, ...grpc.CallOption) *v1.ListBalancesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListBalancesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListBalancesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMPCWallets provides a mock function with given fields: ctx, in, opts
func (_m *MPCWalletServiceClient) ListMPCWallets(ctx context.Context, in *v1.ListMPCWalletsRequest, opts ...grpc.CallOption) (*v1.ListMPCWalletsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ListMPCWalletsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCWalletsRequest, ...grpc.CallOption) (*v1.ListMPCWalletsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCWalletsRequest, ...grpc.CallOption) *v1.ListMPCWalletsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListMPCWalletsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListMPCWalletsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMPCWalletServiceClient creates a new instance of MPCWalletServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMPCWalletServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MPCWalletServiceClient {
	mock := &MPCWalletServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
)

// MPCTransactionServiceClient is an autogenerated mock type for the MPCTransactionServiceClient type
type MPCTransactionServiceClient struct {
	mock.Mock
}

// CreateMPCTransaction provides a mock function with given fields: ctx, in, opts
func (_m *MPCTransactionServiceClient) CreateMPCTransaction(ctx context.Context, in *v1.CreateMPCTransactionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCTransactionRequest, ...grpc.CallOption) (*longrunningpb.Operation, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCTransactionRequest, ...grpc.CallOption) *longrunningpb.Operation); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateMPCTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMPCTransaction provides a mock function with given fields: ctx, in, opts
func (_m *MPCTransactionServiceClient) GetMPCTransaction(ctx context.Context, in *v1.GetMPCTransactionRequest, opts ...grpc.CallOption) (*v1.MPCTransaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.MPCTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCTransactionRequest, ...grpc.CallOption) (*v1.MPCTransaction, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCTransactionRequest, ...grpc.CallOption) *v1.MPCTransaction); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetMPCTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMPCTransactions provides a mock function with given fields: ctx, in, opts
func (_m *MPCTransactionServiceClient) ListMPCTransactions(ctx context.Context, in *v1.ListMPCTransactionsRequest, opts ...grpc.CallOption) (*v1.ListMPCTransactionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ListMPCTransactionsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCTransactionsRequest, ...grpc.CallOption) (*v1.ListMPCTransactionsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCTransactionsRequest, ...grpc.CallOption) *v1.ListMPCTransactionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListMPCTransactionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListMPCTransactionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMPCTransactionServiceClient creates a new instance of MPCTransactionServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMPCTransactionServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MPCTransactionServiceClient {
	mock := &MPCTransactionServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

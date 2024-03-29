// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	typesv1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
)

// ProtocolServiceClient is an autogenerated mock type for the ProtocolServiceClient type
type ProtocolServiceClient struct {
	mock.Mock
}

// BroadcastTransaction provides a mock function with given fields: ctx, in, opts
func (_m *ProtocolServiceClient) BroadcastTransaction(ctx context.Context, in *v1.BroadcastTransactionRequest, opts ...grpc.CallOption) (*typesv1.Transaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.BroadcastTransactionRequest, ...grpc.CallOption) (*typesv1.Transaction, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.BroadcastTransactionRequest, ...grpc.CallOption) *typesv1.Transaction); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.BroadcastTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConstructTransaction provides a mock function with given fields: ctx, in, opts
func (_m *ProtocolServiceClient) ConstructTransaction(ctx context.Context, in *v1.ConstructTransactionRequest, opts ...grpc.CallOption) (*typesv1.Transaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ConstructTransactionRequest, ...grpc.CallOption) (*typesv1.Transaction, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ConstructTransactionRequest, ...grpc.CallOption) *typesv1.Transaction); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ConstructTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConstructTransferTransaction provides a mock function with given fields: ctx, in, opts
func (_m *ProtocolServiceClient) ConstructTransferTransaction(ctx context.Context, in *v1.ConstructTransferTransactionRequest, opts ...grpc.CallOption) (*typesv1.Transaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ConstructTransferTransactionRequest, ...grpc.CallOption) (*typesv1.Transaction, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ConstructTransferTransactionRequest, ...grpc.CallOption) *typesv1.Transaction); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ConstructTransferTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateFee provides a mock function with given fields: ctx, in, opts
func (_m *ProtocolServiceClient) EstimateFee(ctx context.Context, in *v1.EstimateFeeRequest, opts ...grpc.CallOption) (*v1.EstimateFeeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.EstimateFeeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.EstimateFeeRequest, ...grpc.CallOption) (*v1.EstimateFeeResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.EstimateFeeRequest, ...grpc.CallOption) *v1.EstimateFeeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.EstimateFeeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.EstimateFeeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProtocolServiceClient creates a new instance of ProtocolServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProtocolServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProtocolServiceClient {
	mock := &ProtocolServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

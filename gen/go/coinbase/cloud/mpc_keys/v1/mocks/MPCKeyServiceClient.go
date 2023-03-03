// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"

	mock "github.com/stretchr/testify/mock"

	v1 "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// MPCKeyServiceClient is an autogenerated mock type for the MPCKeyServiceClient type
type MPCKeyServiceClient struct {
	mock.Mock
}

// CreateDeviceGroup provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) CreateDeviceGroup(ctx context.Context, in *v1.CreateDeviceGroupRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *longrunningpb.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateDeviceGroupRequest, ...grpc.CallOption) *longrunningpb.Operation); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateDeviceGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMPCKey provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) CreateMPCKey(ctx context.Context, in *v1.CreateMPCKeyRequest, opts ...grpc.CallOption) (*v1.MPCKey, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.MPCKey
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCKeyRequest, ...grpc.CallOption) *v1.MPCKey); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateMPCKeyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSignature provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) CreateSignature(ctx context.Context, in *v1.CreateSignatureRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *longrunningpb.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateSignatureRequest, ...grpc.CallOption) *longrunningpb.Operation); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateSignatureRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) GetDevice(ctx context.Context, in *v1.GetDeviceRequest, opts ...grpc.CallOption) (*v1.Device, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Device
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetDeviceRequest, ...grpc.CallOption) *v1.Device); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetDeviceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceGroup provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) GetDeviceGroup(ctx context.Context, in *v1.GetDeviceGroupRequest, opts ...grpc.CallOption) (*v1.DeviceGroup, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.DeviceGroup
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetDeviceGroupRequest, ...grpc.CallOption) *v1.DeviceGroup); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.DeviceGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetDeviceGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMPCKey provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) GetMPCKey(ctx context.Context, in *v1.GetMPCKeyRequest, opts ...grpc.CallOption) (*v1.MPCKey, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.MPCKey
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCKeyRequest, ...grpc.CallOption) *v1.MPCKey); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetMPCKeyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMPCOperations provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) ListMPCOperations(ctx context.Context, in *v1.ListMPCOperationsRequest, opts ...grpc.CallOption) (*v1.ListMPCOperationsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ListMPCOperationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCOperationsRequest, ...grpc.CallOption) *v1.ListMPCOperationsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListMPCOperationsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListMPCOperationsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDevice provides a mock function with given fields: ctx, in, opts
func (_m *MPCKeyServiceClient) RegisterDevice(ctx context.Context, in *v1.RegisterDeviceRequest, opts ...grpc.CallOption) (*v1.Device, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Device
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RegisterDeviceRequest, ...grpc.CallOption) *v1.Device); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.RegisterDeviceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewMPCKeyServiceClientT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMPCKeyServiceClient creates a new instance of MPCKeyServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMPCKeyServiceClient(t NewMPCKeyServiceClientT) *MPCKeyServiceClient {
	mock := &MPCKeyServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

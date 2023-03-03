// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	mock "github.com/stretchr/testify/mock"

	v1 "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// MPCKeyServiceServer is an autogenerated mock type for the MPCKeyServiceServer type
type MPCKeyServiceServer struct {
	mock.Mock
}

// CreateDeviceGroup provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) CreateDeviceGroup(_a0 context.Context, _a1 *v1.CreateDeviceGroupRequest) (*longrunningpb.Operation, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *longrunningpb.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateDeviceGroupRequest) *longrunningpb.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateDeviceGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMPCKey provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) CreateMPCKey(_a0 context.Context, _a1 *v1.CreateMPCKeyRequest) (*v1.MPCKey, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.MPCKey
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateMPCKeyRequest) *v1.MPCKey); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateMPCKeyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSignature provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) CreateSignature(_a0 context.Context, _a1 *v1.CreateSignatureRequest) (*longrunningpb.Operation, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *longrunningpb.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateSignatureRequest) *longrunningpb.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*longrunningpb.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateSignatureRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) GetDevice(_a0 context.Context, _a1 *v1.GetDeviceRequest) (*v1.Device, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Device
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetDeviceRequest) *v1.Device); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetDeviceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceGroup provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) GetDeviceGroup(_a0 context.Context, _a1 *v1.GetDeviceGroupRequest) (*v1.DeviceGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.DeviceGroup
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetDeviceGroupRequest) *v1.DeviceGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.DeviceGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetDeviceGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMPCKey provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) GetMPCKey(_a0 context.Context, _a1 *v1.GetMPCKeyRequest) (*v1.MPCKey, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.MPCKey
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetMPCKeyRequest) *v1.MPCKey); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MPCKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetMPCKeyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMPCOperations provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) ListMPCOperations(_a0 context.Context, _a1 *v1.ListMPCOperationsRequest) (*v1.ListMPCOperationsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ListMPCOperationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListMPCOperationsRequest) *v1.ListMPCOperationsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListMPCOperationsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListMPCOperationsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDevice provides a mock function with given fields: _a0, _a1
func (_m *MPCKeyServiceServer) RegisterDevice(_a0 context.Context, _a1 *v1.RegisterDeviceRequest) (*v1.Device, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Device
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RegisterDeviceRequest) *v1.Device); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.RegisterDeviceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedMPCKeyServiceServer provides a mock function with given fields:
func (_m *MPCKeyServiceServer) mustEmbedUnimplementedMPCKeyServiceServer() {
	_m.Called()
}

type NewMPCKeyServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMPCKeyServiceServer creates a new instance of MPCKeyServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMPCKeyServiceServer(t NewMPCKeyServiceServerT) *MPCKeyServiceServer {
	mock := &MPCKeyServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MPCKeyServiceClient is the client API for MPCKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MPCKeyServiceClient interface {
	// Registers a new Device. A Device must be registered before it can be added to a DeviceGroup.
	RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*Device, error)
	// Retrieves a Device by resource name.
	GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*Device, error)
	// Creates a DeviceGroup. The DeviceGroup must contain exactly one registered Device, and
	// the Seed in the DeviceGroup must have at least one HardenedChild. After calling this,
	// use ListMPCOperations to poll for the pending CreateDeviceGroup operation, and use the WaaS SDK's
	// ComputeMPCOperation to complete the operation.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Retrieves a DeviceGroup by resource name.
	GetDeviceGroup(ctx context.Context, in *GetDeviceGroupRequest, opts ...grpc.CallOption) (*DeviceGroup, error)
	// Lists the pending MPCOperations awaiting computation associated with the given
	// parent DeviceGroup. Use this API in combination with the WaaS SDK's computeMPCOperation
	// method to complete the operation.
	ListMPCOperations(ctx context.Context, in *ListMPCOperationsRequest, opts ...grpc.CallOption) (*ListMPCOperationsResponse, error)
	// Creates an MPCKey. There must be a HardenedChild in the Seed of the parent
	// DeviceGroup which is a prefix of the derivation path provided in the MPCKey.
	CreateMPCKey(ctx context.Context, in *CreateMPCKeyRequest, opts ...grpc.CallOption) (*MPCKey, error)
	// Retrieves an MPCKey by resource name.
	GetMPCKey(ctx context.Context, in *GetMPCKeyRequest, opts ...grpc.CallOption) (*MPCKey, error)
	// Creates a Signature using an MPCKey. After calling this, use ListMPCOperations
	// to poll for the pending CreateSignature operation, and use the WaaS SDK's
	// computeMPCOperation to complete the operation.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	CreateSignature(ctx context.Context, in *CreateSignatureRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Prepares an archive in the local storage of the given Device. The archive contains cryptographic materials
	// that can be used to export MPCKeys, which have the given DeviceGroup as their parent.
	// The Device specified in the request must be a member of this DeviceGroup and must participate
	// in the associated MPC operation for the archive to be prepared. After calling this,
	// use ListMPCOperations to poll for the pending PrepareDeviceArchive operation, and use the WaaS SDK's
	// ComputeMPCOperation to complete the operation. Once the operation completes, the Device can utilize the
	// WaaS SDK to export the private keys corresponding to each of the MPCKeys under this DeviceGroup.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	PrepareDeviceArchive(ctx context.Context, in *PrepareDeviceArchiveRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Prepares a backup in the given Device. The backup contains certain cryptographic materials
	// that can be used to restore MPCKeys, which have the given DeviceGroup as their parent, on a new Device.
	// The Device specified in the request must be a member of this DeviceGroup and must participate in the associated
	// MPC operation for the backup to be prepared.
	// After calling this RPC, use ListMPCOperations to poll for the pending PrepareDeviceBackup operation,
	// and use the WaaS SDK's ComputeMPCOperation to complete the operation. Once the operation completes,
	// the Device can utilize WaaS SDK to download the backup bundle. We recommend storing this backup bundle securely
	// in a storage provider of your choice. If the user loses access to their existing Device and wants to recover
	// MPCKeys in the given DeviceGroup on a new Device, use AddDevice RPC on the MPCKeyService.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	PrepareDeviceBackup(ctx context.Context, in *PrepareDeviceBackupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Adds a Device to an existing DeviceGroup. Prior to this API being called, the Device must be registered using
	// RegisterDevice RPC. The Device must have access to the backup created with PrepareDeviceBackup RPC to compute this
	// operation. After calling this RPC, use ListMPCOperations to poll for the pending AddDevice operation,
	// and use the WaaS SDK's ComputeAddDeviceMPCOperation to complete the operation.
	// After the operation is computed on WaaS SDK, the Device will have access to cryptographic materials
	// required to process MPCOperations for this DeviceGroup.
	// Once the operation completes on MPCKeyService, the Device will be added to the given DeviceGroup as a new member
	// and all existing Devices in the DeviceGroup will stay functional.
	// Use the RevokeDevice RPC to remove any of the existing Devices from the DeviceGroup.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Revokes a registered Device. This operation removes the registered Device from all the DeviceGroups that it is a
	// part of. Once the Device is revoked, cryptographic materials in your physical Device are invalidated,
	// and the Device can no longer participate in any MPCOperations of the DeviceGroups it was a part of.
	// Use this API in scenarios such as losing the existing Device, switching to a new physical Device, etc.
	// Ensure that a new Device is successfully added to your DeviceGroups using the AddDevice RPC before invoking
	// the RevokeDevice RPC.
	RevokeDevice(ctx context.Context, in *RevokeDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mPCKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMPCKeyServiceClient(cc grpc.ClientConnInterface) MPCKeyServiceClient {
	return &mPCKeyServiceClient{cc}
}

func (c *mPCKeyServiceClient) RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/RegisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/CreateDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) GetDeviceGroup(ctx context.Context, in *GetDeviceGroupRequest, opts ...grpc.CallOption) (*DeviceGroup, error) {
	out := new(DeviceGroup)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/GetDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) ListMPCOperations(ctx context.Context, in *ListMPCOperationsRequest, opts ...grpc.CallOption) (*ListMPCOperationsResponse, error) {
	out := new(ListMPCOperationsResponse)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/ListMPCOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) CreateMPCKey(ctx context.Context, in *CreateMPCKeyRequest, opts ...grpc.CallOption) (*MPCKey, error) {
	out := new(MPCKey)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/CreateMPCKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) GetMPCKey(ctx context.Context, in *GetMPCKeyRequest, opts ...grpc.CallOption) (*MPCKey, error) {
	out := new(MPCKey)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/GetMPCKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) CreateSignature(ctx context.Context, in *CreateSignatureRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/CreateSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) PrepareDeviceArchive(ctx context.Context, in *PrepareDeviceArchiveRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/PrepareDeviceArchive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) PrepareDeviceBackup(ctx context.Context, in *PrepareDeviceBackupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/PrepareDeviceBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCKeyServiceClient) RevokeDevice(ctx context.Context, in *RevokeDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_keys.v1.MPCKeyService/RevokeDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MPCKeyServiceServer is the server API for MPCKeyService service.
// All implementations must embed UnimplementedMPCKeyServiceServer
// for forward compatibility
type MPCKeyServiceServer interface {
	// Registers a new Device. A Device must be registered before it can be added to a DeviceGroup.
	RegisterDevice(context.Context, *RegisterDeviceRequest) (*Device, error)
	// Retrieves a Device by resource name.
	GetDevice(context.Context, *GetDeviceRequest) (*Device, error)
	// Creates a DeviceGroup. The DeviceGroup must contain exactly one registered Device, and
	// the Seed in the DeviceGroup must have at least one HardenedChild. After calling this,
	// use ListMPCOperations to poll for the pending CreateDeviceGroup operation, and use the WaaS SDK's
	// ComputeMPCOperation to complete the operation.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	CreateDeviceGroup(context.Context, *CreateDeviceGroupRequest) (*longrunning.Operation, error)
	// Retrieves a DeviceGroup by resource name.
	GetDeviceGroup(context.Context, *GetDeviceGroupRequest) (*DeviceGroup, error)
	// Lists the pending MPCOperations awaiting computation associated with the given
	// parent DeviceGroup. Use this API in combination with the WaaS SDK's computeMPCOperation
	// method to complete the operation.
	ListMPCOperations(context.Context, *ListMPCOperationsRequest) (*ListMPCOperationsResponse, error)
	// Creates an MPCKey. There must be a HardenedChild in the Seed of the parent
	// DeviceGroup which is a prefix of the derivation path provided in the MPCKey.
	CreateMPCKey(context.Context, *CreateMPCKeyRequest) (*MPCKey, error)
	// Retrieves an MPCKey by resource name.
	GetMPCKey(context.Context, *GetMPCKeyRequest) (*MPCKey, error)
	// Creates a Signature using an MPCKey. After calling this, use ListMPCOperations
	// to poll for the pending CreateSignature operation, and use the WaaS SDK's
	// computeMPCOperation to complete the operation.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	CreateSignature(context.Context, *CreateSignatureRequest) (*longrunning.Operation, error)
	// Prepares an archive in the local storage of the given Device. The archive contains cryptographic materials
	// that can be used to export MPCKeys, which have the given DeviceGroup as their parent.
	// The Device specified in the request must be a member of this DeviceGroup and must participate
	// in the associated MPC operation for the archive to be prepared. After calling this,
	// use ListMPCOperations to poll for the pending PrepareDeviceArchive operation, and use the WaaS SDK's
	// ComputeMPCOperation to complete the operation. Once the operation completes, the Device can utilize the
	// WaaS SDK to export the private keys corresponding to each of the MPCKeys under this DeviceGroup.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	PrepareDeviceArchive(context.Context, *PrepareDeviceArchiveRequest) (*longrunning.Operation, error)
	// Prepares a backup in the given Device. The backup contains certain cryptographic materials
	// that can be used to restore MPCKeys, which have the given DeviceGroup as their parent, on a new Device.
	// The Device specified in the request must be a member of this DeviceGroup and must participate in the associated
	// MPC operation for the backup to be prepared.
	// After calling this RPC, use ListMPCOperations to poll for the pending PrepareDeviceBackup operation,
	// and use the WaaS SDK's ComputeMPCOperation to complete the operation. Once the operation completes,
	// the Device can utilize WaaS SDK to download the backup bundle. We recommend storing this backup bundle securely
	// in a storage provider of your choice. If the user loses access to their existing Device and wants to recover
	// MPCKeys in the given DeviceGroup on a new Device, use AddDevice RPC on the MPCKeyService.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	PrepareDeviceBackup(context.Context, *PrepareDeviceBackupRequest) (*longrunning.Operation, error)
	// Adds a Device to an existing DeviceGroup. Prior to this API being called, the Device must be registered using
	// RegisterDevice RPC. The Device must have access to the backup created with PrepareDeviceBackup RPC to compute this
	// operation. After calling this RPC, use ListMPCOperations to poll for the pending AddDevice operation,
	// and use the WaaS SDK's ComputeAddDeviceMPCOperation to complete the operation.
	// After the operation is computed on WaaS SDK, the Device will have access to cryptographic materials
	// required to process MPCOperations for this DeviceGroup.
	// Once the operation completes on MPCKeyService, the Device will be added to the given DeviceGroup as a new member
	// and all existing Devices in the DeviceGroup will stay functional.
	// Use the RevokeDevice RPC to remove any of the existing Devices from the DeviceGroup.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	AddDevice(context.Context, *AddDeviceRequest) (*longrunning.Operation, error)
	// Revokes a registered Device. This operation removes the registered Device from all the DeviceGroups that it is a
	// part of. Once the Device is revoked, cryptographic materials in your physical Device are invalidated,
	// and the Device can no longer participate in any MPCOperations of the DeviceGroups it was a part of.
	// Use this API in scenarios such as losing the existing Device, switching to a new physical Device, etc.
	// Ensure that a new Device is successfully added to your DeviceGroups using the AddDevice RPC before invoking
	// the RevokeDevice RPC.
	RevokeDevice(context.Context, *RevokeDeviceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMPCKeyServiceServer()
}

// UnimplementedMPCKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMPCKeyServiceServer struct {
}

func (UnimplementedMPCKeyServiceServer) RegisterDevice(context.Context, *RegisterDeviceRequest) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedMPCKeyServiceServer) GetDevice(context.Context, *GetDeviceRequest) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedMPCKeyServiceServer) CreateDeviceGroup(context.Context, *CreateDeviceGroupRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceGroup not implemented")
}
func (UnimplementedMPCKeyServiceServer) GetDeviceGroup(context.Context, *GetDeviceGroupRequest) (*DeviceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceGroup not implemented")
}
func (UnimplementedMPCKeyServiceServer) ListMPCOperations(context.Context, *ListMPCOperationsRequest) (*ListMPCOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMPCOperations not implemented")
}
func (UnimplementedMPCKeyServiceServer) CreateMPCKey(context.Context, *CreateMPCKeyRequest) (*MPCKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMPCKey not implemented")
}
func (UnimplementedMPCKeyServiceServer) GetMPCKey(context.Context, *GetMPCKeyRequest) (*MPCKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMPCKey not implemented")
}
func (UnimplementedMPCKeyServiceServer) CreateSignature(context.Context, *CreateSignatureRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSignature not implemented")
}
func (UnimplementedMPCKeyServiceServer) PrepareDeviceArchive(context.Context, *PrepareDeviceArchiveRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDeviceArchive not implemented")
}
func (UnimplementedMPCKeyServiceServer) PrepareDeviceBackup(context.Context, *PrepareDeviceBackupRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDeviceBackup not implemented")
}
func (UnimplementedMPCKeyServiceServer) AddDevice(context.Context, *AddDeviceRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (UnimplementedMPCKeyServiceServer) RevokeDevice(context.Context, *RevokeDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeDevice not implemented")
}
func (UnimplementedMPCKeyServiceServer) mustEmbedUnimplementedMPCKeyServiceServer() {}

// UnsafeMPCKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MPCKeyServiceServer will
// result in compilation errors.
type UnsafeMPCKeyServiceServer interface {
	mustEmbedUnimplementedMPCKeyServiceServer()
}

func RegisterMPCKeyServiceServer(s grpc.ServiceRegistrar, srv MPCKeyServiceServer) {
	s.RegisterService(&MPCKeyService_ServiceDesc, srv)
}

func _MPCKeyService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).RegisterDevice(ctx, req.(*RegisterDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).GetDevice(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_CreateDeviceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).CreateDeviceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/CreateDeviceGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).CreateDeviceGroup(ctx, req.(*CreateDeviceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_GetDeviceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).GetDeviceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/GetDeviceGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).GetDeviceGroup(ctx, req.(*GetDeviceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_ListMPCOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMPCOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).ListMPCOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/ListMPCOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).ListMPCOperations(ctx, req.(*ListMPCOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_CreateMPCKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMPCKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).CreateMPCKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/CreateMPCKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).CreateMPCKey(ctx, req.(*CreateMPCKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_GetMPCKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMPCKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).GetMPCKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/GetMPCKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).GetMPCKey(ctx, req.(*GetMPCKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_CreateSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).CreateSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/CreateSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).CreateSignature(ctx, req.(*CreateSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_PrepareDeviceArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDeviceArchiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).PrepareDeviceArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/PrepareDeviceArchive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).PrepareDeviceArchive(ctx, req.(*PrepareDeviceArchiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_PrepareDeviceBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDeviceBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).PrepareDeviceBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/PrepareDeviceBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).PrepareDeviceBackup(ctx, req.(*PrepareDeviceBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCKeyService_RevokeDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCKeyServiceServer).RevokeDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_keys.v1.MPCKeyService/RevokeDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCKeyServiceServer).RevokeDevice(ctx, req.(*RevokeDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MPCKeyService_ServiceDesc is the grpc.ServiceDesc for MPCKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MPCKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coinbase.cloud.mpc_keys.v1.MPCKeyService",
	HandlerType: (*MPCKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevice",
			Handler:    _MPCKeyService_RegisterDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _MPCKeyService_GetDevice_Handler,
		},
		{
			MethodName: "CreateDeviceGroup",
			Handler:    _MPCKeyService_CreateDeviceGroup_Handler,
		},
		{
			MethodName: "GetDeviceGroup",
			Handler:    _MPCKeyService_GetDeviceGroup_Handler,
		},
		{
			MethodName: "ListMPCOperations",
			Handler:    _MPCKeyService_ListMPCOperations_Handler,
		},
		{
			MethodName: "CreateMPCKey",
			Handler:    _MPCKeyService_CreateMPCKey_Handler,
		},
		{
			MethodName: "GetMPCKey",
			Handler:    _MPCKeyService_GetMPCKey_Handler,
		},
		{
			MethodName: "CreateSignature",
			Handler:    _MPCKeyService_CreateSignature_Handler,
		},
		{
			MethodName: "PrepareDeviceArchive",
			Handler:    _MPCKeyService_PrepareDeviceArchive_Handler,
		},
		{
			MethodName: "PrepareDeviceBackup",
			Handler:    _MPCKeyService_PrepareDeviceBackup_Handler,
		},
		{
			MethodName: "AddDevice",
			Handler:    _MPCKeyService_AddDevice_Handler,
		},
		{
			MethodName: "RevokeDevice",
			Handler:    _MPCKeyService_RevokeDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coinbase/cloud/mpc_keys/v1/mpc_keys.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MPCWalletServiceClient is the client API for MPCWalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MPCWalletServiceClient interface {
	// Creates an MPCWallet. The Device in the request must have been registered using MPCKeyService's
	// RegisterDevice before this method is called. Under the hood, this calls MPCKeyService's
	// CreateDeviceGroup with the appropriate parameters. After calling this, use MPCKeyService's
	// ListMPCOperations to poll for the pending CreateDeviceGroup operation, and use the WaaS SDK's
	// computeMPCOperation to complete the operation.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	CreateMPCWallet(ctx context.Context, in *CreateMPCWalletRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Retrieves an MPCWallet by resource name.
	GetMPCWallet(ctx context.Context, in *GetMPCWalletRequest, opts ...grpc.CallOption) (*MPCWallet, error)
	// Returns a list of MPCWallets in a Pool.
	ListMPCWallets(ctx context.Context, in *ListMPCWalletsRequest, opts ...grpc.CallOption) (*ListMPCWalletsResponse, error)
	// Generates an Address within an MPCWallet. The Address values generated are identical across
	// Networks of the same protocol family (e.g. EVM). So, for example, calling GenerateAddress twice
	// for networks/ethereum-mainnet, and then calling it twice more for networks/ethereum-goerli, will
	// result in two pairs of identical addresses on Ethereum Mainnet and Goerli.
	GenerateAddress(ctx context.Context, in *GenerateAddressRequest, opts ...grpc.CallOption) (*Address, error)
	// Retrieves an Address by resource name.
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*Address, error)
	// Returns a list of Addresses in an MPCWallet.
	ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error)
	// Returns a list of Balances.
	ListBalances(ctx context.Context, in *ListBalancesRequest, opts ...grpc.CallOption) (*ListBalancesResponse, error)
	// Returns a list of BalanceDetails.
	ListBalanceDetails(ctx context.Context, in *ListBalanceDetailsRequest, opts ...grpc.CallOption) (*ListBalanceDetailsResponse, error)
}

type mPCWalletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMPCWalletServiceClient(cc grpc.ClientConnInterface) MPCWalletServiceClient {
	return &mPCWalletServiceClient{cc}
}

func (c *mPCWalletServiceClient) CreateMPCWallet(ctx context.Context, in *CreateMPCWalletRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/CreateMPCWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) GetMPCWallet(ctx context.Context, in *GetMPCWalletRequest, opts ...grpc.CallOption) (*MPCWallet, error) {
	out := new(MPCWallet)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/GetMPCWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) ListMPCWallets(ctx context.Context, in *ListMPCWalletsRequest, opts ...grpc.CallOption) (*ListMPCWalletsResponse, error) {
	out := new(ListMPCWalletsResponse)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListMPCWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) GenerateAddress(ctx context.Context, in *GenerateAddressRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/GenerateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error) {
	out := new(ListAddressesResponse)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) ListBalances(ctx context.Context, in *ListBalancesRequest, opts ...grpc.CallOption) (*ListBalancesResponse, error) {
	out := new(ListBalancesResponse)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPCWalletServiceClient) ListBalanceDetails(ctx context.Context, in *ListBalanceDetailsRequest, opts ...grpc.CallOption) (*ListBalanceDetailsResponse, error) {
	out := new(ListBalanceDetailsResponse)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListBalanceDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MPCWalletServiceServer is the server API for MPCWalletService service.
// All implementations must embed UnimplementedMPCWalletServiceServer
// for forward compatibility
type MPCWalletServiceServer interface {
	// Creates an MPCWallet. The Device in the request must have been registered using MPCKeyService's
	// RegisterDevice before this method is called. Under the hood, this calls MPCKeyService's
	// CreateDeviceGroup with the appropriate parameters. After calling this, use MPCKeyService's
	// ListMPCOperations to poll for the pending CreateDeviceGroup operation, and use the WaaS SDK's
	// computeMPCOperation to complete the operation.
	// Note: because the creation of MPC operations is asynchronous, ListMPCOperations may return a
	// NOT_FOUND error immediately after calling this. To complete the operation, continue polling
	// ListMPCOperations even after it returns a NOT_FOUND error.
	CreateMPCWallet(context.Context, *CreateMPCWalletRequest) (*longrunning.Operation, error)
	// Retrieves an MPCWallet by resource name.
	GetMPCWallet(context.Context, *GetMPCWalletRequest) (*MPCWallet, error)
	// Returns a list of MPCWallets in a Pool.
	ListMPCWallets(context.Context, *ListMPCWalletsRequest) (*ListMPCWalletsResponse, error)
	// Generates an Address within an MPCWallet. The Address values generated are identical across
	// Networks of the same protocol family (e.g. EVM). So, for example, calling GenerateAddress twice
	// for networks/ethereum-mainnet, and then calling it twice more for networks/ethereum-goerli, will
	// result in two pairs of identical addresses on Ethereum Mainnet and Goerli.
	GenerateAddress(context.Context, *GenerateAddressRequest) (*Address, error)
	// Retrieves an Address by resource name.
	GetAddress(context.Context, *GetAddressRequest) (*Address, error)
	// Returns a list of Addresses in an MPCWallet.
	ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error)
	// Returns a list of Balances.
	ListBalances(context.Context, *ListBalancesRequest) (*ListBalancesResponse, error)
	// Returns a list of BalanceDetails.
	ListBalanceDetails(context.Context, *ListBalanceDetailsRequest) (*ListBalanceDetailsResponse, error)
	mustEmbedUnimplementedMPCWalletServiceServer()
}

// UnimplementedMPCWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMPCWalletServiceServer struct {
}

func (UnimplementedMPCWalletServiceServer) CreateMPCWallet(context.Context, *CreateMPCWalletRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMPCWallet not implemented")
}
func (UnimplementedMPCWalletServiceServer) GetMPCWallet(context.Context, *GetMPCWalletRequest) (*MPCWallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMPCWallet not implemented")
}
func (UnimplementedMPCWalletServiceServer) ListMPCWallets(context.Context, *ListMPCWalletsRequest) (*ListMPCWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMPCWallets not implemented")
}
func (UnimplementedMPCWalletServiceServer) GenerateAddress(context.Context, *GenerateAddressRequest) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAddress not implemented")
}
func (UnimplementedMPCWalletServiceServer) GetAddress(context.Context, *GetAddressRequest) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedMPCWalletServiceServer) ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddresses not implemented")
}
func (UnimplementedMPCWalletServiceServer) ListBalances(context.Context, *ListBalancesRequest) (*ListBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalances not implemented")
}
func (UnimplementedMPCWalletServiceServer) ListBalanceDetails(context.Context, *ListBalanceDetailsRequest) (*ListBalanceDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalanceDetails not implemented")
}
func (UnimplementedMPCWalletServiceServer) mustEmbedUnimplementedMPCWalletServiceServer() {}

// UnsafeMPCWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MPCWalletServiceServer will
// result in compilation errors.
type UnsafeMPCWalletServiceServer interface {
	mustEmbedUnimplementedMPCWalletServiceServer()
}

func RegisterMPCWalletServiceServer(s grpc.ServiceRegistrar, srv MPCWalletServiceServer) {
	s.RegisterService(&MPCWalletService_ServiceDesc, srv)
}

func _MPCWalletService_CreateMPCWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMPCWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).CreateMPCWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/CreateMPCWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).CreateMPCWallet(ctx, req.(*CreateMPCWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_GetMPCWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMPCWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).GetMPCWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/GetMPCWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).GetMPCWallet(ctx, req.(*GetMPCWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_ListMPCWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMPCWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).ListMPCWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListMPCWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).ListMPCWallets(ctx, req.(*ListMPCWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_GenerateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).GenerateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/GenerateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).GenerateAddress(ctx, req.(*GenerateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_ListAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).ListAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).ListAddresses(ctx, req.(*ListAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_ListBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).ListBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).ListBalances(ctx, req.(*ListBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPCWalletService_ListBalanceDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalanceDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCWalletServiceServer).ListBalanceDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.mpc_wallets.v1.MPCWalletService/ListBalanceDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCWalletServiceServer).ListBalanceDetails(ctx, req.(*ListBalanceDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MPCWalletService_ServiceDesc is the grpc.ServiceDesc for MPCWalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MPCWalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coinbase.cloud.mpc_wallets.v1.MPCWalletService",
	HandlerType: (*MPCWalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMPCWallet",
			Handler:    _MPCWalletService_CreateMPCWallet_Handler,
		},
		{
			MethodName: "GetMPCWallet",
			Handler:    _MPCWalletService_GetMPCWallet_Handler,
		},
		{
			MethodName: "ListMPCWallets",
			Handler:    _MPCWalletService_ListMPCWallets_Handler,
		},
		{
			MethodName: "GenerateAddress",
			Handler:    _MPCWalletService_GenerateAddress_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _MPCWalletService_GetAddress_Handler,
		},
		{
			MethodName: "ListAddresses",
			Handler:    _MPCWalletService_ListAddresses_Handler,
		},
		{
			MethodName: "ListBalances",
			Handler:    _MPCWalletService_ListBalances_Handler,
		},
		{
			MethodName: "ListBalanceDetails",
			Handler:    _MPCWalletService_ListBalanceDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coinbase/cloud/mpc_wallets/v1/mpc_wallets.proto",
}

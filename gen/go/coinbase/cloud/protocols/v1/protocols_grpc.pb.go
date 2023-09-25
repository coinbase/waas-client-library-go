// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProtocolServiceClient is the client API for ProtocolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtocolServiceClient interface {
	// Constructs an unsigned transaction. The payloads in the required_signatures of the
	// returned Transaction must be signed before the Transaction is broadcast.
	ConstructTransaction(ctx context.Context, in *ConstructTransactionRequest, opts ...grpc.CallOption) (*v1.Transaction, error)
	// Constructs an unsigned transfer transaction. A transfer transaction is a transaction that
	// moves an Asset from one Address to another. The payloads in the required_signatures of the
	// returned Transaction must be signed before the Transaction is broadcast.
	ConstructTransferTransaction(ctx context.Context, in *ConstructTransferTransactionRequest, opts ...grpc.CallOption) (*v1.Transaction, error)
	// Broadcasts a transaction to a node in the Network. There are two ways of invoking this API:
	// 1. Set the raw_signed_transaction on the Transaction. This is equivalent to the payload used to broadcast transactions
	// via block explorers such as Etherscan. In this case, the TransactionInput does not need to be set on the Transaction.
	// 2. Set the signature(s) in the required_signatures of the Transaction. In this case, the TransactionInput must also be
	// set on the Transaction.
	// The Transaction returned will have the hash set on it.
	BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*v1.Transaction, error)
	// Estimates the current network fee for the specified Network. For EVM Networks, this
	// corresponds to the gas_price, max_fee_per_gas, and max_priority_fee_per_gas.
	EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error)
}

type protocolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtocolServiceClient(cc grpc.ClientConnInterface) ProtocolServiceClient {
	return &protocolServiceClient{cc}
}

func (c *protocolServiceClient) ConstructTransaction(ctx context.Context, in *ConstructTransactionRequest, opts ...grpc.CallOption) (*v1.Transaction, error) {
	out := new(v1.Transaction)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.protocols.v1.ProtocolService/ConstructTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolServiceClient) ConstructTransferTransaction(ctx context.Context, in *ConstructTransferTransactionRequest, opts ...grpc.CallOption) (*v1.Transaction, error) {
	out := new(v1.Transaction)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.protocols.v1.ProtocolService/ConstructTransferTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolServiceClient) BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*v1.Transaction, error) {
	out := new(v1.Transaction)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.protocols.v1.ProtocolService/BroadcastTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolServiceClient) EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error) {
	out := new(EstimateFeeResponse)
	err := c.cc.Invoke(ctx, "/coinbase.cloud.protocols.v1.ProtocolService/EstimateFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtocolServiceServer is the server API for ProtocolService service.
// All implementations must embed UnimplementedProtocolServiceServer
// for forward compatibility
type ProtocolServiceServer interface {
	// Constructs an unsigned transaction. The payloads in the required_signatures of the
	// returned Transaction must be signed before the Transaction is broadcast.
	ConstructTransaction(context.Context, *ConstructTransactionRequest) (*v1.Transaction, error)
	// Constructs an unsigned transfer transaction. A transfer transaction is a transaction that
	// moves an Asset from one Address to another. The payloads in the required_signatures of the
	// returned Transaction must be signed before the Transaction is broadcast.
	ConstructTransferTransaction(context.Context, *ConstructTransferTransactionRequest) (*v1.Transaction, error)
	// Broadcasts a transaction to a node in the Network. There are two ways of invoking this API:
	// 1. Set the raw_signed_transaction on the Transaction. This is equivalent to the payload used to broadcast transactions
	// via block explorers such as Etherscan. In this case, the TransactionInput does not need to be set on the Transaction.
	// 2. Set the signature(s) in the required_signatures of the Transaction. In this case, the TransactionInput must also be
	// set on the Transaction.
	// The Transaction returned will have the hash set on it.
	BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*v1.Transaction, error)
	// Estimates the current network fee for the specified Network. For EVM Networks, this
	// corresponds to the gas_price, max_fee_per_gas, and max_priority_fee_per_gas.
	EstimateFee(context.Context, *EstimateFeeRequest) (*EstimateFeeResponse, error)
	mustEmbedUnimplementedProtocolServiceServer()
}

// UnimplementedProtocolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProtocolServiceServer struct {
}

func (UnimplementedProtocolServiceServer) ConstructTransaction(context.Context, *ConstructTransactionRequest) (*v1.Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConstructTransaction not implemented")
}
func (UnimplementedProtocolServiceServer) ConstructTransferTransaction(context.Context, *ConstructTransferTransactionRequest) (*v1.Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConstructTransferTransaction not implemented")
}
func (UnimplementedProtocolServiceServer) BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*v1.Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTransaction not implemented")
}
func (UnimplementedProtocolServiceServer) EstimateFee(context.Context, *EstimateFeeRequest) (*EstimateFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateFee not implemented")
}
func (UnimplementedProtocolServiceServer) mustEmbedUnimplementedProtocolServiceServer() {}

// UnsafeProtocolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtocolServiceServer will
// result in compilation errors.
type UnsafeProtocolServiceServer interface {
	mustEmbedUnimplementedProtocolServiceServer()
}

func RegisterProtocolServiceServer(s grpc.ServiceRegistrar, srv ProtocolServiceServer) {
	s.RegisterService(&ProtocolService_ServiceDesc, srv)
}

func _ProtocolService_ConstructTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConstructTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).ConstructTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.protocols.v1.ProtocolService/ConstructTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).ConstructTransaction(ctx, req.(*ConstructTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtocolService_ConstructTransferTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConstructTransferTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).ConstructTransferTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.protocols.v1.ProtocolService/ConstructTransferTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).ConstructTransferTransaction(ctx, req.(*ConstructTransferTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtocolService_BroadcastTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).BroadcastTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.protocols.v1.ProtocolService/BroadcastTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).BroadcastTransaction(ctx, req.(*BroadcastTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtocolService_EstimateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).EstimateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinbase.cloud.protocols.v1.ProtocolService/EstimateFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).EstimateFee(ctx, req.(*EstimateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProtocolService_ServiceDesc is the grpc.ServiceDesc for ProtocolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtocolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coinbase.cloud.protocols.v1.ProtocolService",
	HandlerType: (*ProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConstructTransaction",
			Handler:    _ProtocolService_ConstructTransaction_Handler,
		},
		{
			MethodName: "ConstructTransferTransaction",
			Handler:    _ProtocolService_ConstructTransferTransaction_Handler,
		},
		{
			MethodName: "BroadcastTransaction",
			Handler:    _ProtocolService_BroadcastTransaction_Handler,
		},
		{
			MethodName: "EstimateFee",
			Handler:    _ProtocolService_EstimateFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coinbase/cloud/protocols/v1/protocols.proto",
}

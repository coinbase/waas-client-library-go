// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: coinbase/cloud/protocols/v1/protocols.proto

package v1

import (
	v1 "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for ConstructTransaction.
type ConstructTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Network.
	// Format: networks/{network_id}
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// The input indicating the unsigned transaction to construct.
	Input *v1.TransactionInput `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *ConstructTransactionRequest) Reset() {
	*x = ConstructTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstructTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstructTransactionRequest) ProtoMessage() {}

func (x *ConstructTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstructTransactionRequest.ProtoReflect.Descriptor instead.
func (*ConstructTransactionRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_protocols_v1_protocols_proto_rawDescGZIP(), []int{0}
}

func (x *ConstructTransactionRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ConstructTransactionRequest) GetInput() *v1.TransactionInput {
	if x != nil {
		return x.Input
	}
	return nil
}

// The request message for ConstructTransferTransaction.
type ConstructTransferTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Network on which the transfer is happening.
	// Format: networks/{network_id}
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// The resource name of the Asset being transferred.
	// Format: networks/{network_id}/assets/{asset_id}.
	Asset string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	// The address of the sender, e.g. as a 0x-prefixed hex string.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// The address of the recipient, e.g. as a 0x-prefixed hex string.
	Recipient string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// The amount of the asset to transfer, denominated in atomic units of the asset (e.g. Wei),
	// either as a "0x"-prefixed hex string or a base-10 number.
	Amount string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	// The nonce of the from address to be used in transaction construction.
	// This is required only for Account-based networks (e.g. EVM).
	Nonce int64 `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The fee to be paid for the transfer. If not provided, the API will default to a network-based fee estimate.
	Fee *v1.TransactionFee `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *ConstructTransferTransactionRequest) Reset() {
	*x = ConstructTransferTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstructTransferTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstructTransferTransactionRequest) ProtoMessage() {}

func (x *ConstructTransferTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstructTransferTransactionRequest.ProtoReflect.Descriptor instead.
func (*ConstructTransferTransactionRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_protocols_v1_protocols_proto_rawDescGZIP(), []int{1}
}

func (x *ConstructTransferTransactionRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ConstructTransferTransactionRequest) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *ConstructTransferTransactionRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ConstructTransferTransactionRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ConstructTransferTransactionRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ConstructTransferTransactionRequest) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *ConstructTransferTransactionRequest) GetFee() *v1.TransactionFee {
	if x != nil {
		return x.Fee
	}
	return nil
}

// The request message for BroadcastTransaction. For the broadcast to complete successfully,
// the required_signatures' signature values must be populated (i.e. the transaction must be signed).
type BroadcastTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Network.
	// Format: networks/{network_id}
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// The transaction to broadcast. For the broadcast to complete successfully,
	// the required_signatures' signature values must be populated (i.e. the transaction must be signed).
	Transaction *v1.Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *BroadcastTransactionRequest) Reset() {
	*x = BroadcastTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTransactionRequest) ProtoMessage() {}

func (x *BroadcastTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastTransactionRequest.ProtoReflect.Descriptor instead.
func (*BroadcastTransactionRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_protocols_v1_protocols_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastTransactionRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *BroadcastTransactionRequest) GetTransaction() *v1.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

var File_coinbase_cloud_protocols_v1_protocols_proto protoreflect.FileDescriptor

var file_coinbase_cloud_protocols_v1_protocols_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x11, 0xca, 0x3e,
	0x0e, 0xfa, 0x02, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0xe2,
	0x41, 0x01, 0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x3f,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22,
	0xf4, 0x02, 0x0a, 0x23, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x11, 0xca, 0x3e, 0x0e,
	0xfa, 0x02, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0xe2, 0x41,
	0x01, 0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x3a, 0x0a,
	0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xe2, 0x41,
	0x01, 0x02, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02,
	0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x1b, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x11, 0xca, 0x3e, 0x0e, 0xfa,
	0x02, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0xe2, 0x41, 0x01,
	0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x46, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0xd9, 0x0a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xfa, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x81, 0x02, 0x92, 0x41, 0xb5, 0x01, 0x12, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x9b, 0x01, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x20, 0x61, 0x6e, 0x20,
	0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x69, 0x73, 0x20, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0xda, 0x41, 0x0d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x3d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x2a, 0x7d,
	0x3a, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x95, 0x04, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c,
	0x03, 0x92, 0x41, 0xa0, 0x02, 0x12, 0x1e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xfd, 0x01, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x41, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x61,
	0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x61,
	0x74, 0x20, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x20, 0x54, 0x68,
	0x65, 0x20, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x20, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0xda, 0x41, 0x25, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2c, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3a, 0x3a, 0x01, 0x2a, 0x22, 0x35, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x3d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x2a, 0x7d,
	0x3a, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x95, 0x02,
	0x0a, 0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9c, 0x01, 0x92, 0x41, 0x4b, 0x12, 0x15, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x32, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x20,
	0x61, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x20, 0x61, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0xda, 0x41, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x3d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x2a, 0x7d,
	0x3a, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x98, 0x01, 0x92, 0x41, 0x6e, 0x12, 0x6c, 0x41, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x69, 0x6e, 0x67,
	0x20, 0x61, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x65, 0x73, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0xca, 0x41, 0x24, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x42, 0xd6, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x62, 0x68, 0x71,
	0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2f, 0x76, 0x31, 0x92, 0x41, 0x81, 0x01, 0x12, 0x2a, 0x0a, 0x23, 0x43, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x20, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x2d, 0x20, 0x57, 0x61, 0x61,
	0x53, 0x20, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x20, 0x41, 0x50, 0x49, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x1a, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_coinbase_cloud_protocols_v1_protocols_proto_rawDescOnce sync.Once
	file_coinbase_cloud_protocols_v1_protocols_proto_rawDescData = file_coinbase_cloud_protocols_v1_protocols_proto_rawDesc
)

func file_coinbase_cloud_protocols_v1_protocols_proto_rawDescGZIP() []byte {
	file_coinbase_cloud_protocols_v1_protocols_proto_rawDescOnce.Do(func() {
		file_coinbase_cloud_protocols_v1_protocols_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_cloud_protocols_v1_protocols_proto_rawDescData)
	})
	return file_coinbase_cloud_protocols_v1_protocols_proto_rawDescData
}

var file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_cloud_protocols_v1_protocols_proto_goTypes = []interface{}{
	(*ConstructTransactionRequest)(nil),         // 0: coinbase.cloud.protocols.v1.ConstructTransactionRequest
	(*ConstructTransferTransactionRequest)(nil), // 1: coinbase.cloud.protocols.v1.ConstructTransferTransactionRequest
	(*BroadcastTransactionRequest)(nil),         // 2: coinbase.cloud.protocols.v1.BroadcastTransactionRequest
	(*v1.TransactionInput)(nil),                 // 3: coinbase.cloud.types.v1.TransactionInput
	(*v1.TransactionFee)(nil),                   // 4: coinbase.cloud.types.v1.TransactionFee
	(*v1.Transaction)(nil),                      // 5: coinbase.cloud.types.v1.Transaction
}
var file_coinbase_cloud_protocols_v1_protocols_proto_depIdxs = []int32{
	3, // 0: coinbase.cloud.protocols.v1.ConstructTransactionRequest.input:type_name -> coinbase.cloud.types.v1.TransactionInput
	4, // 1: coinbase.cloud.protocols.v1.ConstructTransferTransactionRequest.fee:type_name -> coinbase.cloud.types.v1.TransactionFee
	5, // 2: coinbase.cloud.protocols.v1.BroadcastTransactionRequest.transaction:type_name -> coinbase.cloud.types.v1.Transaction
	0, // 3: coinbase.cloud.protocols.v1.ProtocolService.ConstructTransaction:input_type -> coinbase.cloud.protocols.v1.ConstructTransactionRequest
	1, // 4: coinbase.cloud.protocols.v1.ProtocolService.ConstructTransferTransaction:input_type -> coinbase.cloud.protocols.v1.ConstructTransferTransactionRequest
	2, // 5: coinbase.cloud.protocols.v1.ProtocolService.BroadcastTransaction:input_type -> coinbase.cloud.protocols.v1.BroadcastTransactionRequest
	5, // 6: coinbase.cloud.protocols.v1.ProtocolService.ConstructTransaction:output_type -> coinbase.cloud.types.v1.Transaction
	5, // 7: coinbase.cloud.protocols.v1.ProtocolService.ConstructTransferTransaction:output_type -> coinbase.cloud.types.v1.Transaction
	5, // 8: coinbase.cloud.protocols.v1.ProtocolService.BroadcastTransaction:output_type -> coinbase.cloud.types.v1.Transaction
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coinbase_cloud_protocols_v1_protocols_proto_init() }
func file_coinbase_cloud_protocols_v1_protocols_proto_init() {
	if File_coinbase_cloud_protocols_v1_protocols_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstructTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstructTransferTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_cloud_protocols_v1_protocols_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coinbase_cloud_protocols_v1_protocols_proto_goTypes,
		DependencyIndexes: file_coinbase_cloud_protocols_v1_protocols_proto_depIdxs,
		MessageInfos:      file_coinbase_cloud_protocols_v1_protocols_proto_msgTypes,
	}.Build()
	File_coinbase_cloud_protocols_v1_protocols_proto = out.File
	file_coinbase_cloud_protocols_v1_protocols_proto_rawDesc = nil
	file_coinbase_cloud_protocols_v1_protocols_proto_goTypes = nil
	file_coinbase_cloud_protocols_v1_protocols_proto_depIdxs = nil
}

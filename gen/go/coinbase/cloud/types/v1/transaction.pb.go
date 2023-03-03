// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: coinbase/cloud/types/v1/transaction.proto

package v1

import (
	v1 "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	v11 "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/crypto/types/v1"
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

// A message that wraps protocol-specific transactions, and provides additional
// payloads that enable the transaction to be signed and broadcast.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A user-specified transaction in one of the supported input types.
	Input *TransactionInput `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// The set of signatures required to broadcast this Transaction.
	// For EVM networks, there will be only one signature.
	RequiredSignatures []*RequiredSignature `protobuf:"bytes,2,rep,name=required_signatures,json=requiredSignatures,proto3" json:"required_signatures,omitempty"`
	// The signed transaction to be broadcast. This value can be directly broadcast on-chain.
	// The Signatures in required_signatures need to be transformed into this value.
	// For EVM chains, this is just the R, S, and V values of the signature concatenated.
	RawSignedTransaction []byte `protobuf:"bytes,3,opt,name=raw_signed_transaction,json=rawSignedTransaction,proto3" json:"raw_signed_transaction,omitempty"`
	// The hash of the signed transaction.
	Hash string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_types_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetInput() *TransactionInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Transaction) GetRequiredSignatures() []*RequiredSignature {
	if x != nil {
		return x.RequiredSignatures
	}
	return nil
}

func (x *Transaction) GetRawSignedTransaction() []byte {
	if x != nil {
		return x.RawSignedTransaction
	}
	return nil
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// A message that contains each of the different transaction types we support.
type TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Input:
	//	*TransactionInput_EthereumRlpInput
	//	*TransactionInput_Ethereum_1559Input
	Input isTransactionInput_Input `protobuf_oneof:"input"`
}

func (x *TransactionInput) Reset() {
	*x = TransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInput) ProtoMessage() {}

func (x *TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInput.ProtoReflect.Descriptor instead.
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_types_v1_transaction_proto_rawDescGZIP(), []int{1}
}

func (m *TransactionInput) GetInput() isTransactionInput_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *TransactionInput) GetEthereumRlpInput() *v1.RLPTransaction {
	if x, ok := x.GetInput().(*TransactionInput_EthereumRlpInput); ok {
		return x.EthereumRlpInput
	}
	return nil
}

func (x *TransactionInput) GetEthereum_1559Input() *v1.EIP1559TransactionInput {
	if x, ok := x.GetInput().(*TransactionInput_Ethereum_1559Input); ok {
		return x.Ethereum_1559Input
	}
	return nil
}

type isTransactionInput_Input interface {
	isTransactionInput_Input()
}

type TransactionInput_EthereumRlpInput struct {
	// An Ethereum RLP transaction.
	EthereumRlpInput *v1.RLPTransaction `protobuf:"bytes,1,opt,name=ethereum_rlp_input,json=ethereumRlpInput,proto3,oneof"`
}

type TransactionInput_Ethereum_1559Input struct {
	// An EIP-1559 transaction.
	Ethereum_1559Input *v1.EIP1559TransactionInput `protobuf:"bytes,2,opt,name=ethereum_1559_input,json=ethereum1559Input,proto3,oneof"`
}

func (*TransactionInput_EthereumRlpInput) isTransactionInput_Input() {}

func (*TransactionInput_Ethereum_1559Input) isTransactionInput_Input() {}

// A message representing a payload to be signed and its result.
type RequiredSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payload to be signed. The TransactionInput must be transformed into
	// this payload. For EVM chains, this would be the hash of the type-prefixed
	// RLP encoding of the transaction defined in EIP-2718. This is also the
	// value that should be used in the MPCKeyService's CreateSignature API.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// The signature representing the signed payload. This must be set in order
	// for a Transaction to be broadcast.
	Signature *v11.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *RequiredSignature) Reset() {
	*x = RequiredSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredSignature) ProtoMessage() {}

func (x *RequiredSignature) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredSignature.ProtoReflect.Descriptor instead.
func (*RequiredSignature) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_types_v1_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *RequiredSignature) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *RequiredSignature) GetSignature() *v11.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// A message representing the fee to be paid for a transaction.
type TransactionFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The fee to be paid.
	//
	// Types that are assignable to Fee:
	//	*TransactionFee_EthereumFee
	Fee isTransactionFee_Fee `protobuf_oneof:"fee"`
}

func (x *TransactionFee) Reset() {
	*x = TransactionFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionFee) ProtoMessage() {}

func (x *TransactionFee) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_types_v1_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionFee.ProtoReflect.Descriptor instead.
func (*TransactionFee) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_types_v1_transaction_proto_rawDescGZIP(), []int{3}
}

func (m *TransactionFee) GetFee() isTransactionFee_Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (x *TransactionFee) GetEthereumFee() *v1.DynamicFeeInput {
	if x, ok := x.GetFee().(*TransactionFee_EthereumFee); ok {
		return x.EthereumFee
	}
	return nil
}

type isTransactionFee_Fee interface {
	isTransactionFee_Fee()
}

type TransactionFee_EthereumFee struct {
	// An EIP-1559 fee on EVM networks.
	EthereumFee *v1.DynamicFeeInput `protobuf:"bytes,1,opt,name=ethereum_fee,json=ethereumFee,proto3,oneof"`
}

func (*TransactionFee_EthereumFee) isTransactionFee_Fee() {}

var File_coinbase_cloud_types_v1_transaction_proto protoreflect.FileDescriptor

var file_coinbase_cloud_types_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x31, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x5b, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x12, 0x34, 0x0a, 0x16, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x14, 0x72, 0x61, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xf2, 0x01, 0x0a, 0x10, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x64, 0x0a, 0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x72, 0x6c, 0x70, 0x5f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x4c, 0x50, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x10, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x52, 0x6c, 0x70,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6f, 0x0a, 0x13, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x5f, 0x31, 0x35, 0x35, 0x39, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x49, 0x50, 0x31, 0x35, 0x35,
	0x39, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x48, 0x00, 0x52, 0x11, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x31, 0x35, 0x35,
	0x39, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22,
	0x70, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x73, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x65, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x65, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x48, 0x00, 0x52, 0x0b, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x46, 0x65, 0x65, 0x42,
	0x05, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x62, 0x68, 0x71, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x77, 0x61, 0x61, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_cloud_types_v1_transaction_proto_rawDescOnce sync.Once
	file_coinbase_cloud_types_v1_transaction_proto_rawDescData = file_coinbase_cloud_types_v1_transaction_proto_rawDesc
)

func file_coinbase_cloud_types_v1_transaction_proto_rawDescGZIP() []byte {
	file_coinbase_cloud_types_v1_transaction_proto_rawDescOnce.Do(func() {
		file_coinbase_cloud_types_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_cloud_types_v1_transaction_proto_rawDescData)
	})
	return file_coinbase_cloud_types_v1_transaction_proto_rawDescData
}

var file_coinbase_cloud_types_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coinbase_cloud_types_v1_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),                // 0: coinbase.cloud.types.v1.Transaction
	(*TransactionInput)(nil),           // 1: coinbase.cloud.types.v1.TransactionInput
	(*RequiredSignature)(nil),          // 2: coinbase.cloud.types.v1.RequiredSignature
	(*TransactionFee)(nil),             // 3: coinbase.cloud.types.v1.TransactionFee
	(*v1.RLPTransaction)(nil),          // 4: coinbase.cloud.protocols.ethereum.v1.RLPTransaction
	(*v1.EIP1559TransactionInput)(nil), // 5: coinbase.cloud.protocols.ethereum.v1.EIP1559TransactionInput
	(*v11.Signature)(nil),              // 6: coinbase.crypto.types.v1.Signature
	(*v1.DynamicFeeInput)(nil),         // 7: coinbase.cloud.protocols.ethereum.v1.DynamicFeeInput
}
var file_coinbase_cloud_types_v1_transaction_proto_depIdxs = []int32{
	1, // 0: coinbase.cloud.types.v1.Transaction.input:type_name -> coinbase.cloud.types.v1.TransactionInput
	2, // 1: coinbase.cloud.types.v1.Transaction.required_signatures:type_name -> coinbase.cloud.types.v1.RequiredSignature
	4, // 2: coinbase.cloud.types.v1.TransactionInput.ethereum_rlp_input:type_name -> coinbase.cloud.protocols.ethereum.v1.RLPTransaction
	5, // 3: coinbase.cloud.types.v1.TransactionInput.ethereum_1559_input:type_name -> coinbase.cloud.protocols.ethereum.v1.EIP1559TransactionInput
	6, // 4: coinbase.cloud.types.v1.RequiredSignature.signature:type_name -> coinbase.crypto.types.v1.Signature
	7, // 5: coinbase.cloud.types.v1.TransactionFee.ethereum_fee:type_name -> coinbase.cloud.protocols.ethereum.v1.DynamicFeeInput
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_coinbase_cloud_types_v1_transaction_proto_init() }
func file_coinbase_cloud_types_v1_transaction_proto_init() {
	if File_coinbase_cloud_types_v1_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_cloud_types_v1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_coinbase_cloud_types_v1_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionInput); i {
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
		file_coinbase_cloud_types_v1_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredSignature); i {
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
		file_coinbase_cloud_types_v1_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionFee); i {
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
	file_coinbase_cloud_types_v1_transaction_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TransactionInput_EthereumRlpInput)(nil),
		(*TransactionInput_Ethereum_1559Input)(nil),
	}
	file_coinbase_cloud_types_v1_transaction_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TransactionFee_EthereumFee)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_cloud_types_v1_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_cloud_types_v1_transaction_proto_goTypes,
		DependencyIndexes: file_coinbase_cloud_types_v1_transaction_proto_depIdxs,
		MessageInfos:      file_coinbase_cloud_types_v1_transaction_proto_msgTypes,
	}.Build()
	File_coinbase_cloud_types_v1_transaction_proto = out.File
	file_coinbase_cloud_types_v1_transaction_proto_rawDesc = nil
	file_coinbase_cloud_types_v1_transaction_proto_goTypes = nil
	file_coinbase_cloud_types_v1_transaction_proto_depIdxs = nil
}

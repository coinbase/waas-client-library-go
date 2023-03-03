// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: coinbase/cloud/protocols/ethereum/v1/inputs.proto

package v1

import (
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

// A message representing an EIP-1559 transaction input.
type EIP1559TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chain ID of the transaction either as a "0x"-prefixed hex string or a base-10 number.
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// The nonce of the transaction. This value may be ignored depending on the API.
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The EIP-1559 maximum priority fee per gas either as a "0x"-prefixed hex string or a base-10 number.
	MaxPriorityFeePerGas string `protobuf:"bytes,3,opt,name=max_priority_fee_per_gas,json=maxPriorityFeePerGas,proto3" json:"max_priority_fee_per_gas,omitempty"`
	// The EIP-1559 maximum fee per gas either as a "0x"-prefixed hex string or a base-10 number.
	MaxFeePerGas string `protobuf:"bytes,4,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
	// The maximum amount of gas to use on the transaction.
	Gas uint64 `protobuf:"varint,5,opt,name=gas,proto3" json:"gas,omitempty"`
	// The checksummed address from which the transaction will originate, as a "0x"-prefixed hex string.
	// Note: This is NOT a WaaS Address resource of the form
	// networks/{networkID}/addresses/{addressID}.
	FromAddress string `protobuf:"bytes,6,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// The checksummed address to which the transaction is addressed, as a "0x"-prefixed hex string.
	// Note: This is NOT a WaaS Address resource of the form
	// networks/{networkID}/addresses/{addressID}.
	ToAddress string `protobuf:"bytes,7,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// The native value of the transaction as a "0x"-prefixed hex string or a base-10 number.
	Value string `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	// The data for the transaction.
	Data []byte `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EIP1559TransactionInput) Reset() {
	*x = EIP1559TransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EIP1559TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EIP1559TransactionInput) ProtoMessage() {}

func (x *EIP1559TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EIP1559TransactionInput.ProtoReflect.Descriptor instead.
func (*EIP1559TransactionInput) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescGZIP(), []int{0}
}

func (x *EIP1559TransactionInput) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *EIP1559TransactionInput) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *EIP1559TransactionInput) GetMaxPriorityFeePerGas() string {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return ""
}

func (x *EIP1559TransactionInput) GetMaxFeePerGas() string {
	if x != nil {
		return x.MaxFeePerGas
	}
	return ""
}

func (x *EIP1559TransactionInput) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *EIP1559TransactionInput) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *EIP1559TransactionInput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *EIP1559TransactionInput) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *EIP1559TransactionInput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// A message representing a transaction encoded using Recursive-Length Prefix (RLP)
// Serialization. See https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/
// for more details.
type RLPTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sender address as a hexadecimal string, prefixed with 0x.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// The RLP-encoded transaction as a byte array.
	// See https://github.com/ethereum/go-ethereum/blob/master/rlp/encode.go for an implementation.
	TransactionRlp []byte `protobuf:"bytes,2,opt,name=transaction_rlp,json=transactionRlp,proto3" json:"transaction_rlp,omitempty"`
}

func (x *RLPTransaction) Reset() {
	*x = RLPTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RLPTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RLPTransaction) ProtoMessage() {}

func (x *RLPTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RLPTransaction.ProtoReflect.Descriptor instead.
func (*RLPTransaction) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescGZIP(), []int{1}
}

func (x *RLPTransaction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *RLPTransaction) GetTransactionRlp() []byte {
	if x != nil {
		return x.TransactionRlp
	}
	return nil
}

// A message representing fee information on EIP-1559 supported networks.
type DynamicFeeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The EIP-1559 maximum priority fee per gas either as a "0x"-prefixed hex string or a base-10 number.
	MaxPriorityFeePerGas string `protobuf:"bytes,1,opt,name=max_priority_fee_per_gas,json=maxPriorityFeePerGas,proto3" json:"max_priority_fee_per_gas,omitempty"`
	// The EIP-1559 maximum fee per gas either as a "0x"-prefixed hex string or a base-10 number.
	MaxFeePerGas string `protobuf:"bytes,2,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
}

func (x *DynamicFeeInput) Reset() {
	*x = DynamicFeeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicFeeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicFeeInput) ProtoMessage() {}

func (x *DynamicFeeInput) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicFeeInput.ProtoReflect.Descriptor instead.
func (*DynamicFeeInput) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescGZIP(), []int{2}
}

func (x *DynamicFeeInput) GetMaxPriorityFeePerGas() string {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return ""
}

func (x *DynamicFeeInput) GetMaxFeePerGas() string {
	if x != nil {
		return x.MaxFeePerGas
	}
	return ""
}

var File_coinbase_cloud_protocols_ethereum_v1_inputs_proto protoreflect.FileDescriptor

var file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0xa7, 0x02, 0x0a, 0x17, 0x45, 0x49,
	0x50, 0x31, 0x35, 0x35, 0x39, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67,
	0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x25,
	0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46, 0x65, 0x65, 0x50,
	0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x0e, 0x52, 0x4c, 0x50, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6c, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x6c, 0x70, 0x22, 0x70, 0x0a, 0x0f, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x46, 0x65, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x18, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x61, 0x78,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61,
	0x73, 0x12, 0x25, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46,
	0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x62, 0x68, 0x71, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescOnce sync.Once
	file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescData = file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDesc
)

func file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescGZIP() []byte {
	file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescOnce.Do(func() {
		file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescData)
	})
	return file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDescData
}

var file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_goTypes = []interface{}{
	(*EIP1559TransactionInput)(nil), // 0: coinbase.cloud.protocols.ethereum.v1.EIP1559TransactionInput
	(*RLPTransaction)(nil),          // 1: coinbase.cloud.protocols.ethereum.v1.RLPTransaction
	(*DynamicFeeInput)(nil),         // 2: coinbase.cloud.protocols.ethereum.v1.DynamicFeeInput
}
var file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_init() }
func file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_init() {
	if File_coinbase_cloud_protocols_ethereum_v1_inputs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EIP1559TransactionInput); i {
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
		file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RLPTransaction); i {
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
		file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicFeeInput); i {
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
			RawDescriptor: file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_goTypes,
		DependencyIndexes: file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_depIdxs,
		MessageInfos:      file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_msgTypes,
	}.Build()
	File_coinbase_cloud_protocols_ethereum_v1_inputs_proto = out.File
	file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_rawDesc = nil
	file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_goTypes = nil
	file_coinbase_cloud_protocols_ethereum_v1_inputs_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: coinbase/crypto/types/keys.proto

package types

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

// An enumeration of supported elliptic curve types.
type EllipticCurve int32

const (
	// The default unspecified value.
	EllipticCurve_ELLIPTIC_CURVE_UNSPECIFIED EllipticCurve = 0
	// The curve used by Bitcoin and Ethereum.
	EllipticCurve_SECP256K1 EllipticCurve = 1
)

// Enum value maps for EllipticCurve.
var (
	EllipticCurve_name = map[int32]string{
		0: "ELLIPTIC_CURVE_UNSPECIFIED",
		1: "SECP256K1",
	}
	EllipticCurve_value = map[string]int32{
		"ELLIPTIC_CURVE_UNSPECIFIED": 0,
		"SECP256K1":                  1,
	}
)

func (x EllipticCurve) Enum() *EllipticCurve {
	p := new(EllipticCurve)
	*p = x
	return p
}

func (x EllipticCurve) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EllipticCurve) Descriptor() protoreflect.EnumDescriptor {
	return file_coinbase_crypto_types_keys_proto_enumTypes[0].Descriptor()
}

func (EllipticCurve) Type() protoreflect.EnumType {
	return &file_coinbase_crypto_types_keys_proto_enumTypes[0]
}

func (x EllipticCurve) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EllipticCurve.Descriptor instead.
func (EllipticCurve) EnumDescriptor() ([]byte, []int) {
	return file_coinbase_crypto_types_keys_proto_rawDescGZIP(), []int{0}
}

// A message representing a public key.
type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The public key contents.
	//
	// Types that are assignable to PublicKey:
	//	*PublicKey_EllipticPublicKey
	PublicKey isPublicKey_PublicKey `protobuf_oneof:"public_key"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_types_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_types_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_types_keys_proto_rawDescGZIP(), []int{0}
}

func (m *PublicKey) GetPublicKey() isPublicKey_PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (x *PublicKey) GetEllipticPublicKey() *EllipticPublicKey {
	if x, ok := x.GetPublicKey().(*PublicKey_EllipticPublicKey); ok {
		return x.EllipticPublicKey
	}
	return nil
}

type isPublicKey_PublicKey interface {
	isPublicKey_PublicKey()
}

type PublicKey_EllipticPublicKey struct {
	// An Elliptic Curve Cryptography (ECC) public key.
	EllipticPublicKey *EllipticPublicKey `protobuf:"bytes,1,opt,name=elliptic_public_key,json=ellipticPublicKey,proto3,oneof"`
}

func (*PublicKey_EllipticPublicKey) isPublicKey_PublicKey() {}

// An Elliptic Curve Cryptography (ECC) public key.
type EllipticPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The elliptic curve used by this public key.
	EllipticCurve EllipticCurve `protobuf:"varint,1,opt,name=elliptic_curve,json=ellipticCurve,proto3,enum=coinbase.crypto.types.EllipticCurve" json:"elliptic_curve,omitempty"`
	// The public key in compressed form as defined in the Standards for Efficient
	// Cryptography (SEC) 1, Version 2.0, Section 2.3.3.
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *EllipticPublicKey) Reset() {
	*x = EllipticPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_types_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EllipticPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EllipticPublicKey) ProtoMessage() {}

func (x *EllipticPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_types_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EllipticPublicKey.ProtoReflect.Descriptor instead.
func (*EllipticPublicKey) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_types_keys_proto_rawDescGZIP(), []int{1}
}

func (x *EllipticPublicKey) GetEllipticCurve() EllipticCurve {
	if x != nil {
		return x.EllipticCurve
	}
	return EllipticCurve_ELLIPTIC_CURVE_UNSPECIFIED
}

func (x *EllipticPublicKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

var File_coinbase_crypto_types_keys_proto protoreflect.FileDescriptor

var file_coinbase_crypto_types_keys_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x75, 0x0a, 0x09, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x5a, 0x0a, 0x13, 0x65, 0x6c, 0x6c, 0x69, 0x70, 0x74,
	0x69, 0x63, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6c, 0x6c, 0x69,
	0x70, 0x74, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52,
	0x11, 0x65, 0x6c, 0x6c, 0x69, 0x70, 0x74, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x22, 0x7f, 0x0a, 0x11, 0x45, 0x6c, 0x6c, 0x69, 0x70, 0x74, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x4b, 0x0a, 0x0e, 0x65, 0x6c, 0x6c, 0x69, 0x70, 0x74, 0x69,
	0x63, 0x5f, 0x63, 0x75, 0x72, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6c, 0x6c, 0x69, 0x70, 0x74, 0x69, 0x63, 0x43, 0x75,
	0x72, 0x76, 0x65, 0x52, 0x0d, 0x65, 0x6c, 0x6c, 0x69, 0x70, 0x74, 0x69, 0x63, 0x43, 0x75, 0x72,
	0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x2a, 0x3e, 0x0a, 0x0d, 0x45, 0x6c, 0x6c, 0x69, 0x70, 0x74, 0x69, 0x63, 0x43, 0x75, 0x72,
	0x76, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x4c, 0x4c, 0x49, 0x50, 0x54, 0x49, 0x43, 0x5f, 0x43,
	0x55, 0x52, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x50, 0x32, 0x35, 0x36, 0x4b, 0x31, 0x10,
	0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x62, 0x68, 0x71,
	0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_types_keys_proto_rawDescOnce sync.Once
	file_coinbase_crypto_types_keys_proto_rawDescData = file_coinbase_crypto_types_keys_proto_rawDesc
)

func file_coinbase_crypto_types_keys_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_types_keys_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_types_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_types_keys_proto_rawDescData)
	})
	return file_coinbase_crypto_types_keys_proto_rawDescData
}

var file_coinbase_crypto_types_keys_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_coinbase_crypto_types_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coinbase_crypto_types_keys_proto_goTypes = []interface{}{
	(EllipticCurve)(0),        // 0: coinbase.crypto.types.EllipticCurve
	(*PublicKey)(nil),         // 1: coinbase.crypto.types.PublicKey
	(*EllipticPublicKey)(nil), // 2: coinbase.crypto.types.EllipticPublicKey
}
var file_coinbase_crypto_types_keys_proto_depIdxs = []int32{
	2, // 0: coinbase.crypto.types.PublicKey.elliptic_public_key:type_name -> coinbase.crypto.types.EllipticPublicKey
	0, // 1: coinbase.crypto.types.EllipticPublicKey.elliptic_curve:type_name -> coinbase.crypto.types.EllipticCurve
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_types_keys_proto_init() }
func file_coinbase_crypto_types_keys_proto_init() {
	if File_coinbase_crypto_types_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_types_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_coinbase_crypto_types_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EllipticPublicKey); i {
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
	file_coinbase_crypto_types_keys_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PublicKey_EllipticPublicKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_crypto_types_keys_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_types_keys_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_types_keys_proto_depIdxs,
		EnumInfos:         file_coinbase_crypto_types_keys_proto_enumTypes,
		MessageInfos:      file_coinbase_crypto_types_keys_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_types_keys_proto = out.File
	file_coinbase_crypto_types_keys_proto_rawDesc = nil
	file_coinbase_crypto_types_keys_proto_goTypes = nil
	file_coinbase_crypto_types_keys_proto_depIdxs = nil
}

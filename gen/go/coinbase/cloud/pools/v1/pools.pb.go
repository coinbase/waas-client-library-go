// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: coinbase/cloud/pools/v1/pools.proto

package v1

import (
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

// The Pool resource, which is a top-level container for segregating the
// resources under it via authorization checks.
type Pool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Pool.
	// Format: pools/{poolID}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A user-readable name for the Pool.
	// Example: 'Acme Co. Retail Trading'
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Pool) Reset() {
	*x = Pool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

func (x *Pool) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_pools_v1_pools_proto_rawDescGZIP(), []int{0}
}

func (x *Pool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pool) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// The request message for CreatePool.
type CreatePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Pool to create.
	Pool *Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	// Optional: the ID to use for the Pool, which will become the final component
	// of the Pool's resource name. If not provided, the server will assign a Pool ID
	// automatically.
	PoolId string `protobuf:"bytes,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (x *CreatePoolRequest) Reset() {
	*x = CreatePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePoolRequest) ProtoMessage() {}

func (x *CreatePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePoolRequest.ProtoReflect.Descriptor instead.
func (*CreatePoolRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_pools_v1_pools_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePoolRequest) GetPool() *Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

func (x *CreatePoolRequest) GetPoolId() string {
	if x != nil {
		return x.PoolId
	}
	return ""
}

// The request message for GetPool.
type GetPoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Pool to get.
	// Format: pools/{poolID}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPoolRequest) Reset() {
	*x = GetPoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoolRequest) ProtoMessage() {}

func (x *GetPoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoolRequest.ProtoReflect.Descriptor instead.
func (*GetPoolRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_pools_v1_pools_proto_rawDescGZIP(), []int{2}
}

func (x *GetPoolRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message for ListPools.
type ListPoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of Pools to return. If unspecified, at most 50 Pools
	// will be returned.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous ListPools call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPoolsRequest) Reset() {
	*x = ListPoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoolsRequest) ProtoMessage() {}

func (x *ListPoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoolsRequest.ProtoReflect.Descriptor instead.
func (*ListPoolsRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_pools_v1_pools_proto_rawDescGZIP(), []int{3}
}

func (x *ListPoolsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPoolsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response message for ListPools.
type ListPoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of Pools.
	Pools []*Pool `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// A token, which can be sent as page_token to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPoolsResponse) Reset() {
	*x = ListPoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoolsResponse) ProtoMessage() {}

func (x *ListPoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_cloud_pools_v1_pools_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoolsResponse.ProtoReflect.Descriptor instead.
func (*ListPoolsResponse) Descriptor() ([]byte, []int) {
	return file_coinbase_cloud_pools_v1_pools_proto_rawDescGZIP(), []int{4}
}

func (x *ListPoolsResponse) GetPools() []*Pool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *ListPoolsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_coinbase_cloud_pools_v1_pools_proto protoreflect.FileDescriptor

var file_coinbase_cloud_pools_v1_pools_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x04, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x58, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x44, 0x92, 0x41, 0x3d, 0x2a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x35, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x50, 0x6f,
	0x6f, 0x6c, 0x2e, 0x20, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a, 0x20, 0x70, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x7d, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x5b, 0x92, 0x41, 0x54, 0x2a,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x45, 0x41, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x6e, 0x61,
	0x6d, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x2e,
	0x20, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x20, 0x27, 0x41, 0x63, 0x6d, 0x65, 0x20,
	0x43, 0x6f, 0x2e, 0x20, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x54, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x27, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x3a, 0xb7, 0x01, 0x92, 0x41, 0x7c, 0x0a, 0x7a, 0x2a, 0x04, 0x50, 0x6f,
	0x6f, 0x6c, 0x32, 0x72, 0x54, 0x68, 0x65, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x69, 0x73, 0x20,
	0x61, 0x20, 0x74, 0x6f, 0x70, 0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x20, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x69, 0x74, 0x20, 0x76, 0x69, 0x61,
	0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0xea, 0x41, 0x35, 0x0a, 0x25, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x50, 0x6f, 0x6f, 0x6c,
	0x12, 0x0c, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6f, 0x6c, 0x7d, 0x22, 0xf4,
	0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f,
	0x6c, 0x42, 0x1c, 0x92, 0x41, 0x15, 0x32, 0x13, 0x54, 0x68, 0x65, 0x20, 0x50, 0x6f, 0x6f, 0x6c,
	0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0xce, 0x01, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb4, 0x01, 0x92, 0x41, 0xac, 0x01, 0x32, 0xa9,
	0x01, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x3a, 0x20, 0x74, 0x68, 0x65, 0x20, 0x49,
	0x44, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x77, 0x69, 0x6c,
	0x6c, 0x20, 0x62, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x27, 0x73, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x20, 0x61, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x20, 0x49, 0x44, 0x20, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x2e, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x06,
	0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x3a, 0x3d, 0x92, 0x41, 0x3a, 0x0a, 0x38, 0x2a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x32, 0x23, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x2e, 0x22, 0xdd, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x91, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x7d, 0x92, 0x41, 0x4c, 0x32, 0x3c, 0x54, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x67,
	0x65, 0x74, 0x2e, 0x20, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a, 0x20, 0x70, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x7d, 0xca, 0x3e, 0x0b, 0xfa, 0x02, 0x08,
	0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x27, 0x0a,
	0x25, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x37, 0x92, 0x41,
	0x34, 0x0a, 0x32, 0x2a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x20, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6f, 0x6c, 0x2e, 0x22, 0xd7, 0x02, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x7b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x5e, 0x92,
	0x41, 0x5b, 0x32, 0x59, 0x54, 0x68, 0x65, 0x20, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x20,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x20,
	0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x75, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2c, 0x20, 0x61, 0x74, 0x20, 0x6d, 0x6f,
	0x73, 0x74, 0x20, 0x35, 0x30, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x20, 0x77, 0x69, 0x6c, 0x6c,
	0x20, 0x62, 0x65, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x2e, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x69, 0x92, 0x41,
	0x66, 0x32, 0x64, 0x41, 0x20, 0x70, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c,
	0x20, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61,
	0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x20, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x75, 0x62, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x74, 0x20, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x3a, 0x3b, 0x92, 0x41, 0x38, 0x0a, 0x36, 0x2a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x22, 0x54, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x22,
	0xcd, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x6f, 0x6c, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12, 0x54, 0x68, 0x65, 0x20, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x52, 0x05, 0x70, 0x6f,
	0x6f, 0x6c, 0x73, 0x12, 0xaa, 0x01, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x81, 0x01,
	0x92, 0x41, 0x7e, 0x32, 0x7c, 0x41, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x77, 0x68,
	0x69, 0x63, 0x68, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62, 0x65, 0x20, 0x73, 0x65, 0x6e, 0x74, 0x20,
	0x61, 0x73, 0x20, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x65,
	0x78, 0x74, 0x20, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x69, 0x73, 0x20, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x72, 0x65, 0x20, 0x61, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x20,
	0x73, 0x75, 0x62, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x3a, 0x3d, 0x92, 0x41, 0x3a, 0x0a, 0x38, 0x2a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x23, 0x54, 0x68, 0x65, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x32,
	0xa7, 0x06, 0x0a, 0x0b, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xe1, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x2a,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x87, 0x01, 0x92, 0x41, 0x5e, 0x12,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x1a, 0x50, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x2e, 0x20, 0x43, 0x61, 0x6c,
	0x6c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x20, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e,
	0x79, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x2e, 0xda, 0x41, 0x0c,
	0x70, 0x6f, 0x6f, 0x6c, 0x2c, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x3a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x6f, 0x6c, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12,
	0x27, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x3b, 0x92, 0x41, 0x17, 0x12, 0x07, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x1a, 0x0c, 0x47, 0x65, 0x74, 0x73, 0x20, 0x61, 0x20, 0x50, 0x6f,
	0x6f, 0x6c, 0x2e, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x92, 0x41, 0x19, 0x12, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x1a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x1a, 0x8e, 0x02, 0x92, 0x41, 0xe7, 0x01, 0x12,
	0xe4, 0x01, 0x41, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x50, 0x6f, 0x6f,
	0x6c, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x20, 0x41, 0x20, 0x50, 0x6f,
	0x6f, 0x6c, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x70, 0x2d, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x73, 0x65, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20,
	0x69, 0x74, 0x20, 0x76, 0x69, 0x61, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x20, 0x50, 0x6f, 0x6f,
	0x6c, 0x2d, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x20, 0x61, 0x20, 0x50, 0x6f, 0x6f,
	0x6c, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x74, 0x68, 0x65, 0x79, 0x20, 0x74, 0x68, 0x65, 0x6d,
	0x73, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62, 0x65, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0xca, 0x41, 0x20, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x61, 0x61, 0x73, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x42, 0xc9, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x62, 0x68, 0x71, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x92, 0x41, 0x79, 0x12, 0x26, 0x0a, 0x1f,
	0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x20, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x2d,
	0x20, 0x57, 0x61, 0x61, 0x53, 0x20, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x20, 0x41, 0x50, 0x49, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x1a, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2a, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_cloud_pools_v1_pools_proto_rawDescOnce sync.Once
	file_coinbase_cloud_pools_v1_pools_proto_rawDescData = file_coinbase_cloud_pools_v1_pools_proto_rawDesc
)

func file_coinbase_cloud_pools_v1_pools_proto_rawDescGZIP() []byte {
	file_coinbase_cloud_pools_v1_pools_proto_rawDescOnce.Do(func() {
		file_coinbase_cloud_pools_v1_pools_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_cloud_pools_v1_pools_proto_rawDescData)
	})
	return file_coinbase_cloud_pools_v1_pools_proto_rawDescData
}

var file_coinbase_cloud_pools_v1_pools_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_coinbase_cloud_pools_v1_pools_proto_goTypes = []interface{}{
	(*Pool)(nil),              // 0: coinbase.cloud.pools.v1.Pool
	(*CreatePoolRequest)(nil), // 1: coinbase.cloud.pools.v1.CreatePoolRequest
	(*GetPoolRequest)(nil),    // 2: coinbase.cloud.pools.v1.GetPoolRequest
	(*ListPoolsRequest)(nil),  // 3: coinbase.cloud.pools.v1.ListPoolsRequest
	(*ListPoolsResponse)(nil), // 4: coinbase.cloud.pools.v1.ListPoolsResponse
}
var file_coinbase_cloud_pools_v1_pools_proto_depIdxs = []int32{
	0, // 0: coinbase.cloud.pools.v1.CreatePoolRequest.pool:type_name -> coinbase.cloud.pools.v1.Pool
	0, // 1: coinbase.cloud.pools.v1.ListPoolsResponse.pools:type_name -> coinbase.cloud.pools.v1.Pool
	1, // 2: coinbase.cloud.pools.v1.PoolService.CreatePool:input_type -> coinbase.cloud.pools.v1.CreatePoolRequest
	2, // 3: coinbase.cloud.pools.v1.PoolService.GetPool:input_type -> coinbase.cloud.pools.v1.GetPoolRequest
	3, // 4: coinbase.cloud.pools.v1.PoolService.ListPools:input_type -> coinbase.cloud.pools.v1.ListPoolsRequest
	0, // 5: coinbase.cloud.pools.v1.PoolService.CreatePool:output_type -> coinbase.cloud.pools.v1.Pool
	0, // 6: coinbase.cloud.pools.v1.PoolService.GetPool:output_type -> coinbase.cloud.pools.v1.Pool
	4, // 7: coinbase.cloud.pools.v1.PoolService.ListPools:output_type -> coinbase.cloud.pools.v1.ListPoolsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coinbase_cloud_pools_v1_pools_proto_init() }
func file_coinbase_cloud_pools_v1_pools_proto_init() {
	if File_coinbase_cloud_pools_v1_pools_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_cloud_pools_v1_pools_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pool); i {
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
		file_coinbase_cloud_pools_v1_pools_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePoolRequest); i {
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
		file_coinbase_cloud_pools_v1_pools_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoolRequest); i {
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
		file_coinbase_cloud_pools_v1_pools_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoolsRequest); i {
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
		file_coinbase_cloud_pools_v1_pools_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoolsResponse); i {
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
			RawDescriptor: file_coinbase_cloud_pools_v1_pools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coinbase_cloud_pools_v1_pools_proto_goTypes,
		DependencyIndexes: file_coinbase_cloud_pools_v1_pools_proto_depIdxs,
		MessageInfos:      file_coinbase_cloud_pools_v1_pools_proto_msgTypes,
	}.Build()
	File_coinbase_cloud_pools_v1_pools_proto = out.File
	file_coinbase_cloud_pools_v1_pools_proto_rawDesc = nil
	file_coinbase_cloud_pools_v1_pools_proto_goTypes = nil
	file_coinbase_cloud_pools_v1_pools_proto_depIdxs = nil
}

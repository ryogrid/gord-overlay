// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.6
// source: private.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{0}
}

func (x *Nodes) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{1}
}

func (x *FindRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type PutValueInnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutValueInnerRequest) Reset() {
	*x = PutValueInnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutValueInnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutValueInnerRequest) ProtoMessage() {}

func (x *PutValueInnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutValueInnerRequest.ProtoReflect.Descriptor instead.
func (*PutValueInnerRequest) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{2}
}

func (x *PutValueInnerRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutValueInnerRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutValueInnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PutValueInnerResponse) Reset() {
	*x = PutValueInnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutValueInnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutValueInnerResponse) ProtoMessage() {}

func (x *PutValueInnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutValueInnerResponse.ProtoReflect.Descriptor instead.
func (*PutValueInnerResponse) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{3}
}

func (x *PutValueInnerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetValueInnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueInnerRequest) Reset() {
	*x = GetValueInnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueInnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueInnerRequest) ProtoMessage() {}

func (x *GetValueInnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueInnerRequest.ProtoReflect.Descriptor instead.
func (*GetValueInnerRequest) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{4}
}

func (x *GetValueInnerRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetValueInnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value   string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetValueInnerResponse) Reset() {
	*x = GetValueInnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueInnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueInnerResponse) ProtoMessage() {}

func (x *GetValueInnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueInnerResponse.ProtoReflect.Descriptor instead.
func (*GetValueInnerResponse) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{5}
}

func (x *GetValueInnerResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetValueInnerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteValueInnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteValueInnerRequest) Reset() {
	*x = DeleteValueInnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteValueInnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteValueInnerRequest) ProtoMessage() {}

func (x *DeleteValueInnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteValueInnerRequest.ProtoReflect.Descriptor instead.
func (*DeleteValueInnerRequest) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteValueInnerRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteValueInnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteValueInnerResponse) Reset() {
	*x = DeleteValueInnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteValueInnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteValueInnerResponse) ProtoMessage() {}

func (x *DeleteValueInnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteValueInnerResponse.ProtoReflect.Descriptor instead.
func (*DeleteValueInnerResponse) Descriptor() ([]byte, []int) {
	return file_private_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteValueInnerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_private_proto protoreflect.FileDescriptor

var file_private_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2b, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x1d, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x14,
	0x50, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x15,
	0x50, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x2b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x34, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x9e, 0x05, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x50, 0x72,
	0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x42, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x42,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x18, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65, 0x63, 0x65, 0x64, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0d, 0x50, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x79, 0x6f, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x67, 0x6f, 0x72,
	0x64, 0x2d, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_private_proto_rawDescOnce sync.Once
	file_private_proto_rawDescData = file_private_proto_rawDesc
)

func file_private_proto_rawDescGZIP() []byte {
	file_private_proto_rawDescOnce.Do(func() {
		file_private_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_proto_rawDescData)
	})
	return file_private_proto_rawDescData
}

var file_private_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_private_proto_goTypes = []interface{}{
	(*Nodes)(nil),                    // 0: server.Nodes
	(*FindRequest)(nil),              // 1: server.FindRequest
	(*PutValueInnerRequest)(nil),     // 2: server.PutValueInnerRequest
	(*PutValueInnerResponse)(nil),    // 3: server.PutValueInnerResponse
	(*GetValueInnerRequest)(nil),     // 4: server.GetValueInnerRequest
	(*GetValueInnerResponse)(nil),    // 5: server.GetValueInnerResponse
	(*DeleteValueInnerRequest)(nil),  // 6: server.DeleteValueInnerRequest
	(*DeleteValueInnerResponse)(nil), // 7: server.DeleteValueInnerResponse
	(*Node)(nil),                     // 8: server.Node
	(*emptypb.Empty)(nil),            // 9: google.protobuf.Empty
}
var file_private_proto_depIdxs = []int32{
	8,  // 0: server.Nodes.nodes:type_name -> server.Node
	9,  // 1: server.InternalService.Ping:input_type -> google.protobuf.Empty
	9,  // 2: server.InternalService.Successors:input_type -> google.protobuf.Empty
	9,  // 3: server.InternalService.Predecessor:input_type -> google.protobuf.Empty
	1,  // 4: server.InternalService.FindSuccessorByTable:input_type -> server.FindRequest
	1,  // 5: server.InternalService.FindSuccessorByList:input_type -> server.FindRequest
	1,  // 6: server.InternalService.FindClosestPrecedingNode:input_type -> server.FindRequest
	8,  // 7: server.InternalService.Notify:input_type -> server.Node
	2,  // 8: server.InternalService.PutValueInner:input_type -> server.PutValueInnerRequest
	4,  // 9: server.InternalService.GetValueInner:input_type -> server.GetValueInnerRequest
	6,  // 10: server.InternalService.DeleteValueInner:input_type -> server.DeleteValueInnerRequest
	9,  // 11: server.InternalService.Ping:output_type -> google.protobuf.Empty
	0,  // 12: server.InternalService.Successors:output_type -> server.Nodes
	8,  // 13: server.InternalService.Predecessor:output_type -> server.Node
	8,  // 14: server.InternalService.FindSuccessorByTable:output_type -> server.Node
	8,  // 15: server.InternalService.FindSuccessorByList:output_type -> server.Node
	8,  // 16: server.InternalService.FindClosestPrecedingNode:output_type -> server.Node
	9,  // 17: server.InternalService.Notify:output_type -> google.protobuf.Empty
	3,  // 18: server.InternalService.PutValueInner:output_type -> server.PutValueInnerResponse
	5,  // 19: server.InternalService.GetValueInner:output_type -> server.GetValueInnerResponse
	7,  // 20: server.InternalService.DeleteValueInner:output_type -> server.DeleteValueInnerResponse
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_private_proto_init() }
func file_private_proto_init() {
	if File_private_proto != nil {
		return
	}
	file_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_private_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
		file_private_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_private_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutValueInnerRequest); i {
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
		file_private_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutValueInnerResponse); i {
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
		file_private_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueInnerRequest); i {
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
		file_private_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueInnerResponse); i {
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
		file_private_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteValueInnerRequest); i {
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
		file_private_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteValueInnerResponse); i {
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
			RawDescriptor: file_private_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_private_proto_goTypes,
		DependencyIndexes: file_private_proto_depIdxs,
		MessageInfos:      file_private_proto_msgTypes,
	}.Build()
	File_private_proto = out.File
	file_private_proto_rawDesc = nil
	file_private_proto_goTypes = nil
	file_private_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InternalServiceClient is the client API for InternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Successors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Nodes, error)
	Predecessor(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Node, error)
	FindSuccessorByTable(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	FindSuccessorByList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	FindClosestPrecedingNode(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PutValueInner(ctx context.Context, in *PutValueInnerRequest, opts ...grpc.CallOption) (*PutValueInnerResponse, error)
	GetValueInner(ctx context.Context, in *GetValueInnerRequest, opts ...grpc.CallOption) (*GetValueInnerResponse, error)
	DeleteValueInner(ctx context.Context, in *DeleteValueInnerRequest, opts ...grpc.CallOption) (*DeleteValueInnerResponse, error)
}

type internalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalServiceClient(cc grpc.ClientConnInterface) InternalServiceClient {
	return &internalServiceClient{cc}
}

func (c *internalServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/server.InternalService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Successors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Nodes, error) {
	out := new(Nodes)
	err := c.cc.Invoke(ctx, "/server.InternalService/Successors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Predecessor(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/Predecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindSuccessorByTable(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindSuccessorByTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindSuccessorByList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindSuccessorByList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindClosestPrecedingNode(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindClosestPrecedingNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/server.InternalService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) PutValueInner(ctx context.Context, in *PutValueInnerRequest, opts ...grpc.CallOption) (*PutValueInnerResponse, error) {
	out := new(PutValueInnerResponse)
	err := c.cc.Invoke(ctx, "/server.InternalService/PutValueInner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) GetValueInner(ctx context.Context, in *GetValueInnerRequest, opts ...grpc.CallOption) (*GetValueInnerResponse, error) {
	out := new(GetValueInnerResponse)
	err := c.cc.Invoke(ctx, "/server.InternalService/GetValueInner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) DeleteValueInner(ctx context.Context, in *DeleteValueInnerRequest, opts ...grpc.CallOption) (*DeleteValueInnerResponse, error) {
	out := new(DeleteValueInnerResponse)
	err := c.cc.Invoke(ctx, "/server.InternalService/DeleteValueInner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServiceServer is the server API for InternalService service.
type InternalServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Successors(context.Context, *emptypb.Empty) (*Nodes, error)
	Predecessor(context.Context, *emptypb.Empty) (*Node, error)
	FindSuccessorByTable(context.Context, *FindRequest) (*Node, error)
	FindSuccessorByList(context.Context, *FindRequest) (*Node, error)
	FindClosestPrecedingNode(context.Context, *FindRequest) (*Node, error)
	Notify(context.Context, *Node) (*emptypb.Empty, error)
	PutValueInner(context.Context, *PutValueInnerRequest) (*PutValueInnerResponse, error)
	GetValueInner(context.Context, *GetValueInnerRequest) (*GetValueInnerResponse, error)
	DeleteValueInner(context.Context, *DeleteValueInnerRequest) (*DeleteValueInnerResponse, error)
}

// UnimplementedInternalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInternalServiceServer struct {
}

func (*UnimplementedInternalServiceServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedInternalServiceServer) Successors(context.Context, *emptypb.Empty) (*Nodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Successors not implemented")
}
func (*UnimplementedInternalServiceServer) Predecessor(context.Context, *emptypb.Empty) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predecessor not implemented")
}
func (*UnimplementedInternalServiceServer) FindSuccessorByTable(context.Context, *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorByTable not implemented")
}
func (*UnimplementedInternalServiceServer) FindSuccessorByList(context.Context, *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorByList not implemented")
}
func (*UnimplementedInternalServiceServer) FindClosestPrecedingNode(context.Context, *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindClosestPrecedingNode not implemented")
}
func (*UnimplementedInternalServiceServer) Notify(context.Context, *Node) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedInternalServiceServer) PutValueInner(context.Context, *PutValueInnerRequest) (*PutValueInnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutValueInner not implemented")
}
func (*UnimplementedInternalServiceServer) GetValueInner(context.Context, *GetValueInnerRequest) (*GetValueInnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueInner not implemented")
}
func (*UnimplementedInternalServiceServer) DeleteValueInner(context.Context, *DeleteValueInnerRequest) (*DeleteValueInnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteValueInner not implemented")
}

func RegisterInternalServiceServer(s *grpc.Server, srv InternalServiceServer) {
	s.RegisterService(&_InternalService_serviceDesc, srv)
}

func _InternalService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Successors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Successors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Successors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Successors(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Predecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Predecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Predecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Predecessor(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindSuccessorByTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindSuccessorByTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindSuccessorByTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindSuccessorByTable(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindSuccessorByList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindSuccessorByList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindSuccessorByList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindSuccessorByList(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindClosestPrecedingNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindClosestPrecedingNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindClosestPrecedingNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindClosestPrecedingNode(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_PutValueInner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutValueInnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).PutValueInner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/PutValueInner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).PutValueInner(ctx, req.(*PutValueInnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_GetValueInner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueInnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).GetValueInner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/GetValueInner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).GetValueInner(ctx, req.(*GetValueInnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_DeleteValueInner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteValueInnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).DeleteValueInner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/DeleteValueInner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).DeleteValueInner(ctx, req.(*DeleteValueInnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.InternalService",
	HandlerType: (*InternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _InternalService_Ping_Handler,
		},
		{
			MethodName: "Successors",
			Handler:    _InternalService_Successors_Handler,
		},
		{
			MethodName: "Predecessor",
			Handler:    _InternalService_Predecessor_Handler,
		},
		{
			MethodName: "FindSuccessorByTable",
			Handler:    _InternalService_FindSuccessorByTable_Handler,
		},
		{
			MethodName: "FindSuccessorByList",
			Handler:    _InternalService_FindSuccessorByList_Handler,
		},
		{
			MethodName: "FindClosestPrecedingNode",
			Handler:    _InternalService_FindClosestPrecedingNode_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _InternalService_Notify_Handler,
		},
		{
			MethodName: "PutValueInner",
			Handler:    _InternalService_PutValueInner_Handler,
		},
		{
			MethodName: "GetValueInner",
			Handler:    _InternalService_GetValueInner_Handler,
		},
		{
			MethodName: "DeleteValueInner",
			Handler:    _InternalService_DeleteValueInner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "private.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: pxguide/px_guide.proto

package pxguide

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StartArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V string `protobuf:"bytes,1,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *StartArgs) Reset() {
	*x = StartArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartArgs) ProtoMessage() {}

func (x *StartArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartArgs.ProtoReflect.Descriptor instead.
func (*StartArgs) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{0}
}

func (x *StartArgs) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

type PrepareArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=N,proto3" json:"N,omitempty"`
}

func (x *PrepareArgs) Reset() {
	*x = PrepareArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareArgs) ProtoMessage() {}

func (x *PrepareArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareArgs.ProtoReflect.Descriptor instead.
func (*PrepareArgs) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{1}
}

func (x *PrepareArgs) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type AcceptArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32  `protobuf:"varint,1,opt,name=N,proto3" json:"N,omitempty"`
	V string `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *AcceptArgs) Reset() {
	*x = AcceptArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptArgs) ProtoMessage() {}

func (x *AcceptArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptArgs.ProtoReflect.Descriptor instead.
func (*AcceptArgs) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptArgs) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *AcceptArgs) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

type DecidedArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V string `protobuf:"bytes,1,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *DecidedArgs) Reset() {
	*x = DecidedArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecidedArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecidedArgs) ProtoMessage() {}

func (x *DecidedArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecidedArgs.ProtoReflect.Descriptor instead.
func (*DecidedArgs) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{3}
}

func (x *DecidedArgs) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

type StartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
}

func (x *StartReply) Reset() {
	*x = StartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartReply) ProtoMessage() {}

func (x *StartReply) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartReply.ProtoReflect.Descriptor instead.
func (*StartReply) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{4}
}

func (x *StartReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type PrepareReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	N      int32  `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	Na     int32  `protobuf:"varint,3,opt,name=Na,proto3" json:"Na,omitempty"`
	Va     string `protobuf:"bytes,4,opt,name=Va,proto3" json:"Va,omitempty"`
	IsNull bool   `protobuf:"varint,5,opt,name=IsNull,proto3" json:"IsNull,omitempty"`
	Failed bool   `protobuf:"varint,6,opt,name=Failed,proto3" json:"Failed,omitempty"`
}

func (x *PrepareReply) Reset() {
	*x = PrepareReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareReply) ProtoMessage() {}

func (x *PrepareReply) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareReply.ProtoReflect.Descriptor instead.
func (*PrepareReply) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{5}
}

func (x *PrepareReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *PrepareReply) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *PrepareReply) GetNa() int32 {
	if x != nil {
		return x.Na
	}
	return 0
}

func (x *PrepareReply) GetVa() string {
	if x != nil {
		return x.Va
	}
	return ""
}

func (x *PrepareReply) GetIsNull() bool {
	if x != nil {
		return x.IsNull
	}
	return false
}

func (x *PrepareReply) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

type AcceptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool  `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	N      int32 `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	Failed bool  `protobuf:"varint,3,opt,name=Failed,proto3" json:"Failed,omitempty"`
}

func (x *AcceptReply) Reset() {
	*x = AcceptReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptReply) ProtoMessage() {}

func (x *AcceptReply) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptReply.ProtoReflect.Descriptor instead.
func (*AcceptReply) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *AcceptReply) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *AcceptReply) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

type DecidedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
}

func (x *DecidedReply) Reset() {
	*x = DecidedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pxguide_px_guide_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecidedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecidedReply) ProtoMessage() {}

func (x *DecidedReply) ProtoReflect() protoreflect.Message {
	mi := &file_pxguide_px_guide_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecidedReply.ProtoReflect.Descriptor instead.
func (*DecidedReply) Descriptor() ([]byte, []int) {
	return file_pxguide_px_guide_proto_rawDescGZIP(), []int{7}
}

func (x *DecidedReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_pxguide_px_guide_proto protoreflect.FileDescriptor

var file_pxguide_px_guide_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x70, 0x78, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64,
	0x65, 0x22, 0x19, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0c,
	0x0a, 0x01, 0x56, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x56, 0x22, 0x1b, 0x0a, 0x0b,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x4e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x4e, 0x22, 0x28, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x4e, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x56, 0x22, 0x1b, 0x0a, 0x0b, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x56,
	0x22, 0x1c, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x22, 0x7c,
	0x0a, 0x0c, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x0c,
	0x0a, 0x01, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x4e, 0x12, 0x0e, 0x0a, 0x02,
	0x4e, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x4e, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x56, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x56, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73,
	0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x4f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x4e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f,
	0x6b, 0x32, 0xe8, 0x01, 0x0a, 0x07, 0x50, 0x78, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x70, 0x78, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x70,
	0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72,
	0x67, 0x73, 0x1a, 0x15, 0x2e, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x78, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x12, 0x14, 0x2e,
	0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x41,
	0x72, 0x67, 0x73, 0x1a, 0x15, 0x2e, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x44, 0x65,
	0x63, 0x69, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x48, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x63, 0x6c, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x42, 0x0c, 0x50, 0x78, 0x47, 0x75,
	0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x77, 0x63, 0x6c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x78, 0x67, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pxguide_px_guide_proto_rawDescOnce sync.Once
	file_pxguide_px_guide_proto_rawDescData = file_pxguide_px_guide_proto_rawDesc
)

func file_pxguide_px_guide_proto_rawDescGZIP() []byte {
	file_pxguide_px_guide_proto_rawDescOnce.Do(func() {
		file_pxguide_px_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_pxguide_px_guide_proto_rawDescData)
	})
	return file_pxguide_px_guide_proto_rawDescData
}

var file_pxguide_px_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pxguide_px_guide_proto_goTypes = []interface{}{
	(*StartArgs)(nil),    // 0: pxguide.StartArgs
	(*PrepareArgs)(nil),  // 1: pxguide.PrepareArgs
	(*AcceptArgs)(nil),   // 2: pxguide.AcceptArgs
	(*DecidedArgs)(nil),  // 3: pxguide.DecidedArgs
	(*StartReply)(nil),   // 4: pxguide.StartReply
	(*PrepareReply)(nil), // 5: pxguide.PrepareReply
	(*AcceptReply)(nil),  // 6: pxguide.AcceptReply
	(*DecidedReply)(nil), // 7: pxguide.DecidedReply
}
var file_pxguide_px_guide_proto_depIdxs = []int32{
	0, // 0: pxguide.PxGuide.Start:input_type -> pxguide.StartArgs
	1, // 1: pxguide.PxGuide.Prepare:input_type -> pxguide.PrepareArgs
	2, // 2: pxguide.PxGuide.Accept:input_type -> pxguide.AcceptArgs
	3, // 3: pxguide.PxGuide.Decided:input_type -> pxguide.DecidedArgs
	4, // 4: pxguide.PxGuide.Start:output_type -> pxguide.StartReply
	5, // 5: pxguide.PxGuide.Prepare:output_type -> pxguide.PrepareReply
	6, // 6: pxguide.PxGuide.Accept:output_type -> pxguide.AcceptReply
	7, // 7: pxguide.PxGuide.Decided:output_type -> pxguide.DecidedReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pxguide_px_guide_proto_init() }
func file_pxguide_px_guide_proto_init() {
	if File_pxguide_px_guide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pxguide_px_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartArgs); i {
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
		file_pxguide_px_guide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareArgs); i {
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
		file_pxguide_px_guide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptArgs); i {
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
		file_pxguide_px_guide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecidedArgs); i {
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
		file_pxguide_px_guide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartReply); i {
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
		file_pxguide_px_guide_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareReply); i {
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
		file_pxguide_px_guide_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptReply); i {
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
		file_pxguide_px_guide_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecidedReply); i {
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
			RawDescriptor: file_pxguide_px_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pxguide_px_guide_proto_goTypes,
		DependencyIndexes: file_pxguide_px_guide_proto_depIdxs,
		MessageInfos:      file_pxguide_px_guide_proto_msgTypes,
	}.Build()
	File_pxguide_px_guide_proto = out.File
	file_pxguide_px_guide_proto_rawDesc = nil
	file_pxguide_px_guide_proto_goTypes = nil
	file_pxguide_px_guide_proto_depIdxs = nil
}

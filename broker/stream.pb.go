// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: stream.proto

package broker

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

type StreamRetentionPolicy int32

const (
	StreamRetentionPolicy_TIME_RETENTION_POLICY StreamRetentionPolicy = 0
	StreamRetentionPolicy_SIZE_RETENTION_POLICY StreamRetentionPolicy = 1
)

// Enum value maps for StreamRetentionPolicy.
var (
	StreamRetentionPolicy_name = map[int32]string{
		0: "TIME_RETENTION_POLICY",
		1: "SIZE_RETENTION_POLICY",
	}
	StreamRetentionPolicy_value = map[string]int32{
		"TIME_RETENTION_POLICY": 0,
		"SIZE_RETENTION_POLICY": 1,
	}
)

func (x StreamRetentionPolicy) Enum() *StreamRetentionPolicy {
	p := new(StreamRetentionPolicy)
	*p = x
	return p
}

func (x StreamRetentionPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamRetentionPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_stream_proto_enumTypes[0].Descriptor()
}

func (StreamRetentionPolicy) Type() protoreflect.EnumType {
	return &file_stream_proto_enumTypes[0]
}

func (x StreamRetentionPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamRetentionPolicy.Descriptor instead.
func (StreamRetentionPolicy) EnumDescriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

type StreamRetentionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxAge  string `protobuf:"bytes,1,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	MaxSize int64  `protobuf:"varint,2,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
}

func (x *StreamRetentionOptions) Reset() {
	*x = StreamRetentionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRetentionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRetentionOptions) ProtoMessage() {}

func (x *StreamRetentionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRetentionOptions.ProtoReflect.Descriptor instead.
func (*StreamRetentionOptions) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRetentionOptions) GetMaxAge() string {
	if x != nil {
		return x.MaxAge
	}
	return ""
}

func (x *StreamRetentionOptions) GetMaxSize() int64 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName       string                  `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	RetentionPolicy  StreamRetentionPolicy   `protobuf:"varint,2,opt,name=retention_policy,json=retentionPolicy,proto3,enum=broker.StreamRetentionPolicy" json:"retention_policy,omitempty"`
	RetentionOptions *StreamRetentionOptions `protobuf:"bytes,3,opt,name=retention_options,json=retentionOptions,proto3" json:"retention_options,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *Stream) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *Stream) GetRetentionPolicy() StreamRetentionPolicy {
	if x != nil {
		return x.RetentionPolicy
	}
	return StreamRetentionPolicy_TIME_RETENTION_POLICY
}

func (x *Stream) GetRetentionOptions() *StreamRetentionOptions {
	if x != nil {
		return x.RetentionOptions
	}
	return nil
}

type CreateStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName       string                  `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	RetentionPolicy  StreamRetentionPolicy   `protobuf:"varint,2,opt,name=retention_policy,json=retentionPolicy,proto3,enum=broker.StreamRetentionPolicy" json:"retention_policy,omitempty"`
	RetentionOptions *StreamRetentionOptions `protobuf:"bytes,3,opt,name=retention_options,json=retentionOptions,proto3" json:"retention_options,omitempty"`
}

func (x *CreateStreamRequest) Reset() {
	*x = CreateStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamRequest) ProtoMessage() {}

func (x *CreateStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStreamRequest.ProtoReflect.Descriptor instead.
func (*CreateStreamRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStreamRequest) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *CreateStreamRequest) GetRetentionPolicy() StreamRetentionPolicy {
	if x != nil {
		return x.RetentionPolicy
	}
	return StreamRetentionPolicy_TIME_RETENTION_POLICY
}

func (x *CreateStreamRequest) GetRetentionOptions() *StreamRetentionOptions {
	if x != nil {
		return x.RetentionOptions
	}
	return nil
}

type CreateStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateStreamResponse) Reset() {
	*x = CreateStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamResponse) ProtoMessage() {}

func (x *CreateStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStreamResponse.ProtoReflect.Descriptor instead.
func (*CreateStreamResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStreamResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName string `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
}

func (x *GetStreamRequest) Reset() {
	*x = GetStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamRequest) ProtoMessage() {}

func (x *GetStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamRequest.ProtoReflect.Descriptor instead.
func (*GetStreamRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{4}
}

func (x *GetStreamRequest) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

type GetStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetStreamResponse) Reset() {
	*x = GetStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamResponse) ProtoMessage() {}

func (x *GetStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamResponse.ProtoReflect.Descriptor instead.
func (*GetStreamResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{5}
}

func (x *GetStreamResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x48, 0x0a, 0x10, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x4b, 0x0a, 0x11, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x48, 0x0a, 0x10, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x4b, 0x0a, 0x11, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x4d, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x19,
	0x0a, 0x15, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x54, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x49, 0x5a,
	0x45, 0x5f, 0x52, 0x45, 0x54, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4c, 0x49,
	0x43, 0x59, 0x10, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x69,
	0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stream_proto_goTypes = []interface{}{
	(StreamRetentionPolicy)(0),     // 0: broker.StreamRetentionPolicy
	(*StreamRetentionOptions)(nil), // 1: broker.StreamRetentionOptions
	(*Stream)(nil),                 // 2: broker.Stream
	(*CreateStreamRequest)(nil),    // 3: broker.CreateStreamRequest
	(*CreateStreamResponse)(nil),   // 4: broker.CreateStreamResponse
	(*GetStreamRequest)(nil),       // 5: broker.GetStreamRequest
	(*GetStreamResponse)(nil),      // 6: broker.GetStreamResponse
}
var file_stream_proto_depIdxs = []int32{
	0, // 0: broker.Stream.retention_policy:type_name -> broker.StreamRetentionPolicy
	1, // 1: broker.Stream.retention_options:type_name -> broker.StreamRetentionOptions
	0, // 2: broker.CreateStreamRequest.retention_policy:type_name -> broker.StreamRetentionPolicy
	1, // 3: broker.CreateStreamRequest.retention_options:type_name -> broker.StreamRetentionOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRetentionOptions); i {
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
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStreamRequest); i {
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
		file_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStreamResponse); i {
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
		file_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamRequest); i {
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
		file_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamResponse); i {
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
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		EnumInfos:         file_stream_proto_enumTypes,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.1
// source: plugin/shared/grpc/plugin.proto

package grpc

import (
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

type ActionMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionParameter map[string]*ParameterMetadata `protobuf:"bytes,1,rep,name=actionParameter,proto3" json:"actionParameter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ActionMetadataResponse) Reset() {
	*x = ActionMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMetadataResponse) ProtoMessage() {}

func (x *ActionMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMetadataResponse.ProtoReflect.Descriptor instead.
func (*ActionMetadataResponse) Descriptor() ([]byte, []int) {
	return file_plugin_shared_grpc_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *ActionMetadataResponse) GetActionParameter() map[string]*ParameterMetadata {
	if x != nil {
		return x.ActionParameter
	}
	return nil
}

type ParameterMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParameterMetadata map[string]string `protobuf:"bytes,1,rep,name=parameterMetadata,proto3" json:"parameterMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ParameterMetadata) Reset() {
	*x = ParameterMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterMetadata) ProtoMessage() {}

func (x *ParameterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterMetadata.ProtoReflect.Descriptor instead.
func (*ParameterMetadata) Descriptor() ([]byte, []int) {
	return file_plugin_shared_grpc_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *ParameterMetadata) GetParameterMetadata() map[string]string {
	if x != nil {
		return x.ParameterMetadata
	}
	return nil
}

type MetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle string `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
}

func (x *MetadataResponse) Reset() {
	*x = MetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataResponse) ProtoMessage() {}

func (x *MetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataResponse.ProtoReflect.Descriptor instead.
func (*MetadataResponse) Descriptor() ([]byte, []int) {
	return file_plugin_shared_grpc_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *MetadataResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MetadataResponse) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string             `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Params *ParameterMetadata `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_plugin_shared_grpc_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ExecuteRequest) GetParams() *ParameterMetadata {
	if x != nil {
		return x.Params
	}
	return nil
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_shared_grpc_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_plugin_shared_grpc_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_plugin_shared_grpc_plugin_proto protoreflect.FileDescriptor

var file_plugin_shared_grpc_plugin_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5c, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x5c,
	0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb8, 0x01, 0x0a,
	0x11, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x5d, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x44, 0x0a, 0x16, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5a, 0x0a,
	0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd6, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x14, 0x5a, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_shared_grpc_plugin_proto_rawDescOnce sync.Once
	file_plugin_shared_grpc_plugin_proto_rawDescData = file_plugin_shared_grpc_plugin_proto_rawDesc
)

func file_plugin_shared_grpc_plugin_proto_rawDescGZIP() []byte {
	file_plugin_shared_grpc_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_shared_grpc_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_shared_grpc_plugin_proto_rawDescData)
	})
	return file_plugin_shared_grpc_plugin_proto_rawDescData
}

var file_plugin_shared_grpc_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugin_shared_grpc_plugin_proto_goTypes = []interface{}{
	(*ActionMetadataResponse)(nil), // 0: proto.ActionMetadataResponse
	(*ParameterMetadata)(nil),      // 1: proto.ParameterMetadata
	(*MetadataResponse)(nil),       // 2: proto.MetadataResponse
	(*ExecuteRequest)(nil),         // 3: proto.ExecuteRequest
	(*ExecuteResponse)(nil),        // 4: proto.ExecuteResponse
	nil,                            // 5: proto.ActionMetadataResponse.ActionParameterEntry
	nil,                            // 6: proto.ParameterMetadata.ParameterMetadataEntry
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_plugin_shared_grpc_plugin_proto_depIdxs = []int32{
	5, // 0: proto.ActionMetadataResponse.actionParameter:type_name -> proto.ActionMetadataResponse.ActionParameterEntry
	6, // 1: proto.ParameterMetadata.parameterMetadata:type_name -> proto.ParameterMetadata.ParameterMetadataEntry
	1, // 2: proto.ExecuteRequest.params:type_name -> proto.ParameterMetadata
	1, // 3: proto.ActionMetadataResponse.ActionParameterEntry.value:type_name -> proto.ParameterMetadata
	7, // 4: proto.Plugin.GetActionMetadata:input_type -> google.protobuf.Empty
	7, // 5: proto.Plugin.GetMetadata:input_type -> google.protobuf.Empty
	3, // 6: proto.Plugin.Execute:input_type -> proto.ExecuteRequest
	0, // 7: proto.Plugin.GetActionMetadata:output_type -> proto.ActionMetadataResponse
	2, // 8: proto.Plugin.GetMetadata:output_type -> proto.MetadataResponse
	4, // 9: proto.Plugin.Execute:output_type -> proto.ExecuteResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugin_shared_grpc_plugin_proto_init() }
func file_plugin_shared_grpc_plugin_proto_init() {
	if File_plugin_shared_grpc_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_shared_grpc_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMetadataResponse); i {
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
		file_plugin_shared_grpc_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterMetadata); i {
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
		file_plugin_shared_grpc_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataResponse); i {
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
		file_plugin_shared_grpc_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_plugin_shared_grpc_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
			RawDescriptor: file_plugin_shared_grpc_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_shared_grpc_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_shared_grpc_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_shared_grpc_plugin_proto_msgTypes,
	}.Build()
	File_plugin_shared_grpc_plugin_proto = out.File
	file_plugin_shared_grpc_plugin_proto_rawDesc = nil
	file_plugin_shared_grpc_plugin_proto_goTypes = nil
	file_plugin_shared_grpc_plugin_proto_depIdxs = nil
}

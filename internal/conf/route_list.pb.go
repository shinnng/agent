// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: conf/route_list.proto

package conf

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

type Apis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Apis) Reset() {
	*x = Apis{}
	mi := &file_conf_route_list_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Apis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apis) ProtoMessage() {}

func (x *Apis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_route_list_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apis.ProtoReflect.Descriptor instead.
func (*Apis) Descriptor() ([]byte, []int) {
	return file_conf_route_list_proto_rawDescGZIP(), []int{0}
}

func (x *Apis) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       int32       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	RestEndpoints []*Endpoint `protobuf:"bytes,2,rep,name=rest_endpoints,json=restEndpoints,proto3" json:"rest_endpoints,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_conf_route_list_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_conf_route_list_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_conf_route_list_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Metadata) GetRestEndpoints() []*Endpoint {
	if x != nil {
		return x.RestEndpoints
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
	Url     string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	mi := &file_conf_route_list_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_conf_route_list_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_conf_route_list_proto_rawDescGZIP(), []int{2}
}

func (x *Endpoint) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Endpoint) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_conf_route_list_proto protoreflect.FileDescriptor

var file_conf_route_list_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x22, 0x38, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x22, 0x36, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x1a, 0x5a, 0x18, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_route_list_proto_rawDescOnce sync.Once
	file_conf_route_list_proto_rawDescData = file_conf_route_list_proto_rawDesc
)

func file_conf_route_list_proto_rawDescGZIP() []byte {
	file_conf_route_list_proto_rawDescOnce.Do(func() {
		file_conf_route_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_route_list_proto_rawDescData)
	})
	return file_conf_route_list_proto_rawDescData
}

var file_conf_route_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conf_route_list_proto_goTypes = []any{
	(*Apis)(nil),     // 0: kratos.api.apis
	(*Metadata)(nil), // 1: kratos.api.metadata
	(*Endpoint)(nil), // 2: kratos.api.endpoint
}
var file_conf_route_list_proto_depIdxs = []int32{
	1, // 0: kratos.api.apis.metadata:type_name -> kratos.api.metadata
	2, // 1: kratos.api.metadata.rest_endpoints:type_name -> kratos.api.endpoint
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_conf_route_list_proto_init() }
func file_conf_route_list_proto_init() {
	if File_conf_route_list_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_route_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_route_list_proto_goTypes,
		DependencyIndexes: file_conf_route_list_proto_depIdxs,
		MessageInfos:      file_conf_route_list_proto_msgTypes,
	}.Build()
	File_conf_route_list_proto = out.File
	file_conf_route_list_proto_rawDesc = nil
	file_conf_route_list_proto_goTypes = nil
	file_conf_route_list_proto_depIdxs = nil
}

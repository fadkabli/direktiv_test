// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: pkg/functions/grpc/watch-revisions.proto

package grpc

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

type WatchRevisionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName  *string `protobuf:"bytes,1,opt,name=ServiceName,proto3,oneof" json:"ServiceName,omitempty"`
	RevisionName *string `protobuf:"bytes,2,opt,name=RevisionName,proto3,oneof" json:"RevisionName,omitempty"`
	Scope        *string `protobuf:"bytes,6,opt,name=scope,proto3,oneof" json:"scope,omitempty"`
}

func (x *WatchRevisionsRequest) Reset() {
	*x = WatchRevisionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_revisions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRevisionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRevisionsRequest) ProtoMessage() {}

func (x *WatchRevisionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_revisions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRevisionsRequest.ProtoReflect.Descriptor instead.
func (*WatchRevisionsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_revisions_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRevisionsRequest) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *WatchRevisionsRequest) GetRevisionName() string {
	if x != nil && x.RevisionName != nil {
		return *x.RevisionName
	}
	return ""
}

func (x *WatchRevisionsRequest) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

type WatchRevisionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event     *string          `protobuf:"bytes,1,opt,name=event,proto3,oneof" json:"event,omitempty"`
	Name      *string          `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Namespace *string          `protobuf:"bytes,3,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Workflow  *string          `protobuf:"bytes,4,opt,name=workflow,proto3,oneof" json:"workflow,omitempty"`
	Config    *FunctionsConfig `protobuf:"bytes,5,opt,name=config,proto3,oneof" json:"config,omitempty"`
	Revision  *Revision        `protobuf:"bytes,6,opt,name=revision,proto3,oneof" json:"revision,omitempty"`
}

func (x *WatchRevisionsResponse) Reset() {
	*x = WatchRevisionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_revisions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRevisionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRevisionsResponse) ProtoMessage() {}

func (x *WatchRevisionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_revisions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRevisionsResponse.ProtoReflect.Descriptor instead.
func (*WatchRevisionsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_revisions_proto_rawDescGZIP(), []int{1}
}

func (x *WatchRevisionsResponse) GetEvent() string {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return ""
}

func (x *WatchRevisionsResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *WatchRevisionsResponse) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *WatchRevisionsResponse) GetWorkflow() string {
	if x != nil && x.Workflow != nil {
		return *x.Workflow
	}
	return ""
}

func (x *WatchRevisionsResponse) GetConfig() *FunctionsConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *WatchRevisionsResponse) GetRevision() *Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

var File_pkg_functions_grpc_watch_revisions_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_watch_revisions_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x29,
	0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0c, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0xd7, 0x02, 0x0a, 0x16, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x3d,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x05,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74,
	0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_watch_revisions_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_watch_revisions_proto_rawDescData = file_pkg_functions_grpc_watch_revisions_proto_rawDesc
)

func file_pkg_functions_grpc_watch_revisions_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_watch_revisions_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_watch_revisions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_watch_revisions_proto_rawDescData)
	})
	return file_pkg_functions_grpc_watch_revisions_proto_rawDescData
}

var file_pkg_functions_grpc_watch_revisions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_functions_grpc_watch_revisions_proto_goTypes = []interface{}{
	(*WatchRevisionsRequest)(nil),  // 0: direktiv_functions.WatchRevisionsRequest
	(*WatchRevisionsResponse)(nil), // 1: direktiv_functions.WatchRevisionsResponse
	(*FunctionsConfig)(nil),        // 2: direktiv_functions.FunctionsConfig
	(*Revision)(nil),               // 3: direktiv_functions.Revision
}
var file_pkg_functions_grpc_watch_revisions_proto_depIdxs = []int32{
	2, // 0: direktiv_functions.WatchRevisionsResponse.config:type_name -> direktiv_functions.FunctionsConfig
	3, // 1: direktiv_functions.WatchRevisionsResponse.revision:type_name -> direktiv_functions.Revision
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_watch_revisions_proto_init() }
func file_pkg_functions_grpc_watch_revisions_proto_init() {
	if File_pkg_functions_grpc_watch_revisions_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_watch_revisions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRevisionsRequest); i {
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
		file_pkg_functions_grpc_watch_revisions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRevisionsResponse); i {
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
	file_pkg_functions_grpc_watch_revisions_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_watch_revisions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_watch_revisions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_watch_revisions_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_watch_revisions_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_watch_revisions_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_watch_revisions_proto = out.File
	file_pkg_functions_grpc_watch_revisions_proto_rawDesc = nil
	file_pkg_functions_grpc_watch_revisions_proto_goTypes = nil
	file_pkg_functions_grpc_watch_revisions_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/functions/grpc/registries-get.proto

package grpc

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

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Id   *string `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	User *string `protobuf:"bytes,3,opt,name=user,proto3,oneof" json:"user,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_registries_get_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Registry) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Registry) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

type GetRegistriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *GetRegistriesRequest) Reset() {
	*x = GetRegistriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistriesRequest) ProtoMessage() {}

func (x *GetRegistriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistriesRequest.ProtoReflect.Descriptor instead.
func (*GetRegistriesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_registries_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetRegistriesRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type GetRegistriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registries []*Registry `protobuf:"bytes,1,rep,name=registries,proto3" json:"registries,omitempty"`
}

func (x *GetRegistriesResponse) Reset() {
	*x = GetRegistriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistriesResponse) ProtoMessage() {}

func (x *GetRegistriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistriesResponse.ProtoReflect.Descriptor instead.
func (*GetRegistriesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_registries_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetRegistriesResponse) GetRegistries() []*Registry {
	if x != nil {
		return x.Registries
	}
	return nil
}

var File_pkg_functions_grpc_registries_get_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_registries_get_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2d,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6a, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_registries_get_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_registries_get_proto_rawDescData = file_pkg_functions_grpc_registries_get_proto_rawDesc
)

func file_pkg_functions_grpc_registries_get_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_registries_get_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_registries_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_registries_get_proto_rawDescData)
	})
	return file_pkg_functions_grpc_registries_get_proto_rawDescData
}

var file_pkg_functions_grpc_registries_get_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_functions_grpc_registries_get_proto_goTypes = []interface{}{
	(*Registry)(nil),              // 0: direktiv_functions.Registry
	(*GetRegistriesRequest)(nil),  // 1: direktiv_functions.GetRegistriesRequest
	(*GetRegistriesResponse)(nil), // 2: direktiv_functions.GetRegistriesResponse
}
var file_pkg_functions_grpc_registries_get_proto_depIdxs = []int32{
	0, // 0: direktiv_functions.GetRegistriesResponse.registries:type_name -> direktiv_functions.Registry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_registries_get_proto_init() }
func file_pkg_functions_grpc_registries_get_proto_init() {
	if File_pkg_functions_grpc_registries_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_registries_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_pkg_functions_grpc_registries_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistriesRequest); i {
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
		file_pkg_functions_grpc_registries_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistriesResponse); i {
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
	file_pkg_functions_grpc_registries_get_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_registries_get_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_registries_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_registries_get_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_registries_get_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_registries_get_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_registries_get_proto = out.File
	file_pkg_functions_grpc_registries_get_proto_rawDesc = nil
	file_pkg_functions_grpc_registries_get_proto_goTypes = nil
	file_pkg_functions_grpc_registries_get_proto_depIdxs = nil
}

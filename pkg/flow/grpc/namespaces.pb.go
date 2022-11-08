// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/flow/grpc/namespaces.proto

package grpc

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Name      string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Oid       string               `protobuf:"bytes,4,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Namespace) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Namespace) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

type ResolveNamespaceUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResolveNamespaceUIDRequest) Reset() {
	*x = ResolveNamespaceUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveNamespaceUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNamespaceUIDRequest) ProtoMessage() {}

func (x *ResolveNamespaceUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveNamespaceUIDRequest.ProtoReflect.Descriptor instead.
func (*ResolveNamespaceUIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveNamespaceUIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NamespaceRequest) Reset() {
	*x = NamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceRequest) ProtoMessage() {}

func (x *NamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceRequest.ProtoReflect.Descriptor instead.
func (*NamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NamespaceResponse) Reset() {
	*x = NamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceResponse) ProtoMessage() {}

func (x *NamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceResponse.ProtoReflect.Descriptor instead.
func (*NamespaceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{3}
}

func (x *NamespaceResponse) GetNamespace() *Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

type NamespacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *NamespacesRequest) Reset() {
	*x = NamespacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespacesRequest) ProtoMessage() {}

func (x *NamespacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespacesRequest.ProtoReflect.Descriptor instead.
func (*NamespacesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{4}
}

func (x *NamespacesRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type NamespacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo *PageInfo    `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Results  []*Namespace `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *NamespacesResponse) Reset() {
	*x = NamespacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespacesResponse) ProtoMessage() {}

func (x *NamespacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespacesResponse.ProtoReflect.Descriptor instead.
func (*NamespacesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{5}
}

func (x *NamespacesResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *NamespacesResponse) GetResults() []*Namespace {
	if x != nil {
		return x.Results
	}
	return nil
}

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Idempotent bool   `protobuf:"varint,2,opt,name=idempotent,proto3" json:"idempotent,omitempty"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{6}
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceRequest) GetIdempotent() bool {
	if x != nil {
		return x.Idempotent
	}
	return false
}

type CreateNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *CreateNamespaceResponse) Reset() {
	*x = CreateNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceResponse) ProtoMessage() {}

func (x *CreateNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceResponse.ProtoReflect.Descriptor instead.
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{7}
}

func (x *CreateNamespaceResponse) GetNamespace() *Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Idempotent bool   `protobuf:"varint,2,opt,name=idempotent,proto3" json:"idempotent,omitempty"`
	Recursive  bool   `protobuf:"varint,3,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteNamespaceRequest) GetIdempotent() bool {
	if x != nil {
		return x.Idempotent
	}
	return false
}

func (x *DeleteNamespaceRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type RenameNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Old string `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`
	New string `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *RenameNamespaceRequest) Reset() {
	*x = RenameNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameNamespaceRequest) ProtoMessage() {}

func (x *RenameNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameNamespaceRequest.ProtoReflect.Descriptor instead.
func (*RenameNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{9}
}

func (x *RenameNamespaceRequest) GetOld() string {
	if x != nil {
		return x.Old
	}
	return ""
}

func (x *RenameNamespaceRequest) GetNew() string {
	if x != nil {
		return x.New
	}
	return ""
}

type RenameNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *RenameNamespaceResponse) Reset() {
	*x = RenameNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameNamespaceResponse) ProtoMessage() {}

func (x *RenameNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_namespaces_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameNamespaceResponse.ProtoReflect.Descriptor instead.
func (*RenameNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_namespaces_proto_rawDescGZIP(), []int{10}
}

func (x *RenameNamespaceResponse) GetNamespace() *Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

var File_pkg_flow_grpc_namespaces_proto protoreflect.FileDescriptor

var file_pkg_flow_grpc_namespaces_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa7, 0x01, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4b, 0x0a, 0x11, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4e, 0x0a,
	0x11, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7d, 0x0a,
	0x12, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x4c, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x6a, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x22, 0x51, 0x0a, 0x17, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_flow_grpc_namespaces_proto_rawDescOnce sync.Once
	file_pkg_flow_grpc_namespaces_proto_rawDescData = file_pkg_flow_grpc_namespaces_proto_rawDesc
)

func file_pkg_flow_grpc_namespaces_proto_rawDescGZIP() []byte {
	file_pkg_flow_grpc_namespaces_proto_rawDescOnce.Do(func() {
		file_pkg_flow_grpc_namespaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flow_grpc_namespaces_proto_rawDescData)
	})
	return file_pkg_flow_grpc_namespaces_proto_rawDescData
}

var file_pkg_flow_grpc_namespaces_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_flow_grpc_namespaces_proto_goTypes = []interface{}{
	(*Namespace)(nil),                  // 0: direktiv_flow.Namespace
	(*ResolveNamespaceUIDRequest)(nil), // 1: direktiv_flow.ResolveNamespaceUIDRequest
	(*NamespaceRequest)(nil),           // 2: direktiv_flow.NamespaceRequest
	(*NamespaceResponse)(nil),          // 3: direktiv_flow.NamespaceResponse
	(*NamespacesRequest)(nil),          // 4: direktiv_flow.NamespacesRequest
	(*NamespacesResponse)(nil),         // 5: direktiv_flow.NamespacesResponse
	(*CreateNamespaceRequest)(nil),     // 6: direktiv_flow.CreateNamespaceRequest
	(*CreateNamespaceResponse)(nil),    // 7: direktiv_flow.CreateNamespaceResponse
	(*DeleteNamespaceRequest)(nil),     // 8: direktiv_flow.DeleteNamespaceRequest
	(*RenameNamespaceRequest)(nil),     // 9: direktiv_flow.RenameNamespaceRequest
	(*RenameNamespaceResponse)(nil),    // 10: direktiv_flow.RenameNamespaceResponse
	(*timestamp.Timestamp)(nil),        // 11: google.protobuf.Timestamp
	(*Pagination)(nil),                 // 12: direktiv_flow.Pagination
	(*PageInfo)(nil),                   // 13: direktiv_flow.PageInfo
}
var file_pkg_flow_grpc_namespaces_proto_depIdxs = []int32{
	11, // 0: direktiv_flow.Namespace.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: direktiv_flow.Namespace.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: direktiv_flow.NamespaceResponse.namespace:type_name -> direktiv_flow.Namespace
	12, // 3: direktiv_flow.NamespacesRequest.pagination:type_name -> direktiv_flow.Pagination
	13, // 4: direktiv_flow.NamespacesResponse.pageInfo:type_name -> direktiv_flow.PageInfo
	0,  // 5: direktiv_flow.NamespacesResponse.results:type_name -> direktiv_flow.Namespace
	0,  // 6: direktiv_flow.CreateNamespaceResponse.namespace:type_name -> direktiv_flow.Namespace
	0,  // 7: direktiv_flow.RenameNamespaceResponse.namespace:type_name -> direktiv_flow.Namespace
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_flow_grpc_namespaces_proto_init() }
func file_pkg_flow_grpc_namespaces_proto_init() {
	if File_pkg_flow_grpc_namespaces_proto != nil {
		return
	}
	file_pkg_flow_grpc_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_flow_grpc_namespaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveNamespaceUIDRequest); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceRequest); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceResponse); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespacesRequest); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespacesResponse); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceResponse); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameNamespaceRequest); i {
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
		file_pkg_flow_grpc_namespaces_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameNamespaceResponse); i {
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
			RawDescriptor: file_pkg_flow_grpc_namespaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_flow_grpc_namespaces_proto_goTypes,
		DependencyIndexes: file_pkg_flow_grpc_namespaces_proto_depIdxs,
		MessageInfos:      file_pkg_flow_grpc_namespaces_proto_msgTypes,
	}.Build()
	File_pkg_flow_grpc_namespaces_proto = out.File
	file_pkg_flow_grpc_namespaces_proto_rawDesc = nil
	file_pkg_flow_grpc_namespaces_proto_goTypes = nil
	file_pkg_flow_grpc_namespaces_proto_depIdxs = nil
}

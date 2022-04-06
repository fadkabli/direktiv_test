// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/functions/grpc/functions-common.proto

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

type BaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Namespace     *string `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Workflow      *string `protobuf:"bytes,3,opt,name=workflow,proto3,oneof" json:"workflow,omitempty"`
	Image         *string `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Cmd           *string `protobuf:"bytes,5,opt,name=cmd,proto3,oneof" json:"cmd,omitempty"`
	Size          *int32  `protobuf:"varint,6,opt,name=size,proto3,oneof" json:"size,omitempty"`
	MinScale      *int32  `protobuf:"varint,7,opt,name=minScale,proto3,oneof" json:"minScale,omitempty"`
	NamespaceName *string `protobuf:"bytes,8,opt,name=namespaceName,proto3,oneof" json:"namespaceName,omitempty"`
	Path          *string `protobuf:"bytes,9,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Revision      *string `protobuf:"bytes,10,opt,name=revision,proto3,oneof" json:"revision,omitempty"`
}

func (x *BaseInfo) Reset() {
	*x = BaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseInfo) ProtoMessage() {}

func (x *BaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseInfo.ProtoReflect.Descriptor instead.
func (*BaseInfo) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_common_proto_rawDescGZIP(), []int{0}
}

func (x *BaseInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BaseInfo) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *BaseInfo) GetWorkflow() string {
	if x != nil && x.Workflow != nil {
		return *x.Workflow
	}
	return ""
}

func (x *BaseInfo) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *BaseInfo) GetCmd() string {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return ""
}

func (x *BaseInfo) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *BaseInfo) GetMinScale() int32 {
	if x != nil && x.MinScale != nil {
		return *x.MinScale
	}
	return 0
}

func (x *BaseInfo) GetNamespaceName() string {
	if x != nil && x.NamespaceName != nil {
		return *x.NamespaceName
	}
	return ""
}

func (x *BaseInfo) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *BaseInfo) GetRevision() string {
	if x != nil && x.Revision != nil {
		return *x.Revision
	}
	return ""
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Status  *string `protobuf:"bytes,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Reason  *string `protobuf:"bytes,3,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	Message *string `protobuf:"bytes,4,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_common_proto_rawDescGZIP(), []int{1}
}

func (x *Condition) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Condition) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Condition) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *Condition) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type FunctionsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maxscale *int32 `protobuf:"varint,1,opt,name=maxscale,proto3,oneof" json:"maxscale,omitempty"`
}

func (x *FunctionsConfig) Reset() {
	*x = FunctionsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsConfig) ProtoMessage() {}

func (x *FunctionsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsConfig.ProtoReflect.Descriptor instead.
func (*FunctionsConfig) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_common_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionsConfig) GetMaxscale() int32 {
	if x != nil && x.Maxscale != nil {
		return *x.Maxscale
	}
	return 0
}

type FunctionsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info        *BaseInfo    `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
	Status      *string      `protobuf:"bytes,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Conditions  []*Condition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
	ServiceName *string      `protobuf:"bytes,4,opt,name=serviceName,proto3,oneof" json:"serviceName,omitempty"`
}

func (x *FunctionsInfo) Reset() {
	*x = FunctionsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsInfo) ProtoMessage() {}

func (x *FunctionsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsInfo.ProtoReflect.Descriptor instead.
func (*FunctionsInfo) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_common_proto_rawDescGZIP(), []int{3}
}

func (x *FunctionsInfo) GetInfo() *BaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *FunctionsInfo) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *FunctionsInfo) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *FunctionsInfo) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

type PodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Status          *string `protobuf:"bytes,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	ServiceName     *string `protobuf:"bytes,3,opt,name=serviceName,proto3,oneof" json:"serviceName,omitempty"`
	ServiceRevision *string `protobuf:"bytes,4,opt,name=serviceRevision,proto3,oneof" json:"serviceRevision,omitempty"`
}

func (x *PodsInfo) Reset() {
	*x = PodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodsInfo) ProtoMessage() {}

func (x *PodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodsInfo.ProtoReflect.Descriptor instead.
func (*PodsInfo) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_common_proto_rawDescGZIP(), []int{4}
}

func (x *PodsInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PodsInfo) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *PodsInfo) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *PodsInfo) GetServiceRevision() string {
	if x != nil && x.ServiceRevision != nil {
		return *x.ServiceRevision
	}
	return ""
}

type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            *string      `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Image           *string      `protobuf:"bytes,2,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Cmd             *string      `protobuf:"bytes,3,opt,name=cmd,proto3,oneof" json:"cmd,omitempty"`
	Size            *int32       `protobuf:"varint,4,opt,name=size,proto3,oneof" json:"size,omitempty"`
	MinScale        *int32       `protobuf:"varint,5,opt,name=minScale,proto3,oneof" json:"minScale,omitempty"`
	Generation      *int64       `protobuf:"varint,6,opt,name=generation,proto3,oneof" json:"generation,omitempty"`
	Created         *int64       `protobuf:"varint,7,opt,name=created,proto3,oneof" json:"created,omitempty"`
	Status          *string      `protobuf:"bytes,8,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Conditions      []*Condition `protobuf:"bytes,9,rep,name=conditions,proto3" json:"conditions,omitempty"`
	Traffic         *int64       `protobuf:"varint,10,opt,name=traffic,proto3,oneof" json:"traffic,omitempty"`
	DesiredReplicas int64        `protobuf:"varint,11,opt,name=desiredReplicas,proto3" json:"desiredReplicas,omitempty"`
	ActualReplicas  int64        `protobuf:"varint,12,opt,name=actualReplicas,proto3" json:"actualReplicas,omitempty"`
	Rev             *string      `protobuf:"bytes,13,opt,name=rev,proto3,oneof" json:"rev,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_common_proto_rawDescGZIP(), []int{5}
}

func (x *Revision) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Revision) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *Revision) GetCmd() string {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return ""
}

func (x *Revision) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *Revision) GetMinScale() int32 {
	if x != nil && x.MinScale != nil {
		return *x.MinScale
	}
	return 0
}

func (x *Revision) GetGeneration() int64 {
	if x != nil && x.Generation != nil {
		return *x.Generation
	}
	return 0
}

func (x *Revision) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

func (x *Revision) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Revision) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Revision) GetTraffic() int64 {
	if x != nil && x.Traffic != nil {
		return *x.Traffic
	}
	return 0
}

func (x *Revision) GetDesiredReplicas() int64 {
	if x != nil {
		return x.DesiredReplicas
	}
	return 0
}

func (x *Revision) GetActualReplicas() int64 {
	if x != nil {
		return x.ActualReplicas
	}
	return 0
}

func (x *Revision) GetRev() string {
	if x != nil && x.Rev != nil {
		return *x.Rev
	}
	return ""
}

var File_pkg_functions_grpc_functions_common_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_functions_common_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xac, 0x03, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0d,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x6d, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x69,
	0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa8,
	0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x0f, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x08,
	0x6d, 0x61, 0x78, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x08, 0x6d, 0x61, 0x78, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x6d, 0x61, 0x78, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x0d, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x25, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x50,
	0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x04, 0x0a, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x63, 0x6d, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04,
	0x52, 0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x07, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a,
	0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x08,
	0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0f,
	0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x15,
	0x0a, 0x03, 0x72, 0x65, 0x76, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x03, 0x72,
	0x65, 0x76, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x6d, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x69,
	0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x76,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_functions_common_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_functions_common_proto_rawDescData = file_pkg_functions_grpc_functions_common_proto_rawDesc
)

func file_pkg_functions_grpc_functions_common_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_functions_common_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_functions_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_functions_common_proto_rawDescData)
	})
	return file_pkg_functions_grpc_functions_common_proto_rawDescData
}

var file_pkg_functions_grpc_functions_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_functions_grpc_functions_common_proto_goTypes = []interface{}{
	(*BaseInfo)(nil),        // 0: direktiv_functions.BaseInfo
	(*Condition)(nil),       // 1: direktiv_functions.Condition
	(*FunctionsConfig)(nil), // 2: direktiv_functions.FunctionsConfig
	(*FunctionsInfo)(nil),   // 3: direktiv_functions.FunctionsInfo
	(*PodsInfo)(nil),        // 4: direktiv_functions.PodsInfo
	(*Revision)(nil),        // 5: direktiv_functions.Revision
}
var file_pkg_functions_grpc_functions_common_proto_depIdxs = []int32{
	0, // 0: direktiv_functions.FunctionsInfo.info:type_name -> direktiv_functions.BaseInfo
	1, // 1: direktiv_functions.FunctionsInfo.conditions:type_name -> direktiv_functions.Condition
	1, // 2: direktiv_functions.Revision.conditions:type_name -> direktiv_functions.Condition
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_functions_common_proto_init() }
func file_pkg_functions_grpc_functions_common_proto_init() {
	if File_pkg_functions_grpc_functions_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_functions_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseInfo); i {
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
		file_pkg_functions_grpc_functions_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_pkg_functions_grpc_functions_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsConfig); i {
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
		file_pkg_functions_grpc_functions_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsInfo); i {
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
		file_pkg_functions_grpc_functions_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodsInfo); i {
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
		file_pkg_functions_grpc_functions_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
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
	file_pkg_functions_grpc_functions_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_functions_common_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_functions_common_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_functions_common_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_functions_common_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_functions_common_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_functions_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_functions_common_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_functions_common_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_functions_common_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_functions_common_proto = out.File
	file_pkg_functions_grpc_functions_common_proto_rawDesc = nil
	file_pkg_functions_grpc_functions_common_proto_goTypes = nil
	file_pkg_functions_grpc_functions_common_proto_depIdxs = nil
}

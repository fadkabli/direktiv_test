// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FunctionsServiceClient is the client API for FunctionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FunctionsServiceClient interface {
	StoreRegistry(ctx context.Context, in *StoreRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetRegistries(ctx context.Context, in *GetRegistriesRequest, opts ...grpc.CallOption) (*GetRegistriesResponse, error)
	DeleteRegistry(ctx context.Context, in *DeleteRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReconstructFunction(ctx context.Context, in *ReconstructFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateFunction(ctx context.Context, in *UpdateFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateFunction(ctx context.Context, in *CreateFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error)
	GetFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*GetFunctionResponse, error)
	DeleteFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetFunctionsTraffic(ctx context.Context, in *SetTrafficRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateFunctionsPod(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (*CreatePodResponse, error)
	CancelFunctionsPod(ctx context.Context, in *CancelPodRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteRevision(ctx context.Context, in *DeleteRevisionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	WatchFunctions(ctx context.Context, in *WatchFunctionsRequest, opts ...grpc.CallOption) (FunctionsService_WatchFunctionsClient, error)
	WatchPods(ctx context.Context, in *WatchPodsRequest, opts ...grpc.CallOption) (FunctionsService_WatchPodsClient, error)
	WatchRevisions(ctx context.Context, in *WatchRevisionsRequest, opts ...grpc.CallOption) (FunctionsService_WatchRevisionsClient, error)
	WatchLogs(ctx context.Context, in *WatchLogsRequest, opts ...grpc.CallOption) (FunctionsService_WatchLogsClient, error)
	ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	Build(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BuildResponse, error)
}

type functionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFunctionsServiceClient(cc grpc.ClientConnInterface) FunctionsServiceClient {
	return &functionsServiceClient{cc}
}

func (c *functionsServiceClient) StoreRegistry(ctx context.Context, in *StoreRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/StoreRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) GetRegistries(ctx context.Context, in *GetRegistriesRequest, opts ...grpc.CallOption) (*GetRegistriesResponse, error) {
	out := new(GetRegistriesResponse)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/GetRegistries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) DeleteRegistry(ctx context.Context, in *DeleteRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/DeleteRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) ReconstructFunction(ctx context.Context, in *ReconstructFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/ReconstructFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) UpdateFunction(ctx context.Context, in *UpdateFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/UpdateFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) CreateFunction(ctx context.Context, in *CreateFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/CreateFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) DeleteFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/DeleteFunctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error) {
	out := new(ListFunctionsResponse)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/ListFunctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) GetFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*GetFunctionResponse, error) {
	out := new(GetFunctionResponse)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/GetFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) DeleteFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/DeleteFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) SetFunctionsTraffic(ctx context.Context, in *SetTrafficRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/SetFunctionsTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) CreateFunctionsPod(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (*CreatePodResponse, error) {
	out := new(CreatePodResponse)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/CreateFunctionsPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) CancelFunctionsPod(ctx context.Context, in *CancelPodRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/CancelFunctionsPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) DeleteRevision(ctx context.Context, in *DeleteRevisionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/DeleteRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) WatchFunctions(ctx context.Context, in *WatchFunctionsRequest, opts ...grpc.CallOption) (FunctionsService_WatchFunctionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FunctionsService_ServiceDesc.Streams[0], "/direktiv_functions.FunctionsService/WatchFunctions", opts...)
	if err != nil {
		return nil, err
	}
	x := &functionsServiceWatchFunctionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FunctionsService_WatchFunctionsClient interface {
	Recv() (*WatchFunctionsResponse, error)
	grpc.ClientStream
}

type functionsServiceWatchFunctionsClient struct {
	grpc.ClientStream
}

func (x *functionsServiceWatchFunctionsClient) Recv() (*WatchFunctionsResponse, error) {
	m := new(WatchFunctionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *functionsServiceClient) WatchPods(ctx context.Context, in *WatchPodsRequest, opts ...grpc.CallOption) (FunctionsService_WatchPodsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FunctionsService_ServiceDesc.Streams[1], "/direktiv_functions.FunctionsService/WatchPods", opts...)
	if err != nil {
		return nil, err
	}
	x := &functionsServiceWatchPodsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FunctionsService_WatchPodsClient interface {
	Recv() (*WatchPodsResponse, error)
	grpc.ClientStream
}

type functionsServiceWatchPodsClient struct {
	grpc.ClientStream
}

func (x *functionsServiceWatchPodsClient) Recv() (*WatchPodsResponse, error) {
	m := new(WatchPodsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *functionsServiceClient) WatchRevisions(ctx context.Context, in *WatchRevisionsRequest, opts ...grpc.CallOption) (FunctionsService_WatchRevisionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FunctionsService_ServiceDesc.Streams[2], "/direktiv_functions.FunctionsService/WatchRevisions", opts...)
	if err != nil {
		return nil, err
	}
	x := &functionsServiceWatchRevisionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FunctionsService_WatchRevisionsClient interface {
	Recv() (*WatchRevisionsResponse, error)
	grpc.ClientStream
}

type functionsServiceWatchRevisionsClient struct {
	grpc.ClientStream
}

func (x *functionsServiceWatchRevisionsClient) Recv() (*WatchRevisionsResponse, error) {
	m := new(WatchRevisionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *functionsServiceClient) WatchLogs(ctx context.Context, in *WatchLogsRequest, opts ...grpc.CallOption) (FunctionsService_WatchLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FunctionsService_ServiceDesc.Streams[3], "/direktiv_functions.FunctionsService/WatchLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &functionsServiceWatchLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FunctionsService_WatchLogsClient interface {
	Recv() (*WatchLogsResponse, error)
	grpc.ClientStream
}

type functionsServiceWatchLogsClient struct {
	grpc.ClientStream
}

func (x *functionsServiceWatchLogsClient) Recv() (*WatchLogsResponse, error) {
	m := new(WatchLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *functionsServiceClient) ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/ListPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsServiceClient) Build(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/direktiv_functions.FunctionsService/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FunctionsServiceServer is the server API for FunctionsService service.
// All implementations must embed UnimplementedFunctionsServiceServer
// for forward compatibility
type FunctionsServiceServer interface {
	StoreRegistry(context.Context, *StoreRegistryRequest) (*empty.Empty, error)
	GetRegistries(context.Context, *GetRegistriesRequest) (*GetRegistriesResponse, error)
	DeleteRegistry(context.Context, *DeleteRegistryRequest) (*empty.Empty, error)
	ReconstructFunction(context.Context, *ReconstructFunctionRequest) (*empty.Empty, error)
	UpdateFunction(context.Context, *UpdateFunctionRequest) (*empty.Empty, error)
	CreateFunction(context.Context, *CreateFunctionRequest) (*empty.Empty, error)
	DeleteFunctions(context.Context, *ListFunctionsRequest) (*empty.Empty, error)
	ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error)
	GetFunction(context.Context, *GetFunctionRequest) (*GetFunctionResponse, error)
	DeleteFunction(context.Context, *GetFunctionRequest) (*empty.Empty, error)
	SetFunctionsTraffic(context.Context, *SetTrafficRequest) (*empty.Empty, error)
	CreateFunctionsPod(context.Context, *CreatePodRequest) (*CreatePodResponse, error)
	CancelFunctionsPod(context.Context, *CancelPodRequest) (*empty.Empty, error)
	DeleteRevision(context.Context, *DeleteRevisionRequest) (*empty.Empty, error)
	WatchFunctions(*WatchFunctionsRequest, FunctionsService_WatchFunctionsServer) error
	WatchPods(*WatchPodsRequest, FunctionsService_WatchPodsServer) error
	WatchRevisions(*WatchRevisionsRequest, FunctionsService_WatchRevisionsServer) error
	WatchLogs(*WatchLogsRequest, FunctionsService_WatchLogsServer) error
	ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	Build(context.Context, *empty.Empty) (*BuildResponse, error)
	mustEmbedUnimplementedFunctionsServiceServer()
}

// UnimplementedFunctionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFunctionsServiceServer struct {
}

func (UnimplementedFunctionsServiceServer) StoreRegistry(context.Context, *StoreRegistryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRegistry not implemented")
}
func (UnimplementedFunctionsServiceServer) GetRegistries(context.Context, *GetRegistriesRequest) (*GetRegistriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistries not implemented")
}
func (UnimplementedFunctionsServiceServer) DeleteRegistry(context.Context, *DeleteRegistryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegistry not implemented")
}
func (UnimplementedFunctionsServiceServer) ReconstructFunction(context.Context, *ReconstructFunctionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReconstructFunction not implemented")
}
func (UnimplementedFunctionsServiceServer) UpdateFunction(context.Context, *UpdateFunctionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFunction not implemented")
}
func (UnimplementedFunctionsServiceServer) CreateFunction(context.Context, *CreateFunctionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFunction not implemented")
}
func (UnimplementedFunctionsServiceServer) DeleteFunctions(context.Context, *ListFunctionsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFunctions not implemented")
}
func (UnimplementedFunctionsServiceServer) ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFunctions not implemented")
}
func (UnimplementedFunctionsServiceServer) GetFunction(context.Context, *GetFunctionRequest) (*GetFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFunction not implemented")
}
func (UnimplementedFunctionsServiceServer) DeleteFunction(context.Context, *GetFunctionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFunction not implemented")
}
func (UnimplementedFunctionsServiceServer) SetFunctionsTraffic(context.Context, *SetTrafficRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFunctionsTraffic not implemented")
}
func (UnimplementedFunctionsServiceServer) CreateFunctionsPod(context.Context, *CreatePodRequest) (*CreatePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFunctionsPod not implemented")
}
func (UnimplementedFunctionsServiceServer) CancelFunctionsPod(context.Context, *CancelPodRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelFunctionsPod not implemented")
}
func (UnimplementedFunctionsServiceServer) DeleteRevision(context.Context, *DeleteRevisionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRevision not implemented")
}
func (UnimplementedFunctionsServiceServer) WatchFunctions(*WatchFunctionsRequest, FunctionsService_WatchFunctionsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchFunctions not implemented")
}
func (UnimplementedFunctionsServiceServer) WatchPods(*WatchPodsRequest, FunctionsService_WatchPodsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPods not implemented")
}
func (UnimplementedFunctionsServiceServer) WatchRevisions(*WatchRevisionsRequest, FunctionsService_WatchRevisionsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchRevisions not implemented")
}
func (UnimplementedFunctionsServiceServer) WatchLogs(*WatchLogsRequest, FunctionsService_WatchLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchLogs not implemented")
}
func (UnimplementedFunctionsServiceServer) ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPods not implemented")
}
func (UnimplementedFunctionsServiceServer) Build(context.Context, *empty.Empty) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (UnimplementedFunctionsServiceServer) mustEmbedUnimplementedFunctionsServiceServer() {}

// UnsafeFunctionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FunctionsServiceServer will
// result in compilation errors.
type UnsafeFunctionsServiceServer interface {
	mustEmbedUnimplementedFunctionsServiceServer()
}

func RegisterFunctionsServiceServer(s grpc.ServiceRegistrar, srv FunctionsServiceServer) {
	s.RegisterService(&FunctionsService_ServiceDesc, srv)
}

func _FunctionsService_StoreRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).StoreRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/StoreRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).StoreRegistry(ctx, req.(*StoreRegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_GetRegistries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).GetRegistries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/GetRegistries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).GetRegistries(ctx, req.(*GetRegistriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_DeleteRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).DeleteRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/DeleteRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).DeleteRegistry(ctx, req.(*DeleteRegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_ReconstructFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReconstructFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).ReconstructFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/ReconstructFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).ReconstructFunction(ctx, req.(*ReconstructFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_UpdateFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).UpdateFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/UpdateFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).UpdateFunction(ctx, req.(*UpdateFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_CreateFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).CreateFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/CreateFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).CreateFunction(ctx, req.(*CreateFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_DeleteFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFunctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).DeleteFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/DeleteFunctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).DeleteFunctions(ctx, req.(*ListFunctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_ListFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFunctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).ListFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/ListFunctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).ListFunctions(ctx, req.(*ListFunctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_GetFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).GetFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/GetFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).GetFunction(ctx, req.(*GetFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_DeleteFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).DeleteFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/DeleteFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).DeleteFunction(ctx, req.(*GetFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_SetFunctionsTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).SetFunctionsTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/SetFunctionsTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).SetFunctionsTraffic(ctx, req.(*SetTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_CreateFunctionsPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).CreateFunctionsPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/CreateFunctionsPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).CreateFunctionsPod(ctx, req.(*CreatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_CancelFunctionsPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).CancelFunctionsPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/CancelFunctionsPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).CancelFunctionsPod(ctx, req.(*CancelPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_DeleteRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).DeleteRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/DeleteRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).DeleteRevision(ctx, req.(*DeleteRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_WatchFunctions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchFunctionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunctionsServiceServer).WatchFunctions(m, &functionsServiceWatchFunctionsServer{stream})
}

type FunctionsService_WatchFunctionsServer interface {
	Send(*WatchFunctionsResponse) error
	grpc.ServerStream
}

type functionsServiceWatchFunctionsServer struct {
	grpc.ServerStream
}

func (x *functionsServiceWatchFunctionsServer) Send(m *WatchFunctionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FunctionsService_WatchPods_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPodsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunctionsServiceServer).WatchPods(m, &functionsServiceWatchPodsServer{stream})
}

type FunctionsService_WatchPodsServer interface {
	Send(*WatchPodsResponse) error
	grpc.ServerStream
}

type functionsServiceWatchPodsServer struct {
	grpc.ServerStream
}

func (x *functionsServiceWatchPodsServer) Send(m *WatchPodsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FunctionsService_WatchRevisions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRevisionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunctionsServiceServer).WatchRevisions(m, &functionsServiceWatchRevisionsServer{stream})
}

type FunctionsService_WatchRevisionsServer interface {
	Send(*WatchRevisionsResponse) error
	grpc.ServerStream
}

type functionsServiceWatchRevisionsServer struct {
	grpc.ServerStream
}

func (x *functionsServiceWatchRevisionsServer) Send(m *WatchRevisionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FunctionsService_WatchLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunctionsServiceServer).WatchLogs(m, &functionsServiceWatchLogsServer{stream})
}

type FunctionsService_WatchLogsServer interface {
	Send(*WatchLogsResponse) error
	grpc.ServerStream
}

type functionsServiceWatchLogsServer struct {
	grpc.ServerStream
}

func (x *functionsServiceWatchLogsServer) Send(m *WatchLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FunctionsService_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).ListPods(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FunctionsService_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/direktiv_functions.FunctionsService/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServiceServer).Build(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FunctionsService_ServiceDesc is the grpc.ServiceDesc for FunctionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FunctionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "direktiv_functions.FunctionsService",
	HandlerType: (*FunctionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreRegistry",
			Handler:    _FunctionsService_StoreRegistry_Handler,
		},
		{
			MethodName: "GetRegistries",
			Handler:    _FunctionsService_GetRegistries_Handler,
		},
		{
			MethodName: "DeleteRegistry",
			Handler:    _FunctionsService_DeleteRegistry_Handler,
		},
		{
			MethodName: "ReconstructFunction",
			Handler:    _FunctionsService_ReconstructFunction_Handler,
		},
		{
			MethodName: "UpdateFunction",
			Handler:    _FunctionsService_UpdateFunction_Handler,
		},
		{
			MethodName: "CreateFunction",
			Handler:    _FunctionsService_CreateFunction_Handler,
		},
		{
			MethodName: "DeleteFunctions",
			Handler:    _FunctionsService_DeleteFunctions_Handler,
		},
		{
			MethodName: "ListFunctions",
			Handler:    _FunctionsService_ListFunctions_Handler,
		},
		{
			MethodName: "GetFunction",
			Handler:    _FunctionsService_GetFunction_Handler,
		},
		{
			MethodName: "DeleteFunction",
			Handler:    _FunctionsService_DeleteFunction_Handler,
		},
		{
			MethodName: "SetFunctionsTraffic",
			Handler:    _FunctionsService_SetFunctionsTraffic_Handler,
		},
		{
			MethodName: "CreateFunctionsPod",
			Handler:    _FunctionsService_CreateFunctionsPod_Handler,
		},
		{
			MethodName: "CancelFunctionsPod",
			Handler:    _FunctionsService_CancelFunctionsPod_Handler,
		},
		{
			MethodName: "DeleteRevision",
			Handler:    _FunctionsService_DeleteRevision_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _FunctionsService_ListPods_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _FunctionsService_Build_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchFunctions",
			Handler:       _FunctionsService_WatchFunctions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchPods",
			Handler:       _FunctionsService_WatchPods_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchRevisions",
			Handler:       _FunctionsService_WatchRevisions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchLogs",
			Handler:       _FunctionsService_WatchLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/functions/grpc/protocol.proto",
}

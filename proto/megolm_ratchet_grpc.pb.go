// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: megolm_ratchet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MegolmRatchetService_PostData_FullMethodName   = "/megolmratchet.MegolmRatchetService/PostData"
	MegolmRatchetService_GetData_FullMethodName    = "/megolmratchet.MegolmRatchetService/GetData"
	MegolmRatchetService_AppendData_FullMethodName = "/megolmratchet.MegolmRatchetService/AppendData"
)

// MegolmRatchetServiceClient is the client API for MegolmRatchetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MegolmRatchetServiceClient interface {
	PostData(ctx context.Context, in *PostDataRequest, opts ...grpc.CallOption) (*PostDataResponse, error)
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
	AppendData(ctx context.Context, in *AppendDataRequest, opts ...grpc.CallOption) (*AppendDataResponse, error)
}

type megolmRatchetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMegolmRatchetServiceClient(cc grpc.ClientConnInterface) MegolmRatchetServiceClient {
	return &megolmRatchetServiceClient{cc}
}

func (c *megolmRatchetServiceClient) PostData(ctx context.Context, in *PostDataRequest, opts ...grpc.CallOption) (*PostDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostDataResponse)
	err := c.cc.Invoke(ctx, MegolmRatchetService_PostData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *megolmRatchetServiceClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, MegolmRatchetService_GetData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *megolmRatchetServiceClient) AppendData(ctx context.Context, in *AppendDataRequest, opts ...grpc.CallOption) (*AppendDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppendDataResponse)
	err := c.cc.Invoke(ctx, MegolmRatchetService_AppendData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MegolmRatchetServiceServer is the server API for MegolmRatchetService service.
// All implementations must embed UnimplementedMegolmRatchetServiceServer
// for forward compatibility.
type MegolmRatchetServiceServer interface {
	PostData(context.Context, *PostDataRequest) (*PostDataResponse, error)
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	AppendData(context.Context, *AppendDataRequest) (*AppendDataResponse, error)
	mustEmbedUnimplementedMegolmRatchetServiceServer()
}

// UnimplementedMegolmRatchetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMegolmRatchetServiceServer struct{}

func (UnimplementedMegolmRatchetServiceServer) PostData(context.Context, *PostDataRequest) (*PostDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (UnimplementedMegolmRatchetServiceServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedMegolmRatchetServiceServer) AppendData(context.Context, *AppendDataRequest) (*AppendDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendData not implemented")
}
func (UnimplementedMegolmRatchetServiceServer) mustEmbedUnimplementedMegolmRatchetServiceServer() {}
func (UnimplementedMegolmRatchetServiceServer) testEmbeddedByValue()                              {}

// UnsafeMegolmRatchetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MegolmRatchetServiceServer will
// result in compilation errors.
type UnsafeMegolmRatchetServiceServer interface {
	mustEmbedUnimplementedMegolmRatchetServiceServer()
}

func RegisterMegolmRatchetServiceServer(s grpc.ServiceRegistrar, srv MegolmRatchetServiceServer) {
	// If the following call pancis, it indicates UnimplementedMegolmRatchetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MegolmRatchetService_ServiceDesc, srv)
}

func _MegolmRatchetService_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MegolmRatchetServiceServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MegolmRatchetService_PostData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MegolmRatchetServiceServer).PostData(ctx, req.(*PostDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MegolmRatchetService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MegolmRatchetServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MegolmRatchetService_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MegolmRatchetServiceServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MegolmRatchetService_AppendData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MegolmRatchetServiceServer).AppendData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MegolmRatchetService_AppendData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MegolmRatchetServiceServer).AppendData(ctx, req.(*AppendDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MegolmRatchetService_ServiceDesc is the grpc.ServiceDesc for MegolmRatchetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MegolmRatchetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "megolmratchet.MegolmRatchetService",
	HandlerType: (*MegolmRatchetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostData",
			Handler:    _MegolmRatchetService_PostData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _MegolmRatchetService_GetData_Handler,
		},
		{
			MethodName: "AppendData",
			Handler:    _MegolmRatchetService_AppendData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "megolm_ratchet.proto",
}

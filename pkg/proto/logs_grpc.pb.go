// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/logs.proto

package logs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LogsClient is the client API for Logs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogsClient interface {
	Insert(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error)
}

type logsClient struct {
	cc grpc.ClientConnInterface
}

func NewLogsClient(cc grpc.ClientConnInterface) LogsClient {
	return &logsClient{cc}
}

func (c *logsClient) Insert(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/log.Logs/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServer is the server API for Logs service.
// All implementations should embed UnimplementedLogsServer
// for forward compatibility
type LogsServer interface {
	Insert(context.Context, *LogRequest) (*Empty, error)
}

// UnimplementedLogsServer should be embedded to have forward compatible implementations.
type UnimplementedLogsServer struct {
}

func (UnimplementedLogsServer) Insert(context.Context, *LogRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

// UnsafeLogsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogsServer will
// result in compilation errors.
type UnsafeLogsServer interface {
	mustEmbedUnimplementedLogsServer()
}

func RegisterLogsServer(s grpc.ServiceRegistrar, srv LogsServer) {
	s.RegisterService(&Logs_ServiceDesc, srv)
}

func _Logs_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.Logs/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).Insert(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Logs_ServiceDesc is the grpc.ServiceDesc for Logs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "log.Logs",
	HandlerType: (*LogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Logs_Insert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logs.proto",
}
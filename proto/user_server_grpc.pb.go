// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: user_server.proto

package __

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

// UserNameClient is the client API for UserName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserNameClient interface {
	GetUserByName(ctx context.Context, in *Username, opts ...grpc.CallOption) (*User, error)
}

type userNameClient struct {
	cc grpc.ClientConnInterface
}

func NewUserNameClient(cc grpc.ClientConnInterface) UserNameClient {
	return &userNameClient{cc}
}

func (c *userNameClient) GetUserByName(ctx context.Context, in *Username, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_server.UserName/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserNameServer is the server API for UserName service.
// All implementations must embed UnimplementedUserNameServer
// for forward compatibility
type UserNameServer interface {
	GetUserByName(context.Context, *Username) (*User, error)
	mustEmbedUnimplementedUserNameServer()
}

// UnimplementedUserNameServer must be embedded to have forward compatible implementations.
type UnimplementedUserNameServer struct {
}

func (UnimplementedUserNameServer) GetUserByName(context.Context, *Username) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedUserNameServer) mustEmbedUnimplementedUserNameServer() {}

// UnsafeUserNameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserNameServer will
// result in compilation errors.
type UnsafeUserNameServer interface {
	mustEmbedUnimplementedUserNameServer()
}

func RegisterUserNameServer(s grpc.ServiceRegistrar, srv UserNameServer) {
	s.RegisterService(&UserName_ServiceDesc, srv)
}

func _UserName_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNameServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserName/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNameServer).GetUserByName(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

// UserName_ServiceDesc is the grpc.ServiceDesc for UserName service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserName_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_server.UserName",
	HandlerType: (*UserNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByName",
			Handler:    _UserName_GetUserByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_server.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authsvc

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	AuthLogin(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginReply, error)
	AuthLogout(ctx context.Context, in *AuthLogoutRequest, opts ...grpc.CallOption) (*AuthLogoutReply, error)
	AuthVerifyToken(ctx context.Context, in *AuthVerifyTokenRequest, opts ...grpc.CallOption) (*AuthVerifyTokenReply, error)
	AuthChangeTimezone(ctx context.Context, in *AuthChangeTimezoneRequest, opts ...grpc.CallOption) (*AuthChangeTimezoneReply, error)
	AuthGetTimezone(ctx context.Context, in *AuthGetTimezoneRequest, opts ...grpc.CallOption) (*AuthGetTimezoneReply, error)
	AuthServiceStatus(ctx context.Context, in *AuthServiceStatusRequest, opts ...grpc.CallOption) (*AuthServiceStatusReply, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) AuthLogin(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginReply, error) {
	out := new(AuthLoginReply)
	err := c.cc.Invoke(ctx, "/pb.Auth/AuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthLogout(ctx context.Context, in *AuthLogoutRequest, opts ...grpc.CallOption) (*AuthLogoutReply, error) {
	out := new(AuthLogoutReply)
	err := c.cc.Invoke(ctx, "/pb.Auth/AuthLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthVerifyToken(ctx context.Context, in *AuthVerifyTokenRequest, opts ...grpc.CallOption) (*AuthVerifyTokenReply, error) {
	out := new(AuthVerifyTokenReply)
	err := c.cc.Invoke(ctx, "/pb.Auth/AuthVerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthChangeTimezone(ctx context.Context, in *AuthChangeTimezoneRequest, opts ...grpc.CallOption) (*AuthChangeTimezoneReply, error) {
	out := new(AuthChangeTimezoneReply)
	err := c.cc.Invoke(ctx, "/pb.Auth/AuthChangeTimezone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthGetTimezone(ctx context.Context, in *AuthGetTimezoneRequest, opts ...grpc.CallOption) (*AuthGetTimezoneReply, error) {
	out := new(AuthGetTimezoneReply)
	err := c.cc.Invoke(ctx, "/pb.Auth/AuthGetTimezone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthServiceStatus(ctx context.Context, in *AuthServiceStatusRequest, opts ...grpc.CallOption) (*AuthServiceStatusReply, error) {
	out := new(AuthServiceStatusReply)
	err := c.cc.Invoke(ctx, "/pb.Auth/AuthServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	AuthLogin(context.Context, *AuthLoginRequest) (*AuthLoginReply, error)
	AuthLogout(context.Context, *AuthLogoutRequest) (*AuthLogoutReply, error)
	AuthVerifyToken(context.Context, *AuthVerifyTokenRequest) (*AuthVerifyTokenReply, error)
	AuthChangeTimezone(context.Context, *AuthChangeTimezoneRequest) (*AuthChangeTimezoneReply, error)
	AuthGetTimezone(context.Context, *AuthGetTimezoneRequest) (*AuthGetTimezoneReply, error)
	AuthServiceStatus(context.Context, *AuthServiceStatusRequest) (*AuthServiceStatusReply, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) AuthLogin(context.Context, *AuthLoginRequest) (*AuthLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthLogin not implemented")
}
func (UnimplementedAuthServer) AuthLogout(context.Context, *AuthLogoutRequest) (*AuthLogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthLogout not implemented")
}
func (UnimplementedAuthServer) AuthVerifyToken(context.Context, *AuthVerifyTokenRequest) (*AuthVerifyTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthVerifyToken not implemented")
}
func (UnimplementedAuthServer) AuthChangeTimezone(context.Context, *AuthChangeTimezoneRequest) (*AuthChangeTimezoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthChangeTimezone not implemented")
}
func (UnimplementedAuthServer) AuthGetTimezone(context.Context, *AuthGetTimezoneRequest) (*AuthGetTimezoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthGetTimezone not implemented")
}
func (UnimplementedAuthServer) AuthServiceStatus(context.Context, *AuthServiceStatusRequest) (*AuthServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthServiceStatus not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_AuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/AuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthLogin(ctx, req.(*AuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/AuthLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthLogout(ctx, req.(*AuthLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthVerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthVerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthVerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/AuthVerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthVerifyToken(ctx, req.(*AuthVerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthChangeTimezone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthChangeTimezoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthChangeTimezone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/AuthChangeTimezone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthChangeTimezone(ctx, req.(*AuthChangeTimezoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthGetTimezone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthGetTimezoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthGetTimezone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/AuthGetTimezone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthGetTimezone(ctx, req.(*AuthGetTimezoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/AuthServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthServiceStatus(ctx, req.(*AuthServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthLogin",
			Handler:    _Auth_AuthLogin_Handler,
		},
		{
			MethodName: "AuthLogout",
			Handler:    _Auth_AuthLogout_Handler,
		},
		{
			MethodName: "AuthVerifyToken",
			Handler:    _Auth_AuthVerifyToken_Handler,
		},
		{
			MethodName: "AuthChangeTimezone",
			Handler:    _Auth_AuthChangeTimezone_Handler,
		},
		{
			MethodName: "AuthGetTimezone",
			Handler:    _Auth_AuthGetTimezone_Handler,
		},
		{
			MethodName: "AuthServiceStatus",
			Handler:    _Auth_AuthServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authsvc.proto",
}

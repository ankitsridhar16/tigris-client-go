// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: server/v1/auth.proto

package api

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
	// Gets an access token.
	// supports two grant_type: [refresh_token, client_credentials]
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error)
	// Create an application.
	CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error)
	// Delete an application.
	DeleteApplication(ctx context.Context, in *DeleteApplicationsRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error)
	// Lists all application visible to requesting actor.
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
	// RotateApplicationRequest returns the new application with rotated secret
	RotateApplicationSecret(ctx context.Context, in *RotateApplicationSecretRequest, opts ...grpc.CallOption) (*RotateApplicationSecretResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error) {
	out := new(GetAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.auth.v1.Auth/getAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error) {
	out := new(CreateApplicationResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.auth.v1.Auth/createApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteApplication(ctx context.Context, in *DeleteApplicationsRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error) {
	out := new(DeleteApplicationResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.auth.v1.Auth/deleteApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.auth.v1.Auth/listApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RotateApplicationSecret(ctx context.Context, in *RotateApplicationSecretRequest, opts ...grpc.CallOption) (*RotateApplicationSecretResponse, error) {
	out := new(RotateApplicationSecretResponse)
	err := c.cc.Invoke(ctx, "/tigrisdata.auth.v1.Auth/rotateApplicationSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// Gets an access token.
	// supports two grant_type: [refresh_token, client_credentials]
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	// Create an application.
	CreateApplication(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error)
	// Delete an application.
	DeleteApplication(context.Context, *DeleteApplicationsRequest) (*DeleteApplicationResponse, error)
	// Lists all application visible to requesting actor.
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
	// RotateApplicationRequest returns the new application with rotated secret
	RotateApplicationSecret(context.Context, *RotateApplicationSecretRequest) (*RotateApplicationSecretResponse, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedAuthServer) CreateApplication(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedAuthServer) DeleteApplication(context.Context, *DeleteApplicationsRequest) (*DeleteApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (UnimplementedAuthServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedAuthServer) RotateApplicationSecret(context.Context, *RotateApplicationSecretRequest) (*RotateApplicationSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateApplicationSecret not implemented")
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

func _Auth_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.auth.v1.Auth/getAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.auth.v1.Auth/createApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateApplication(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.auth.v1.Auth/deleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteApplication(ctx, req.(*DeleteApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.auth.v1.Auth/listApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RotateApplicationSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateApplicationSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RotateApplicationSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tigrisdata.auth.v1.Auth/rotateApplicationSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RotateApplicationSecret(ctx, req.(*RotateApplicationSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tigrisdata.auth.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAccessToken",
			Handler:    _Auth_GetAccessToken_Handler,
		},
		{
			MethodName: "createApplication",
			Handler:    _Auth_CreateApplication_Handler,
		},
		{
			MethodName: "deleteApplication",
			Handler:    _Auth_DeleteApplication_Handler,
		},
		{
			MethodName: "listApplications",
			Handler:    _Auth_ListApplications_Handler,
		},
		{
			MethodName: "rotateApplicationSecret",
			Handler:    _Auth_RotateApplicationSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/v1/auth.proto",
}

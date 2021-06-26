// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SyncTreeClient is the client API for SyncTree service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncTreeClient interface {
	Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Response, error)
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error)
	UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Response, error)
	UserSend(ctx context.Context, in *UserSendRequest, opts ...grpc.CallOption) (*Response, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UserCreateTrade(ctx context.Context, in *UserTradeRequest, opts ...grpc.CallOption) (*Response, error)
	UserCancelTrade(ctx context.Context, in *UserCancelTradeRequest, opts ...grpc.CallOption) (*Response, error)
	UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...grpc.CallOption) (*Response, error)
	UserWithdrawal(ctx context.Context, in *UserWithDrawalRequest, opts ...grpc.CallOption) (*Response, error)
	UserSearch(ctx context.Context, in *UserSearchRequest, opts ...grpc.CallOption) (*Markets, error)
	MarketCraete(ctx context.Context, in *MarketCreateRequest, opts ...grpc.CallOption) (*Response, error)
	MarketUpdate(ctx context.Context, in *MarketUpdateRequest, opts ...grpc.CallOption) (*Response, error)
	MarketInfo(ctx context.Context, in *MarketInfoRequest, opts ...grpc.CallOption) (*MarketInfoResponse, error)
	MarketDeposit(ctx context.Context, in *MarketDepositRequest, opts ...grpc.CallOption) (*Response, error)
	MarketWithDrawal(ctx context.Context, in *MarketWithDrawalRequest, opts ...grpc.CallOption) (*Response, error)
}

type syncTreeClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncTreeClient(cc grpc.ClientConnInterface) SyncTreeClient {
	return &syncTreeClient{cc}
}

func (c *syncTreeClient) Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserSend(ctx context.Context, in *UserSendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserCreateTrade(ctx context.Context, in *UserTradeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserCreateTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserCancelTrade(ctx context.Context, in *UserCancelTradeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserCancelTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserWithdrawal(ctx context.Context, in *UserWithDrawalRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserSearch(ctx context.Context, in *UserSearchRequest, opts ...grpc.CallOption) (*Markets, error) {
	out := new(Markets)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketCraete(ctx context.Context, in *MarketCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketCraete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketUpdate(ctx context.Context, in *MarketUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketInfo(ctx context.Context, in *MarketInfoRequest, opts ...grpc.CallOption) (*MarketInfoResponse, error) {
	out := new(MarketInfoResponse)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketDeposit(ctx context.Context, in *MarketDepositRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketWithDrawal(ctx context.Context, in *MarketWithDrawalRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketWithDrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncTreeServer is the server API for SyncTree service.
// All implementations must embed UnimplementedSyncTreeServer
// for forward compatibility
type SyncTreeServer interface {
	Message(context.Context, *MessageRequest) (*Response, error)
	UserCreate(context.Context, *UserCreateRequest) (*Response, error)
	UserUpdate(context.Context, *UserUpdateRequest) (*Response, error)
	UserSend(context.Context, *UserSendRequest) (*Response, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	UserCreateTrade(context.Context, *UserTradeRequest) (*Response, error)
	UserCancelTrade(context.Context, *UserCancelTradeRequest) (*Response, error)
	UserDeposit(context.Context, *UserDepositRequest) (*Response, error)
	UserWithdrawal(context.Context, *UserWithDrawalRequest) (*Response, error)
	UserSearch(context.Context, *UserSearchRequest) (*Markets, error)
	MarketCraete(context.Context, *MarketCreateRequest) (*Response, error)
	MarketUpdate(context.Context, *MarketUpdateRequest) (*Response, error)
	MarketInfo(context.Context, *MarketInfoRequest) (*MarketInfoResponse, error)
	MarketDeposit(context.Context, *MarketDepositRequest) (*Response, error)
	MarketWithDrawal(context.Context, *MarketWithDrawalRequest) (*Response, error)
	mustEmbedUnimplementedSyncTreeServer()
}

// UnimplementedSyncTreeServer must be embedded to have forward compatible implementations.
type UnimplementedSyncTreeServer struct {
}

func (UnimplementedSyncTreeServer) Message(context.Context, *MessageRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (UnimplementedSyncTreeServer) UserCreate(context.Context, *UserCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedSyncTreeServer) UserUpdate(context.Context, *UserUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedSyncTreeServer) UserSend(context.Context, *UserSendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSend not implemented")
}
func (UnimplementedSyncTreeServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSyncTreeServer) UserCreateTrade(context.Context, *UserTradeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreateTrade not implemented")
}
func (UnimplementedSyncTreeServer) UserCancelTrade(context.Context, *UserCancelTradeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCancelTrade not implemented")
}
func (UnimplementedSyncTreeServer) UserDeposit(context.Context, *UserDepositRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDeposit not implemented")
}
func (UnimplementedSyncTreeServer) UserWithdrawal(context.Context, *UserWithDrawalRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserWithdrawal not implemented")
}
func (UnimplementedSyncTreeServer) UserSearch(context.Context, *UserSearchRequest) (*Markets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSearch not implemented")
}
func (UnimplementedSyncTreeServer) MarketCraete(context.Context, *MarketCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketCraete not implemented")
}
func (UnimplementedSyncTreeServer) MarketUpdate(context.Context, *MarketUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketUpdate not implemented")
}
func (UnimplementedSyncTreeServer) MarketInfo(context.Context, *MarketInfoRequest) (*MarketInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketInfo not implemented")
}
func (UnimplementedSyncTreeServer) MarketDeposit(context.Context, *MarketDepositRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketDeposit not implemented")
}
func (UnimplementedSyncTreeServer) MarketWithDrawal(context.Context, *MarketWithDrawalRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketWithDrawal not implemented")
}
func (UnimplementedSyncTreeServer) mustEmbedUnimplementedSyncTreeServer() {}

// UnsafeSyncTreeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncTreeServer will
// result in compilation errors.
type UnsafeSyncTreeServer interface {
	mustEmbedUnimplementedSyncTreeServer()
}

func RegisterSyncTreeServer(s grpc.ServiceRegistrar, srv SyncTreeServer) {
	s.RegisterService(&SyncTree_ServiceDesc, srv)
}

func _SyncTree_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).Message(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserUpdate(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserSend(ctx, req.(*UserSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserCreateTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserCreateTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserCreateTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserCreateTrade(ctx, req.(*UserTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserCancelTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCancelTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserCancelTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserCancelTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserCancelTrade(ctx, req.(*UserCancelTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserDeposit(ctx, req.(*UserDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserWithDrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserWithdrawal(ctx, req.(*UserWithDrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserSearch(ctx, req.(*UserSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketCraete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketCraete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketCraete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketCraete(ctx, req.(*MarketCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketUpdate(ctx, req.(*MarketUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketInfo(ctx, req.(*MarketInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketDeposit(ctx, req.(*MarketDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketWithDrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketWithDrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketWithDrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketWithDrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketWithDrawal(ctx, req.(*MarketWithDrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncTree_ServiceDesc is the grpc.ServiceDesc for SyncTree service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncTree_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SyncTree",
	HandlerType: (*SyncTreeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Message",
			Handler:    _SyncTree_Message_Handler,
		},
		{
			MethodName: "UserCreate",
			Handler:    _SyncTree_UserCreate_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _SyncTree_UserUpdate_Handler,
		},
		{
			MethodName: "UserSend",
			Handler:    _SyncTree_UserSend_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _SyncTree_UserInfo_Handler,
		},
		{
			MethodName: "UserCreateTrade",
			Handler:    _SyncTree_UserCreateTrade_Handler,
		},
		{
			MethodName: "UserCancelTrade",
			Handler:    _SyncTree_UserCancelTrade_Handler,
		},
		{
			MethodName: "UserDeposit",
			Handler:    _SyncTree_UserDeposit_Handler,
		},
		{
			MethodName: "UserWithdrawal",
			Handler:    _SyncTree_UserWithdrawal_Handler,
		},
		{
			MethodName: "UserSearch",
			Handler:    _SyncTree_UserSearch_Handler,
		},
		{
			MethodName: "MarketCraete",
			Handler:    _SyncTree_MarketCraete_Handler,
		},
		{
			MethodName: "MarketUpdate",
			Handler:    _SyncTree_MarketUpdate_Handler,
		},
		{
			MethodName: "MarketInfo",
			Handler:    _SyncTree_MarketInfo_Handler,
		},
		{
			MethodName: "MarketDeposit",
			Handler:    _SyncTree_MarketDeposit_Handler,
		},
		{
			MethodName: "MarketWithDrawal",
			Handler:    _SyncTree_MarketWithDrawal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

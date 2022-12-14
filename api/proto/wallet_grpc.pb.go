// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/proto/src/wallet.proto

package walletApi

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

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	Create(ctx context.Context, in *WalletCreateReq, opts ...grpc.CallOption) (*Wallet, error)
	Deposit(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*Wallet, error)
	Withdrawal(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*Wallet, error)
	ReturnByPhoneNumber(ctx context.Context, in *WalletReturnByPhoneReq, opts ...grpc.CallOption) (*Wallet, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) Create(ctx context.Context, in *WalletCreateReq, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/wallet_api.WalletService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) Deposit(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/wallet_api.WalletService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) Withdrawal(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/wallet_api.WalletService/Withdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) ReturnByPhoneNumber(ctx context.Context, in *WalletReturnByPhoneReq, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/wallet_api.WalletService/ReturnByPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	Create(context.Context, *WalletCreateReq) (*Wallet, error)
	Deposit(context.Context, *NewTransaction) (*Wallet, error)
	Withdrawal(context.Context, *NewTransaction) (*Wallet, error)
	ReturnByPhoneNumber(context.Context, *WalletReturnByPhoneReq) (*Wallet, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) Create(context.Context, *WalletCreateReq) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWalletServiceServer) Deposit(context.Context, *NewTransaction) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedWalletServiceServer) Withdrawal(context.Context, *NewTransaction) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdrawal not implemented")
}
func (UnimplementedWalletServiceServer) ReturnByPhoneNumber(context.Context, *WalletReturnByPhoneReq) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnByPhoneNumber not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.WalletService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).Create(ctx, req.(*WalletCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.WalletService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).Deposit(ctx, req.(*NewTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_Withdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).Withdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.WalletService/Withdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).Withdrawal(ctx, req.(*NewTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_ReturnByPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletReturnByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).ReturnByPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.WalletService/ReturnByPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).ReturnByPhoneNumber(ctx, req.(*WalletReturnByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet_api.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WalletService_Create_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _WalletService_Deposit_Handler,
		},
		{
			MethodName: "Withdrawal",
			Handler:    _WalletService_Withdrawal_Handler,
		},
		{
			MethodName: "ReturnByPhoneNumber",
			Handler:    _WalletService_ReturnByPhoneNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/src/wallet.proto",
}

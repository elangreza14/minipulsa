// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package minipulsa

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

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	LoginRegister(ctx context.Context, in *LoginRegisterRequest, opts ...grpc.CallOption) (*LoginRegisterResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) LoginRegister(ctx context.Context, in *LoginRegisterRequest, opts ...grpc.CallOption) (*LoginRegisterResponse, error) {
	out := new(LoginRegisterResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.AuthenticationService/LoginRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.AuthenticationService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	LoginRegister(context.Context, *LoginRegisterRequest) (*LoginRegisterResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) LoginRegister(context.Context, *LoginRegisterRequest) (*LoginRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginRegister not implemented")
}
func (UnimplementedAuthenticationServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_LoginRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).LoginRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.AuthenticationService/LoginRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).LoginRegister(ctx, req.(*LoginRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.AuthenticationService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minipulsa.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginRegister",
			Handler:    _AuthenticationService_LoginRegister_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _AuthenticationService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minipulsa.proto",
}

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	GetProducts(context.Context, *Empty) (*GetProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetProducts(context.Context, *Empty) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minipulsa.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minipulsa.proto",
}

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateORUpdateOrderResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*CreateORUpdateOrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateORUpdateOrderResponse, error) {
	out := new(CreateORUpdateOrderResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*CreateORUpdateOrderResponse, error) {
	out := new(CreateORUpdateOrderResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.OrderService/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.OrderService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateORUpdateOrderResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*CreateORUpdateOrderResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateORUpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*CreateORUpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.OrderService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.OrderService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minipulsa.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrderService_GetOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minipulsa.proto",
}

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetWalletDetail(ctx context.Context, in *GetWalletDetailRequest, opts ...grpc.CallOption) (*GetWalletDetailResponse, error)
	UseWallet(ctx context.Context, in *UseWalletRequest, opts ...grpc.CallOption) (*UseWalletResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetWalletDetail(ctx context.Context, in *GetWalletDetailRequest, opts ...grpc.CallOption) (*GetWalletDetailResponse, error) {
	out := new(GetWalletDetailResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.WalletService/GetWalletDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) UseWallet(ctx context.Context, in *UseWalletRequest, opts ...grpc.CallOption) (*UseWalletResponse, error) {
	out := new(UseWalletResponse)
	err := c.cc.Invoke(ctx, "/minipulsa.WalletService/UseWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	GetWalletDetail(context.Context, *GetWalletDetailRequest) (*GetWalletDetailResponse, error)
	UseWallet(context.Context, *UseWalletRequest) (*UseWalletResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) GetWalletDetail(context.Context, *GetWalletDetailRequest) (*GetWalletDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletDetail not implemented")
}
func (UnimplementedWalletServiceServer) UseWallet(context.Context, *UseWalletRequest) (*UseWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseWallet not implemented")
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

func _WalletService_GetWalletDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWalletDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.WalletService/GetWalletDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWalletDetail(ctx, req.(*GetWalletDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_UseWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).UseWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minipulsa.WalletService/UseWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).UseWallet(ctx, req.(*UseWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minipulsa.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWalletDetail",
			Handler:    _WalletService_GetWalletDetail_Handler,
		},
		{
			MethodName: "UseWallet",
			Handler:    _WalletService_UseWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minipulsa.proto",
}

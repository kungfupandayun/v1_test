// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package productrpc

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Fetch all existing products in the system.
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	// Fetch one product by id
	FetchProduct(ctx context.Context, in *FetchProductsRequest, opts ...grpc.CallOption) (*Product, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, "/product.Service/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FetchProduct(ctx context.Context, in *FetchProductsRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product.Service/FetchProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Fetch all existing products in the system.
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	// Fetch one product by id
	FetchProduct(context.Context, *FetchProductsRequest) (*Product, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedServiceServer) FetchProduct(context.Context, *FetchProductsRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchProduct not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Service/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListProducts(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FetchProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FetchProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Service/FetchProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FetchProduct(ctx, req.(*FetchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _Service_ListProducts_Handler,
		},
		{
			MethodName: "FetchProduct",
			Handler:    _Service_FetchProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/rpc/product.proto",
}

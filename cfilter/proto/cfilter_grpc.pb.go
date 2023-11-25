// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/cfilter.proto

package proto

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

const (
	CFilter_Filter_FullMethodName = "/cfilter.examples.sicky.dev.CFilter/Filter"
)

// CFilterClient is the client API for CFilter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CFilterClient interface {
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
}

type cFilterClient struct {
	cc grpc.ClientConnInterface
}

func NewCFilterClient(cc grpc.ClientConnInterface) CFilterClient {
	return &cFilterClient{cc}
}

func (c *cFilterClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, CFilter_Filter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CFilterServer is the server API for CFilter service.
// All implementations must embed UnimplementedCFilterServer
// for forward compatibility
type CFilterServer interface {
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	mustEmbedUnimplementedCFilterServer()
}

// UnimplementedCFilterServer must be embedded to have forward compatible implementations.
type UnimplementedCFilterServer struct {
}

func (UnimplementedCFilterServer) Filter(context.Context, *FilterRequest) (*FilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (UnimplementedCFilterServer) mustEmbedUnimplementedCFilterServer() {}

// UnsafeCFilterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CFilterServer will
// result in compilation errors.
type UnsafeCFilterServer interface {
	mustEmbedUnimplementedCFilterServer()
}

func RegisterCFilterServer(s grpc.ServiceRegistrar, srv CFilterServer) {
	s.RegisterService(&CFilter_ServiceDesc, srv)
}

func _CFilter_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CFilterServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CFilter_Filter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CFilterServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CFilter_ServiceDesc is the grpc.ServiceDesc for CFilter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CFilter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cfilter.examples.sicky.dev.CFilter",
	HandlerType: (*CFilterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Filter",
			Handler:    _CFilter_Filter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cfilter.proto",
}

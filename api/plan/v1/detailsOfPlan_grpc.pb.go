// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api/plan/v1/detailsOfPlan.proto

package v1

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

// DetailsOfPlanClient is the client API for DetailsOfPlan service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DetailsOfPlanClient interface {
	GetPlan(ctx context.Context, in *GetPlanRequest, opts ...grpc.CallOption) (*Reply, error)
	ListPlan(ctx context.Context, in *ListPlanRequest, opts ...grpc.CallOption) (*Reply, error)
}

type detailsOfPlanClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailsOfPlanClient(cc grpc.ClientConnInterface) DetailsOfPlanClient {
	return &detailsOfPlanClient{cc}
}

func (c *detailsOfPlanClient) GetPlan(ctx context.Context, in *GetPlanRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/api.detailsOfPlan.v1.DetailsOfPlan/GetPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detailsOfPlanClient) ListPlan(ctx context.Context, in *ListPlanRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/api.detailsOfPlan.v1.DetailsOfPlan/ListPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailsOfPlanServer is the server API for DetailsOfPlan service.
// All implementations must embed UnimplementedDetailsOfPlanServer
// for forward compatibility
type DetailsOfPlanServer interface {
	GetPlan(context.Context, *GetPlanRequest) (*Reply, error)
	ListPlan(context.Context, *ListPlanRequest) (*Reply, error)
	mustEmbedUnimplementedDetailsOfPlanServer()
}

// UnimplementedDetailsOfPlanServer must be embedded to have forward compatible implementations.
type UnimplementedDetailsOfPlanServer struct {
}

func (UnimplementedDetailsOfPlanServer) GetPlan(context.Context, *GetPlanRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlan not implemented")
}
func (UnimplementedDetailsOfPlanServer) ListPlan(context.Context, *ListPlanRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlan not implemented")
}
func (UnimplementedDetailsOfPlanServer) mustEmbedUnimplementedDetailsOfPlanServer() {}

// UnsafeDetailsOfPlanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DetailsOfPlanServer will
// result in compilation errors.
type UnsafeDetailsOfPlanServer interface {
	mustEmbedUnimplementedDetailsOfPlanServer()
}

func RegisterDetailsOfPlanServer(s grpc.ServiceRegistrar, srv DetailsOfPlanServer) {
	s.RegisterService(&DetailsOfPlan_ServiceDesc, srv)
}

func _DetailsOfPlan_GetPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsOfPlanServer).GetPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.detailsOfPlan.v1.DetailsOfPlan/GetPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsOfPlanServer).GetPlan(ctx, req.(*GetPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DetailsOfPlan_ListPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsOfPlanServer).ListPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.detailsOfPlan.v1.DetailsOfPlan/ListPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsOfPlanServer).ListPlan(ctx, req.(*ListPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DetailsOfPlan_ServiceDesc is the grpc.ServiceDesc for DetailsOfPlan service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DetailsOfPlan_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.detailsOfPlan.v1.DetailsOfPlan",
	HandlerType: (*DetailsOfPlanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlan",
			Handler:    _DetailsOfPlan_GetPlan_Handler,
		},
		{
			MethodName: "ListPlan",
			Handler:    _DetailsOfPlan_ListPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/plan/v1/detailsOfPlan.proto",
}
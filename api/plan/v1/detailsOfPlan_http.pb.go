// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.21.7
// source: api/plan/v1/detailsOfPlan.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDetailsOfPlanGetPlan = "/api.detailsOfPlan.v1.DetailsOfPlan/GetPlan"
const OperationDetailsOfPlanListPlan = "/api.detailsOfPlan.v1.DetailsOfPlan/ListPlan"

type DetailsOfPlanHTTPServer interface {
	GetPlan(context.Context, *GetPlanRequest) (*Reply, error)
	ListPlan(context.Context, *ListPlanRequest) (*Reply, error)
}

func RegisterDetailsOfPlanHTTPServer(s *http.Server, srv DetailsOfPlanHTTPServer) {
	r := s.Route("/")
	r.POST("/greeyun/fmz-home/plan/getPlan", _DetailsOfPlan_GetPlan0_HTTP_Handler(srv))
	r.POST("/greeyun/fmz-home/plan/listPlan", _DetailsOfPlan_ListPlan0_HTTP_Handler(srv))
}

func _DetailsOfPlan_GetPlan0_HTTP_Handler(srv DetailsOfPlanHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPlanRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDetailsOfPlanGetPlan)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPlan(ctx, req.(*GetPlanRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Reply)
		return ctx.Result(200, reply)
	}
}

func _DetailsOfPlan_ListPlan0_HTTP_Handler(srv DetailsOfPlanHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPlanRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDetailsOfPlanListPlan)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPlan(ctx, req.(*ListPlanRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Reply)
		return ctx.Result(200, reply)
	}
}

type DetailsOfPlanHTTPClient interface {
	GetPlan(ctx context.Context, req *GetPlanRequest, opts ...http.CallOption) (rsp *Reply, err error)
	ListPlan(ctx context.Context, req *ListPlanRequest, opts ...http.CallOption) (rsp *Reply, err error)
}

type DetailsOfPlanHTTPClientImpl struct {
	cc *http.Client
}

func NewDetailsOfPlanHTTPClient(client *http.Client) DetailsOfPlanHTTPClient {
	return &DetailsOfPlanHTTPClientImpl{client}
}

func (c *DetailsOfPlanHTTPClientImpl) GetPlan(ctx context.Context, in *GetPlanRequest, opts ...http.CallOption) (*Reply, error) {
	var out Reply
	pattern := "/greeyun/fmz-home/plan/getPlan"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDetailsOfPlanGetPlan))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DetailsOfPlanHTTPClientImpl) ListPlan(ctx context.Context, in *ListPlanRequest, opts ...http.CallOption) (*Reply, error) {
	var out Reply
	pattern := "/greeyun/fmz-home/plan/listPlan"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDetailsOfPlanListPlan))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

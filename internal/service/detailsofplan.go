package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "gy_home/api/login/v1"
	"gy_home/customizeLog"
	"gy_home/internal/biz"
	"gy_home/internal/errorcode"

	pb "gy_home/api/plan/v1"
)

type DetailsOfPlanService struct {
	pb.UnimplementedDetailsOfPlanServer
	uc *biz.PlanUseCase
	log *log.Helper

}

func NewDetailsOfPlanService(uc *biz.PlanUseCase,logger log.Logger) *DetailsOfPlanService {
	return &DetailsOfPlanService{uc:uc,log: log.NewHelper(logger)}
}

func (s *DetailsOfPlanService) GetPlan(ctx context.Context, req *pb.GetPlanRequest) (*pb.Reply, error) {
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.log.WithContext(ctx).Infow("msg", fmt.Sprintf("GetPlan Received cid:%v", req.Cid),
		"ext",customizeLog.Ext{
			Ip: ip.(string),
			Name: token.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Get solution details",
			Action: "get",
		},
		"category", 1)
	var (
		err error
		res interface{}
	)
	res,err = s.uc.GetPlan(ctx,&req.Cid,req.Id,req.Listid)
	if err != nil{
		s.log.WithContext(ctx).Errorw("msg",err.Error(),
			"category",2)
		return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
	}

	data,err := dataToStruct("detailsOfPlan",res)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg",err.Error(),
			"category",2)
		return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
	}

	return &pb.Reply{
		Code: errorcode.OK,
		Data: data,
		Message: "ok",
	},nil
}

func (s *DetailsOfPlanService) ListPlan(ctx context.Context, req *pb.ListPlanRequest) (*pb.Reply, error) {
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.log.WithContext(ctx).Infow("msg",fmt.Sprintf("GetPlan Received cid:%v",req.Cid),
		"ext",customizeLog.Ext{
			Ip: ip.(string),
			Name: token.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Get scheme list",
			Action: "get",
		},
		"category",1)

	res,err := s.uc.ListPlan(ctx,&req.Cid)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("s.uc.GetPlan error:%v",err),
			"category",2)
		return nil,errors.New(errorcode.DATA_ERROR,"",err.Error())
	}

	data,err := dataToStruct("listPlan",res)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg",err.Error(),
			"category",2)
		return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
	}

	return &pb.Reply{
		Code: 200,
		Data: data,
		Message: "ok",
	}, nil
}

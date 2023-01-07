package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Plan struct{
	ID int32 `json:"id"`
	CID int32 `json:"cid"`
	ListID int32 `json:"list_id"`
 	CName string `json:"c_name"`
	Name string `json:"name"`
	Url string `json:"url"`
	Utime string `json:"utime"`
}

type DataPlan struct{
	Cap int `json:",omitempty"`
	Parent []*Plan `json:"parent"`
	Recommend []*Plan `json:"recommend"`
}

type PlanRepo interface{
	GetPlan(context.Context,*int32,*int32,*int32)(*Plan,error)
	ListPlan(ctx context.Context,cid *int32)(*DataPlan,error)
}

type PlanUseCase struct {
	repo PlanRepo
	log *log.Helper
}

func NewPlanUseCase(repo PlanRepo, logger log.Logger)*PlanUseCase{
	return &PlanUseCase{repo: repo,log: log.NewHelper(logger)}
}

func (uc *PlanUseCase)GetPlan(ctx context.Context,cid *int32,id *int32,listid *int32)(*Plan,error){
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("GetPlan: cid:%v",*cid),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.GetPlan(ctx,cid,id,listid)
}

func (uc *PlanUseCase)ListPlan(ctx context.Context,cid *int32)(*DataPlan,error){
	return uc.repo.ListPlan(ctx,cid)
}
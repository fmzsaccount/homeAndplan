package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gy_home/internal/biz"
	"gy_home/internal/data/ent"
	"gy_home/internal/data/ent/plan"
)

type planRepo struct {
	data *Data
	log  *log.Helper
}
// NewPlanRepo .
func NewPlanRepo(data *Data, logger log.Logger) biz.PlanRepo {
	return &planRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *planRepo)GetPlan(ctx context.Context,cid *int32,id *int32,listid *int32)(*biz.Plan,error){

	res,err := r.data.db.Plan.Query().
		Where(plan.CategoryID(*cid)).
		Order(ent.Asc(plan.FieldUtime)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if id != nil{
		for k,v := range res{
			//fmt.Println("id: ",v.ID)
			if *id == v.ID{
				return &biz.Plan{
					ID: v.ID,
					CID: v.CategoryID,
					ListID: v.CategoryID+int32(k),
					CName: v.Category,
					Name: v.Name,
					Url: v.URL,
					Utime: v.Utime.Format("2006-01-02 15:04:05"),
				}, nil
			}
		}
		return nil, fmt.Errorf("No plan with id:%v was found ",*id)
	}
	if listid == nil{
		return &biz.Plan{
			ID: res[0].ID,
			CID: res[0].CategoryID,
			ListID: res[0].CategoryID,
			CName: res[0].Category,
			Name: res[0].Name,
			Utime: res[0].Utime.Format("2006-01-02 15:04:05"),
			Url: res[0].URL,
		},nil
	}else{
		index := *listid-res[0].CategoryID
		return &biz.Plan{
			ID: res[index].ID,
			CID: res[index].CategoryID,
			ListID: *listid,
			CName: res[index].Category,
			Name: res[index].Name,
			Utime: res[index].Utime.Format("2006-01-02 15:04:05"),
			Url: res[index].URL,
		},nil
	}

}

func (r *planRepo)ListPlan(ctx context.Context,cid *int32)(*biz.DataPlan,error){
	var(
		parent = make([]*biz.Plan,0)
		recommend = make([]*biz.Plan,0)
	)
	resAll,err := r.data.db.Plan.Query().
		Where(plan.CategoryID(*cid)).
		Order(ent.Asc(plan.FieldUtime)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	resRecommend,err := r.data.db.Plan.Query().
		Where(plan.Recommend(1)).
		Order(ent.Desc(plan.FieldUtime)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	for k,v := range resAll{
		listid := v.CategoryID
		parent = append(parent,&biz.Plan{
			ID: v.ID,
			CID: v.CategoryID,
			ListID: listid+int32(k),
			CName: v.Category,
			Url: v.URL,
			Name: v.Name,
			Utime: v.Utime.Format("2006-01-02 15:04:05"),
		})
	}
	for _,v := range resRecommend{
		recommend = append(recommend,&biz.Plan{
			ID: v.ID,
			CID: v.CategoryID,
			CName: v.Category,
			Url: v.URL,
			Name: v.Name,
			Utime: v.Utime.Format("2006-01-02 15:04:05"),
		})
	}

	for _,v := range recommend{
		v.ListID = v.CID
		for _,v1 := range parent{
			if v1.CID == v.CID && v1.Name==v.Name{
				v.ListID = v1.ListID
			}
		}
	}
	return &biz.DataPlan{
		Cap: len(parent),
		Parent: parent[0:5],
		Recommend: recommend[0:5],
	}, nil
}
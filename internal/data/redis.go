package data

import (
	"context"
	"encoding/json"
	"fmt"
	"gy_home/internal/biz"
)

func (r *homePageRepo)Keys(ctx context.Context,rediskey *string)([]string,error){
	r.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Keys Received:%v",*rediskey),
		"evt:","",
		"category:",2)
	return r.data.rdb.Keys(ctx,*rediskey).Result()
}


func (r *homePageRepo)Set(ctx context.Context,search *biz.Search,redisKey *string)error{
	r.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Set Received: search:%v rediskey: %v",*search,*redisKey),
		"evt:","",
		"category:",2)
	bs,_ := json.Marshal(*search)
	return r.data.rdb.Set(ctx,*redisKey,bs,0).Err()
}

func (r *homePageRepo) List(ctx context.Context,redisKeys *string)([]*biz.Search,error){
	r.log.WithContext(ctx).Infow("msg",fmt.Sprintf("List Received:%v",*redisKeys),
		"evt:","",
		"category:",2)
	var (
		searchInfo []*biz.Search
	)
	resKeys,keyserr := r.data.rdb.Keys(ctx,*redisKeys).Result()
	if keyserr != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("redis keys error:%v",keyserr),
			"evt:","",
			"category:",2)
		return nil, keyserr
	}
	for _,v := range resKeys{
		var search biz.Search
		resGet,err := r.data.rdb.Get(ctx,v).Result()
		if err != nil {
			return nil,err
		}
		resGetbyte := []byte(resGet)
		err1 := json.Unmarshal(resGetbyte,&search)
		if err1 != nil{
			r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("json unmarshal error:%v",err1),
				"evt:","",
				"category:",2)
			return nil, err1
		}
		searchInfo = append(searchInfo,&search)
	}
	return searchInfo,nil
}

func (r *homePageRepo) Get(ctx context.Context,redisKey *string)(*biz.Search,error){
	r.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Get Received:%v",*redisKey),
		"evt:","",
		"category:",2)
	var search biz.Search
	resGet,err := r.data.rdb.Get(ctx,*redisKey).Result()
	if err != nil{
		return nil, err
	}
	resGetbyte := []byte(resGet)
	err1 := json.Unmarshal(resGetbyte,&search)
	if err1 != nil{
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("json unmarshal error:%v",err1),
			"evt:","",
			"category:",2)
		return nil, err1
	}
	return &search,nil
}

func(r *homePageRepo)Del(ctx context.Context,redisKey ...string)error{
	r.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Del Received:%v",redisKey),
		"evt:","",
		"category:",2)
	return r.data.rdb.Del(ctx,redisKey...).Err()
}

func(r *homePageRepo)FlushDb(ctx context.Context)error{
	r.log.WithContext(ctx).Infow("msg","FlushDb run",
		"evt:","",
		"category:",2)
	return r.data.rdb.FlushDB(ctx).Err()
}
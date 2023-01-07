package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/structpb"
	pb "gy_home/api/homepage/v1"
	v1 "gy_home/api/login/v1"
	"gy_home/customizeLog"
	"gy_home/internal/biz"
	"gy_home/internal/errorcode"
	"strconv"
	"time"
)

type HomepageService struct {
	pb.UnimplementedHomepageServer
	servicelog *log.Helper
	uc  *biz.HomepageUsecase
}

const (
	SearchProKey = "ts-greeyun-searchPro"
	SearchAllKey = "ts-greeyun-searchAll"
	SearchUserKey = "ts-greeyun-searchUser"
)

func NewHomepageService(uc *biz.HomepageUsecase,logger log.Logger) *HomepageService {
	return &HomepageService{uc: uc,servicelog: log.NewHelper(logger)}
}

func (s *HomepageService) GetHomepage(ctx context.Context, req *pb.GetHomepageRequest) (*pb.Reply, error) {
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.servicelog.WithContext(ctx).Infow(
		"msg","GetHomepage Run",
		"is_get",0,
		"ext",&customizeLog.Ext{
			Ip: ip.(string),
			Name: token.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Get the home page",
			Action: "Get",
		},
		"evt","get homepage",
		"category",1)

	var (
		g = &biz.Homepage{
			Banner:    nil,
			PlanCategory: nil,
			Pro:          nil,
		}
	)
	homepagep, err := s.uc.CreateHomepage(ctx, g)
	if err != nil {
		s.servicelog.WithContext(ctx).Errorw(
			"msg",fmt.Sprintf("s.uc.CreateHomepage() error: %v", err),
			"evt","biz.CreateHomepage error",
			"category",2)
		return nil, err
	}
	data,err := dataToStruct("homepage",homepagep)
	if err != nil {
		return nil, err
	}
	return &pb.Reply{
		Code: errorcode.OK,
		Data: data,
		Message: "ok",
	}, nil
}
func (s *HomepageService) GetHomepagePro(ctx context.Context, req *pb.GetHomepageProRequest) (*pb.Reply, error) {
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.servicelog.WithContext(ctx).Infow(
		"msg","GetHomepagePro Run",
		"ext",&customizeLog.Ext{
			Ip: ip.(string),
			Name: token.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Get the Products and Services drop-down box on the home page",
			Action: "Get",
		},
		"evt","get homepage",
		"category",1)

	var (
		g = &biz.HomepagePro{
			Proc:          nil,
			HistorySearch: nil,
		}
	)
	searchProkeys := SearchProKey + "*"
	homepageprop, err := s.uc.CreateHomepagePro(ctx, g,searchProkeys)
	if err != nil {
		s.servicelog.WithContext(ctx).Errorw(
			"msg",fmt.Sprintf("s.uc.CreateHomepagePro() error: %v", err),
			"evt","biz.CreateHomepagePro error",
			"category",2)
		return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
	}

	BubbleSortCount(homepageprop.HistorySearch)

	data,err := dataToStruct("homepagePro",homepageprop)
	if err != nil {
		return nil, err
	}

	return &pb.Reply{
		Code: errorcode.OK,
		Data: data,
		Message: "ok",
	}, nil

}


func (s *HomepageService) ListSearch(ctx context.Context, req *pb.ListSearchRequest) (*pb.Reply, error) {
	authorization := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.servicelog.WithContext(ctx).Infow("msg",fmt.Sprintf("ListSearch Received Type:%v",req.GetType()),
		"ext",customizeLog.Ext{
			Ip: ip.(string),
			Name: authorization.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Get a list of search keywords",
			Action: "list",
		},
		"category",1)

	var (
		usernamestr string
	)
	usernamestr = authorization.(*v1.RefreshAuthorizationReply).GetMail()
	if req.Type == 1 {
		redisKeysAll := SearchAllKey + "*"
		listSearch, err := s.uc.List(ctx, &redisKeysAll)
		if err != nil {
			s.servicelog.WithContext(ctx).Errorw("msg", fmt.Sprintf("ListSearch:List() err:%v", err),
				"category", 2)
			return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
		}

		redisKeyUser := SearchUserKey + usernamestr + "*"
		listUserSearch, err := s.uc.List(ctx, &redisKeyUser)
		if err != nil {
			s.servicelog.WithContext(ctx).Errorw("msg", fmt.Sprintf("listUserSearch:List() err:%v", err),
				"category", 2)
			return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
		}

		BubbleSortCount(listSearch)
		BubbleSortUtime(listUserSearch)

		data,err := dataToStruct("keywordList",&biz.KeywordList{
			Recent: listUserSearch,
			Recommend: listSearch,
		})
		if err != nil {
			s.servicelog.WithContext(ctx).Errorw("msg",err.Error(),
				"category",2)
			return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
		}

		return &pb.Reply{
			Code: errorcode.OK,
			Data: data,
			Message: "ok",
		},nil
	}
	return &pb.Reply{
		Code: errorcode.OK,
		Data: nil,
		Message: "ok",
	},nil
}

func (s *HomepageService) SetSearch(ctx context.Context, req *pb.SetSearchRequest) (*pb.Reply, error) {
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.servicelog.WithContext(ctx).Infow(
		"msg",fmt.Sprintf("SetSearch Received keyword:%v Id:%v Type:%v",req.GetKeyword(),req.GetId(),req.GetType()),
		"ext",customizeLog.Ext{
			Ip: ip.(string),
			Name: token.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Set the search keyword",
			Action: "set",
		},
		"category",1,
		)

	var datastr string
	searchstr :=  req.Keyword + strconv.Itoa(int(req.Id)) + strconv.Itoa(int(req.Type))
	usernamestr := token.(*v1.RefreshAuthorizationReply).GetMail()
	redisKeyUser := SearchUserKey + usernamestr + searchstr
	//fmt.Println("redisKeyUser:",redisKeyUser)
	rediskeyAll := SearchAllKey + searchstr
	rediskeyPro := SearchProKey + searchstr
	if req.Type == 1{
		datastr = setKey(s,ctx,redisKeyUser,req)
		//fmt.Println("set Userkey ok")
		datastr = datastr+"All:"+setKey(s,ctx,rediskeyAll,req)
	}else {
		datastr = setKey(s,ctx,rediskeyPro,req)
	}

	data,err := dataToStruct("setKeyword",datastr)
	if err != nil{
		s.servicelog.WithContext(ctx).Errorw("msg",err.Error(),
			"category",2)
		return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
	}
	return &pb.Reply{
		Code: errorcode.OK,
		Data: data,
		Message: "ok",
	},nil
}

func (s *HomepageService) DelSearch(ctx context.Context, req *pb.DelSearchRequest) (*pb.Reply, error) {
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.servicelog.WithContext(ctx).Infow("msg",fmt.Sprintf("DelSearch Received isdelete:%v",req.GetIsdelete()),
		"ext",customizeLog.Ext{
		Ip: ip.(string),
		Name: token.(*v1.RefreshAuthorizationReply).GetName(),
		Module: "Delete the search keyword list",
		Action: "delete",
		},
		"category",1)

	usernamestr := token.(*v1.RefreshAuthorizationReply).GetMail()
	rediskeys := SearchUserKey + usernamestr+"*"
	//fmt.Println(rediskeys)
	if req.Isdelete == 1{
		keyslice,err := s.uc.Keys(ctx,&rediskeys)
		if err != nil {
			s.servicelog.WithContext(ctx).Errorw("msg",fmt.Sprintf("Delsearch error:%v",err),
				"category",2)
			return nil,err
		}
		for _,v := range keyslice{
			temp := v
			err = s.uc.Del(ctx,temp)
			if err != nil {
				break
			}
		}
	}
	return &pb.Reply{
		Code: 200,
		Data: nil,
		Message: "ok",
	},nil
}


func (s *HomepageService)GetSite(ctx context.Context,req *pb.GetSiteReq)(*pb.Reply,error){
	token := ctx.Value("RefreshAuthorizationReply")
	ip := ctx.Value("ip")
	s.servicelog.WithContext(ctx).Infow("msg","GetSite Run",
		"ext",customizeLog.Ext{
			Ip: ip.(string),
			Name: token.(*v1.RefreshAuthorizationReply).GetName(),
			Module: "Get site",
			Action: "get",
		},
		"category",1)

	res,err := s.uc.GetSite(ctx)
	if err != nil{
		s.servicelog.WithContext(ctx).Errorw("msg",fmt.Sprintf("GetSite: err:%v",err),
			"category",2)
		return nil, err
	}
	data,err := dataToStruct("category",res)
	if err != nil {
		s.servicelog.WithContext(ctx).Errorw("msg",err.Error(),
			"category",2)
		return nil, errors.New(errorcode.DATA_ERROR,"",err.Error())
	}
	return &pb.Reply{
		Code: errorcode.OK,
		Data: data,
		Message: "ok",
	},nil
}


func BubbleSortCount(data []*biz.Search){

	for i:=0;i<len(data);i++{
		for j:=0;j<len(data)-i-1;j++{
			if data[j].Count < data[j+1].Count{
				data[j],data[j+1] = data[j+1],data[j]
			}
		}
	}
}

func BubbleSortUtime(data []*biz.Search){
	for i:=0;i<len(data);i++{
		for j:=0;j<len(data)-i-1;j++{
			if data[j].Utime.Nanosecond() < data[j+1].Utime.Nanosecond(){
				data[j],data[j+1] = data[j+1],data[j]
			}
		}
	}
}

func setKey(s *HomepageService,ctx context.Context,redisKey string,req *pb.SetSearchRequest)string{
	var datastr string
	resGetAll,err := s.uc.Get(ctx,&redisKey)
	if err != nil {
		s.servicelog.WithContext(ctx).Errorw("msg",fmt.Sprintf("SetSearch:Get keyword error:%v",err),
			"category",2)
		//return ""
	}
	if resGetAll == nil{
		err1 := s.uc.Set(ctx,&biz.Search{
			ID: req.Id,
			Keyword: req.Keyword,
			Type: req.Type,
			Count: 0,
			Utime: time.Now(),
			Ctime: time.Now(),
		},&redisKey)
		datastr = "you set key: " + redisKey
		if err1 != nil {
			s.servicelog.WithContext(ctx).Errorw("msg",fmt.Sprintf("SetSearch:Set Allkey err:%v",err1),
				"category",2)
			return ""
		}
	}else {
		err2 := s.uc.Set(ctx,&biz.Search{
			ID: resGetAll.ID,
			Keyword: resGetAll.Keyword,
			Type: resGetAll.Type,
			Count: resGetAll.Count+1,
			Utime: time.Now(),
			Ctime: resGetAll.Ctime,
		},&redisKey)
		if err2 != nil {
			s.servicelog.WithContext(ctx).Errorw("msg",fmt.Sprintf("SetSearch:Set Allkey err:%v",err2),
				"category",2)
			return ""
		}
		datastr = "you set key: " + redisKey
	}
	return datastr
}

func dataToStruct(key string,req interface{})(*structpb.Struct,error){
	res,err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var i interface{}
	err1 := json.Unmarshal(res,&i)
	if err != nil {
		return nil, err1
	}
	mapdata := map[string]interface{}{
		key:i,
	}
	data,err := structpb.NewStruct(mapdata)
	return data,nil
}
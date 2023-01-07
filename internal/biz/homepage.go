package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "gy_home/api/homepage/v1"
	"time"
)

var (
	// ErrHomepageNotFound is Homepage not found.
	ErrHomepageNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "Homepage not found")
	ErrContentMissing   = errors.NotFound(v1.ErrorReason_CONTENT_MISSING.String(), "Homepage content missing")
)

// Homepage is a Homepage model.
type Homepage struct {
	Banner    []*Banner `json:"banner,omitempty"`
	Pro          []*Product `json:"product,omitempty"`
	PlanCategory []*PlanCategory `json:"plan_category,omitempty"`
}

type HomepagePro struct {
	Proc          []*ProductCategory `json:"product_category,omitempty"`
	HistorySearch []*Search `json:"keywordList,omitempty"`
}

type KeywordList struct{
	Recent []*Search `json:"recent"`
	Recommend []*Search `json:"recommend"`
}

type Banner struct{
	ID int32 `json:"id,omitempty"`
	Sort int32 `json:"sort,omitempty"`
	Title string `json:"title,omitempty"`
	Type string `json:"type,omitempty"`
	Bucket string `json:"-"`
	URL string `json:"url,omitempty"`
	LinkID string `json:"link_id,omitempty"`
}
type ProductCategory struct {
	Status   int `json:"-"`
	ID       int32 `json:"id,omitempty"`
	PID      int32 `json:"pid,omitempty"`
	Sort     int32 `json:"-"`
	Name     string `json:"name,omitempty"`
	Children []*Level2 `json:"children,omitempty"`
}

type Level2 struct {
	Status   int `json:"-"`
	ID       int32 `json:"id,omitempty"`
	PID      int32 `json:"pid,omitempty"`
	Sort     int32 `json:"-"`
	Name     string `json:"name,omitempty"`
	Children []*Product `json:"children,omitempty"`
}

type Product struct {
	Status    int `json:"-"`
	Recommend int `json:"-"`
	ID        int32 `json:"id,omitempty"`
	CID       int32 `json:"cid,omitempty"`
	History   int32 `json:"history,omitempty"`
	CName     string `json:"cname,omitempty"`
	ProName   string `json:"name,omitempty"`
	ProInfo   string `json:"content,omitempty"`

}


type PlanCategory struct {
	ID int32 `json:"id,omitempty"`
	Sort     int32 `json:"sort,omitempty"`
	PlanName string `json:"name,omitempty"`
	PlanInfo string `json:"content,omitempty"`
	Bucket string `json:"-,omitempty"`
	URL string `json:"url,omitempty"`
}

type Search struct {
	ID int32 `json:"id,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	Count int32 `json:"count,omitempty"`
	Type int32 `json:"type,omitempty"`
	Utime time.Time `json:"-"`
	Ctime time.Time `json:"-"`
}
type SiteCategory struct {
	Name string `json:"name"`
	Site []*Site `json:"site"`
}

type Site struct {
	ID int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
	Category string `json:"category,omitempty"`
}
// HomepageRepo is a Homepage repo.
type HomepageRepo interface {
	//db
	GetHomepage(context.Context, *Homepage) (*Homepage, error)
	GetHomepagePro(ctx context.Context, g *HomepagePro,redisKey string) (*HomepagePro, error)
	GetSite(ctx context.Context)([]*SiteCategory,error)
	//redis
	Keys(ctx context.Context,rediskey *string)([]string,error)
	Set(ctx context.Context,search *Search,redisKey *string)error
	List(ctx context.Context,redisKeys *string)([]*Search,error)
	Get(ctx context.Context,redisKey *string)(*Search,error)
	Del(ctx context.Context,redisKey ...string)error
	FlushDb(ctx context.Context)error
}

// HomepageUsecase is a Homepage usecase.
type HomepageUsecase struct {
	repo HomepageRepo
	log  *log.Helper
}

// NewHomepageUsecase new a Homepage usecase.
func NewHomepageUsecase(repo HomepageRepo, logger log.Logger) *HomepageUsecase {
	return &HomepageUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateHomepage creates a Homepage, and returns the new Homepage.
func (uc *HomepageUsecase) CreateHomepage(ctx context.Context, g *Homepage) (*Homepage, error) {
	//uc.log.WithContext(ctx).Infow("msg","CreateHomepage Run",
	//"evt","",
	//"category",2)
	return uc.repo.GetHomepage(ctx, g)
}

// CreateHomepagepro creates a Homepagepro, and returns the new Homepagepro.
func (uc *HomepageUsecase) CreateHomepagePro(ctx context.Context, g *HomepagePro,redisKey string) (*HomepagePro, error) {
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("CreateHomepagePro:  redisKey:%v",redisKey),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.GetHomepagePro(ctx, g,redisKey)
}

func (uc *HomepageUsecase)GetSite(ctx context.Context)([]*SiteCategory,error){
	//uc.log.WithContext(ctx).Infow("msg","GetSite run",
	//	"category",2)
	return uc.repo.GetSite(ctx)
}

func (uc *HomepageUsecase)Keys(ctx context.Context,rediskey *string)([]string,error){
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Keys() Received keys path: %v",*rediskey),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.Keys(ctx,rediskey)
}

func (uc *HomepageUsecase) Set(ctx context.Context, g *Search,redisKey *string)error{
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Set Received redisKey: %v value: %v", *redisKey,*g),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.Set(ctx,g,redisKey)
}

func (uc *HomepageUsecase)  Get(ctx context.Context,redisKey *string)(*Search,error){
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Get Received redisKey: %v", *redisKey),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.Get(ctx,redisKey)
}

func (uc *HomepageUsecase) List(ctx context.Context,redisKeys *string)([]*Search,error){
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("List Received redisKey: %v", *redisKeys),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.List(ctx,redisKeys)
}

func (uc *HomepageUsecase)  Del(ctx context.Context,redisKey ...string)error{
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("Del Received redisKey: %v", redisKey),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.Del(ctx,redisKey...)
}

func(uc *HomepageUsecase) FlushDB(ctx context.Context)error{
	//uc.log.WithContext(ctx).Infow("msg",fmt.Sprintf("FlushDb Received: %v",ctx),
	//	"evt:","",
	//	"category:",2)
	return uc.repo.FlushDb(ctx)
}

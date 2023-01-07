package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gy_home/internal/biz"
	"gy_home/internal/data/ent"
	"gy_home/internal/data/ent/banner"
	"gy_home/internal/data/ent/plancategory"
	"gy_home/internal/data/ent/product"
	"gy_home/internal/data/ent/productcategory"
	"gy_home/internal/data/ent/site"
)

var _ biz.HomepageRepo = (*homePageRepo)(nil)

type homePageRepo struct {
	data *Data
	log  *log.Helper
}

// NewHomepageRepo .
func NewHomepageRepo(data *Data, logger log.Logger) biz.HomepageRepo {
	return &homePageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *homePageRepo) GetHomepage(ctx context.Context, g *biz.Homepage) (*biz.Homepage, error) {
	var (
		homepage biz.Homepage
	)

	resQuerybanner, err := r.data.db.Banner.Query().
		Where(banner.Status(1)).
		Order(ent.Asc(banner.FieldSort)).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select url from banner error: %v", err),
			"category",2)
		return g, err
	}
	for _, v := range resQuerybanner {
		if 5-len(homepage.Banner) == 0{
			break
		}
		homepage.Banner = append(homepage.Banner, &biz.Banner{
			ID: v.ID,
			Title: v.Title,
			Type: v.Type,
			URL: v.URL,
			LinkID: v.LinkID,
			Sort: v.Sort,
			//Bucket: v.Bucket,
		})
	}

	proclient, err1 := r.data.db.Product.
		Query().
		Where(product.Status(1)).
		Order(ent.Desc(product.FieldHistory)).
		All(ctx)
	if err1 != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select name content from product error: %v", err1),
			"evt:","",
			"category:",2)
		return g, err1
	}

	for _, v := range proclient {
		if 8-len(homepage.Pro)==0{
			break
		}
		if v.Recommend == 1{
			homepage.Pro = append(homepage.Pro, &biz.Product{
				ID: v.ID,
				CID: v.CategoryID,
				CName: v.CategoryName,
				History: v.History,
				ProName: v.Name,
				ProInfo: v.Content,
				Recommend: v.Recommend,
			})
		}
	}
	for _,v := range proclient{
		if 8-len(homepage.Pro)==0{
			break
		}
		homepage.Pro = append(homepage.Pro, &biz.Product{
			ID: v.ID,
			CID: v.CategoryID,
			CName: v.CategoryName,
			History: v.History,
			ProName: v.Name,
			ProInfo: v.Content,
			Recommend: v.Recommend,
		})
	}

	plancategoryClient, err2 := r.data.db.PlanCategory.
		Query().
		Order(ent.Asc(plancategory.FieldSort)).
		All(ctx)
	if err2 != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select name content from plan_category error: %v", err2),
			"evt:","",
			"category:",2)
		return g, err2
	}
	for _, v := range plancategoryClient {
		if 8-len(homepage.PlanCategory) == 0 {
			break
		}
		homepage.PlanCategory = append(homepage.PlanCategory, &biz.PlanCategory{
			ID: v.ID,
			Sort: v.Sort,
			//Bucket: v.Bucket,
			URL: v.URL,
			PlanName: v.Name,
			PlanInfo: v.Content,

		})
	}
	g = &homepage
	return g, nil
}

func (r *homePageRepo) GetHomepagePro(ctx context.Context, g *biz.HomepagePro,redisKey string) (*biz.HomepagePro, error) {
	var (
		homepagepro biz.HomepagePro
	)

	level1, err := r.data.db.ProductCategory.Query().
		Where(productcategory.Status(1), productcategory.Level(1)).
		Order(ent.Asc(productcategory.FieldSort)).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select name id pid from product_category where status level1 error: %v", err),
			"evt:","",
			"category:",2)
		return g, err
	}
	level2, err1 := r.data.db.ProductCategory.Query().
		Where(productcategory.Level(2), productcategory.Status(1)).
		Order(ent.Asc(productcategory.FieldSort)).
		All(ctx)
	if err1 != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select name id pid from product_category where status level2 error: %v", err1),
			"evt:","",
			"category:",2)
		return g, err1
	}
	proAll, err2 := r.data.db.Product.Query().
		Where(product.Status(1)).
		Order(ent.Desc(product.FieldHistory)).
		All(ctx)
	if err2 != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select * from product where status recommend error: %v", err2),
			"evt:","",
			"category:",2)
		return g, err2
	}
	for _, v1 := range level1 {
		var level1data = biz.ProductCategory{
			ID:   v1.ID,
			PID:  v1.ParentID,
			Name: v1.Name,
		}
		homepagepro.Proc = append(homepagepro.Proc, &level1data)
		for _, v2 := range level2 {
			if v1.ID == v2.ParentID {
				var level2data = biz.Level2{
					ID:   v2.ID,
					PID:  v2.ParentID,
					Name: v2.Name,
				}
				level1data.Children = append(level1data.Children, &level2data)
				for _, prov := range proAll {
					slink := make([]*string, 0)
					temp := prov
					if prov.CategoryID == v2.ID {
						slink = append(slink, &temp.ServiceLink)
						level2data.Children = append(level2data.Children, &biz.Product{
							ID:      prov.ID,
							CID:     prov.CategoryID,
							CName:   prov.CategoryName,
							ProName: prov.Name,
						})
					}
				}
			}
		}
	}

	prodb, err3 := r.List(ctx,&redisKey)
	if err3 != nil {
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("select name from product order by history desc error: %v", err3),
			"evt:","",
			"category:",2)
		return g, err3
	}

	for _, v := range prodb {
		homepagepro.HistorySearch = append(homepagepro.HistorySearch, &biz.Search{
			ID: v.ID,
			Count: v.Count,
			Keyword: v.Keyword,
			Utime: v.Utime,
		})
	}
	g = &homepagepro
	return g, nil
}


func (r *homePageRepo)GetSite(ctx context.Context)([]*biz.SiteCategory,error){
	var (
		siteCategorySlice = make([]*biz.SiteCategory,0)
	)
	resQueryCategory,err := r.data.db.Site.Query().
		Where(site.Status(1)).
		Order(ent.Asc(site.FieldSort)).
		Select(site.FieldCategory).
		All(ctx)
	if err !=nil{
		r.log.WithContext(ctx).Errorw("msg",fmt.Sprintf("query site error:%v",err),
			"category",2)
		return nil, err
	}

	for _,v := range resQueryCategory{
		//fmt.Printf("data/resQueryCategory data:%v\n",*v)
		temp := v
		resQuerySite,err:=r.data.db.Site.Query().
			Where(site.Category(temp.Category)).
			Order(ent.Asc(site.FieldCategory)).
			All(ctx)
		if err != nil{
			return nil, err
		}
		var siteSlice = make([]*biz.Site,0)
		for _,v1:=range resQuerySite{
			//fmt.Printf("data/resQuerySite data:%v\n",*v1)
			siteSlice = append(siteSlice,&biz.Site{
				ID: v1.ID,
				Name: v1.Name,
				Link: v1.Link,
				Category: v1.Category,
			})
		}
		//for _,v2:=range siteSlice{
		//	fmt.Printf("data/siteSlice data:%v\n",*v2)
		//}
		siteCategorySlice = append(siteCategorySlice,&biz.SiteCategory{
			Name: v.Category,
			Site: siteSlice,
		})
	}
	//for _,v := range siteCategorySlice{
	//	fmt.Printf("data/siteCategory data:%v\n",*v)
	//}

	return siteCategorySlice,err
}
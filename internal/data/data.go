package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gy_home/internal/conf"
	"gy_home/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewHomepageRepo,
	NewEntClient,
	//NewRedisClient,
	NewPlanRepo,
	)

// Data .
type Data struct {
	// TODO wrapped database client
	db *ent.Client
	rdb *redis.Client
}

// NewEntClient create new client of the ent
func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	l := log.NewHelper(log.With(logger))

	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		l.Fatalw("msg",fmt.Sprintf("failed opening connection to db:%v", err),
			"evt:","",
			"category:",2)
	}
	return client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger,entClient *ent.Client) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger))

	rdb := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		Password: c.Redis.Password,
		DB: 0,
		DialTimeout: c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout: c.Redis.ReadTimeout.AsDuration(),
	})

	d := &Data{
		db: entClient,
		rdb: rdb,
	}
	return d, func() {
		log.Infow("msg","closing the data resourse",
			"evt:","",
			"category:",2)
		if err := d.db.Close(); err != nil {
			log.Errorw("msg",fmt.Sprintf("db close error:%v",err),
				"evt:","",
				"category:",2)
		}
		if err := d.rdb.Close();err!=nil{
			log.Errorw("msg",fmt.Sprintf("rdb close error:%v",err),
				"evt:","",
				"category:",2)
		}
	}, nil
}

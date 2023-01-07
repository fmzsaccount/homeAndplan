package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"gy_home/internal/errorcode"

	"github.com/go-kratos/kratos/v2/middleware"
	//"github.com/go-kratos/kratos/v2/transport/grpc"
	googlegrpc "google.golang.org/grpc"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v12 "gy_home/api/homepage/v1"
	v13 "gy_home/api/plan/v1"
	v1 "gy_home/api/login/v1"
	"gy_home/customizeLog"
	"gy_home/internal/conf"
	"gy_home/internal/service"
	nhttp "net/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,logger log.Logger, homepage *service.HomepageService,plan *service.DetailsOfPlanService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			customizeLog.Server(logger),
			MiddlewareCors(logger,c),
			logMiddleware(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	opts=append(opts,http.Filter(handlers.CORS(
		//域名配置
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type","Authorization","Username"}),
		handlers.AllowedMethods([]string{"GET","POST","DELETE","PUT","OPTIONS","PATCH"}),
	)))

	opts = append(opts, http.ErrorEncoder(func(w nhttp.ResponseWriter, r *nhttp.Request, err error) {
		se := errors.FromError(err)
		codec ,_ := http.CodecForRequest(r,"Accept")
		body, err := codec.Marshal(se)
		if err != nil {
			w.WriteHeader(nhttp.StatusInternalServerError)
			return
		}
		//fmt.Println(string(body))
		//w.Header().Set("Content-Type", Contenttype(codec.Name()))
		// 设置HTTP Status Code
		w.WriteHeader(200)
		w.Write(body)
	}),
	)
	srv := http.NewServer(opts...)
	v12.RegisterHomepageHTTPServer(srv, homepage)
	v13.RegisterDetailsOfPlanHTTPServer(srv,plan)
	return srv
}

func MiddlewareCors(logger log.Logger,c *conf.Server)middleware.Middleware{
	return func(handler middleware.Handler)middleware.Handler{
		return func (ctx context.Context,req interface{})(interface{},error){
			if ts,ok:= transport.FromServerContext(ctx);ok{

					token := ts.RequestHeader().Get("Authorization")
					ip := ts.RequestHeader().Get("X-Real-IP")
					ctx = context.WithValue(ctx,"ip",ip)
					//v12.
					conn,err := googlegrpc.Dial(c.Other.Addr,googlegrpc.WithInsecure())
					if err != nil{
						log.NewHelper(logger).WithContext(ctx).Errorw("msg",fmt.Sprintf("failed to connect %v",c.Other.Addr),
							"category",2)
						return nil, err
					}
					defer conn.Close()
					client := v1.NewLoginClient(conn)
					res,err := client.RefreshAuthorization(ctx,&v1.RefreshAuthorizationRequest{
						Authorization: token,
						Acl: 1,
					})
					ctx = context.WithValue(ctx,"RefreshAuthorizationReply",res)
					if res.Code != 200 || err!=nil{
						log.NewHelper(logger).WithContext(ctx).Errorw("msg","Invalid identity credentials",
							"category",2)
						return nil,errors.New(errorcode.NOT_FOUND_USER,"",res.GetMessage())
					}
			}else {
				log.NewHelper(logger).WithContext(ctx).Errorw("msg","Invalid identity credentials",
					"category",2)
				return nil,errors.New(errorcode.NOT_FOUND_USER,"","authorization expired")
			}
			return handler(ctx,req)
		}
	}
}

func logMiddleware(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			//level, _ := extractError(err)
			_ = log.WithContext(ctx, logger)
			return
		}
	}
}
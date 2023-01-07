package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "gy_home/api/homepage/v1"
	v12 "gy_home/api/plan/v1"
	"gy_home/customizeLog"
	"gy_home/internal/conf"
	"gy_home/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.HomepageService, plan *service.DetailsOfPlanService,logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			customizeLog.Server(logger),
			logMiddleware(logger),
			MiddlewareCors(logger,c),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterHomepageServer(srv, greeter)
	v12.RegisterDetailsOfPlanServer(srv, plan)
	return srv
}

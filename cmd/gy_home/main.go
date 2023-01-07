package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gy_home/customizeLog"
	"gy_home/internal/conf"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	//path := "./configs/config.yaml"
	//c := config.New(
	//		config.WithSource(
	//				file.NewSource(path),
	//			),
	//	)
	defer c.Close()

	if err := c.Load(); err != nil {
		toany := any(err)
		panic(toany)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		toany := any(err)
		panic(toany)
	}

	logger := log.With(customizeLog.Logger(bc.Zaplog))
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		toany := any(err)
		panic(toany)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		toany := any(err)
		panic(toany)
	}
}


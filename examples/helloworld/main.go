package main

import (
	"context"

	"github.com/go-hare/haremicro"
	"github.com/go-hare/haremicro/config"
	cetcd "github.com/go-hare/haremicro/config/source/etcd"
	greeter "github.com/go-hare/haremicro/examples/helloworld/proto"
	"github.com/go-hare/haremicro/logger"
	mlogrus "github.com/go-hare/haremicro/logger/logrus"
	"github.com/go-hare/haremicro/registry"
	"github.com/go-hare/haremicro/registry/etcd"
	"github.com/sirupsen/logrus"
)

var (
	cfg config.Config
)

type TestInfo struct {
	Test string `json:"test"`
}

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	rsp.Greeting = "Hello " + req.Name
	if cfg != nil {
		logger.Info("config data:", cfg.Map())
		// config in etcd:
		// key: helloworld/test
		// value: {"test": "test"}
		var t1, t2 TestInfo
		cfg.Get("test").Scan(&t1)
		cfg.Get("1", "t").Scan(&t2)
		logger.Info("test : ", t1)
		logger.Info("test : ", t2)
	}
	return nil
}

func main() {
	serviceName := "helloworld"

	logger.DefaultLogger = mlogrus.NewLogger(mlogrus.WithJSONFormatter(&logrus.JSONFormatter{}))
	logger.Init(logger.WithLevel(logger.TraceLevel))

	logger.Logf(logger.InfoLevel, "Example Name: %s", serviceName)

	etcdAddress := "127.0.0.1:2379"

	var err error
	cfg, err = config.NewConfig(config.WithSource(
		cetcd.NewSource(
			cetcd.WithAddress(etcdAddress),
			cetcd.WithPrefix(serviceName),
			cetcd.StripPrefix(true),
			cetcd.WithPrefixCreate(true),
		),
	))
	if err != nil {
		logger.Fatal(err)
		return
	}

	service := haremicro.NewService(
		haremicro.Name(serviceName),
		haremicro.Registry(etcd.NewRegistry(registry.Addrs([]string{etcdAddress}...))),
	)

	service.Init()

	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

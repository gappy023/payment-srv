package main

import (
	"fmt"

	"github.com/gappy023/payment-srv/handler"

	payment "github.com/gappy023/payment-srv/proto/payment"

	"github.com/gappy023/basic"
	"github.com/gappy023/payment-srv/model"

	config "github.com/gappy023/basic/config"
	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	etcd "github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.payment"),
		micro.Version("latest"),
		micro.Register(micReg),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) error {
			model.Init()
			handler.Init()
			return nil
		}),
	)

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	//	micro.RegisterSubscriber("mu.micro.book.srv.payment", service.Server(), new(subscriber.Payment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}

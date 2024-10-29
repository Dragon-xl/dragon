package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "getcaptcha"
	version = "latest"
)

func main() {
	cons := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Registry(cons),
		micro.Address(":12341"),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

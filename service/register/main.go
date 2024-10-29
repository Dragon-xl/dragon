package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"register/handler"
	"register/model"
	register "register/proto"
)

var (
	service = "register"
	version = "latest"
)

func main() {
	// Create service
	model.InitRedis()
	err := model.InitDB()
	if err != nil {
		fmt.Println("service register initDB failed err:", err)
		return
	}
	conReg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(conReg),
		micro.Address(":12342"))
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := register.RegisterRegisterHandler(srv.Server(), new(handler.Register)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

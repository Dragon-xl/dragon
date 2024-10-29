package utils

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
)

// 初始化 micro
func InitMicro() micro.Service {
	consulReg := consul.NewRegistry()
	return micro.NewService(micro.Registry(consulReg))
}

package main

import (
	"dragon/controller"
	"dragon/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化连接池redis
	//model.InitRedis()
	err := model.InitDB()
	if err != nil {
		fmt.Println("main InitDB err:", err)
		return
	}

	router := gin.Default()
	router.Static("/home", "view")

	v1 := router.Group("api/v1")
	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSmscd)
		v1.POST("/users", controller.PostRet)
		v1.GET("/areas")
	}

	router.Run(":8080")

}

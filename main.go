package main

import (
	"dragon/controller"
	"dragon/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化连接池redis
	model.InitRedis()
	err := model.InitDB()
	if err != nil {
		fmt.Println("main InitDB err:", err)
		return
	}

	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "192.168.101.45:6379", "123", []byte("123456"))
	router.Use(sessions.Sessions("dragon", store))
	router.Static("/home", "view")

	v1 := router.Group("api/v1")
	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSmscd)
		v1.POST("/users", controller.PostRet)
		v1.GET("/areas", controller.GetArea)
		v1.POST("/sessions", controller.PostLogin)
		v1.DELETE("/session", controller.DeleteSession)
		v1.GET("/user", controller.GetUserInfo)
		v1.PUT("/user/name", controller.PutUserInfo)
	}

	router.Run(":8080")

}

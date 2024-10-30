package controller

import (
	"context"
	"dragon/proto/getCaptcha"
	"dragon/proto/register"
	"dragon/utils"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
)

func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	//指定服务发现为consul
	srv := utils.InitMicro()
	cli := getCaptcha.NewGetCaptchaService("getcaptcha", srv.Client())
	rsp, err := cli.Call(context.TODO(), &getCaptcha.CallRequest{Uuid: uuid})
	if err != nil {
		fmt.Println("rpc:getIMg:", err)
		return
	}
	var img captcha.Image
	json.Unmarshal(rsp.Img, &img)
	png.Encode(ctx.Writer, img)
	//ctx.String(http.StatusOK, uuid)
}
func GetSmscd(ctx *gin.Context) {
	phone := ctx.Param("phone")
	capCode := ctx.Query("text")
	uuid := ctx.Query("id")
	srv := utils.InitMicro()
	cli := register.NewRegisterService("register", srv.Client())
	rsp, err := cli.SendSms(context.TODO(), &register.CallRequest{Uuid: uuid, CapCode: capCode, Phone: phone})
	if err != nil {
		fmt.Println("rpc:getSendSms err :", err)
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

//发送注册信息

func PostRet(ctx *gin.Context) {
	var reqDate struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	err := ctx.Bind(&reqDate)
	if err != nil {
		fmt.Println("func PostRet bind failed \n")
		return
	}
	srv := utils.InitMicro()
	cli := register.NewRegisterService("register", srv.Client())
	rsp, err := cli.Register(context.TODO(), &register.RegRequest{
		Mobile:   reqDate.Mobile,
		Password: reqDate.Password,
		SmsCode:  reqDate.SmsCode,
	})
	if err != nil {
		fmt.Println("rpc:postReg err :", err)
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

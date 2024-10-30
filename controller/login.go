package controller

import (
	"dragon/model"
	"dragon/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	rsp := make(map[string]interface{})
	userName := session.Get("userName")
	if userName == nil {
		//用户未登录
		rsp["errno"] = utils.RECODE_SESSIONERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		name := make(map[string]string)
		name["name"] = userName.(string)
		rsp["data"] = name
		rsp["errno"] = utils.RECODE_OK
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, rsp)

}
func PostLogin(ctx *gin.Context) {
	var LoginData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	err := ctx.Bind(&LoginData)
	if err != nil {
		fmt.Println("bind failed err:", err)
		return
	}
	rsp := make(map[string]string)
	name, err := model.LoginRetName(LoginData.Mobile, LoginData.Password)
	if err != nil {
		fmt.Println("post login failed err:", err)
		rsp["errno"] = utils.RECODE_LOGINERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	//将获取的用户名写入session
	v := sessions.Default(ctx)
	v.Set("userName", name)
	v.Save()
	rsp["errno"] = utils.RECODE_OK
	rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	ctx.JSON(http.StatusOK, rsp)
}

// 退出 Logout
func DeleteSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("userName")
	err := session.Save()
	rsp := make(map[string]string)
	if err != nil {
		rsp["errno"] = utils.RECODE_IOERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		rsp["errno"] = utils.RECODE_OK
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, rsp)
}

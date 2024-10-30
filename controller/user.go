package controller

import (
	"dragon/model"
	"dragon/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userName := session.Get("userName")
	rsp := make(map[string]interface{})
	if userName == nil {
		//用户未登录 恶意访问
		rsp["errno"] = utils.RECODE_SESSIONERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		//查MYSQL 提取用户信息
		user, err := model.GetUserInfo(userName.(string))
		if err != nil {
			rsp["errno"] = utils.RECODE_SESSIONERR
			rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		} else {
			data := make(map[string]interface{})
			data["user_id"] = user.ID
			data["name"] = user.Name
			data["mobile"] = user.Mobile
			data["read_name"] = user.Real_name
			data["id_card"] = user.Id_card
			data["avatar_url"] = user.Avatar_url
			rsp["data"] = data
			rsp["errno"] = utils.RECODE_OK
			rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		}
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 更新用户名
func PutUserInfo(ctx *gin.Context) {
	//获取当前用户名
	session := sessions.Default(ctx)
	userName := session.Get("userName")
	//获取新用户名
	var name struct {
		Name string `json:"name"`
	}
	rsp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, rsp)
	ctx.Bind(&name)
	//判断用户是否存在
	if userName == nil {
		rsp["errno"] = utils.RECODE_SESSIONERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	//更新用户名
	err := model.UpdateUserName(userName.(string), name.Name)
	if err != nil {
		rsp["errno"] = utils.RECODE_DBERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	session.Set("userName", name.Name)
	err = session.Save()
	if err != nil {
		rsp["errno"] = utils.RECODE_SESSIONERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	rsp["errno"] = utils.RECODE_OK
	rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	rsp["data"] = name
}

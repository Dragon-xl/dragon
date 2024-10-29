package controller

import (
	"dragon/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSession(ctx *gin.Context) {
	resp := make(map[string]string)
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)
}

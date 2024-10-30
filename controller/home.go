package controller

import (
	"dragon/model"
	"dragon/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

// 获取地域信息
func GetArea(ctx *gin.Context) {
	//现在redis中查数据
	conn := model.RedisPool.Get()
	defer conn.Close()
	areaDate, _ := redis.Bytes(conn.Do("get", "areaData"))
	var areas []model.Area
	if len(areaDate) == 0 {
		//mysql取出数据
		model.DB.Find(&areas)
		//写入redis
		areaJson, _ := json.Marshal(areas)
		conn.Do("set", "areaData", areaJson)
	} else {
		json.Unmarshal(areaDate, &areas)
	}
	rsp := make(map[string]interface{})
	rsp["errno"] = utils.RECODE_OK
	rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	rsp["data"] = areas
	ctx.JSON(http.StatusOK, rsp)
}

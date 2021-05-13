package app

import (
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type AppGin struct {
	C *gin.Context
}

type AppJsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *AppGin) GetQueryWithDefault(key string, defaultValue string) string {
	val, b := g.C.GetQuery(key)
	if b {
		return val
	} else {
		return defaultValue
	}
}

func (g *AppGin) GetIntQueryWithDefault(key string, defaultValue int) int {
	if val, b := g.C.GetQuery(key); b {
		if ival, b1 := cast.ToIntE(val); b1 == nil {
			return ival
		}
	}
	return defaultValue
}

func (g *AppGin) MakeResponse(httpCode int, statusCode int, data interface{}) {
	g.C.JSON(httpCode, AppJsonResponse{
		Code:    statusCode,
		Message: e.GetMessage(statusCode),
		Data:    data,
	})
	return
}

func (g *AppGin) MakeEmptyResponse(httpCode int, statusCode int) {
	g.C.JSON(httpCode, AppJsonResponse{
		Code:    statusCode,
		Message: e.GetMessage(statusCode),
		Data:    nil,
	})
	return
}

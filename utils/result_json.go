package utils

import "github.com/gin-gonic/gin"

type ResultJson struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Obj     any    `json:"obj,omitempty"`
}

const (
	SUCCESS = iota
	ERROR
)

var Codes = []int{200, 500}
var Messages = []string{"请求成功", "请求失败"}

func respSuccess(ctx *gin.Context, resultJson ResultJson) {
	ctx.AbortWithStatusJSON(resultJson.Code, resultJson)
}
func respError(ctx *gin.Context, resultJson ResultJson) {
	ctx.AbortWithStatusJSON(resultJson.Code, resultJson)
}
func RespSuccess(ctx *gin.Context, message string, obj any) {
	if message == "" {
		message = Messages[SUCCESS]
	}
	resultJson := &ResultJson{Codes[SUCCESS], message, obj}
	respSuccess(ctx, *resultJson)
}
func RespError(ctx *gin.Context, message string) {
	if message == "" {
		message = Messages[ERROR]
	}
	resultJson := &ResultJson{Codes[ERROR], message, nil}
	respError(ctx, *resultJson)
}

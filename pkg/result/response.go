package result

import (
	"ChatServer/consts" // 你的错误码定义包
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	Code    int32    `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string `json:"trace_id"`
}

// Result 返回响应
func Result(c *gin.Context, data interface{}, message string, code int32) {
	traceId := c.GetString("trace_id")
	if message == ""{
		message = consts.GetMessage(code)
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
		TraceId: traceId,
	})
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	Result(c, data, "", consts.CodeSuccess)
}

// Fail 返回失败响应
func Fail(c *gin.Context, data interface{}, code int32) {
	Result(c, data, "", code)
}

// SuccessWithMessage 返回成功响应并自定义消息
func SuccessWithMessage(c *gin.Context, data interface{}, message string) {
	Result(c, data, message, consts.CodeSuccess)
}

// FailWithMessage 返回失败响应并自定义消息
func FailWithMessage(c *gin.Context, data interface{}, message string, code int32) {
	Result(c, data, message, code)
}


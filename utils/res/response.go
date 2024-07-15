package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"` // code
	Msg  string `json:"msg"`  // description message
	Data any    `json:"data"` // returns data
}

// constant
const (
	SUCCESS = 0
	ERROR   = 7
)

// Result Response encapsulation
func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Ok returns success data
func Ok(msg string, data any, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

// OkWithMsg returns success information info from which description message is available
func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, msg, map[string]interface{}{}, c)
}

// OkWithData returns success information for which data is available
func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, "请求成功", data, c)
}

// Fail returns error data
func Fail(msg string, data any, c *gin.Context) {
	Result(ERROR, msg, data, c)
}

// FailWithMsg returns error information for which description message is available
func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, msg, map[string]interface{}{}, c)
}

// FailErrorCode Error code is returned after request fails.Business: parameter binding failed,etc
func FailErrorCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), msg, map[string]interface{}{}, c)
		return
	}
	Result(int(code), "系统未知错误，请反馈给管理员", map[string]interface{}{}, c)
}

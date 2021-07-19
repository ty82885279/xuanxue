package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResCode int64

const (
	CodeSuccess ResCode = 2000 + iota
	CodeSuccessOne
	CodeSuccessMuti
	CodeInvalidParam ResCode = 3000 + iota
	CodeNoData
	CodeServiceBusy
)

func (r ResCode) Msg() string {

	msg, ok := codeMsgMap[r]
	if !ok {
		return codeMsgMap[CodeServiceBusy]
	}
	return msg
}

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "请求成功",
	CodeSuccessOne:   "请求成功",
	CodeSuccessMuti:  "请求成功",
	CodeInvalidParam: "参数错误",
	CodeNoData:       "没有记录",
	CodeServiceBusy:  "服务繁忙",
}

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
func ResponseSuccessWithCode(c *gin.Context, code ResCode, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}
func ResponseErr(c *gin.Context, errCode ResCode) {

	c.JSON(http.StatusOK, &ResponseData{
		Code: errCode,
		Msg:  errCode.Msg(),
		Data: nil,
	})
}

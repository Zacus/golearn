package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code":10001,	//错误码
	"msg":"XX"		//提示信息
	"data"：{}		//数据
}
*/

type ResposeData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResposeError(c *gin.Context, code ResCode) {
	// rd := &ResposeData{
	// 	Code: code,
	// 	Msg:  code.Msg(),
	// 	Data: nil,
	// }
	c.JSON(http.StatusOK, &ResposeData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResposeErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResposeData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResposeSuccess(c *gin.Context, data interface{}) {
	rd := &ResposeData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

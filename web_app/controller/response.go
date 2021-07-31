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

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	// rd := &ResposeData{
	// 	Code: code,
	// 	Msg:  code.Msg(),
	// 	Data: nil,
	// }
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

package controller

import (
	"golearn/web_app/logic"
	"golearn/web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//处理注册请求的函数
func SignUpHandler(c *gin.Context) {

	//1.参数校验.
	var p = new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {

		zap.L().Error("SignUp with invaild param", zap.Error(err))
		//请求参数有误，直接返回响应
		//判断是否为validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": err.Error(),
			// })
			ResposeErrorWithMsg(c, CodeInvaildParam, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": removeTopStruct(errs.Translate(trans)),
		// })
		ResposeErrorWithMsg(c, CodeInvaildParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// //手动校验
	// if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	// 	//请求参数有误，直接返回响应
	// 	zap.L().Error("SignUp with invaild param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }
	//fmt.Println(p)
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUp with invaild param", zap.Error(err))
		//请求参数有误，直接返回响应
		//判断是否为validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": err.Error(),
			// })
			ResposeErrorWithMsg(c, CodeInvaildParam, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": removeTopStruct(errs.Translate(trans)),
		// })
		ResposeErrorWithMsg(c, CodeInvaildParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//3.返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "Registered successfully",
	// })
	ResposeSuccess(c, "Registered successfully")

}

func LoginHandler(c *gin.Context) {

	//1.获取请求参数，参数校验.
	var p = new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invaild param", zap.Error(err))
		//请求参数有误，直接返回响应
		//判断是否为validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResposeErrorWithMsg(c, CodeInvaildParam, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResposeErrorWithMsg(c, CodeInvaildParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.逻辑处理
	if err := logic.Login(p); err != nil {
		ResposeError(c, CodeInvaildLogin)
		return
	}
	//3.返回响应
	ResposeSuccess(c, "login sucess")
}

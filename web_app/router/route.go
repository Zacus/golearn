package router

import (
	"golearn/web_app/controller"
	"golearn/web_app/logger"
	"golearn/web_app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	r.POST("/api/v1/signup", controller.SignUpHandler)

	//登录业务路由
	r.POST("/api/v1/login", controller.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//如果是登录的用户，判断请求头中是否有 有效的JWT
		c.String(http.StatusOK, "pong")

	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r

}

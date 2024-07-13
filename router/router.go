package router

import (
	"github.com/gin-gonic/gin"
	"go_project/bluebell/controller"
	"go_project/bluebell/logger"
	"go_project/bluebell/middlewares"
	"net/http"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//c.Request.Header.Get("Authorization")
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

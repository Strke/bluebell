package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/bluebell/logic"
	"go_project/bluebell/models"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	// 1、获取参数参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "err-params-request",
		})
		return
	}
	// 2、业务处理
	logic.SignUp()
	// 3、返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

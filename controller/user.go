package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
	"go_project/bluebell/dao/mysql"
	"go_project/bluebell/logic"
	"go_project/bluebell/models"
)

func SignUpHandler(c *gin.Context) {
	// 1、获取参数参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		// 判断当前的错误是不是校验错误，有些时候可能会是其他错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 2、业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, CodeSuccess)
}

func LoginHandler(c *gin.Context) {
	//1、获取参数并且参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	//2、业务处理
	atoken, _, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}

	//3、返回响应
	ResponseSuccess(c, atoken)
}

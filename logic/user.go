package logic

import (
	"go_project/bluebell/dao/mysql"
	"go_project/bluebell/models"
	"go_project/bluebell/pkg/jwt"
	snowFlake "go_project/bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户存在不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	//生成UID
	userID := snowFlake.GenID()
	//密码加密
	U := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(&U)
}

func Login(p *models.ParamLogin) (atoken, rtoken string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.CheckUserPassword(user); err != nil {
		return "", "", err
	}
	return jwt.GenToken(user.UserID)
}

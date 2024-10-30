package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	snowFlake "bluebell/pkg/snowflake"
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

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{

		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.CheckUserPassword(user); err != nil {
		return nil, err
	}
	aToken, rToken, err := jwt.GenToken(user.UserID)
	if err != nil {
		return nil, err
	}
	user.AToken = aToken
	user.RToken = rToken
	return

}

package mysql

import "errors"

var (
	ErrorUserExist    = errors.New("用户已存在")
	ErrorUserNoExist  = errors.New("用户不存在")
	ErrorUserPassword = errors.New("账号密码不匹配")
	ErrorInvalidID    = errors.New("无效的ID")
)

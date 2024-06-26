package logic

import (
	"go_project/bluebell/dao/mysql"
	snowFlake "go_project/bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp() {
	//判断用户存在不存在
	mysql.QueryUserByUserName()
	//生成UID
	snowFlake.GenID()
	//密码加密

	//保存进数据库
	mysql.InsertUser()

}

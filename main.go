package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_project/bluebell/dao/mysql"
	"go_project/bluebell/logger"
	"go_project/bluebell/setting"
)

func main() {
	if err := setting.Init("../config.yaml"); err != nil {
		fmt.Println("load config failed, err:#{err}\n")
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, "debug"); err != nil {
		fmt.Println("init logger failed, err:#{err}\n")
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:#{err}\n")
		return
	}
	defer mysql.Close()
	//if err := redis.Init(setting.Conf.RedisConfig); err != nil {
	//	fmt.Println("init redis failed, err:#{err}\n")
	//	return
	//}
	//defer redis.Close()
	//if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
	//	fmt.Println("init snowflake failed, err:#{err}\n")
	//}
}

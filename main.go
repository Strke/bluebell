package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_project/bluebell/controller"
	"go_project/bluebell/dao/mysql"
	"go_project/bluebell/dao/redis"
	"go_project/bluebell/logger"
	snowflake "go_project/bluebell/pkg/snowflake"
	"go_project/bluebell/router"
	"go_project/bluebell/setting"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: bluebell config.yaml")
	//	return
	//}
	// 加载配置
	// linux: /mnt/d/go_project/bluebell/conf/config.yaml
	// windows: D:/go_project/bluebell/conf/config.yaml
	if err := setting.Init("/mnt/d/go_project/bluebell/conf/config.yaml"); err != nil {
		fmt.Printf("load config failed, err1:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	//// 初始化gin框架内置的校验器使用的翻译器
	//if err := controller.InitTrans("zh"); err != nil {
	//	fmt.Printf("init validator trans failed, err:%v\n", err)
	//	return
	//}
	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans faild, err:%v", err)
		return
	}

	// 注册路由
	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

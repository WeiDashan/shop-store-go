package cmd

import (
	"fmt"
	"github.com/WeiDashan/shop-go/conf"
	"github.com/WeiDashan/shop-go/global"
	"github.com/WeiDashan/shop-go/global/constants"
	"github.com/WeiDashan/shop-go/router"
	"github.com/WeiDashan/shop-go/utils"
)

func Start() {
	var initErr error
	// 初始化系统配置文件
	conf.InitConfig()

	// 初始化日志组件
	global.Logger = conf.InitLogger()

	// 初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db

	// 初始化Redis连接
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化RabbitMQ连接
	rbMQClient, err := conf.InitRabbitMQ()
	global.RabbitMQClient = rbMQClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	// 初始化过程中遇到错误的处理
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// 初始化过滤器配置文件
	constants.InitFilter()

	// 初始化系统路由
	router.InitRouter()

}
func Clean() {
	fmt.Println("===========Clean==============")
}

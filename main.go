package main

import (
	"fmt"
	Mysql "gin/databases"
	"gin/global"
	"gin/middleware"
	"gin/redis"
	Router "gin/router"
	"github.com/gin-gonic/gin"
)

var VERSION = "0.0.1"

func main() {
	app := gin.Default()
	app.Use(middleware.Logger())
	// 注册数据库
	Mysql.InitDB()
	// 注册接口列表
	Router.Register(app)
	// 执行 Redis
	redis.InitRedis()
	// 注册 zookeeper
	middleware.ZookeeperRegister()

	fmt.Printf(`
🎉 欢迎使用蛙蛙低代码客户端
当前版本：%s

`, VERSION)

	err := app.Run(global.ADDR)

	if err != nil {
		fmt.Printf(`
❌ 蛙蛙低代码客户端报错
`)
	}
}

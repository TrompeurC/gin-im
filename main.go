package main

import (
	"gin-im/db"
	"gin-im/router"
	"gin-im/utils"
)

func main() {
	// 初始化配置文件
	utils.InitConfig()
	// 初始化数据库
	db.InitMysql()
	r := router.InitRouter()

	r.Run(":8080")
}

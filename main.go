package main

import (
	"gin-im/db"
	"gin-im/router"
)

func main() {
	// 初始化数据库
	db.InitMysql()
	r := router.InitRouter()

	r.Run(":8080")
}

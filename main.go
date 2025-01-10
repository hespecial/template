package main

import (
	"blog/global"
	"blog/internal/router"
)

func main() {
	global.InitConfig()
	global.InitLogger()
	global.InitMysql()
	//global.InitRedis()
	router.StartServer()
}

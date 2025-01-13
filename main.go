package main

import (
	"template/global"
	"template/internal/router"
)

func main() {
	global.InitConfig()
	global.InitLogger()
	global.InitMysql()
	//global.InitRedis()
	router.StartServer()
}

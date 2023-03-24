package main

import (
	"gwa_server/core"
	"gwa_server/global"
)

func main() {
	//读取配置文件
	core.InitConf()

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

	//初始化server
	core.RunWindowsServer()
}

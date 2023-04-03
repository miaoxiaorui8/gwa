package main

import (
	"gwa_server/core"
	_ "gwa_server/docs"
	"gwa_server/global"
)

// @title Gwa_server API文档
// @version 1.0
// @description Gwa_server API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()

	//初始化日志
	global.GWA_Log = core.InitLogger()

	//连接数据库
	global.GWA_DB = core.InitGorm()

	//初始化表
	if global.GWA_DB != nil {
		core.RegisterTables()
		// 程序结束前关闭数据库链接
		dB, _ := global.GWA_DB.DB()
		defer dB.Close()
	}

	//初始化server
	core.RunWindowsServer()
}

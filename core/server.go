package core

import (
	"fmt"
	"gwa_server/global"
	"gwa_server/initialize"
)

func RunWindowsServer() {
	fmt.Printf(`
	欢迎使用 Gwa
	当前版本:v1.0.0
	默认运行地址:http://127.0.0.1:8080
`)

	//初始化路由
	router := initialize.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}

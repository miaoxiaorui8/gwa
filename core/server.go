package core

import (
	"fmt"
	"gwa_server/global"
	"gwa_server/initialize"
	"text/template"
)

func RunWindowsServer() {
	fmt.Printf(`
	欢迎使用 Gwa
	当前版本:v1.0.0
	默认运行地址:http://127.0.0.1:8080
	swagApi文档地址:http://127.0.0.1:8080/swagger/index.html
`)

	//初始化路由
	router := initialize.Routers()
	//自定义模板函数 注意要把这个函数放在加载模板之前
	router.SetFuncMap(template.FuncMap{})
	//加载模板 放在配置路由前面
	router.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录 第一个参数表示路由,第二个表示映射的目录
	router.Static("/static", "./static")
	addr := global.GWA_Config.System.Addr()
	err := router.Run(addr)
	if err != nil {
		return
	}
}

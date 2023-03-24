package initialize

import (
	"github.com/gin-gonic/gin"
	"gwa_server/global"
	"gwa_server/routers"
)

// 初始化总路由
func Routers() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	systemRouter := routers.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		systemRouter.InitLoginRouter(PublicGroup) //登陆路由
	}
	return Router
}

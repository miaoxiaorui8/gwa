package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gwa_server/global"
	"gwa_server/middleware"
	"gwa_server/routers"
)

// 初始化总路由
func Routers() *gin.Engine {
	gin.SetMode(global.GWA_Config.System.Env)
	Router := gin.Default()
	systemRouter := routers.RouterGroupApp.System
	PublicGroup := Router.Group("")
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//swag文档
	PublicGroup.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	{
		systemRouter.InitLoginRouter(PublicGroup, PrivateGroup) //登陆路由
	}
	return Router
}

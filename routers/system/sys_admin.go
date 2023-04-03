package system

import (
	"github.com/gin-gonic/gin"
	v1 "gwa_server/api/v1"
	"gwa_server/render"
)

type AdminRouter struct{}

func (s *AdminRouter) InitLoginRouter(PubRouter, Router *gin.RouterGroup) {
	htmlRenderApi := render.RenderGroupHtml.AdminRenderGroup.SysAdminRender
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SysAdminApi
	apiPubHtml := PubRouter.Group("gwa")
	{
		apiPubHtml.GET("/login", htmlRenderApi.Login)
	}
	apiHtml := Router.Group("gwa")
	{
		apiHtml.GET("", htmlRenderApi.Main)
		//仪表盘
		apiHtml.GET("/main", htmlRenderApi.Main)
		//角色管理
		apiHtml.GET("/authorities", htmlRenderApi.Role)
		//菜单管理
		apiHtml.GET("/menu", htmlRenderApi.Menu)
		//API管理
		apiHtml.GET("/api", htmlRenderApi.Api)
		//系统
		//apiHtml.GET("/system", htmlRenderApi.System)
	}
	apiPubRouter := PubRouter.Group("api/gwa")
	{
		//管理员登陆
		apiPubRouter.POST("/login", apiRouterApi.Login)
	}
	apiRouter := Router.Group("api/gwa")
	{
		apiRouter.GET("/test", apiRouterApi.Home)
		apiRouter.GET("/menu", apiRouterApi.GetMenu)
		apiRouter.GET("/menuList", apiRouterApi.GetMenuList)
		apiRouter.GET("/authoritiesList", apiRouterApi.GetAuthoritiesList)
		apiRouter.GET("apiList", apiRouterApi.GetApiList)
	}

}

package system

import (
	"github.com/gin-gonic/gin"
	v1 "gwa_server/api/v1"
	"gwa_server/render"
)

type AdminRouter struct{}

func (s *AdminRouter) InitLoginRouter(Router *gin.RouterGroup) {
	htmlRenderApi := render.RenderGroupHtml.AdminRenderGroup.SysAdminRender
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SysAdminApi
	apiHtml := Router.Group("gwa")
	{
		apiHtml.GET("/login", htmlRenderApi.Login)
	}
	apiRouter := Router.Group("api/gwa")
	{
		apiRouter.POST("/login", apiRouterApi.Login)
	}
}

package system

import (
	"github.com/gin-gonic/gin"
	v1 "gwa_server/api/v1"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.USysTestApi
	{
		apiRouter.GET("", apiRouterApi.GetTest)
	}
}

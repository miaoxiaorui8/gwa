package middleware

import (
	"gwa_server/global"
	res "gwa_server/models/response"
	"gwa_server/service"
	"gwa_server/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GWA_Config.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			//获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, "")
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // 判断策略中是否存在
			//sub权限ID obj域名path如/gwa act
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				res.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

package middleware

import (
	"gwa_server/global"
	res "gwa_server/models/response"
	"gwa_server/service"
	"gwa_server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		reqtoken := c.Request.Header.Get("Cookie")
		if reqtoken == "" {
			c.HTML(200, "admin/login/login.html", gin.H{})
			c.Abort()
			return
		}
		token := strings.Split(reqtoken, "=")[1]

		//if jwtService.IsBlacklist(token) {
		//	res.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		//	c.Abort()
		//	return
		//}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				res.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			res.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	res.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GWA_Config.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			//if global.GWA_CONFIG.System.UseMultipoint {
			//	RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
			//	if err != nil {
			//		global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
			//	} else { // 当之前的取成功时才进行拉黑操作
			//		_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
			//	}
			//	// 无论如何都要记录当前的活跃状态
			//	_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			//}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

package admin

import "github.com/gin-gonic/gin"

type SysAdminRender struct {
}

func (sysAdminRender *SysAdminRender) Login(c *gin.Context) {
	c.HTML(200, "admin/login/login.html", gin.H{})
	return
}

package admin

import (
	"github.com/gin-gonic/gin"
)

type SysAdminRender struct {
}

// 登陆
func (sysAdminRender *SysAdminRender) Login(c *gin.Context) {
	c.HTML(200, "admin/login/login.html", gin.H{})
	return
}

// Home
func (sysAdminRender *SysAdminRender) Home(c *gin.Context) {
	c.HTML(200, "admin/home/home.html", gin.H{})
	return
}

// 仪表盘
func (sysAdminRender *SysAdminRender) Main(c *gin.Context) {
	c.HTML(200, "admin/main/main.html", gin.H{})
	return
}

// 角色管理
func (sysAdminRender *SysAdminRender) Role(c *gin.Context) {
	c.HTML(200, "admin/authorities/authorities.html", gin.H{})
	return
}

// 菜单管理
func (sysAdminRender *SysAdminRender) Menu(c *gin.Context) {
	c.HTML(200, "admin/menu/menu.html", gin.H{})
	return
}

// API管理
func (sysAdminRender *SysAdminRender) Api(c *gin.Context) {
	c.HTML(200, "admin/api/api.html", gin.H{})
	return
}

// 系统详情
func (sysAdminRender *SysAdminRender) System(c *gin.Context) {
	c.HTML(200, "admin/system/system.html", gin.H{})
	return
}

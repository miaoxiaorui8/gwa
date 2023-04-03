package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gwa_server/global"
	res "gwa_server/models/response"
	"gwa_server/models/system"
	"gwa_server/models/system/request"
	sysRes "gwa_server/models/system/response"
	"gwa_server/service"
	"gwa_server/utils"
)

type SysAdminApi struct {
}

// Login 管理员登陆
// @Tags 管理员管理
// @Summary 管理员登陆
// @Param data body request.Login	true	"用户名,密码"
// @Router /api/gwa/login [post]
// @Success 200 {object} res.Response{data=request.Login,msg=string}  "返回包括用户信息,token,过期时间"
func (sysLogin *SysAdminApi) Login(c *gin.Context) {
	var loginReq request.Login
	err := c.ShouldBind(&loginReq)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var admin system.SysAdmin
	err = global.GWA_DB.Where("username = ? AND password = ?", loginReq.Username, loginReq.Password).First(&admin).Error
	if err != nil {
		global.GWA_Log.Warn(err)
		res.FailWithMessage("账号或密码错误!", c)
		return
	}

	j := &utils.JWT{SigningKey: []byte(global.GWA_Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        admin.UUID,
		ID:          admin.ID,
		NickName:    admin.NickName,
		Username:    loginReq.Username,
		AuthorityId: admin.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GWA_Log.Warn(err)
		res.FailWithMessage("Token生成失败!", c)
		return
	}
	res.OkWithDetailed(sysRes.LoginResponse{
		Admin:     admin,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}

// Home 后台首页
// @Tags 后台首页
// @Summary 后台首页
// @Param data body request.Login	true	"用户名"
// @Router /api/gwa/home [post]
// @Success 200 {object} res.Response{data=request.Login,msg=string}  "返回包括用户信息,token,过期时间"
func (sysLogin *SysAdminApi) Home(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	fmt.Println(waitUse.AuthorityId)
	res.Ok(c)
}

// GetMenu 获取menu列表
// @Tags 获取menu列表
// @Summary 获取menu列表
// @Param data body request.Login	true	"用户名"
// @Router /api/gwa/menu [post]
// @Success 200 {object} res.Response{data=request.Login,msg=string}  "返回包括用户信息,token,过期时间"
func (sysLogin *SysAdminApi) GetMenu(c *gin.Context) {
	menuService := service.ServiceGroupApp.SystemServiceGroup.MenuService
	menus, err := menuService.GetMenuTreeMap(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GWA_Log.Error("获取失败!", err)
		res.FailWithMessage("获取失败", c)
	}
	res.OkWithData(menus, c)
}

// GetMenuList 获取所有menu列表
// @Tags 获取menu列表
// @Summary 获取menu列表
// @Param data body
// @Router /api/gwa/menu [get]
// @Success 200 {object} res.Response{data=request.Login,msg=string}  "返回包括用户信息,token,过期时间"
func (sysLogin *SysAdminApi) GetMenuList(c *gin.Context) {
	menuService := service.ServiceGroupApp.SystemServiceGroup.MenuService
	menus, err := menuService.GetMenu(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GWA_Log.Error("获取失败!", err)
		res.FailWithMessage("获取失败", c)
	}
	res.OkWithData(menus, c)
}

// GetAuthoritiesList 获取所有menu列表
// @Tags 获取menu列表
// @Summary 获取menu列表
// @Param data body
// @Router /api/gwa/menu [get]
// @Success 200 {object} res.Response{data=request.Login,msg=string}  "返回包括用户信息,token,过期时间"
func (sysLogin *SysAdminApi) GetAuthoritiesList(c *gin.Context) {
	authoritiesService := service.ServiceGroupApp.SystemServiceGroup.AuthoritiesService
	authoritiesList, err := authoritiesService.GetAuthoritiesList()
	if err != nil {
		global.GWA_Log.Error("获取失败!", err)
		res.FailWithMessage("获取失败", c)
	}
	res.OkWithData(authoritiesList, c)
}

// GetApiList 获取所有API列表
// @Tags 获取menu列表
// @Summary 获取menu列表
// @Param data body
// @Router /api/gwa/menu [get]
// @Success 200 {object} res.Response{data=request.Login,msg=string}  "返回包括用户信息,token,过期时间"
func (sysLogin *SysAdminApi) GetApiList(c *gin.Context) {
	apiService := service.ServiceGroupApp.SystemServiceGroup.ApiService
	apiList, err := apiService.GetApiList()
	if err != nil {
		global.GWA_Log.Error("获取失败!", err)
		res.FailWithMessage("获取失败", c)
	}
	res.OkWithData(apiList, c)
}

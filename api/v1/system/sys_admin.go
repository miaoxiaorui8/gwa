package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	res "gwa_server/models/response"
)

type SysAdminApi struct {
}

type loginRequest struct {
	Username string `json:"username" form:"username"` // 用户名
	Password string `json:"password" form:"password"` // 密码
}

func (sysLogin *SysAdminApi) Login(c *gin.Context) {
	var loginReq loginRequest
	err := c.ShouldBind(&loginReq)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(loginReq)
	res.OkWithData(loginReq, c)
	return
}

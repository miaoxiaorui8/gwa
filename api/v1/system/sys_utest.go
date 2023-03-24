package system

import (
	"github.com/gin-gonic/gin"
	res "gwa_server/models/response"
)

type USysTestApi struct {
}

func (sysTestApi *USysTestApi) GetTest(c *gin.Context) {
	res.FailWithMessage("1", c)
	return
}

package system

import "github.com/gin-gonic/gin"

type USysTestApi struct {
}

func (sysTestApi *USysTestApi) GetTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "xxx",
	})
	return
}

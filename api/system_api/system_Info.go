package system_api

import "github.com/gin-gonic/gin"

func (SystemApi) SystemsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "xxx"})
}

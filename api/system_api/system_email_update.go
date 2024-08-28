package system_api

import (
	"github.com/gin-gonic/gin"
	"myServer/config"
	"myServer/global"
	"myServer/setup"
	"myServer/utils/encapsulation/resp"
)

func (SystemApi) SystemsEmailUpdateView(c *gin.Context) {
	var email config.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	global.Config.Email = email
	err = setup.SetYaml()
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg(err.Error(), c)
		return
	}
	resp.OkWithC(c)
}

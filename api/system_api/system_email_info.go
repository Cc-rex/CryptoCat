package system_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/utils/encapsulation/resp"
)

func (SystemApi) SystemsEmailInfoView(c *gin.Context) {
	resp.OkWithData(global.Config.Email, c)
}

package system_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/utils/encapsulation/resp"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SystemApi) SystemsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		resp.OkWithData(global.Config.SiteInfo, c)
	case "email":
		resp.OkWithData(global.Config.Email, c)
	case "qq":
		resp.OkWithData(global.Config.QQ, c)
	case "qiniu":
		resp.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		resp.OkWithData(global.Config.Jwt, c)
	default:
		resp.FailWithMsg("没有对应的配置信息", c)
	}
}

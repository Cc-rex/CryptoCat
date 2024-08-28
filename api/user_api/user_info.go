package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"myServer/global"
	"myServer/models"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// UserInfoView 用户信息
// @Tags 用户管理
// @Summary 用户信息
// @Description 用户信息
// @Router /api/user_info [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=models.UserModel}
func (UserApi) UserInfoView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var userInfo models.UserModel
	err := global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		return
	}
	resp.OkWithData(filter.Select("info", userInfo), c)

}

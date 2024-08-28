package user_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/service"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// UserLogoutView
// @Tags User Management
// @Summary User logout
// @Description Logs out a user by invalidating the JWT token.
// @Accept json
// @Produce json
// @Router /api/users/logout [post]
// @Success 200 {object} resp.Response{} "Successfully logged out"
// @Failure 400 {object} resp.Response{} "Invalid request or failed to logout"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	_token, _ := c.Get("token")
	token := _token.(string)

	err := service.MyService.UserService.UserLogout(claims, token)

	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("注销失败", c)
		return
	}

	resp.OkWithMsg("注销成功", c)

}

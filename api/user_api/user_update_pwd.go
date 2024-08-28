package user_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
	"myServer/utils/pwd"
)

// UserUpdatePassword 修改登录人的密码
// @Tags User Management
// @Summary Update user password
// @Description Updates the password for the currently authenticated user.
// @Accept json
// @Produce json
// @Param body body ctype.UpdatePasswordRequest true "Password update request body"
// @Security BearerAuth
// @Router /api/users/set_pwd [put]
// @Success 200 {object} resp.Response{} "Password updated successfully"
// @Failure 400 {object} resp.Response{} "Invalid request parameters"
// @Failure 401 {object} resp.Response{} "Unauthorized or incorrect old password"
// @Failure 500 {object} resp.Response{} "Failed to update password"
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr ctype.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		resp.FailWithMsg("密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.NewPwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("密码修改失败", c)
		return
	}
	resp.OkWithMsg("密码修改成功", c)
	return
}

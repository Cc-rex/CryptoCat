package user_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// UserUpdateStatusView 用户权限变更
// @Tags User Management
// @Summary Update user status and nickname
// @Description Updates the status and nickname for a specified user by user ID.
// @Accept json
// @Produce json
// @Param body body ctype.UserStatusRequest true "User status update request body"
// @Security BearerAuth
// @Router /api/users/set_status [put]
// @Success 200 {object} resp.Response{} "Status and nickname updated successfully"
// @Failure 400 {object} resp.Response{} "Invalid request parameters"
// @Failure 401 {object} resp.Response{} "Unauthorized or user does not exist"
// @Failure 500 {object} resp.Response{} "Failed to update user status"
func (UserApi) UserUpdateStatusView(c *gin.Context) {
	var cr ctype.UserStatusRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户id错误，用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"status":    cr.Status,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("修改权限失败", c)
		return
	}
	resp.OkWithMsg("修改权限成功", c)
}

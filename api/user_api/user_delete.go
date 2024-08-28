package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// UserDeleteView
// @Tags User Management
// @Summary Delete users by their IDs
// @Description Deletes specified users and their related data from the system. This includes removing users from message tables, comment tables, user collections, and published articles.
// @Accept json
// @Produce json
// @Param req body ctype.DeleteRequest true "Request body containing the list of user IDs to be deleted"
// @Router /api/users [delete]
// @Success 200 {object} resp.Response{} "Confirms that the specified users have been successfully deleted from the system."
// @Failure 400 {object} resp.Response{} "Invalid request body or user IDs"
// @Failure 404 {object} resp.Response{} "Specified users not found"
// @Failure 500 {object} resp.Response{} "Failed to delete users due to internal server error"
func (UserApi) UserDeleteView(c *gin.Context) {
	var cr ctype.DeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		resp.FailWithMsg("用户不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("删除用户失败", c)
		return
	}
	resp.OkWithMsg(fmt.Sprintf("共删除 %d 个用户", count), c)

}

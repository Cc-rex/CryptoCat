package user_api

import (
	"github.com/gin-gonic/gin"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/common"
	"myServer/utils/encapsulation/hide"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// UserListView
// @Tags User Management
// @Summary Retrieve a paginated list of users
// @Description Fetches a list of users with optional hiding of sensitive information based on user permissions. Admins can see all details, while non-admin users have sensitive data hidden.
// @Accept json
// @Produce json
// @Param page query ctype.PageInfo true "Pagination information including page number and page size"
// @Router /api/users [get]
// @Success 200 {object} resp.Response{} "Returns a list of users with sensitive information hidden based on permissions"
// @Failure 400 {object} resp.Response{} "Invalid request parameters"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员
	myClaim, _ := c.Get("claims")
	claims := myClaim.(*jwt.CustomClaims)

	var page ctype.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ListQuery(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if claims.Status != ctype.PermissionAdmin {
			// 非管理员
			user.UserName = ""
		}
		user.Tel = hide.TelHide(user.Tel)
		user.Email = hide.EmailHide(user.Email)
		// 脱敏
		users = append(users, user)
	}

	resp.OkWithList(users, count, c)

}

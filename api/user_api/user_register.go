package user_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/pwd"
)

// UserRegisterView
// @Tags User Management
// @Summary User registration
// @Description Registers a new user by providing user details and hashed password.
// @Accept json
// @Produce json
// @Param body body ctype.RegisterRequest true "User registration request body"
// @Router /api/users/register [post]
// @Success 200 {object} resp.Response{} "User registered successfully"
// @Failure 400 {object} resp.Response{} "Invalid request parameters"
// @Failure 500 {object} resp.Response{} "Failed to register user"
func (UserApi) UserRegisterView(c *gin.Context) {
	var cr ctype.RegisterRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}

	hashedPwd := pwd.HashPwd(cr.Password)
	newUser := models.UserModel{
		UserName: cr.UserName,
		Password: hashedPwd,
		Email:    cr.Email,
		NickName: cr.NickName,
	}

	if err := global.DB.Create(&newUser).Error; err != nil {
		global.Log.Error("Failed to register new user: ", err)
		resp.FailWithMsg("Failed to register", c)
		return
	}

	resp.OkWithMsg("User registered successfully", c)
}

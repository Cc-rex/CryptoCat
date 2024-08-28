package user_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
	"myServer/utils/key"
	"myServer/utils/pwd"
)

// UserLoginView
// @Tags User Management
// @Summary User login
// @Description Authenticates a user with their username or email and password, and generates a JWT token upon successful login.
// @Accept json
// @Produce json
// @Param loginRequest body ctype.LoginRequest true "Login request payload containing username or email and password"
// @Router /api/users/login [post]
// @Success 200 {object} resp.Response{} "Returns a JWT token upon successful authentication"
// @Failure 400 {object} resp.Response{} "Invalid request parameters or authentication failed"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (UserApi) UserLoginView(c *gin.Context) {
	var cr ctype.LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		global.Log.Warn("The username does not exist!")
		resp.FailWithMsg("The User does not exist!", c)
		return
	}
	//check the pwd
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("The username or password is wrong!")
		resp.FailWithMsg("The username or password is wrong!", c)
		return
	}
	//generate token
	privateKey, err := key.LoadPrivateKey()
	if err != nil {
		global.Log.Info(err)
	}

	token, err := jwt.GenerateToken(jwt.MyPayLoad{
		UserID:   userModel.ID,
		Status:   userModel.Status,
		Username: userModel.UserName,
		NickName: userModel.NickName,
	}, privateKey)
	//fmt.Println("login token=", token)
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("Failed to generate token!", c)
		return
	}
	resp.OkWithData(token, c)
}

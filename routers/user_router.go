package routers

import (
	"myServer/api"
	"myServer/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.MyApiGroup.UserApi
	router.POST("users/login", userApi.UserLoginView)
	router.POST("users", userApi.UserRegisterView)
	router.POST("users/logout", middleware.JwtAuth(), userApi.UserLogoutView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.GET("user_info", middleware.JwtAuth(), userApi.UserInfoView)
	router.PUT("users/set_status", middleware.JwtAdmin(), userApi.UserUpdateStatusView)
	router.PUT("users/set_pwd", middleware.JwtAuth(), userApi.UserUpdatePassword)
	router.DELETE("users", middleware.JwtAdmin(), userApi.UserDeleteView)

}

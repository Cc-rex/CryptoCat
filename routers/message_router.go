package routers

import (
	"myServer/api"
	"myServer/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.MyApiGroup.MessageApi
	router.POST("messages", middleware.JwtAuth(), messageApi.MessageCreateView)
	router.GET("messages_all", messageApi.MessageListAllView)
	router.GET("messages/list", middleware.JwtAuth(), messageApi.MessageListView)
	router.GET("messages/history", middleware.JwtAuth(), messageApi.MessageHistoryView)

}

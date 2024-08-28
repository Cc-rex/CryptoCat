package routers

import "myServer/api"

func (router RouterGroup) ChatRouter() {
	chatApi := api.MyApiGroup.ChatApi
	router.GET("chat_groups", chatApi.ChatGroupView)
	router.GET("chat_groups/records", chatApi.ChatListView)

}

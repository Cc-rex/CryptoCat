package api

import (
	"myServer/api/article_api"
	"myServer/api/chat_api"
	"myServer/api/comment_api"
	"myServer/api/images_api"
	"myServer/api/menu_api"
	"myServer/api/message_api"
	"myServer/api/system_api"
	"myServer/api/tag_api"
	"myServer/api/user_api"
)

type ApiGroup struct {
	SystemsApi system_api.SystemApi
	ImagesApi  images_api.ImagesApi
	MenuApi    menu_api.MenuApi
	TagApi     tag_api.TagApi
	UserApi    user_api.UserApi
	MessageApi message_api.MessageApi
	ArticleApi article_api.ArticleApi
	CommentApi comment_api.CommentApi
	ChatApi    chat_api.ChatApi
}

var MyApiGroup = new(ApiGroup)

package routers

import (
	"myServer/api"
	"myServer/middleware"
)

func (router RouterGroup) CommentRouter() {
	commentApi := api.MyApiGroup.CommentApi
	router.POST("comments", middleware.JwtAuth(), commentApi.CommentCreateView)
	router.GET("comments", commentApi.CommentListView)
	router.GET("comments/:id", commentApi.CommentLike)
	router.DELETE("comments/:id", commentApi.CommentDeleteView)

}

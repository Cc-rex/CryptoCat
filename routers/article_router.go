package routers

import (
	"myServer/api"
	"myServer/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.MyApiGroup.ArticleApi
	router.GET("articles", articleApi.ArticleListView)
	router.GET("articles/detail", articleApi.ArticleDetailViewByTitle)
	router.GET("articles/:id", articleApi.ArticleDetailViewByID)
	router.GET("articles/calendar", articleApi.ArticleCalendarView)
	router.GET("articles/tags", articleApi.ArticleTagListView)
	router.GET("articles/collect", middleware.JwtAuth(), articleApi.ArticleCollListView)
	router.POST("articles", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.POST("articles/like", articleApi.ArticleLikeView)
	router.POST("articles/collect", middleware.JwtAuth(), articleApi.ArticleCollectCreateView)
	router.PUT("articles", articleApi.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAuth(), articleApi.ArticleDeleteView)
	router.DELETE("articles/collect", middleware.JwtAuth(), articleApi.ArticleCollBatchDeleteView)

}

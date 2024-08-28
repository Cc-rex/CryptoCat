package routers

import "myServer/api"

func (router RouterGroup) TagsRouter() {
	tagApi := api.MyApiGroup.TagApi
	router.POST("tags", tagApi.TagCreateView)
	router.GET("tags", tagApi.TagListView)
	router.PUT("tags/:id", tagApi.TagUpdateView)
	router.DELETE("tags", tagApi.TagDeleteView)

}

package routers

import "myServer/api"

func (router RouterGroup) ImagesRouter() {
	imageApi := api.MyApiGroup.ImagesApi
	router.GET("images", imageApi.ImageListView)
	router.GET("image_names", imageApi.ImageNameList)
	router.POST("images", imageApi.ImageUploadView)
	router.DELETE("images", imageApi.ImageDeleteView)
	router.PUT("images", imageApi.ImageUpdateView)
}

package routers

import "myServer/api"

func (router RouterGroup) MenuRouter() {
	menuApi := api.MyApiGroup.MenuApi
	router.POST("menus", menuApi.MenuCreateView)
	router.GET("menus", menuApi.MenuImageListView)
	router.GET("menu_names", menuApi.MenuNameListView)
	router.PUT("menus/:id", menuApi.MenuUpdateView)
	router.DELETE("menus", menuApi.MenuDeleteView)
	router.GET("menus/:id", menuApi.MenuDetailView)
}

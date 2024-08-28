package routers

import (
	"myServer/api"
)

func (router RouterGroup) SystemRouter() {
	systemsApi := api.MyApiGroup.SystemsApi
	router.GET("/settings/:name", systemsApi.SystemsInfoView)
	router.PUT("/settings/:name", systemsApi.SystemUpdateView)
}

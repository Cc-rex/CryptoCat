package routers

import (
	"myServer/api"
)

func (router RouterGroup) SystemRouter() {
	systemsApi := api.MyApigroup.SystemsApi
	router.GET("", systemsApi.SystemsInfoView)
}

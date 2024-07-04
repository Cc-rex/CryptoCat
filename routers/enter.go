package routers

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	myRouterGroup := RouterGroup{apiRouterGroup}
	myRouterGroup.SystemRouter()
	return router
}

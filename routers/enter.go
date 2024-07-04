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
	apiGroup := router.Group("api")
	apiRouterGroup := RouterGroup{apiGroup}
	apiRouterGroup.SystemRouter()
	return router
}

package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"myServer/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("api")
	apiRouterGroup := RouterGroup{apiGroup}
	apiRouterGroup.SystemRouter()
	apiRouterGroup.ImagesRouter()
	apiRouterGroup.MenuRouter()
	apiRouterGroup.TagsRouter()
	apiRouterGroup.UserRouter()
	apiRouterGroup.MessageRouter()
	apiRouterGroup.ArticleRouter()
	apiRouterGroup.CommentRouter()
	apiRouterGroup.ChatRouter()

	return router
}

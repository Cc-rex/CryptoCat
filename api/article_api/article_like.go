package article_api

import (
	"github.com/gin-gonic/gin"
	"myServer/models/ctype"
	"myServer/service/redis_service"
	"myServer/utils/encapsulation/resp"
)

// ArticleLikeView
// @Tags Article Management
// @Summary Like an article
// @Description Registers a 'like' for an article by its ID, updating the like count in Redis.
// @Accept json
// @Produce json
// @Param req body ctype.ESIDRequest true "Request body containing the ID of the article to be liked"
// @Router /api/articles/like [post]
// @Success 200 {object} resp.Response{} "Confirms that the like has been registered successfully."
// @Failure 400 {object} resp.Response{} "Invalid request body or argument errors"
// @Failure 500 {object} resp.Response{} "Failed to process the like due to internal server error"
func (ArticleApi) ArticleLikeView(c *gin.Context) {
	var cr ctype.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	// 对长度校验
	// 查es
	like := redis_service.NewLike()
	like.Set(cr.ID)
	resp.OkWithMsg("文章点赞成功", c)
}

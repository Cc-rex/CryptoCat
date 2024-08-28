package article_api

import (
	"github.com/gin-gonic/gin"
	"myServer/models/ctype"
	"myServer/service/es_service"
	"myServer/service/redis_service"
	"myServer/utils/encapsulation/resp"
)

// ArticleDetailViewByID
// @Tags Article Management
// @Summary Retrieve article details by ID
// @Description Gets detailed information for a specified article ID, including incrementing the view count via Redis.
// @Accept json
// @Produce json
// @Param id path string true "The ID of the article to retrieve"
// @Router /api/articles/{id} [get]
// @Success 200 {object} models.ArticleModel "Detailed information about the article, including content and metadata"
// @Failure 400 {object} resp.Response{} "Invalid article ID specified"
// @Failure 404 {object} resp.Response{} "Article not found"
// @Failure 500 {object} resp.Response{} "Failed to retrieve article details due to internal server error"
func (ArticleApi) ArticleDetailViewByID(c *gin.Context) {
	var cr ctype.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	look := redis_service.NewArticleLook()
	look.Set(cr.ID)

	model, err := es_service.DetailQueryByID(cr.ID)
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	resp.OkWithData(model, c)
}

// ArticleDetailViewByTitle
// @Tags Article Management
// @Summary Retrieve article details by title
// @Description Gets detailed information for a specified article title, including incrementing the view count via Redis.
// @Accept json
// @Produce json
// @Param id path string true "The ID of the article to retrieve"
// @Router /api/articles/detail [get]
// @Success 200 {object} ctype.ArticleDetailRequest "Detailed information about the article, including content and metadata"
// @Failure 400 {object} resp.Response{} "Invalid article ID specified"
// @Failure 404 {object} resp.Response{} "Article not found"
// @Failure 500 {object} resp.Response{} "Failed to retrieve article details due to internal server error"
func (ArticleApi) ArticleDetailViewByTitle(c *gin.Context) {
	var cr ctype.ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	model, err := es_service.DetailQueryByKeyword(cr.Title)
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	resp.OkWithData(model, c)
}

package article_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/es_service"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// ArticleCollectCreateView
// @Tags Article Management
// @Summary Create article collection
// @Description Allows a user to either collect or uncollect an article based on its current collection status. If the article is not already collected, it will be collected; if it is already collected, the collection will be cancelled.
// @Accept json
// @Produce json
// @Param req body ctype.ESIDRequest true "Request body containing the ID of the article to be collected or uncollected"
// @Router /api/articles/collect [post]
// @Success 200 {object} resp.Response{} "Indicates successful collection of the article"
// @Success 200 {object} resp.Response{} "Indicates successful uncollection of the article"
// @Failure 400 {object} resp.Response{} "Invalid request body or argument errors"
// @Failure 404 {object} resp.Response{} "Article not found or no longer available"
// @Failure 500 {object} resp.Response{} "Failed to update the collection status due to internal server error"
func (ArticleApi) ArticleCollectCreateView(c *gin.Context) {
	var cr ctype.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	model, err := es_service.DetailQueryByID(cr.ID)
	if err != nil {
		resp.FailWithMsg("文章不存在", c)
		return
	}

	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	}
	// 取消收藏
	// 文章数 -1
	global.DB.Delete(&coll)

	// 更新文章收藏数
	err = es_service.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		resp.OkWithMsg("收藏文章成功", c)
	} else {
		resp.OkWithMsg("取消收藏成功", c)
	}
}

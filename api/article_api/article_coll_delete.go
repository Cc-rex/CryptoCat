package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/es_service"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// ArticleCollBatchDeleteView
// @Tags Article Management
// @Summary Batch delete article collections
// @Description Deletes a batch of article collections based on the list of article IDs provided, and updates the collect count for each article.
// @Accept json
// @Produce json
// @Param req body ctype.ESIDListRequest true "Request body containing a list of article IDs to be uncollected"
// @Router /api/articles/collect [delete]
// @Success 200 {object} resp.Response{} "Indicates the number of articles successfully uncollected."
// @Failure 400 {object} resp.Response{} "Invalid request body or argument errors"
// @Failure 403 {object} resp.Response{} "Unauthorized request or action not allowed"
// @Failure 500 {object} resp.Response{} "Failed to delete collections due to internal server error"
func (ArticleApi) ArticleCollBatchDeleteView(c *gin.Context) {
	var cr ctype.ESIDListRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var collects []models.UserCollectModel
	var articleIDList []string
	global.DB.Find(&collects, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).
		Select("article_id").
		Scan(&articleIDList)
	if len(articleIDList) == 0 {
		resp.FailWithMsg("请求非法", c)
		return
	}
	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}
	// 更新文章数
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		count := article.CollectsCount - 1
		err = es_service.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	global.DB.Delete(&collects)
	resp.OkWithMsg(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), c)

}

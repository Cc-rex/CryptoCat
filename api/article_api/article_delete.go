package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// ArticleDeleteView
// @Tags Article Management
// @Summary Delete a new article
// @Description Delete the article from the ES based on the article ID
// @Accept json
// @Produce json
// @Param req body ctype.DeleteIDList true "Request body for deleting a new article including title, content, and optional banner ID"
// @Router /api/articles [delete]
// @Success 200 {object} resp.Response{} "Confirms that the article has been successfully added to the database."
// @Failure 400 {object} resp.Response{} "Invalid request body or failure in content sanitization"
// @Failure 404 {object} resp.Response{} "Banner or user not found"
// @Failure 409 {object} resp.Response{} "Duplicate article entries found"
// @Failure 500 {object} resp.Response{} "Failed to create article due to internal server error"
func (ArticleApi) ArticleDeleteView(c *gin.Context) {
	var cr ctype.DeleteIDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("删除失败", c)
		return
	}
	resp.OkWithMsg(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
	return
}

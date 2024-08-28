package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/common"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

// ArticleCollListView
// @Tags Article Management
// @Summary Show the list of article collections
// @Description Retrieves a list of articles collected by the user, including the article details and the date each was collected.
// @Accept json
// @Produce json
// @Param page query int false "Page number of the collection list"
// @Param size query int false "Number of collections per page"
// @Router /api/articles/collect [get]
// @Success 200 {array} CollResponse "List of collected articles with pagination details"
// @Failure 400 {object} resp.Response{} "Invalid query parameters or binding errors"
// @Failure 403 {object} resp.Response{} "Unauthorized access or invalid credentials"
// @Failure 500 {object} resp.Response{} "Failed to retrieve collections due to internal server error"
func (ArticleApi) ArticleCollListView(c *gin.Context) {

	var cr ctype.PageInfo

	c.ShouldBindQuery(&cr)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var articleIDList []interface{}

	list, count, err := common.ListQuery(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})

	var collTimeMap = map[string]string{}

	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collTimeMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)

	var collList = make([]CollResponse, 0)

	// 传id列表，查es
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	//fmt.Println(result.Hits.TotalHits.Value, articleIDList)

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		collList = append(collList, CollResponse{
			ArticleModel: article,
			CreatedAt:    collTimeMap[hit.Id],
		})
	}
	resp.OkWithList(collList, count, c)
}

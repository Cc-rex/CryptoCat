package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// ArticleTagListView
// @Tags Article Management
// @Summary List article tags with article counts
// @Description Retrieves a list of article tags, each with a count of articles associated and a list of keywords from those articles. Pagination is applied to the tags.
// @Accept json
// @Produce json
// @Param page query int false "Page number for tag listing"
// @Param size query int false "Number of tags to return per page"
// @Router /api/articles/tags [get]
// @Success 200 {array} ctype.TagsResponse "Paginated list of tags with associated article counts and article keywords"
// @Failure 400 {object} resp.Response{} "Invalid query parameters or binding errors"
// @Failure 500 {object} resp.Response{} "Failed to retrieve tags due to internal server error"
func (ArticleApi) ArticleTagListView(c *gin.Context) {

	var cr ctype.PageInfo
	_ = c.ShouldBindQuery(&cr)

	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		Size(0).
		Do(context.Background())

	cTag, _ := result.Aggregations.Cardinality("tags")

	count := int64(*cTag.Value)

	agg := elastic.NewTermsAggregation().Field("tags")

	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

	query := elastic.NewBoolQuery()

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg(err.Error(), c)
		return
	}
	var tagType ctype.TagsType
	var tagList = make([]*ctype.TagsResponse, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	var tagStringList []string
	for _, bucket := range tagType.Buckets {

		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, &ctype.TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
		tagStringList = append(tagStringList, bucket.Key)
	}

	var tagModelList []models.TagModel
	global.DB.Find(&tagModelList, "title in ?", tagStringList)
	var tagDate = map[string]string{}
	for _, model := range tagModelList {
		tagDate[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	for _, response := range tagList {
		response.CreatedAt = tagDate[response.Tag]
	}

	resp.OkWithList(tagList, count, c)
}

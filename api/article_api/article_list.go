package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/es_service"
	"myServer/utils/encapsulation/resp"
)

// ArticleListView
// @Tags Article Management
// @Summary List articles with pagination and filters
// @Description Retrieves a paginated list of articles based on search criteria such as tags and fields like title and content. The response filters out empty values.
// @Accept json
// @Produce json
// @Param page query int false "Page number of the results"
// @Param size query int false "Number of results per page"
// @Param tag query string false "Tag to filter articles by"
// @Router /api/articles [get]
// @Success 200 {array} models.ArticleModel "A list of articles with pagination details"
// @Failure 400 {object} resp.Response{} "Invalid query parameters or binding errors"
// @Failure 500 {object} resp.Response{} "Failed to retrieve articles due to internal server error"
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ctype.ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	list, count, err := es_service.ListQuery(es_service.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		resp.OkWithMsg("查询失败", c)
		return
	}

	//json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		resp.OkWithList(list, int64(count), c)
		return
	}
	resp.OkWithList(data, int64(count), c)
}

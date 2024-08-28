package tag_api

import (
	"github.com/gin-gonic/gin"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/common"
	"myServer/utils/encapsulation/resp"
)

// TagListView
// @Tags Tag Management
// @Summary List tags
// @Description Retrieves a paginated list of tags based on given pagination parameters.
// @Accept json
// @Produce json
// @Param page query int false "Page number of the tag list"
// @Param pageSize query int false "Number of tags per page"
// @Router /api/tags [get]
// @Success 200 {object} resp.Response{} "Successfully retrieved list of tags with pagination."
// @Failure 400 {object} resp.Response{} "Invalid query parameters."
func (TagApi) TagListView(c *gin.Context) {
	var cr ctype.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	list, count, _ := common.ListQuery(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	resp.OkWithList(list, count, c)
}

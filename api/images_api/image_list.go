package images_api

import (
	"github.com/gin-gonic/gin"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/common"
	"myServer/utils/encapsulation/resp"
)

// ImageListView
// @Tags Image Management
// @Summary Image List
// @Description Used for users to query images
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Router /api/images [get]
// @Success 200 {object} resp.Response{data=resp.ListResponse[models.BannerModel]} "Successful operation"
// @Failure 400 {object} resp.Response{} "Invalid request"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr ctype.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	list, count, err := common.ListQuery(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})

	resp.OkWithList(list, count, c)

	return
}

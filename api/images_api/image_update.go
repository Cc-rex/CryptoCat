package images_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// ImageUpdateView
// @Tags Image Management
// @Summary Image Update
// @Description Used for users to update image information
// @Accept json
// @Produce json
// @Param data body ctype.ImageUpdateRequest true "Image update parameters"
// @Router /api/images [put]
// @Success 200 {object} resp.Response{} "Image update successful"
// @Failure 400 {object} resp.Response{} "Invalid request"
// @Failure 404 {object} resp.Response{} "Image not found"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ctype.ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		resp.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	resp.OkWithMsg("图片名称修改成功", c)
	return

}

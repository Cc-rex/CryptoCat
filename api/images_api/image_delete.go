package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// ImageDeleteView
// @Tags Image Management
// @Summary Images Delete
// @Description Used for users to delete images
// @Accept json
// @Produce json
// @Param data body ctype.DeleteRequest true "IDs of images to delete"
// @Router /api/images [delete]
// @Success 200 {object} resp.Response{} "Images deleted successfully"
// @Failure 400 {object} resp.Response{} "Invalid request"
// @Failure 404 {object} resp.Response{} "Images not found"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (ImagesApi) ImageDeleteView(c *gin.Context) {
	var cr ctype.DeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		resp.FailWithMsg("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	resp.OkWithMsg(fmt.Sprintf("共删除 %d 张图片", count), c)
}

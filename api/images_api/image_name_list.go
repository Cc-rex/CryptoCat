package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// ImageNameList
// @Tags Image Management
// @Summary Image Name List
// @Description Return the name list of images
// @Produce json
// @Router /api/images_name [get]
// @Success 200 {object} resp.Response{data=[]ctype.ImageNameResponse} "Successful operation"
func (ImagesApi) ImageNameList(c *gin.Context) {
	var imageList []ctype.ImageNameResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	fmt.Println(imageList)
	resp.OkWithData(imageList, c)
	return
}

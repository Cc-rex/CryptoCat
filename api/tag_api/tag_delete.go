package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// TagDeleteView
// @Tags Tag Management
// @Summary Delete tags
// @Description Deletes tags based on a list of tag IDs provided in the request.
// @Accept json
// @Produce json
// @Param tag body ctype.DeleteRequest true "Details of the tags to be deleted by ID"
// @Router /api/tags [delete]
// @Success 200 {object} resp.Response{} "Successfully deleted tags with a count of how many were deleted."
// @Failure 400 {object} resp.Response{} "Invalid input."
// @Failure 404 {object} resp.Response{} "Tag does not exist."
func (TagApi) TagDeleteView(c *gin.Context) {
	var cr ctype.DeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		resp.FailWithMsg("Tag does not exist", c)
		return
	}

	global.DB.Delete(&tagList)
	resp.OkWithMsg(fmt.Sprintf("Deleted %d Tags", count), c)
}

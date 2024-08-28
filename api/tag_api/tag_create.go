package tag_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// TagCreateView
// @Tags Tag Management
// @Summary Create a new tag
// @Description Creates a new tag with a unique title.
// @Accept json
// @Produce json
// @Param tag body ctype.TagRequest true "Details of the tag to be created"
// @Router /api/tags [post]
// @Success 200 {object} resp.Response{} "Successfully created tag."
// @Failure 400 {object} resp.Response{} "Invalid input."
// @Failure 409 {object} resp.Response{} "The tag already exists."
// @Failure 500 {object} resp.Response{} "Failed to create tag."
func (TagApi) TagCreateView(c *gin.Context) {
	var cr ctype.TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		resp.FailWithMsg("The tag already exists", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("Failed to create tag", c)
		return
	}

	resp.OkWithMsg("Successfully created tag", c)
}

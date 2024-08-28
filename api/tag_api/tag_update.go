package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// TagUpdateView updates the details of a specific tag.
// @Tags Tag Management
// @Summary Update tag details
// @Description Updates the details of an existing tag based on the tag ID and provided new tag details.
// @Accept json
// @Produce json
// @Param id path string true "The ID of the tag to update"
// @Param tag body ctype.TagRequest true "New details for the tag"
// @Router /api/tags/{id} [put]
// @Success 200 {object} resp.Response{} "msg: Modify the tag successfully."
// @Failure 400 {object} resp.Response{} "Invalid input."
// @Failure 404 {object} resp.Response{} "The tag does not exist."
// @Failure 500 {object} resp.Response{} "Failed to modify tag."
func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr ctype.TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}

	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		resp.FailWithMsg("The tag dose not exist", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("Failed to modify tag", c)
		return
	}

	resp.OkWithMsg("Modify the tag successfully", c)

}

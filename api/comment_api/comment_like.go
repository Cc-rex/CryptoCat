package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/redis_service"
	"myServer/utils/encapsulation/resp"
)

// CommentLike
// @Tags Comment Management
// @Summary Like a comment
// @Description Registers a 'like' for a specific comment using its ID, updating the like count in Redis.
// @Accept json
// @Produce json
// @Param id path string true "The ID of the comment to be liked"
// @Router /api/comments/{id} [get]
// @Success 200 {object} resp.Response{} "Confirms that the like has been successfully registered for the comment."
// @Failure 400 {object} resp.Response{} "Invalid comment ID or URI binding errors"
// @Failure 404 {object} resp.Response{} "Comment not found"
// @Failure 500 {object} resp.Response{} "Failed to like the comment due to internal server error"
func (CommentApi) CommentLike(c *gin.Context) {
	var cr ctype.CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		resp.FailWithMsg("评论不存在", c)
		return
	}

	redis_service.NewCommentLike().Set(fmt.Sprintf("%d", cr.ID))

	resp.OkWithMsg("评论点赞成功", c)
	return

}

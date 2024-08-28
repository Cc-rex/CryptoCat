package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/redis_service"
	"myServer/utils"
	"myServer/utils/encapsulation/resp"
)

// CommentDeleteView
// @Tags Comment Management
// @Summary Delete a comment
// @Description Deletes a comment and all its sub-comments from an article. Adjusts comment counts accordingly in the database and cache.
// @Accept json
// @Produce json
// @Param id path string true "The ID of the comment to be deleted"
// @Router /api/comments/{id} [delete]
// @Success 200 {object} resp.Response{} "Confirms that the comment and any sub-comments have been successfully deleted."
// @Failure 400 {object} resp.Response{} "Invalid comment ID or URI binding errors"
// @Failure 404 {object} resp.Response{} "Comment not found"
// @Failure 500 {object} resp.Response{} "Failed to delete comments due to internal server error"
func (CommentApi) CommentDeleteView(c *gin.Context) {
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
	// 统计评论下的子评论数 再把自己算上去
	subCommentList := FindSubCommentCount(commentModel)
	count := len(subCommentList) + 1
	redis_service.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	// 判断是否是子评论
	if commentModel.ParentCommentID != nil {
		// 子评论
		// 找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// 删除子评论以及当前评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	// 反转，然后一个一个删
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	resp.OkWithMsg(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
	return
}

func FindSubCommentCount(model models.CommentModel) (subCommentList []models.CommentModel) {
	findSubCommentList(model, &subCommentList)
	return subCommentList
}

// findSubCommentList 递归查评论下的子评论
func findSubCommentList(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return
}

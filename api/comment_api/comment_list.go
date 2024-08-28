package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/redis_service"
	"myServer/utils/encapsulation/resp"
)

// CommentListView
// @Tags Comment Management
// @Summary List comments for an article
// @Description Retrieves all root-level comments for a specified article ID, with an option to filter and sort based on query parameters.
// @Accept json
// @Produce json
// @Param article_id query string true "The ID of the article for which comments are being retrieved"
// @Router /api/comments [get]
// @Success 200 {array} models.CommentModel "List of root comments for the specified article"
// @Failure 400 {object} resp.Response{} "Invalid query parameters or binding errors"
// @Failure 500 {object} resp.Response{} "Failed to retrieve comments due to internal server error"
func (CommentApi) CommentListView(c *gin.Context) {
	var cr ctype.CommentListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	resp.OkWithData(filter.Select("c", rootCommentList), c)
	return

}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	likeInfo := redis_service.NewCommentLike().GetAllInfo()
	for _, model := range RootCommentList {
		var subCommentList, newSubCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		for _, commentModel := range subCommentList {
			like := likeInfo[fmt.Sprintf("%d", commentModel.ID)]
			commentModel.LikeCount = commentModel.LikeCount + like
			newSubCommentList = append(newSubCommentList, commentModel)
		}
		modelLike := likeInfo[fmt.Sprintf("%d", model.ID)]
		model.LikeCount = model.LikeCount + modelLike
		model.SubComments = newSubCommentList
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return
}

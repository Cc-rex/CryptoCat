package comment_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/es_service"
	"myServer/service/redis_service"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// CommentCreateView
// @Tags Comment Management
// @Summary Create a new comment
// @Description Adds a new comment to an article, with the ability to specify if the comment is a reply to another comment.
// @Accept json
// @Produce json
// @Param req body ctype.CommentRequest true "Request body for creating a new comment including article ID, content, and optional parent comment ID"
// @Router /api/comments [post]
// @Success 200 {object} resp.Response{} "Confirms that the comment has been successfully added."
// @Failure 400 {object} resp.Response{} "Invalid request body or article/comment ID mismatch"
// @Failure 404 {object} resp.Response{} "Article or parent comment not found"
// @Failure 500 {object} resp.Response{} "Failed to create comment due to internal server error"
func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr ctype.CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 文章是否存在
	_, err = es_service.DetailQueryByID(cr.ArticleID)
	if err != nil {
		resp.FailWithMsg("文章不存在", c)
		return
	}

	// 判断是否是子评论
	if cr.ParentCommentID != nil {
		// 子评论
		// 给父评论数 +1
		// 父评论id
		var parentComment models.CommentModel
		// 找父评论
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			resp.FailWithMsg("父评论不存在", c)
			return
		}
		// 判断父评论的文章是否和当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			resp.FailWithMsg("评论文章不一致", c)
			return
		}
		// 给父评论评论数+1
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	// 添加评论
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})
	// 拿到文章数，新的文章评论数存缓存里
	//newCommentCount := article.CommentCount + 1
	// 给文章评论数 +1
	redis_service.NewCommentCount().Set(cr.ArticleID)
	resp.OkWithMsg("文章评论成功", c)
}

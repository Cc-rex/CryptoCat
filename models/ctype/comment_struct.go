package ctype

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

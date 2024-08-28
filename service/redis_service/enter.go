package redis_service

const (
	articleLookPrefix         = "article_look"
	articleCommentCountPrefix = "article_comment"
	articleLikePrefix         = "article_like"
	commentLikePrefix         = "comment_like"
)

func NewLike() CountDB {
	return CountDB{
		Index: articleLikePrefix,
	}
}
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentCountPrefix,
	}
}
func NewCommentLike() CountDB {
	return CountDB{
		Index: commentLikePrefix,
	}
}

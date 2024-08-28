package main

import (
	"myServer/global"
	"myServer/models"
	"myServer/setup"
)

func main4() {
	setup.InitConf()
	global.Log = setup.InitLogger()
	global.DB = setup.InitGorm()
	FindArticleCommentList("")

}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	for _, model := range RootCommentList {
		var subCommentList []models.CommentModel
		RecursionSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
	}
	return
}

func RecursionSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		RecursionSubComment(sub, subCommentList)
	}
	return
}

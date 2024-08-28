package cron_service

import (
	"gorm.io/gorm"
	"myServer/global"
	"myServer/models"
	"myServer/service/redis_service"
)

func SyncCommentData() {
	commentLikeInfo := redis_service.NewCommentLike().GetAllInfo()
	for i, count := range commentLikeInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, i).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).
			Update("like_count", gorm.Expr("like_count + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 点赞数更新成功", comment.Content)
	}
	redis_service.NewCommentLike().Clear()
}

package cron_service

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"myServer/global"
	"myServer/models"
	"myServer/service/redis_service"
)

func SyncArticleData() {
	// 查询es中全部数据
	result, err := global.ESClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	// 拿到redis中缓存数据
	likeInfo := redis_service.NewLike().GetAllInfo()
	lookInfo := redis_service.NewArticleLook().GetAllInfo()
	commentInfo := redis_service.NewCommentCount().GetAllInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		like := likeInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		newLike := article.LikeCount + like
		newLook := article.LookCount + look
		newcomment := article.CommentCount + comment

		if like == 0 && look == 0 && comment == 0 {
			global.Log.Infof("%s数据无变化", article.Title)
			continue
		}

		_, err := global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"look_count":    newLook,
				"comment_count": newcomment,
				"like_count":    newLike,
			}).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s数据更新成功 点赞数：%d, 浏览量：%d, 评论数：%d", article.Title, newLike, newLook, newcomment)
	}
	// 清除redis中的数据
	redis_service.NewLike().Clear()
	redis_service.NewArticleLook().Clear()
	redis_service.NewCommentCount().Clear()
}

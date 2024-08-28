package redis

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"myServer/global"
	"myServer/models"
	"myServer/service/redis_service"
	"myServer/setup"
)

func RedisToES() {
	// 读取配置文件
	setup.InitConf()
	// 初始化日志
	global.Log = setup.InitLogger()

	global.Redis = setup.ConnectRedis()
	global.ESClient = setup.EsConnect()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	likeInfo := redis_service.NewLike().GetAllInfo()
	lookInfo := redis_service.NewArticleLook().GetAllInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)

		like := likeInfo[hit.Id]
		look := lookInfo[hit.Id]
		newLike := article.LikeCount + like
		newLook := article.LookCount + look
		if article.LikeCount == newLike && article.LookCount == newLook {
			logrus.Info(article.Title, "点赞数和浏览数无变化")
			continue
		}
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"like_count": newLike,
				"look_count": newLook,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Infof("%s, 点赞数据同步成功， 点赞数: %d 浏览数: %d", article.Title, newLike, newLook)
	}
	redis_service.NewLike().Clear()
	redis_service.NewArticleLook().Clear()
}

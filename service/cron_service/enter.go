package cron_service

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	c.AddFunc("0 0 0 * * *", SyncArticleData)
	c.AddFunc("0 0 0 * * *", SyncCommentData)
	c.Start()
}

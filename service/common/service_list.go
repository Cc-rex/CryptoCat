package common

import (
	"gorm.io/gorm"
	"myServer/global"
	"myServer/models/ctype"
)

type Option struct {
	ctype.PageInfo
	Debug bool
}

func ListQuery[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}

	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Limit == 0 {
		option.Limit = -1
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}

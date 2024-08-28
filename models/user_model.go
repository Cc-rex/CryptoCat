package models

import (
	"gorm.io/gorm"
	"myServer/models/ctype"
)

type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:42" json:"nick_name,select(c|info)"`
	UserName   string           `gorm:"size:36" json:"user_name,select(info)"`
	Password   string           `gorm:"size:64" json:"-"`
	Avatar     string           `gorm:"size:256" json:"avatar,select(c|info)"`
	Email      string           `gorm:"size:128" json:"email"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr,select(c)"`
	Token      string           `gorm:"size:512" json:"token"`
	IP         string           `gorm:"size:20" json:"ip,select(c)"`
	Status     ctype.Status     `gorm:"size:4;default:2" json:"status,select(info)"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
}

const defaultAvatar = "/uploadFile/avatar/default.png"

func (u *UserModel) BeforeSave(tx *gorm.DB) (err error) {
	if u.Avatar == "" {
		u.Avatar = defaultAvatar
	}
	return
}

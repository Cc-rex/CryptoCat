package ctype

import (
	"github.com/gorilla/websocket"
	"time"
)

type MsgType int

const (
	InRoomMsg  MsgType = 1 //进入聊天室
	TextMsg    MsgType = 2 //发送文本消息
	ImageMsg   MsgType = 3 //发送图片
	SystemMsg  MsgType = 4 //发送系统消息
	OutRoomMsg MsgType = 5 //退出聊天室
)

type GroupRequest struct {
	Content string  `json:"content"`  // 聊天的内容
	MsgType MsgType `json:"msg_type"` // 聊天类型
}
type GroupResponse struct {
	NickName        string    `json:"nick_name"`         // 前端自己生成
	Avatar          string    `json:"avatar"`            // 头像
	MsgType         MsgType   `json:"msg_type"`          // 聊天类型
	Content         string    `json:"content"`           // 聊天的内容
	OnlineUserCount int       `json:"online_user_count"` //在线人数
	Date            time.Time `json:"date"`              // 消息的时间
}

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

package ctype

import (
	"encoding/json"
)

type SignStatus int
type Status int

const (
	SignQQ    SignStatus = 1 // QQ
	SignGitee SignStatus = 2 // gitee
	SignEmail SignStatus = 3 // email
)

const (
	PermissionAdmin         Status = 1 // Admin
	PermissionUser          Status = 2 // User
	PermissionVisitor       Status = 3 // Visitor
	PermissionForbiddenUser Status = 4 // Forbidden User
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "gitee"
	case SignEmail:
		str = "email"
	default:
		str = "others"
	}
	return str
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(s))
}

func (s Status) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "Admin"
	case PermissionUser:
		str = "User"
	case PermissionVisitor:
		str = "Visitor"
	case PermissionForbiddenUser:
		str = "ForbiddenUser"
	default:
		str = "others"
	}
	return str
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"please enter the username"`
	Password string `json:"password" binding:"required" msg:"please enter the password"`
}

type RegisterRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"please enter the username"`
	Password string `json:"password" binding:"required" msg:"please enter the password"`
	Email    string `json:"email" binding:"required,email" msg:"please enter the email"`
	NickName string `json:"nick_name" binding:"required" msg:"please enter the nickname"`
}

type UserStatusRequest struct {
	Status   Status `json:"status" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint   `json:"user_id" binding:"required" msg:"用户id错误"`
}

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"` // 旧密码
	NewPwd string `json:"new_pwd"` // 新密码
}

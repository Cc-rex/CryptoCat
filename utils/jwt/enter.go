package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"myServer/models/ctype"
)

// JwtPayLoad jwt中payload数据

type MyPayLoad struct {
	Username string       `json:"username"`  // 用户名
	NickName string       `json:"nick_name"` // 昵称
	Status   ctype.Status `json:"status"`    // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint         `json:"user_id"`   // 用户id
}

type CustomClaims struct {
	MyPayLoad
	jwt.RegisteredClaims
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models/ctype"
	"myServer/service/redis_service"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
	"myServer/utils/key"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头部提取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			resp.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}

		// 检查Bearer token格式
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader { // TrimPrefix未作更改，说明格式不正确或者token不存在
			resp.FailWithMsg("token格式错误", c)
			c.Abort()
			return
		}
		publicKey, err := key.LoadPublicKey()
		if err != nil {
			global.Log.Info(err)
		}
		claims, err := jwt.ParseToken(token, publicKey)
		if err != nil {
			resp.FailWithMsg("token错误", c)
			c.Abort()
			return
		}

		//判断是否在redis中
		if redis_service.CheckLogout(token) {
			resp.FailWithMsg("token已失效", c)
			c.Abort()
			return
		}

		// 登录的用户
		c.Set("claims", claims)
		c.Set("token", token)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头部提取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			resp.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}

		// 检查Bearer token格式
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader { // TrimPrefix未作更改，说明格式不正确或者token不存在
			resp.FailWithMsg("token格式错误", c)
			c.Abort()
			return
		}
		publicKey, err := key.LoadPublicKey()
		if err != nil {
			global.Log.Info(err)
		}
		claims, err := jwt.ParseToken(token, publicKey)
		if err != nil {
			resp.FailWithMsg("token错误", c)
			c.Abort()
			return
		}

		//判断是否在redis中
		if redis_service.CheckLogout(token) {
			resp.FailWithMsg("token已失效", c)
			c.Abort()
			return
		}

		if claims.Status != ctype.PermissionAdmin {
			resp.FailWithMsg("权限错误", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
		c.Set("token", token)
	}
}


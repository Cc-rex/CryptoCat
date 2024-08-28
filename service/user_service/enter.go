package user_service

import (
	"myServer/service/redis_service"
	"myServer/utils/jwt"
	"time"
)

type UserService struct {
}

func (UserService) UserLogout(claims *jwt.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_service.Logout(token, diff)
}

package jwt

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"myServer/global"
	"time"
)

func GenerateToken(user MyPayLoad, rsaPrivateKey *rsa.PrivateKey) (string, error) {
	// 创建带有自定义声明的JWT
	claims := CustomClaims{
		MyPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), // Token 过期时间
			Issuer:    "CryptoCat",                                                                              // 发行者
		},
	}

	// 使用RS256算法创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// 使用私钥签名Token
	signedToken, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

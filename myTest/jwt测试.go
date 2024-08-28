package main

import (
	"fmt"
	"myServer/global"
	"myServer/setup"
	"myServer/utils/jwt"
	"myServer/utils/key"
)

func main1() {
	setup.InitConf()
	global.Log = setup.InitLogger()

	privateKey, err := key.LoadPrivateKey()
	if err != nil {
		global.Log.Info(err)
	}
	token, err := jwt.GenerateToken(jwt.MyPayLoad{
		UserID:   1,
		Status:   1,
		Username: "cc",
		NickName: "xxx",
	}, privateKey)

	publicKey, err := key.LoadPublicKey()
	if err != nil {
		global.Log.Info(err)
	}
	claims, err := jwt.ParseToken(token, publicKey)
	if err != nil {
		global.Log.Error(err)
	}
	fmt.Println("claims.Username:", claims.Username)

}

package main

import (
	_ "myServer/docs"
	"myServer/flag"
	"myServer/global"
	"myServer/routers"
	"myServer/service/cron_service"
	"myServer/setup"
	"net/http"
)

// @title CryptoCat
// @version 1.0
// @description CryptoCat API docs
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	//Load the configuration file
	setup.InitConf()
	//connect the database
	global.Log = setup.InitLogger()
	global.DB = setup.InitGorm()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
	}

	//connect the redis
	global.Redis = setup.ConnectRedis()
	//connect the es
	global.ESClient = setup.EsConnect()

	go cron_service.CronInit()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("CryptoCat_server运行在： %s", addr)
	global.Log.Infof("CryptoCat API文档运行在： http://%s/swagger/index.html#", addr)
	router.StaticFS("uploadFile", http.Dir("uploadFile"))
	// 第一个uploads是web中访问的路径
	// 第二个uploads是项目根路径下的uploads
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}

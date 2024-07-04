package main

import (
	"myServer/global"
	"myServer/routers"
	"myServer/setup"
)

func main() {
	//Load the configuration file
	setup.InitConf()
	//connect the database
	global.Log = setup.InitLogger()
	global.DB = setup.InitGorm()
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("server运行在： %s", addr)
	router.Run(addr)
}

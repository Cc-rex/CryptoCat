package main

import (
	"myServer/global"
	"myServer/setup"
)

func main() {
	//Load the configuration file
	setup.InitConf()
	//connect the database
	global.Log = setup.InitLogger()
	global.Log.Warnln("xixi")
	global.Log.Error("xixi")
	global.Log.Infof("xixi")
	global.DB = setup.InitGorm()
}

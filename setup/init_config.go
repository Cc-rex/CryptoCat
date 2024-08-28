package setup

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
	"myServer/config"
	"myServer/global"
	"myServer/utils/key"
)

const ConfigFile = "settings.yaml"

func InitConf() {
	config := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConfig error : %s", err))
	}
	err = yaml.Unmarshal(yamlConf, config)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = config
	err = InitKey()
	if err != nil {
		global.Log.Error(err)
	}
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}

func InitKey() error {
	privatePath := global.Config.Jwt.PrivateKeyPath
	publicPath := global.Config.Jwt.PublicKeyPath
	err := key.GenerateKeyPair(privatePath, publicPath)
	if err != nil {
		global.Log.Warn(err)
		return err
	}

	return nil
}

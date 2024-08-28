package main

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"myServer/setup"
)

var client *elastic.Client

func init() {
	setup.InitConf()
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"

	client, err = elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("ES连接失败 %s", err.Error())
	}
}

type DemoModel struct {
}

func (DemoModel) Index() string {
	return "demo_index"
}

func main() {

}

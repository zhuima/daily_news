package main

import (
	"daily_news/config"
	"daily_news/robot"
	"daily_news/alert"
	"strings"

	"github.com/sirupsen/logrus"

	//"errors"

)

var contents string

func main() {
	logrus.Info("Start!")


	// 获取引用变量
	var c config.Config
	conf := c.GetConf()

	d := strings.Split(conf.Urllist, " ")

	// 获取报警信息
	Title := conf.Title
	dingdingurl := conf.DingdingUrl

	// 遍历url实现业务信息汇集
	for url := range d {
		content, err := robot.Bot(d[url])
		if err != nil {
			logrus.Error("Can't Get info from robot!")
		}
		contents = contents + " \n" + content
	}


	if contents != "" {
		logrus.Info("send info is", contents)
		_, err := alert.SendMsg(dingdingurl, Title, contents)
		if err != nil {
			logrus.Fatal("Notify faild!")
		}
		logrus.Info("Done!")

	} else {
		logrus.Info("Wow, Can't get anything!")
	}

}

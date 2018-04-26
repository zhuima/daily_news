/*
解析yaml文件，获取所需
参考文档:
https://codebe.org/2018/01/04/go/15/
https://blog.csdn.net/yuyinghua0302/article/details/78794397
*/

package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//  定义结构体来实现配置文件对应相关事宜
type Config struct {
	Title       string `yaml: "title"`
	DingdingUrl string `yaml: "dingdingurl"`
	Urllist     string `yaml: "urllist"`
}

// 读取配置文件并返回一个结构体，全局使用
func (c *Config) GetConf() *Config {
	buffer, err := ioutil.ReadFile("./conf/config.yaml")

	if err != nil {
		logrus.Fatal("get yaml", err.Error())
	}

	err = yaml.Unmarshal(buffer, c)

	if err != nil {
		logrus.Fatal("get yaml", err.Error())
	}
	return c
}

package core

import (
	"backend/config"
	"backend/global"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const configFile = "settings.yaml"

// 读取配置操作
// 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}
	yamlConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("read config file error, %v", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatal("unmarshal config file error, %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

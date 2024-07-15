package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

const yamlFile = "settings.yaml"

// 初始化配置文件
func Init_conf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		panic(fmt.Sprintf("read yamlfile error %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yaml Unmarshal err %s", err)
	}
	log.Println("yaml Unmarshal success")
	global.Config = c
}

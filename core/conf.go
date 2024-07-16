package core

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

const yamlFile = "settings.yaml"

// Init_conf init configuration file
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

// Set_yaml update configuration file
func Set_yaml() error {
	bytes, err := yaml.Marshal(global.Config)
	if err != nil {
		return errors.New("configuration file update err " + err.Error())
	}
	err = ioutil.WriteFile(yamlFile, bytes, fs.ModePerm)
	if err != nil {
		return errors.New("configuration file update err " + err.Error())
	}
	return nil
}

package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gwa_server/config"
	"gwa_server/global"
	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "config.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("获取yaml文件失败: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("configyaml初始化拆分失败: %v", err)
	}
	log.Println("yaml文件初始化成功")
	global.Config = c
}

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Read(configs ...string) {
	conf := append(configs, "conf/app.yml")[0]
	YamlConfigLoad(conf)
}

// 配置文件名，conf 目录下找起
func YamlConfigLoad(config string) {
	yamlFile, err := ioutil.ReadFile(config)
	if err != nil {
		panic(fmt.Sprintf("读取自定义配置文件失败：%s", err))
	}
	if yamlFile == nil {
		panic("没有找到配置文件")
	}

	// 数据库配置文件加载
	if err = yaml.Unmarshal(yamlFile, &C); err != nil {
		panic(fmt.Sprintf("加载自定义配置失败：%s", err))
	}
}

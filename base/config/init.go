package config

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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

func InitLog() {
	err := logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s/%s.log", "maxdays":15}`, C.Beego.LogDir, web.BConfig.AppName))
	if err != nil {
		panic(fmt.Sprintf("初始化日志文件失败：%s", err))
	}
}

func InitSwaggerAPI() {
	web.BConfig.WebConfig.DirectoryIndex = true
	web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
}

func InitDB() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", C.DB.Platform, C.DB.DSN())
	if err != nil {
		panic(fmt.Sprintf("初始化数据库连接失败 - %s", err))
	}
}

func TableCreate() {
	// 默认的表名
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(fmt.Sprintf("创建数据表失败：%s", err))
	}
}

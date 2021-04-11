package main

import (
	"beego-learn/base/config"
	"beego-learn/modules"
	_ "beego-learn/routers" // 必要
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // 必要
)

func init() {
	config.YamlConfigLoad("conf/application.yml")
}

func main() {
	// 配置初始化
	config.InitLog()
	config.InitDB()
	config.TableCreate()
	if web.BConfig.RunMode == "dev" {
		config.InitSwaggerAPI()
		orm.Debug = config.C.Beego.OrmDebug // sql 打印
	}

	// 自定义模块初始化
	modules.Init()

	// 启动项目
	logs.Info(fmt.Sprintf("%s 服务以【%s】模式运行...", web.BConfig.AppName, web.BConfig.RunMode))
	web.Run()
}

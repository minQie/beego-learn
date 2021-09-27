package main

import (
	"beego-learn/modules"
	"beego-learn/modules/config"
	_ "beego-learn/routers" // 必要
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	bConfig "github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 配置文件配置项到对应配置实体的映射
	config.Read("conf/app_dev.yml")
}

func main() {
	// 基础项初始化
	initLog()
	initDB()
	initSwagger()

	// 自定义模块初始化
	modules.Init()

	// 启动项目

	// web.BConfig.AppName
	// web.BConfig.RunMode

	// appName, err := bConfig.String("appname")
	// if err != nil {
	// 	panic(err)
	// }
	// runMode, err := bConfig.String("runmode")
	// if err != nil {
	// 	panic(err)
	// }
	appName, err := web.AppConfig.String("appname")
	if err != nil {
		panic(err)
	}
	runMode, err := web.AppConfig.String("runmode")
	if err != nil {
		panic(err)
	}

	logs.Info(fmt.Sprintf("%s 服务以【%s】模式运行...", appName, runMode))
	web.Run()
}

// beego 默认有个默认的日志输出实例【目的：控制台】【级别：Debug】
func initLog() {
	// 是否打印 sql；通过事务对象或者新的 orm 事务写法，不会打印 sql；orm 一定是 client/orm 包
	orm.Debug = config.C.Beego.OrmDebug

	// web.BConfig.AppName
	appName, err := bConfig.String("appname")
	if err != nil {
		panic(err)
	}
	logName := fmt.Sprintf(`{"filename":"%s/%s.log", "maxdays":15}`, config.C.Beego.LogDir, appName)
	if err := logs.SetLogger(logs.AdapterFile, logName); err != nil {
		panic(fmt.Sprintf("初始化日志文件失败：%s", err))
	}
}

func initSwagger() {
	if web.BConfig.RunMode == web.DEV {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func initDB() {
	// NOTE import _ "github.com/go-sql-driver/mysql"
	if err := orm.RegisterDataBase("default", config.C.DB.Platform, config.C.DB.DSN()); err != nil {
		panic(fmt.Sprintf("初始化数据库连接失败 - %s", err))
	}
}

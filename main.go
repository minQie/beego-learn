package main

import "C"
import (
	"beego-learn/modules"
	"beego-learn/modules/config"
	_ "beego-learn/routers" // 必要
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // 必要
)

// TODO json 类库已经有 v3

// TODO 不能将 conf 包省略引用，因为直接通过 C 调用变量树型，会导致 goland 自动引入一个不知道什么鬼的 "C"
func init() {
	// TODO beego 的配置文件到实体的映射是在什么时候
	// TODO 当参数实体字段在配置文件中没有对应的，会报错么（应当报错）
	// 配置文件配置项到对应配置实体的映射
	config.Read("conf/app.yml")
}

func main() {
	// 基础项初始化
	BeegoInit()

	// 自定义模块初始化
	modules.Init()

	// 启动项目
	logs.Info(fmt.Sprintf("%s 服务以【%s】模式运行...", web.BConfig.AppName, web.BConfig.RunMode))
	web.Run()
}

func BeegoInit() {
	InitLog()
	InitDB()
	TableCreate()
	InitSwaggerAPI()
}

// beego 默认有个默认的日志输出实例【目的：控制台】【级别：Debug】
func InitLog() {
	err := logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"target/%s/%s.log", "maxdays":15}`, C.Beego.LogDir, web.BConfig.AppName))
	if err != nil {
		panic(fmt.Sprintf("初始化日志文件失败：%s", err))
	}
}

func InitSwaggerAPI() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
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
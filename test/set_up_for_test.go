package test

import (
	"beego-learn/conf"
	"beego-learn/modules"
	"beego-learn/modules/config"
	_ "beego-learn/routers"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.YamlConfigLoad("../conf/application_test.yaml")

	conf.InitLog()
	conf.InitDB()
	conf.TableCreate()
	conf.InitSwaggerAPI()
	orm.Debug = config.C.Beego.OrmDebug

	modules.Init()
}

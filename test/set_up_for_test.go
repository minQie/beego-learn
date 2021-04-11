package test

import (
	"beego-learn/base/config"
	"beego-learn/modules"
	_ "beego-learn/routers"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.YamlConfigLoad("../conf/application_test.yaml")

	config.InitLog()
	config.InitDB()
	config.TableCreate()
	config.InitSwaggerAPI()
	orm.Debug = config.C.Beego.OrmDebug

	modules.Init()
}

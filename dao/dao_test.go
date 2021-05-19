package dao

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // 重要
	"testing"
)

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	dsn := "root:123@tcp(127.0.0.1:8090)/hellobeego?charset=utf8mb4&loc=Local"
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		panic(fmt.Sprintf("初始化数据库连接失败 - %s", err))
	}

	// 打印 sql
	orm.Debug = true
}

// 查询里边有一个比较大的注意点 - 查询如果一条都没有查到，方法会返回 err：<QuerySeter> no row found
func TestDao(t *testing.T) {
	// saveOneTest()
	// saveListTest()
	// findOneTest()
	findListTest()
	// paramTest()
	// multiParamTest()
	// resultMappingTest()
	// selectCountTest()
	// testOrmDelete()
	// jsonTimeFieldTest()
	// transactionTest()
	//transaction2Test()
}

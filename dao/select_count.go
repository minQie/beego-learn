package dao

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func selectCountTest() {
	// 1、接收 COUNT 函数的结果
	sql := `
    SELECT
        COUNT(id)
    FROM 
        user`

	var count int
	if err := orm.NewOrm().Raw(sql).QueryRow(&count); err != nil {
		logs.Error(err)
		return
	}
	logs.Info(count)

	// 2、判断查询数据结果的条数（当然，这样写是不行的，你需要给接收数据的参数，不然运行时直接panic）
	sql = `SELECT id FROM user`
	sum, err := orm.NewOrm().Raw(sql).QueryRows()
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(sum)
}

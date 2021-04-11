package dao

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type userDTO struct {
	Id       int64 `orm:"column(id)"`
	Username string
}

func resultMappingTest() {
	sql := `
    SELECT
        id, username
    FROM
        user
    WHERE 
        id = 1`

	// 测试1：支持小蛇转大驼峰
	//      单表起别名不会更改 MySQL 查询结果的字段名
	// 测试2：username → Username（ok）、username → UserName（凉凉）
	// 测试3：可以通过 `orm:"column(字段名)"` 进行结果集映射（优先级大于结构体字段名），但是这个标签同时也是表字段到 GO 结构体字段的映射依据，所以是不现实的
	var userDTO userDTO
	if err := orm.NewOrm().Raw(sql).QueryRow(&userDTO); err != nil {
		logs.Error(err)
		return
	}
	logs.Info(userDTO)
}

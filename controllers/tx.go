package controllers

import (
	"beego-learn/utils"
	adapterOrm "github.com/beego/beego/v2/adapter/orm"
	clientOrm "github.com/beego/beego/v2/client/orm"
)

type TestController struct {
	BaseController
}

// @Title 测试满足哪些条件，才会有 SQL 打印
// @router /tx1 [GET]
func (c *TestController) TestTxOne() {
	// 1、配置文件中配置 orm_debug = true（x）
	// 2、x
	adapterOrm.Debug = true
	// 3、✔
	clientOrm.Debug = false
	var (
		o   = adapterOrm.NewOrm()
		err error
	)

	// 查
	var theMap = make([]interface{}, 0)
	if _, err = o.Raw(`SELECT * FROM a`).QueryRows(&theMap, &theMap); err != nil {
		c.ResponseJson(nil, err)
		return
	}

	// 事务 写写
	_ = o.Begin()
	utils.HandleTransaction(o, err)

	q := `INSERT a (name) VALUES (?)`
	if _, err = o.Raw(q, "a4").Exec(); err != nil {
		c.ResponseJson(nil, err)
		return
	}
	if _, err = o.Raw(q, "a5").Exec(); err != nil {
		c.ResponseJson(nil, err)
		return
	}
}

// @Title 测试满足哪些条件，才会有 SQL 打印
// @router /tx2 [GET]
func (c *TestController) TestTxTwo() {
	// 1、配置文件中配置 orm_debug = true（x）
	// 2、x
	adapterOrm.Debug = true
	// 3、✔
	clientOrm.Debug = true
	var (
		o   = clientOrm.NewOrm()
		err error
	)


	// 查
	var theMap = make([]interface{}, 0)
	if _, err = o.Raw(`SELECT * FROM a`).QueryRows(&theMap, &theMap); err != nil {
		c.ResponseJson(nil, err)
		return
	}

	// 事务 写写
	tx, _ := o.Begin()
	// utils.HandleTransaction(o, err)

	q := `INSERT a (name) VALUES (?)`
	if _, err = tx.Raw(q, "a4").Exec(); err != nil {
		c.ResponseJson(nil, err)
		return
	}
	if _, err = tx.Raw(q, "a5").Exec(); err != nil {
		c.ResponseJson(nil, err)
		return
	}
	_ = tx.Commit()
}

package dao

import (
	"beego-learn/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func paramTest() {
	q := `
    SELECT 
		*
	FROM
		user
    WHERE
        id = ? AND is_deleted = ?`

	// <√>  []interface{}
	// <√> *[]interface{}
	// <×> **[]interface{}：type: unsupported type []interface {}, a slice of interface

	// var param = []interface{}{1, 0}
	// var paramAddress = &param

	var user models.User
	if err := orm.NewOrm().Raw(q, 1, 0).QueryRow(&user); err != nil {
		logs.Error(err)
		return
	}
	logs.Info(user)
}

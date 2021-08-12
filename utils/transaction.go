package utils

import "github.com/beego/beego/v2/adapter/orm"

// HandleTx example
// var err error
// o := orm.NewOrm()
// _ = o.Begin()
// defer help.HandleTx(o, &err)
// ...
// repo.NewXxx(o).Xxx
func HandleTx(o orm.Ormer, e *error) {
	if *e == nil {
		_ = o.Commit()
	} else {
		_ = o.Rollback()
	}
}

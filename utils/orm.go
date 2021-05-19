package utils

import "github.com/beego/beego/v2/adapter/orm"

func HandleTransaction(o orm.Ormer, err *error) {
	if *err != nil {
		_ = o.Rollback()
	} else {
		_ = o.Commit()
	}
}
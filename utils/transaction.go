package utils

import "github.com/beego/beego/v2/client/orm"

// var success = true
// tx, _ = r.Begin()
// defer help.TxControl(tx, &success)
// ...
func TxControl(tx orm.TxOrmer, success *bool) {
	if *success {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
}

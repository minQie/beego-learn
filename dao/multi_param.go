package dao

import (
	"beego-learn/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func multiParamTest() {
	ids := []int64{1, 2, 3}

	sql := `
    SELECT 
        * 
    FROM 
        user
    WHERE
        is_deleted = ? AND id IN (?,?,?)`

	var user models.User
	// 并不需要将参数都展开，这样写也是支持的
	if err := orm.NewOrm().Raw(sql, 0, ids).QueryRow(&user); err != nil {
		logs.Error(err)
		return
	}
	logs.Info(user)
}

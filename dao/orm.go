package dao

import (
	"beego-learn/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/prometheus/common/log"
)

func testOrmDelete() {
	// Delete 方法参数只能是 *类型，用 **类型：
	// panic: <Ormer> table: `.` not found, make sure it was registered with `RegisterModel()`

	// 且实体一定要注册（orm.RegisterModel(new(User))），否则
	// panic: <Ormer> table: `beego-learn/models.User` not found, make sure it was registered with `RegisterModel()`

	user := new(models.User)
	user.Id = 1

	// Delete 后边的参数类型为 string 表示实体中哪些字段作为数据匹配条件（不传默认 id，主动传了值，就没有 id了）
	if _, err := orm.NewOrm().Delete(user, "username"); err != nil {
		log.Error(err)
		return
	}

	logs.Info(user)
}

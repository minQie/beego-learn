package dao

import (
	"beego-learn/models"
	"beego-learn/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"strings"
	"time"
)

func saveListTest() {
	// 插入参数
	users := []*models.User{{
		RoleId:   1,
		Account:  "7",
		Password: "123",
		Username: "七",
		Age:      10,
		Gender:   1,
		Birthday: time.Now(),
	}, {
		RoleId:   1,
		Account:  "8",
		Password: "123",
		Username: "八",
		Age:      10,
		Gender:   1,
		Birthday: time.Now(),
	}}

	var b strings.Builder
	b.Grow(100 + 15*len(users))
	params := make([]interface{}, 0, 7*len(users))

	b.WriteString(`
	INSERT user(
        role_id, account, password, username, age, gender, birthday
    ) VALUES `)

	oneParamStr := utils.OrmJoinRepeat(7)
	for index, user := range users {
		// 拼接 sql
		b.WriteString(oneParamStr)
		if index != len(users)-1 {
			b.WriteString(`,`)
		}

		// 拼接 参数
		params = append(params, user.RoleId, user.Account, user.Password, user.Username, user.Age, user.Gender, user.Birthday)
	}

	// params 后边加 ... 也行
	if _, err := orm.NewOrm().Raw(b.String(), params).Exec(); err != nil {
		logs.Error(err)
	}
}

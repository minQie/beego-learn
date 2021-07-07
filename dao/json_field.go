package dao

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// beego 会如何序列化 【Go time.Time 类型字段】 到 【MySQL json 中的一个字段】
// time.Time：                                    【2021-03-05 12:55:29】（直接 println 的结果同下）
// time.Time.String()：                           【2021-03-05 12:55:29.8899722 +0800 CST m=+0.006027401】
// time.Time.Format("2006-01-02T15:04:05Z07:00")：【2021-03-05T12:55:29+08:00】

// 但是存在问题，时间存储格式如上，是无法正常反序列化的，所以存储到 MySQL 时
// 不能存储 time.now()、time.now().String()
// 应该存储：time.now().Format("2006-01-02T15:04:05Z07:00")
func timeFieldTest() {
	q := `
	INSERT
		hellobeego.test_json_field
	(name, extra) 
	VALUES
		(?, JSON_OBJECT('refresh_at', ?))`

	params := []interface{}{"time1", time.Now()}
	if _, err := orm.NewOrm().Raw(q, params).Exec(); err != nil {
		panic(err)
	}
}

type R struct {
	J J `orm:"column(j)" json:"j"`
}
type J struct {
	Name string `orm:"column(name)" json:"name"`
}
// 结论：查不出来，将 R.J 的类型改为 string 就查询的出来了
func jsonFieldTest() {
	q := `
	SELECT
		j
	FROM hellobeego.test_json_field
	WHERE id = 1`

	r := new(R)
	if err := orm.NewOrm().Raw(q).QueryRow(r); err != nil {
		panic(err)
	}
	logs.Info(*r)
}

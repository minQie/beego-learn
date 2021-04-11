package dao

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// 注意 beego 有一个参数如果可为空的坑在里边，如果某个参数可以为空，你把他放到 interface{} 类型的切片中作为参数，就等着收到 panic
// 你需要将 (*Xxx)(nil) 转成 (interface{})(nil) 即 nil，可以通过如下的方法，sPtr 的类型不能改为 interface{}（详见 interface_nil_test.go）
func BeegoOrmNilSafe(sPtr *string) interface{} {
	if sPtr == nil {
		return nil
	}
	return sPtr
}

func saveOneTest() {
	sql := `INSERT user(
        role_id,
        account,
        password, 
        username, 
        age, 
        gender, 
        birthday
    ) VALUES (?, ?, ?, ?, ?, ?, ?)`

	// 测试：sql: expected 7 arguments, got 1
	// 不能直接拿参数实体作为 sql 的参数

	// 标准的插入写法如下
	params := []interface{}{1, "qq", "123", "丘丘", 100, 1, time.Now()}
	if _, err := orm.NewOrm().Raw(sql, params).Exec(); err != nil {
		logs.Error(err)
	}
}

package dao

import (
	"beego-learn/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func findOneTest() {
	sql := `
    SELECT 
        *
    FROM 
        user
    WHERE 
        id = ? AND is_deleted = 0`

	// 测试（失败）
	// 1、实体指针为 nil
	// panic: reflect: call of reflect.Value.Type on zero Value
	// 2、传实体、传容量为n的切片、传容量为n的切片指针地址
	// panic: <RawSeter.QueryRow> all args must be use ptr
	// 3、传容量为n的切片地址
	// sql: expected 15 destination arguments in Scan, not 1

	// 测试（成功）
	// 方式一：xxxPtr := new(Xxx) | QueryRow(xxxPtr)
	//                          | QueryRow(&xxxPtr) 这都行，也就行到这里了，再多一层就不行了
	// 方式二：var xxx Xxx        | QueryRow(&xxx)
	// userPtr := new(models.User)

	// 结论：框架会将查询的结果赋值到你给的容器对象地址对应的容器对象中，所以要求对应容器对象地址不能为空
	// 你希望通过结构体切片来接收指定的切片，结果会是预料之外的，因为框架会以为你希望将每一列结果分别封装到切片中

	userPtr := new(models.User)
	// userPtrAdd := &userPtr
	// var user models.User

	logs.Info("%p", userPtr)
	if err := orm.NewOrm().Raw(sql, 1).QueryRow(&userPtr); err != nil {
		logs.Error(err)
		return
	}
	logs.Info("%p", userPtr)

	logs.Info(*userPtr)
}

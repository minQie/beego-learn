package dao

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func findListTest() {
	sql := `
    SELECT
        id, username
    FROM
        user
    WHERE 
        id = 1 OR id = 2`

	// 测试1：容器允许为 nil
	// 测试2：测试前后切片类型的容器的底层数组的地址都变了
	// 测试3：切片的元素类型是否是指针类型都支持
	// 测试4：QueryRows 的形参必须得是 切片的指针类型
	// 结论：开发者不应该自己初始化了存储结果的容器，beego 不会使用（为 nil 都行），自己初始化相当于白白产生内存垃圾
	// beego\v2@v2.0.1\client\orm\orm_raw.go - 605 位置处的源码是 beego 使用内部创建的 切片 直接替换 入参容器
	// 这里反射赋值的概念和 Go 标准库 的 json.Unmarshal 方法的理念是不同的，json.Unmarshal 方法会使用入参的容器
	// 其实也可以理解，毕竟这里的大部门应用场景，一般都不能够预测数据条数；不支持是不是考虑了并发安全，json.Unmarshal 中是会检测并发的，检测到了直接 panic
	// 总之，还是像 json.Unmarshal 那样考虑一下比较好，假如 sql 含有 limit 子句，那么开发者完全可以预先设置 容器 slice 的 cap 来达到最好的性能

	var (
		userDTOPtrs     []*userDTO
		userDTOs        []userDTO
		userDTOPtrsWith = make([]*userDTO, 2)
		userDTOsWith    = make([]userDTO, 2)
	)

	logs.Info("%p", userDTOPtrsWith)
	if _, err := orm.NewOrm().Raw(sql).QueryRows(&userDTOPtrsWith); err != nil {
		logs.Error(err)
		return
	}
	logs.Info("%p", userDTOPtrsWith)

	logs.Info(userDTOPtrs)
	// logs.Info(*userDTOPtrs[0], *userDTOPtrs[1])
	logs.Info(userDTOs)
	logs.Info(userDTOPtrsWith)
	// logs.Info(*userDTOPtrsWith[0], *userDTOPtrsWith[1])
	logs.Info(userDTOsWith)
}

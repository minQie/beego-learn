package modules

import "beego-learn/modules/attachment"

// 因为模块可能存在，需要根据配置文件中的配置进行初始化的，所以不能通过 go 的 init 的初始化机制来初始化
func Init() {
	attachment.Init()
}

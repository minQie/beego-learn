package local

import (
	"github.com/beego/beego/v2/core/logs"
	"testing"
)

// 源码原理很简单，了解一下是有好处的：
// 一个方面是动态参数占位符 %v %s 有了解切实的应用场景
// 一个是了解 Beego 日志是怎么做到简单易用的

// 简单表达源码的含义，方法形参有两个：常量内容（可带有模板参数）、动态参数
// 如果常量包含 动态参数占位符(%)，就不对常量内容进行处理，否则就依据动态参数的个数给常量内容添加 %v，最后交给 fmt.Sprintf 处理
// 也就是说不要指望指定了两个占位符，但是带有三个动态参数，能正确的输出日志（自己写的话，就判断 % 的数量和动态参数的数量是否匹配，少了就加 v%）
func TestLogOutput(t *testing.T) {
	logs.Info("%d %d", 1, 2)
	logs.Info(1, 2)
}
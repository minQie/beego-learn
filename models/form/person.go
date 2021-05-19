package form

// 用于测试接收嵌套的参数结构（主要是 json）
type Person struct {
	Age       int       `form:"age"`
	Name      string    `form:"name"`
	Addresses []Address `form:"addresses"`
}

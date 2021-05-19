package form

// 用于测试接收简单的参数结构
type UserForm struct {
	Age  int    `form:"age"`
	Name string `form:"name"`
}

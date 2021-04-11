package trial_param

// 用于测试接收简单的参数结构
type UserForm struct {
	Age  int    `form:"age"  json:"age"`
	Name string `form:"name" json:"name"`
}

// 用于测试接收嵌套的参数结构（主要是 json）
type Person struct {
	Age       int       `form:"age"       json:"age"`
	Name      string    `form:"name"      json:"name"`
	Addresses []Address `form:"addresses" json:"addresses"`
}

type Address struct {
	Name string `form:"name" json:"name"`
	Tag  string `form:"tag"  json:"tag"`
}

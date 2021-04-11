package trial_param

import "encoding/json"

// @router /complex_form [POST]
func (c *ConstructController) ComplexForm() {
	var (
		person Person
		err    error
	)
	err = c.ParseForm(&person)
	c.ResponseJson(person, err)
}

// @router /complex_json [POST]
func (c *ConstructController) ComplexJson() {
	var (
		person Person
		err    error
	)
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &person)
	c.ResponseJson(person, err)
}

// 测试结论：PostForm 的参数封装并不支持嵌套的 json 参数结构，也就是想要支持嵌套结构就要借助 json 结构，那就不如参数格式就使用 json

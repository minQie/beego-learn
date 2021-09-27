package controllers

import "encoding/json"

/* 当前文件的核心结构体，内置 beego 的 Controller 基类，就能够使用其内部的 接口、... */
type HelloController struct {
	BaseController
}

/* 实现 beego.Controller.ControllerInterface.URLMapping 方法 */
func (c *HelloController) URLMapping() {
	c.Mapping("Hello", c.Hello)
}

/* 实现 beego.Controller 中 ControllerInterface接口的 Get 方法 */
func (c *HelloController) Get() {
	c.Hello()
}

// @router /:key [GET]
func (c *HelloController) Hello() {
	c.ResponseJson("hello", nil)
}

// @router /postForm [POST]
func (c *HelloController) PostForm() {
	name := c.GetString("name")

	c.ResponseJson(name, nil)
}

// @router /postJson [POST]
func (c *HelloController) PostJson() {
	m := map[string]string{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &m)

	c.ResponseJson(m, err)
}
package trial_param

import (
	"beego-learn/controllers"
	"beego-learn/models/form"
	"encoding/json"
)

/* 将请求参数解析到结构体中 */

type ConstructController struct {
	controllers.BaseController
}

// [Params √] [form-data ×] [x-www-form-urlencoded ×] [json ×]
// @router /form [GET]
func (c *ConstructController) FormGet() {
	var (
		userForm form.UserForm
		err      error
	)
	err = c.ParseForm(&userForm)
	c.ResponseJson(userForm, err)
}

// [Params √] [form-data √] [x-www-form-urlencoded √] [json ×]
// @router /form [POST]
func (c *ConstructController) FormPost() {
	var (
		userForm form.UserForm
		err      error
	)
	err = c.ParseForm(&userForm)
	c.ResponseJson(userForm, err)
}

// [Params ×] [form-data ×] [x-www-form-urlencoded ×] [json ×]
// @router /json [GET]
func (c *ConstructController) JsonGet() {
	var (
		userForm form.UserForm
		err      error
	)
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &userForm)
	c.ResponseJson(userForm, err)
}

// [Params ×] [form-data ×] [x-www-form-urlencoded ×] [json √]
// @router /json [POST]
func (c *ConstructController) JsonPost() {
	var (
		userForm form.UserForm
		err      error
	)
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &userForm)
	c.ResponseJson(userForm, err)
}

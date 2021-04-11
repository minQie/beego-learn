package trial_param

import (
	"beego-learn/controllers"
	"strconv"
)

type BasicController struct {
	controllers.BaseController
}

// [Params √] [form-data ×] [x-www-form-urlencoded ×] [json ×]
// @router /simple_get [GET]
func (c *BasicController) SimpleGet() {
	age, _ := c.GetInt("age")
	name := c.GetString("name")

	userForm := UserForm{
		Age:  age,
		Name: name,
	}
	c.ResponseJson(userForm, nil)
}

// [Params √] [form-data √] [x-www-form-urlencoded √] [json ×]
// @router /simple_post [POST]
func (c *BasicController) SimplePost() {
	age, _ := c.GetInt("age")
	name := c.GetString("name")

	userForm := UserForm{
		Age:  age,
		Name: name,
	}
	c.ResponseJson(userForm, nil)
}

// [Params ✔] [form-data ×] [x-www-form-urlencoded ×] [json ×]
// @router /query_get [GET]
func (c *BasicController) QueryGet() {
	ageStr := c.Ctx.Input.Query("age")
	name := c.Ctx.Input.Query("name")
	age, _ := strconv.Atoi(ageStr)

	userForm := UserForm{
		Age:  age,
		Name: name,
	}
	c.ResponseJson(userForm, nil)
}

// [Params ✔] [form-data ✔] [x-www-form-urlencoded ✔] [json ×]
// @router /query_post [POST]
func (c *BasicController) QueryPost() {
	ageStr := c.Ctx.Input.Query("age")
	name := c.Ctx.Input.Query("name")
	age, _ := strconv.Atoi(ageStr)

	userForm := UserForm{
		Age:  age,
		Name: name,
	}
	c.ResponseJson(userForm, nil)
}

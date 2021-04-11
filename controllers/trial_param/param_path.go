package trial_param

import (
	"beego-learn/controllers"
)

type PathParamController struct {
	controllers.BaseController
}

// 注意，路由的写法，只有这样写路由确保必须带有路径参数才能访问到该接口方法，有路径参数也访问不到没有定义路径参数的接口
// @router /pathParam/:id [GET]
func (c *PathParamController) Get() {
	id, err := c.GetInt64(":id")
	c.ResponseJson(id, err)
}

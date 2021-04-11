package trial_param

import "beego-learn/controllers"

type CookieController struct {
	controllers.BaseController
}

// @title 浏览器请求该接口将得到一个固定的 Cookie
// @router /set [GET]
func (c *CookieController) Set() {
	c.Ctx.SetCookie("token", "123", 60)
	c.ResponseJson(nil, nil)
}

// @title 浏览器请求清除指定的 Cookie
// @router /clear [POST]
func (c *CookieController) Clear() {
	c.Ctx.SetCookie("token", "", -1)
	c.ResponseJson(nil, nil)
}

// @title 浏览器请求接口，会带上上一个接口设置好的 Cookie，这里会测试假如同名的多个 Cookie 会得到怎样的值
// @router /get [GET]
func (c *CookieController) Get() {
	token := c.Ctx.GetCookie("token")
	c.ResponseJson(token, nil)
}

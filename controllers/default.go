package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type DefaultController struct {
	web.Controller
}

// @router / [GET]
func (c *DefaultController) Get() {
	c.Data["Website"] = "www.wangmincong.com"
	c.Data["Email"] = "mincong.wang@shanghairanking.com"
	c.TplName = "index.tpl"
}

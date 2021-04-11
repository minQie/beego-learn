package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.wangmincong.com"
	c.Data["Email"] = "mincong.wang@shanghairanking.com"
	c.TplName = "index.tpl"
}

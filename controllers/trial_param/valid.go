package trial_param

import (
	"beego-learn/controllers"
	"beego-learn/models/form"
	"github.com/beego/beego/v2/core/logs"
)

type TController struct {
	controllers.BaseController
}

// @router / [GET]
func (c *TController) P() {
	var (
		f   = new(form.Valid)
		err error
	)
	if err = c.ParseFormAndValidate(f); err != nil {
		c.ResponseJson(nil, err)
		return
	}

	logs.Info("正常解析到参数：%v", *f)
	c.ResponseJson(nil, nil)
}

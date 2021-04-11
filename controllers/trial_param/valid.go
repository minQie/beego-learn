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
		f   = new(form.TForm)
		err error
	)
	if err = c.ParseForm(f); err != nil {
		logs.Info("解析参数出错：%v", *f)
		c.ResponseJson(nil, err)
		return
	}
	if err = c.Validate(f); err != nil {
		logs.Error("校验参数出错：%s", err)
		c.ResponseJson(nil, err)
		return
	}

	logs.Info("正常解析到参数：%v", *f)
}

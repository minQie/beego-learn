package trial_param

import (
	"beego-learn/controllers"
	"github.com/beego/beego/v2/core/logs"
)

type MultiParamController struct {
	controllers.BaseController
}

// @Title Get 请求中的多参数（注意是 GET 请求）
// @router / [GET]
func (c *MultiParamController) Get() {

	// GET - Url：[小明 小刚 小光]√
	names1 := c.GetStrings("names")
	logs.Info(names1)

	// GET - Url：[小明 小刚 小光]√
	names2 := c.Ctx.Request.Form["names"]
	logs.Info(names2)

	// GET - Url：nil×
	names3 := c.Ctx.Input.GetData("names")
	logs.Info(names3)

	// GET - Url：小明×
	names4 := c.Ctx.Input.Query("names")
	logs.Info(names4)

	// 一、GET - Url：如上
	// 二、POST - FormData：收不到
	// 三、POST - X-WWW-FORM-URL-ENCODED：收不到

	// 结论：c.GetStrings 和 c.Ctx.Request.Form 都行，推荐 c.GetStrings
}

// @Title POST 请求 form-data 方式的多参数（注意是 POST 请求）
// @router / [POST]
func (c *MultiParamController) Post() {
	names1 := c.GetStrings("names")
	logs.Info(names1)

	names2 := c.Ctx.Request.Form["names"]
	logs.Info(names2)

	names3 := c.Ctx.Input.GetData("names")
	logs.Info(names3)

	names4 := c.Ctx.Input.Query("names")
	logs.Info(names4)

	// 一、GET - Url：同上
	// 二、POST - FormData：同上
	// 三、POST - X-WWW-FORM-URL-ENCODED：同上

	// 结论：c.GetStrings 和 c.Ctx.Request.Form 都行，推荐 c.GetStrings
}

// 公司项目可以通过 c.GetStrings 来解析 names=value1,values2,values3
// 是因为在 BaseController 中封装了这样的方法

// func (c *BaseController) GetStrings(key string, def ...[]string) ([]string, error) {
// 	sVal := c.Ctx.Input.Query(key)
// 	if len(sVal) == 0 && len(def) > 0 {
// 		return def[0], nil
// 	}
// 	return strings.Split(sVal, ","), nil
// }

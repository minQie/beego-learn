package trial_redirect

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type RedirectController struct {
	web.Controller
}

// 重定向
func (c *RedirectController) Redirect() {
	var redirectTo = "https://www.baidu.com"

	// 方式一：原生
	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, redirectTo, http.StatusFound)

	// 方式二：beego
	c.Ctx.Redirect(http.StatusFound, redirectTo)
}

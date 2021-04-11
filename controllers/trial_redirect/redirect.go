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
	// 方式一：原生
	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, "redirect to url", http.StatusFound)

	// 方式二：beego
	c.Ctx.Redirect(http.StatusFound, "redirect to url")
}

package filter

import (
	"beego-learn/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func WhiteIP(ctx *context.Context) {
	intercepted := false
	ip := utils.IP(ctx)

	// 本地回环 和 内网 IP 不拦截
	if utils.IsInternalIP(ip) {
		return
	}

	switch ctx.Request.URL.String() {
	case "/v1/xxx/xxx":
		// intercepted = !slice.String(config.C.IPWhite.Xxx).Contain(ip)
	}

	if intercepted {
		logs.Error("没有访问权限", "not allow request ip", ip)
		_ = ctx.Output.JSON(map[string]interface{}{
			"Code": -1,
			"Msg":  "没有访问权限",
		}, web.BConfig.RunMode != web.PROD, false)
	}
}

package filter

import (
	"beego-learn/filter/model"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strconv"
)

// 权限认证的核心方法
func Authenticate(c *context.Context) {
	requestAuthorization := model.ParseRequest(c)

	if err := requestAuthorization.Authenticate(); err != nil {
		logs.Error(err)
		c.Abort(http.StatusUnauthorized, strconv.Itoa(http.StatusUnauthorized))
		return
	}
}

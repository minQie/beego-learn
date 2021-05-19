package trial_redirect

import (
	"beego-learn/controllers"
	"bytes"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"net/http"
)

type DispatcherController struct {
	controllers.BaseController
}

/* 转发：大致的想法和实现如下，未经实际测试，应该也有更好的实现 */
func (c *DispatcherController) dispatcher() {
	var (
		oriReq     = c.Ctx.Request
		oriReqBody []byte

		client      = http.DefaultClient
		disReq      *http.Request
		disResp     *http.Response
		disRespBody []byte
		err         error
	)

	// 请求体
	if oriReqBody, err = ioutil.ReadAll(oriReq.Body); err != nil {
		logs.Error("Read origin request body ", err)
		return
	}
	// 创建转发请求
	if disReq, err = http.NewRequest(oriReq.Method, oriReq.URL.String(), bytes.NewReader(oriReqBody)); err != nil {
		logs.Error("http.NewRequest ", err)
		return
	}
	// 复制请求头
	for k, v := range oriReq.Header {
		disReq.Header.Set(k, v[0])
	}
	// 发起转发请求
	if disResp, err = client.Do(disReq); err != nil {
		logs.Error("转发请求发生错误 ", err)
		return
	}
	defer disResp.Body.Close()

	// 将转发请求的结果做为原请求的响应
	if disRespBody, err = ioutil.ReadAll(disResp.Body); err != nil {
		logs.Error("Read origin request body ", err)
		return
	}
	if _, err = c.Ctx.ResponseWriter.Write(disRespBody); err != nil {
		logs.Error("Response ", err)
		return
	}
}

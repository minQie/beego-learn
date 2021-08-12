package controllers

import (
	"beego-learn/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type BaseController struct {
	web.Controller
}

/* 接口返回数据的数据实体 */
type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type BaseControllerReq interface {
	// GetUrl 获取用户实际发送请求的完整 url 地址
	GetUrl() string

	ParseFormAndValidate(param interface{}) (map[string]string, error)
	ParseJsonAndValidate(param interface{}) (map[string]string, error)
	Validate(param interface{}) (map[string]string, error)

	ParseJson(param interface{}) error
	GetLoginUserId() int64
}

type BaseControllerResp interface {
	ResponseJson(data interface{}, err error)
	RespVJson(data interface{}, err error)
	RespOkJson(data interface{}) error
	RespFailJson(err error) error
	RespErrJson(errMsg string) error

	ResponseFile(reader io.Reader, filename string)
}

func (c *BaseController) GetUrl() string {
	scheme := "http://"
	if c.Ctx.Request.TLS != nil {
		scheme = "https://"
	}

	return strings.Join([]string{scheme, c.Ctx.Request.Host, c.Ctx.Request.RequestURI}, "")
}

func (c *BaseController) ResponseJson(data interface{}, err error) {
	if err != nil {
		err = c.RespFailJson(err)
	} else {
		err = c.RespOkJson(data)
	}

	if err != nil {
		logs.Error("BaseController 响应 json 发生错误：%s", err)
	}
}

func (c *BaseController) RespVJson(data interface{}, err error) {
	c.Data["json"] = CommonResponse{
		Code: -1,
		Msg:  err.Error(),
		Data: data,
	}
	if err = c.ServeJSON(); err != nil {
		logs.Error("BaseController 响应 json 发生错误：%s", err)
	}
}

func (c *BaseController) RespOkJson(data interface{}) error {
	c.Data["json"] = CommonResponse{
		Code: http.StatusOK,
		Data: data,
	}
	return c.ServeJSON()
}

func (c *BaseController) RespFailJson(err error) error {
	c.Data["json"] = CommonResponse{
		Code: -1,
		Msg:  err.Error(),
	}
	return c.ServeJSON()
}

func (c *BaseController) RespErrJson(errMsg string) error {
	resp := CommonResponse{
		Code: -1,
		Msg:  errMsg,
	}
	c.Data["json"] = resp
	return c.ServeJSON()
}

// MDN：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Disposition
// Golang 中国：https://www.golangtc.com/t/54d9ca47421aa9170200000f
// stackoverflow：https://stackoverflow.com/questions/93551/how-to-encode-the-filename-parameter-of-content-disposition-header-in-http
// https://blog.robotshell.org/2012/deal-with-http-header-encoding-for-file-download/
// 1、Postman 只取 filename 的值，不理会 filename* 的值（找不到相关资料）
// 2、Go 默认 UTF-8 不支持其他编码，所以无法通过转成 ISO8859 来解决 Postman 下载的文件名乱码问题
func (c *BaseController) ResponseFile(file io.Reader, filename string) {
	// 这里的做法是上面的资料，以及下边的源码中提及到的最佳的做法
	// 详见：c.Ctx.Output.Download(string, ...string) 方法
	name := url.PathEscape(filename)

	c.Ctx.Output.Header("Content-Type", "application/octet-stream")
	c.Ctx.Output.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"; filename*=UTF-8''%s`, name, name))
	_, err := io.Copy(c.Ctx.ResponseWriter, file)
	if err != nil {
		logs.Error("响应文件失败", err)
	}
}

func (c *BaseController) ParseFormAndValidate(param interface{}) (map[string]string, error) {
	if err := c.ParseForm(param); err != nil {
		return nil, err
	}
	return c.Validate(param)
}

func (c *BaseController) ParseJsonAndValidate(param interface{}) (map[string]string, error) {
	if err := c.ParseJson(param); err != nil {
		return nil, err
	}
	return c.Validate(param)
}

func (c *BaseController) ParseJson(param interface{}) error {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, param); err != nil {
		c.ResponseJson(nil, fmt.Errorf("读取请求体中的 json 参数失败：%s", err))
		return err
	}
	return nil
}

func (c *BaseController) Validate(param interface{}) (map[string]string, error) {
	valid := validation.Validation{}
	noErr, err := valid.Valid(param)

	if err != nil {
		errMsg := fmt.Sprintf("存在非法参数：%s", err)
		return nil, utils.LogError(errMsg)
	}
	if !noErr {
		msgMap := make(map[string]string, len(valid.Errors))

		for _, vErr := range valid.Errors {
			msgMap[vErr.Key] = vErr.Message
		}
		return msgMap, errors.New(msgMap[valid.Errors[0].Key])
	}
	return nil, nil
}

// 用户没有登录，将返回 -1
func (c *BaseController) GetLoginUserId() int64 {
	userId, _ := c.GetInt64("loginUserId", -1)
	return userId
}

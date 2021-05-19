package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"net/http"
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

type BaseControllerInterface interface {
	// 获取用户实际发送请求的完整 url 地址
	GetUrl() string
	// 请求
	ParseFormAndValidate(param interface{}) error
	ParseJsonAndValidate(param interface{}) error
	ParseJson(param interface{}) error
	Validate(param interface{}) error
	// 响应
	ResponseJson(data interface{}, err error)
	RespOkJson(data interface{}) error
	RespFailJson(err error) error
	RespErrJson(errMsg string) error

	GetLoginUserId() int64
}

func (c *BaseController) GetUrl() string {
	scheme := "http://"
	if c.Ctx.Request.TLS != nil {
		scheme = "https://"
	}

	return fmt.Sprintf(strings.Join([]string{scheme, c.Ctx.Request.Host, c.Ctx.Request.RequestURI}, ""))
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

func (c *BaseController) RespOkJson(data interface{}) error {
	c.Data["json"] = CommonResponse{
		Code: http.StatusOK,
		Data: data,
	}
	return c.ServeJSON()
}

func (c *BaseController) RespFailJson(err error) error {
	c.Data["json"] = CommonResponse{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
	}
	return c.ServeJSON()
}

func (c *BaseController) RespErrJson(errMsg string) error {
	resp := CommonResponse{
		Code: http.StatusInternalServerError,
		Msg:  errMsg,
	}
	c.Data["json"] = resp
	return c.ServeJSON()
}

func (c *BaseController) ParseFormAndValidate(param interface{}) error {
	if err := c.ParseForm(param); err != nil {
		return err
	}
	return c.Validate(param)
}

func (c *BaseController) ParseJsonAndValidate(param interface{}) error {
	if err := c.ParseJson(param); err != nil {
		return err
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

func (c *BaseController) Validate(param interface{}) error {
	var (
		valid = validation.Validation{}
		noErr bool
		err   error
	)
	if noErr, err = valid.Valid(param); err != nil {
		return fmt.Errorf("校验请求体中的 json 参数失败：%s", err)
	}
	if !noErr {
		errMsgMap := make(map[string]string, len(valid.Errors))
		for _, value := range valid.Errors {
			errMsgMap[value.Key] = value.Message
		}
		return errors.New("存在非法参数")
	}
	return nil
}

// 用户没有登录，将返回 -1
func (c *BaseController) GetLoginUserId() int64 {
	userId, _ := c.GetInt64("loginUserId", -1)
	return userId
}

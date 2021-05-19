package utils

import (
	"errors"
	"github.com/beego/beego/v2/core/logs"
)

// Param msg 作为返回值的错误信息内容
// Param v   一般传递实际发生错误对应的 error 实例
func LogError(msg string, v ...interface{}) error {
	logs.Error(msg, v)
	return errors.New(msg)
}

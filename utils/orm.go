package utils

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	cliOrm "github.com/beego/beego/v2/client/orm"
)

func HandleTransaction(o orm.Ormer, err error) {
	if err != nil {
		_ = o.Rollback()
	} else {
		_ = o.Commit()
	}
}

// 包装一下查询的错误结果，防止将太露骨的错误直接返回给前端
func PackRowsErr(err error, modelName string) error {
	if err == nil {
		return nil
	}
	switch err {
	case cliOrm.ErrNoRows:
		return fmt.Errorf("没有找到有效的%s", modelName)
	case cliOrm.ErrMultiRows:
		return fmt.Errorf("找到多条%s", modelName)
	default:
		return err
	}
}

func IsNoRow(e error) bool {
	return errors.Is(e, cliOrm.ErrNoRows)
}

func IsMultiRows(e error) bool {
	return errors.Is(e, cliOrm.ErrMultiRows)
}

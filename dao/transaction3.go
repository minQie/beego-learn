package dao

import (
	"context"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type test3Repo struct {
	orm.Ormer
}

func NewTest3Repo(o orm.Ormer) *test2Repo {
	if o == nil {
		o = orm.NewOrm()
	}
	return &test2Repo{o}
}

func GetDML(r orm.Ormer, tx []orm.TxOrmer) orm.DML {
	var dml orm.DML
	if len(tx) != 0 && tx[0] != nil {
		dml = tx[0]
	}

	return dml
}

func (r *test3Repo) save1(tx ...orm.TxOrmer) error {
	var dml = GetDML(r, tx)

	q := `
	INSERT user 
	(role_id, account, password, username, age, birthday)
	VALUES (1, 'a', '123', 'ping1', '1', '2021-03-01 10:08:48')`

	if _, err := dml.Raw(q).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *test3Repo) save2(tx ...orm.TxOrmer) error {
	var dml = GetDML(r, tx)

	q := `
	INSERT user 
	(role_id, account, password, username, age, birthday)
	VALUES (1, 'b', '123', 'ping2', '2', '2021-03-01 10:09:39')`

	if _, err := dml.Raw(q).Exec(); err != nil {
		return err
	}
	return errors.New("保存失败")
}

// 优化的写法，且看写法，就知道不可复用
// 注意：Repo 在日常开发中，不要和 orm.DML 中的接口方法取相同的名字，这样会造成方法覆盖，从而 Repo 不算实现了接口
func transaction3Test() {
	var (
		r   = NewTest3Repo(nil)
		err error
	)

	err = r.DoTx(func(_ context.Context, txOrm orm.TxOrmer) error {
		if err = r.save1(); err != nil {
			logs.Error("save1 发生错误")
			return err
		}
		if err = r.save2(); err != nil {
			logs.Error("save2 发生错误")
			return err
		}
		return nil
	})

	// err 已经处理了
	return
}

package dao

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type testRepo struct {
	orm.TxOrmer
}

func NewTestRepo(tx orm.TxOrmer) *testRepo {
	if tx == nil {
		tx, _ = orm.NewOrm().Begin()
	}
	return &testRepo{tx}
}

func (r *testRepo) save1() error {
	q := `
	INSERT user 
	(role_id, account, password, username, age, birthday)
	VALUES (1, 'a', '123', 'ping1', '1', '2021-03-01 10:08:48')`

	if _, err := r.Raw(q).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *testRepo) save2() error {
	q := `
	INSERT user 
	(role_id, account, password, username, age, birthday)
	VALUES (1, 'b', '123', 'ping2', '2', '2021-03-01 10:09:39')`

	// panic("123")

	if _, err := r.Raw(q).Exec(); err != nil {
		return err
	}
	return errors.New("保存失败")
}

// Repo 中内置 orm.TxOrmer，dao 方法1 正常运行，dao 方法2 panic 或 err，都能够正常回滚
// 注意，Repo 实例无法复用：
//     已经提交事务，方法再次调用 或 再次提交事务：transaction has already been committed or rolled back
//     即使上面允许，开发者也无法主动调用 Begin 方法，即，无法执行多条 SQL 的同事务操作
func transactionTest() {
	var (
		r      = NewTestRepo(nil)
		result = true
		err    error
	)

	if err = r.save1(); err != nil {
		logs.Error("save1 发生错误")
		result = false
	}
	if err = r.save2(); err != nil {
		logs.Error("save2 发生错误")
		result = false
	}
	if result {
		_ = r.Commit()
	} else {
		_ = r.Rollback()
	}

	// 已经提交或者回滚的事务不能再次开启 - 继续使用
	if err = r.save1(); err != nil {
		logs.Error("again1", err)
	}
	_ = r.Commit()
}

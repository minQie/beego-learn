package dao

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type test2Repo struct {
	orm.Ormer
}

func NewTest2Repo(o orm.Ormer) *test2Repo {
	if o == nil {
		o = orm.NewOrm()
	}
	return &test2Repo{o}
}

func (r *test2Repo) save1() error {
	q := `
	INSERT user 
	(role_id, account, password, username, age, birthday)
	VALUES (1, 'a', '123', 'ping1', '1', '2021-03-01 10:08:48')`

	if _, err := r.Raw(q).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *test2Repo) save2() error {
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

// 事务回滚失败
// 失败原因：并没有通过开启的事务对象来执行 sql
func transaction2Test() {
	var (
		r   = NewTest2Repo(nil)
		err error
	)

	tx, _ := r.Begin()
	defer func() {
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	// 开始执行处于一个事务的两个方法
	if err = r.save1(); err != nil {
		logs.Error("save1 发生错误")
		return
	}
	if err = r.save2(); err != nil {
		logs.Error("save2 发生错误")
		return
	}

	// 复用
	// if err = r.save1(); err != nil {
	// 	logs.Error("again", err)
	// }
}

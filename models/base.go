package models

import "time"

type BaseEntity struct {
	Id         int64     `orm:"column(id)"              json:"id"` // 主键id
	CreateTime time.Time `orm:"column(create_time)"     json:"-"`  // 创建时间
	UpdateTime time.Time `orm:"column(update_time)"     json:"-"`  // 更新时间
}

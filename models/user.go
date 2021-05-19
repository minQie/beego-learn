package models

import (
	"beego-learn/models/const/gender"
	"beego-learn/models/const/status"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

/*
-- DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id`              BIGINT UNSIGNED           NOT NULL KEY AUTO_INCREMENT                                    COMMENT '主键id',
  `role_id`         INT UNSIGNED              NOT NULL                                                       COMMENT '角色id',
  `account`         VARCHAR(20)               NOT NULL                                                       COMMENT '账号、昵称',
  `password`        VARCHAR(128)              NOT NULL                                                       COMMENT '密码',
  `username`        VARCHAR(10)               NOT NULL                                                       COMMENT '名称',
  `age`             TINYINT UNSIGNED          NOT NULL                                                       COMMENT '年龄',
  `gender`          TINYINT UNSIGNED          NOT NULL DEFAULT 0                                             COMMENT '性别（0=未知，1=男，2=女）',
  `birthday`        DATETIME                  NOT NULL                                                       COMMENT '生日',
  `status`          TINYINT UNSIGNED          NOT NULL DEFAULT 0                                             COMMENT '状态（0=正常，1=停用）',
  `last_login_ip`   VARCHAR(15)               NULL                                                           COMMENT '上次登录ip',
  `last_login_time` DATETIME                  NULL                                                           COMMENT '上次登录时间',
  `create_time`     DATETIME                  NOT NULL DEFAULT c_TIMESTAMP                             COMMENT '创建时间',
  `update_time`     DATETIME                  NOT NULL DEFAULT c_TIMESTAMP ON UPDATE c_TIMESTAMP COMMENT '更新时间',
  `delete_by`       BIGINT UNSIGNED           NULL                                                           COMMENT '删除当前用户的用户id',
  `is_deleted`      TINYINT UNSIGNED          NOT NULL DEFAULT 0                                             COMMENT '删除标识（0=正常，1=删除）'
) COMMENT = '用户表';

CREATE UNIQUE INDEX `uk_account` ON `user` (account);
CREATE UNIQUE INDEX `uk_username` ON `user` (username);
*/

// 定义用户数据表实体
type User struct {
	Id            int64       `orm:"column(id)"              json:"id"`
	RoleId        int64       `orm:"column(role_id)"         json:"role_id"`   // 角色id
	Account       string      `orm:"column(account)"         json:"account"`   // 昵称、账号
	Password      string      `orm:"column(password)"        json:"-"`         // 密码
	Username      string      `orm:"column(username)"        json:"name"`      // 名称
	Age           int         `orm:"column(age)"             json:"age"`       // 年龄
	Gender        gender.Enum `orm:"column(gender)"          json:"gender"`    // 性别
	Birthday      time.Time   `orm:"column(birthday)"        json:"birthday"`  // 生日
	Status        status.Enum `orm:"column(status)"          json:"status"`    // 状态
	LastLoginIp   string      `orm:"column(last_login_ip)"   json:"-"`         // 上次登录id
	LastLoginTime time.Time   `orm:"column(last_login_time)" json:"-"`         // 上次登录时间
	CreateTime    time.Time   `orm:"column(create_time)"     json:"-"`         // 创建时间
	CreatedBy     int64       `orm:"column(created_by)"      json:"createdBy"` // 创建人
	UpdateTime    time.Time   `orm:"column(update_time)"     json:"-"`         // 更新时间
	UpdatedBy     int64       `orm:"column(updated_by)"      json:"updatedBy"` // 更新人
	DeleteTime    *time.Time  `orm:"column(delete_time)"     json:"-"`         // 删除标识
	DeleteBy      int64       `orm:"column(delete_by)"       json:"-"`         // 删除人
}

func init() {
	orm.RegisterModel(new(User))
}

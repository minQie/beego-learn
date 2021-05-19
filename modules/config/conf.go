/*
conf 包是 git 配置成忽略的，故当前的 base/config 作为存放配置基建模块的位置
*/
package config

import "beego-learn/modules/config/child"

// yaml 配置文件配置项的结构体映射
var C config

type config struct {
	Beego         child.Beego
	DB            child.DB
	Client        child.Client
	ResponseCache child.ResponseCache `yaml:"response_cache"`
	Attachment    child.Attachment    `yaml:"attachment"`
}

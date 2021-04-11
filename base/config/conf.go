package config

import "beego-learn/base/config/child"

// yaml 配置文件配置项的结构体映射
var C Config

type Config struct {
	Beego         child.BeegoConfig
	DB            child.DBConfig
	Client        child.ClientConfig
	ResponseCache child.ResponseCacheConfig `yaml:"response_cache"`
	Attachment    child.AttachmentConfig    `yaml:"attachment"`
}

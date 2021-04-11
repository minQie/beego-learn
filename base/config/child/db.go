package child

import "fmt"

type DBConfig struct {
	Platform string `yaml:"platform"` // 数据库平台，例如，mysql（实际上目前也就支持这个）
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// DSN DataBase Source Name
func (db DBConfig) DSN() string {
	dsnTemplate := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local"
	return fmt.Sprintf(dsnTemplate, db.Username, db.Password, db.IP, db.Port, db.DbName)
}

package persistence

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	InitDb    bool
	AliasName string
	UserName  string
	Password  string
	Host      string
	Port      int
	DbName    string
}

func (c *MysqlConfig) Build() (b BaseConfig) {

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.UserName, c.Password, c.Host, c.Port, c.DbName)
	b = BaseConfig{
		dataSource: dataSource,
		aliasName:  c.AliasName,
		driver: DriverMysql,
		initDb:     c.InitDb,
	}

	//b.params = append(b.params, c.MaxIdle, c.MaxConn)
	return
}

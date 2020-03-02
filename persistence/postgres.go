package persistence

import (
	"fmt"
	//_ "github.com/lib/pq"
)

type PostgresConfig struct {
	InitDb     bool
	AliasName  string
	UserName   string
	Password   string
	Host       string
	Port       int
	DbName     string
	MaxIdle    int
	MaxConn    int
}

func (c *PostgresConfig) Build() (b BaseConfig) {

	dataSource := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", c.UserName, c.Password, c.DbName, c.Host, c.Port)
	b = BaseConfig{
		dataSource: dataSource,
		aliasName: c.AliasName,
		driver: DriverPostgres,
		initDb: c.InitDb,
	}

	b.params = append(b.params, c.MaxIdle, c.MaxConn)
	return
}

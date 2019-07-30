package persistence

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	//_ "github.com/lib/pq"
)

type PostgresConfig struct {

	UserName string

	Password string
	Host string
	Port int
	DbName string
	MaxIdle int
	MaxConn int
	InitDb bool
}

//Initialize
func InitPostgresOrm(config PostgresConfig) (context OrmContext, err error) {

	context = OrmContext{}
	connInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", config.UserName, config.Password, config.DbName, config.Host, config.Port)
	err = orm.RegisterDataBase("default", "postgres", connInfo, config.MaxIdle, config.MaxConn)
	if err == nil {
		if config.InitDb {
			err = orm.RunSyncdb("default", false, true)
		}
		context.Context = orm.NewOrm()
	}
	return context, err
}
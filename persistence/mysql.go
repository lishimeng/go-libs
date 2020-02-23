package persistence

import (
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Database string
}

func InitMysqlOrm(config SqliteConfig) (context OrmContext, err error) {
	context = OrmContext{}
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err == nil {
		err = orm.RegisterDataBase("default", "mysql", config.Database)
		if err == nil {
			err = orm.RunSyncdb("default", false, true)
			context.Context = orm.NewOrm()
		}
	}

	return context, err
}
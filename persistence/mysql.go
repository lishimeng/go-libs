package persistence

import (
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	BaseConfig
	Database string
}

func InitMysqlOrm(config MysqlConfig) (context OrmContext, err error) {
	context = OrmContext{}
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err == nil {
		err = orm.RegisterDataBase("default", "mysql", config.Database)
		if err == nil {
			if config.BaseConfig.ForceDdl {
				err = orm.RunSyncdb("default", false, true)
			}
			context.Context = orm.NewOrm()
		}
	}

	return context, err
}
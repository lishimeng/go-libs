package persistence

import (
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
)

type SqliteConfig struct {
	Database string
}

func InitSqliteOrm(config SqliteConfig) (context OrmContext, err error) {
	context = OrmContext{}
	err = orm.RegisterDriver("sqlite", orm.DRSqlite)
	if err == nil {
		err = orm.RegisterDataBase("default", "sqlite3", config.Database)
		if err == nil {
			err = orm.RunSyncdb("default", false, true)
			context.Context = orm.NewOrm()
		}
	}

	return context, err
}

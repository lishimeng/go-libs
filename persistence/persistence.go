package persistence

import (
	"github.com/astaxie/beego/orm"
)

//RegisterModel
type DriverType orm.DriverType

type Driver struct {
	name string
	t    orm.DriverType
}

type BaseConfig struct {
	initDb bool
	aliasName string
	driver Driver
	dataSource string
	params []int
	models []interface{}
}

var DriverMysql = Driver{"mysql", orm.DRMySQL}
var DriverSqlite = Driver{"sqlite3", orm.DRSqlite}
var DriverOracle = Driver{"oracle", orm.DROracle}
var DriverPostgres = Driver{"postgres", orm.DRPostgres}
var DriverTiDB = Driver{"tidb", orm.DRTiDB}

func (b *BaseConfig) RegisterModel(models ...interface{}) {
	b.models = append(b.models, models...)
}

func InitOrm(config BaseConfig) (context OrmContext, err error) {

	context = OrmContext{}
	err = orm.RegisterDriver(config.driver.name, config.driver.t)
	if err == nil {
		if len(config.models) > 0 {
			orm.RegisterModel(config.models...)
		}
		err = orm.RegisterDataBase(config.aliasName, config.driver.name, config.dataSource, config.params...)
		if err == nil {
			if config.initDb {
				err = orm.RunSyncdb(config.aliasName, false, true)
			}
			context.Context = orm.NewOrm()
		}
	}

	return context, err
}

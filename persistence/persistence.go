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

var DriverMysql = Driver{"mysql", orm.DRMySQL}
var DriverSqlite = Driver{"sqlite3", orm.DRSqlite}
var DriverOracle = Driver{"oracle", orm.DROracle}
var DriverPostgres = Driver{"postgres", orm.DRPostgres}
var DriverTiDB = Driver{"tidb", orm.DRTiDB}

func RegisterDriver(driver Driver) (err error) {
	err = orm.RegisterDriver(driver.name, driver.t)
	return
}

func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

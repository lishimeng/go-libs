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
var DriverSqlite = Driver{"mysql", orm.DRSqlite}
var DriverOracle = Driver{"mysql", orm.DROracle}
var DriverPostgres = Driver{"mysql", orm.DRPostgres}
var DriverTiDB = Driver{"mysql", orm.DRTiDB}

func RegisterDriver(driver Driver) (err error) {
	err = orm.RegisterDriver(driver.name, driver.t)
	return
}

func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

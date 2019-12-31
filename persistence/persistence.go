package persistence

import (
	"github.com/astaxie/beego/orm"
)

//RegisterModel
type DriverType orm.DriverType

const (
	_          DriverType = iota // int enum type
	DRMySQL                      // mysql
	DRSqlite                     // sqlite
	DROracle                     // oracle
	DRPostgres                   // pgsql
	DRTiDB                       // TiDB
)

func RegisterDriver(driver string, t DriverType) (err error) {
	err = orm.RegisterDriver(driver, orm.DriverType(t))
	return
}

func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

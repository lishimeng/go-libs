package persistence

import (
	"github.com/astaxie/beego/orm"
)

//RegisterModel
func RegisterDriver(driver string, t int) (err error) {
	err = orm.RegisterDriver(driver, orm.DriverType(t))
	return
}

func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

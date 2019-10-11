package persistence

import (
	"github.com/astaxie/beego/orm"
)

//RegisterModel
func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

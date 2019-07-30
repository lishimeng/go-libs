package persistence

import "github.com/astaxie/beego/orm"

type OrmContext struct {
	Context orm.Ormer
}

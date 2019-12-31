package persistence

import "github.com/astaxie/beego/orm"

type OrmContext struct {
	Context orm.Ormer
}

func (o *OrmContext) Transaction(h func(ctx OrmContext) error) (err error) {

	if h == nil {
		return
	}
	err = o.Context.Begin()

	if err != nil {
		return
	}

	err = h(*o)

	if err != nil {
		_ = o.Context.Rollback()
	} else {
		err = o.Context.Commit()
	}

	return
}

package persistence

import "strings"

func CheckErr(err error) error {

	if err != nil {
		if !strings.Contains(err.Error(), "LastInsertId") {
			return err
		}
	}
	return nil
}

type Pager struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
}

func (p *Pager) GetLimit() (limit int, start int) {
	limit = p.PageSize
	start = (p.PageNo - 1) * p.PageSize
	return limit, start
}

func BuildPager(count int, pageNo int, pageSize int) Pager {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Pager{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
	}
}

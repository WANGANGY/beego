package models

import "github.com/astaxie/beego/orm"

type Utsmessage struct {
	claimNo string
	processId string
	taskassign string
	taskassignName string
}

func  init()  {
	print("初始化 Utsmessage")
}
//get user list
func GetUtsmessagelist(page int64, page_size int64, sort string) (utsmessages []orm.Params, count int64) {
	o := orm.NewOrm()
	utsmessage := new(Utsmessage)
	qs := o.QueryTable(utsmessage)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&utsmessages)
	count, _ = qs.Count()
	return utsmessages, count
}
package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// ban表
type Ban struct {
	Id int
	// 被ban的类型
	Type string
	// 被ban的目标
	Target string
	// 生成时间
	Time time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Ban) Insert() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

func (m *Ban) Delete() (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

package models

import (
	"github.com/astaxie/beego/orm"
)

// 权限表
type Auth struct {
	Id         int
	Pid        int
	AuthName   string
	AuthUrl    string
	Sort       int
	Icon       string
	IsShow     int
	UserId     int
	CreateId   int
	UpdateId   int
	Status     int
	CreateTime int64
	UpdateTime int64
}

// 获取表名称
func (a *Auth) TableName() string {
	return TableName("uc_auth")
}

// 获取权限列表
func (this *Auth) AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(this.TableName())

	// 筛选条件
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	// 统计总数
	count, _ := query.Count()

	// 查询数据
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)
	return list, count
}

package models

import (
	"github.com/astaxie/beego/orm"
)

// 角色管理
type Role struct {
	Id         int64
	RoleName   string
	Detail     string
	CreateId   int64
	UpdateId   int64
	Status     int
	CreateTime int64
	UpdateTime int64
}

// 获取表名称
func (this *Role) TableName() string {
	return TableName("uc_role")
}

// 计算Offset
func (this *Role) LimitOffset(page, pageSize int) int {
	return LimitOffset(page, pageSize)
}

/**
 * 获取角色列表
 * @param page		int // 分页页数
 * @param pageSize	int	// 分页大小
 * @param filters   map[string]interface{} // 过滤参数
 * @return []*Role
 */
func (this *Role) RoleGetList(page, pageSize int, filters map[string]interface{}) (role []*Role, count int64, err error) {
	query := orm.NewOrm().QueryTable(this.TableName())
	for field, fieldValue := range filters {
		query.Filter(field, fieldValue)
	}

	// 统计条数
	count, err = query.Count()
	if err != nil {
		return role, 0, err
	}

	// 查询数据
	_, err = query.Limit(pageSize, this.LimitOffset(page, pageSize)).OrderBy("id").All(&role)
	if err != nil {
		return role, 0, err
	}
	return role, count, nil
}

/**
 * 按照id获取角色信息
 * @param id int 	// 角色表主键ID
 * @return Role
 */
func (this *Role) RoleGetInfo(id int) (*Role, error) {
	role := new(Role)
	err := orm.NewOrm().QueryTable(this.TableName()).One(role)
	if err != nil {
		return role, err
	}
	return role, nil
}

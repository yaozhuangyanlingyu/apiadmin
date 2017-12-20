package models

import (
	"github.com/astaxie/beego/orm"
)

// 管理员表
type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

// 拼接表名称
func (this *Admin) TableName() string {
	return TableName("uc_admin")
}

/**
 * 根据用户名获得用户信息
 * @param userName string
 */
func (this *Admin) AdminGetByName(userName string) (*Admin, error) {
	o := orm.NewOrm()
	o.Using("default")
	admin := new(Admin)
	err := o.QueryTable(this.TableName()).Filter("login_name", userName).One(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

/**
 * 根据用户ID获得用户信息
 * @param userId int
 */
func (this *Admin) AdminGetByUserId(userId int) (*Admin, error) {
	admin := new(Admin)
	err := orm.NewOrm().QueryTable(this.TableName()).Filter("id", userId).One(admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// 更新admin表数据
func (this *Admin) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(this, fields...)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 获取admin用户列表
 * @param page		int // 分页页数
 * @param pageSize	int	// 分页大小
 * @return list []*Admin, count int, err error
 */
func (this *Admin) AdminGetList(page, pageSize int, filters map[string]interface{}) (list []*Admin, count int64, err error) {
	list = make([]*Admin, 0)
	offset := (page - 1) * pageSize

	// query对象
	query := orm.NewOrm().QueryTable(this.TableName())

	// 参数过滤
	for field, fieldValue := range filters {
		query.Filter(field, fieldValue)
	}

	// 数据条数
	count, err = query.Count()
	if err != nil {
		return list, 0, err
	}

	// 查询数据
	_, err = query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	if err != nil {
		return list, 0, err
	}

	return list, count, nil
}

/**
 * 新增数据
 * @param userId int
 */
func (this *Admin) AdminCreate(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

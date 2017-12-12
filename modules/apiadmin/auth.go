package apiadmin

import (
	"github.com/apiadmin/models"
)

/**
 * 获取角色列表
 * @return role []*Role, count int64, err error
 */
func GetRoleList() (role []*models.Role, count int64, err error) {
	authObj := models.Role{}
	filters := make(map[string]interface{})
	filters["status"] = 1
	list, count, err := authObj.RoleGetList(1, 1000, filters)
	if err != nil {
		return role, count, err
	}
	return list, count, nil
}

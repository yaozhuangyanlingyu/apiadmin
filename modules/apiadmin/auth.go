package apiadmin

import (
	"github.com/apiadmin/models"
)

/**
 * 获取角色列表
 * @return role []*Role, count int64, err error
 */
func GetRoleList(page, limit int) (role []*models.Role, count int64, err error) {
	authObj := models.Role{}
	filters := make(map[string]interface{})
	filters["status"] = 1
	list, count, err := authObj.RoleGetList(page, limit, filters)
	if err != nil {
		return role, count, err
	}
	return list, count, nil
}

/**
 * 按照ID获取角色信息
 * @return role []*Role, count int64, err error
 */
func GetRoleInfoById(id int) (role *models.Role, err error) {
	modelObj := models.Role{}
	return modelObj.RoleGetInfo(id)
}

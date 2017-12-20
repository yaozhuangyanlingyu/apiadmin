package dataformat

import (
	"strconv"

	"github.com/apiadmin/models"
)

// admin格式化

/**
 * 角色列表数据格式化
 * @param roleList []*models.Role 	// 角色列表数据
 * @param roleIds  []string			// 用户所属角色ID
 * @return []map[string]interface{}
 */
func GetRoleListFormat(roleList []*models.Role, roleIds []string) []map[string]interface{} {
	list := make([]map[string]interface{}, len(roleList))
	for k, v := range roleList {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(roleIds); i++ {
			roleId, _ := strconv.ParseInt(roleIds[i], 10, 64)

			if roleId == v.Id {
				row["checked"] = 1
			}
		}
		row["Id"] = v.Id
		row["RoleName"] = v.RoleName
		list[k] = row
	}
	return list
}

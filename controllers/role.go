package controllers

import (
	"fmt"

	"github.com/apiadmin/modules/apiadmin"
)

// 角色控制器
type RoleController struct {
	BaseController
}

// 角色列表 - 页面
func (this *RoleController) List() {

	// 展示模板
	this.Display()
}

// 角色列表 - 数据
func (this *RoleController) Table() {
	// 初始化参数
	page, _ := this.GetInt("page", 1)
	limit, _ := this.GetInt("limit", 30)

	// 查询角色数据
	roleList, count, err := apiadmin.GetRoleList(page, limit)
	if err != nil {
		this.AjaxList(err.Error(), MSG_ERR, roleList, count)
	}

	// 返回数据
	this.AjaxList("success", MSG_OK, roleList, count)
}

// 编辑角色页面
func (this *RoleController) Edit() {
	// 获取参数
	id, err := this.GetInt("id")
	if err != nil {
		this.AjaxMsg("id参数不正确", MSG_ERR)
	}

	// 获取角色数据
	roleInfo, err := apiadmin.GetRoleInfoById(id)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	fmt.Println(roleInfo)

	// 显示模板
	this.Data["role"] = roleInfo
	this.Display()
}

package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/apiadmin/modules/apiadmin"
	"github.com/astaxie/beego"
)

type AdminController struct {
	BaseController
}

// 管理员列表页面
func (this *AdminController) List() {

	// 显示模板
	this.Display()
}

// 管理员列表数据
func (this *AdminController) Table() {
	// 初始化参数
	page, err := this.GetInt("page")
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	// 查询数据
	filters := make(map[string]interface{})
	adminList, count, err := apiadmin.GetAdminList(page, limit, filters)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	// 返回数据
	this.AjaxList("success", MSG_OK, adminList, count)
}

// 管理员信息编辑页面
func (this *AdminController) Edit() {
	// 初始化参数
	uid, err := this.GetInt("id", 0)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	// 查询管理员信息
	userInfo, err := apiadmin.GetUserInfoById(uid)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	// 查询角色
	roleIds := strings.Split(userInfo.RoleIds, ",")
	result, _, err := apiadmin.GetRoleList()
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
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

	// 展示模板
	this.Data["admin"] = userInfo
	this.Data["role"] = list
	this.Data["default_passwd"] = beego.AppConfig.String("site.default_passwd")
	this.Display()
}

// 修改管理员信息
func (this *AdminController) AjaxSave() {
	// 初始化参数
	id, err := this.GetInt("id", 0)
	if err != nil {
		this.AjaxMsg("id参数不正确", MSG_ERR)
	}
	loginName := strings.TrimSpace(this.GetString("login_name"))
	realName := strings.TrimSpace(this.GetString("real_name"))
	phone := strings.TrimSpace(this.GetString("phone"))
	email := strings.TrimSpace(this.GetString("email"))
	roleIds := strings.TrimSpace(this.GetString("roleids"))
	resetPwd, err := this.GetInt("reset_pwd")
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	// 修改数据
	adminInfo, err := apiadmin.GetUserInfoById(id)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	adminInfo.LoginName = loginName
	adminInfo.RealName = realName
	adminInfo.Phone = phone
	adminInfo.Email = email
	adminInfo.RoleIds = roleIds
	adminInfo.UpdateTime = time.Now().Unix()
	adminInfo.UpdateId = this.UserId

	// 重置密码
	if resetPwd == 1 {
		defaultPasswd := beego.AppConfig.String("site.default_passwd")
		md5Sum := md5.New()
		md5Sum.Write([]byte(defaultPasswd + adminInfo.Salt))
		newPasswd := hex.EncodeToString(md5Sum.Sum(nil))
		adminInfo.Password = newPasswd
	}
	err = adminInfo.Update()
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	// 返回提示信息
	this.AjaxMsg("操作成功", MSG_OK)
}

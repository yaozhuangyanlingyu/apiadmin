package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/apiadmin/libs"
	"github.com/apiadmin/models"
	"github.com/apiadmin/modules/apiadmin"
	"github.com/apiadmin/modules/dataformat"
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
	filters["status"] = 1
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
	result, _, err := apiadmin.GetRoleList(1, 1000)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	role := dataformat.GetRoleListFormat(result, roleIds)

	// 展示模板
	this.Data["admin"] = userInfo
	this.Data["role"] = role
	this.Data["default_passwd"] = beego.AppConfig.String("site.default_passwd")
	this.Display()
}

// 添加管理员页面
func (this *AdminController) Add() {
	// 角色数据
	result, _, err := apiadmin.GetRoleList(1, 1000)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	var roleIds []string
	role := dataformat.GetRoleListFormat(result, roleIds)

	// 默认密码
	this.Data["role"] = role
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

// 增加管理员信息
func (this *AdminController) AjaxAdd() {
	// 初始化参数
	admin := new(models.Admin)
	admin.LoginName = strings.TrimSpace(this.GetString("login_name"))
	admin.RealName = strings.TrimSpace(this.GetString("real_name"))
	admin.Phone = strings.TrimSpace(this.GetString("phone"))
	admin.Email = strings.TrimSpace(this.GetString("email"))
	admin.RoleIds = strings.TrimSpace(this.GetString("roleids"))

	// 参数验证
	if len(admin.LoginName) == 0 {
		this.AjaxMsg("登陆账号不能为空", MSG_ERR)
	}
	if len(admin.RealName) == 0 {
		this.AjaxMsg("真实姓名不能为空", MSG_ERR)
	}
	if len(admin.Phone) == 0 {
		this.AjaxMsg("手机号码不能为空", MSG_ERR)
	}
	if len(admin.Email) == 0 {
		this.AjaxMsg("电子邮箱不能为空", MSG_ERR)
	}

	// 检查用户名是否已经存在
	_, err := apiadmin.GetUserInfoByName(admin.LoginName)
	if err == nil {
		this.AjaxMsg("登陆名已经存在", MSG_ERR)
	}

	// 增加用户
	pwd, salt := libs.GetPasswdString("")
	admin.Password = pwd
	admin.UpdateId = this.UserId
	admin.CreateId = this.UserId
	admin.Status = 1
	admin.Salt = salt
	admin.CreateTime = time.Now().Unix()
	admin.UpdateTime = time.Now().Unix()
	if _, err := apiadmin.AdminCreate(admin); err != nil {
		this.AjaxMsg("管理员新增失败", MSG_ERR)
	}
	this.AjaxMsg("操作成功", MSG_OK)
}

// 删除管理员
func (this *AdminController) AjaxDel() {
	// 获取参数
	adminId, err := this.GetInt("id")
	if err != nil {
		this.AjaxMsg("id参数不正确", MSG_ERR)
	}

	// 修改数据
	adminInfo, err := apiadmin.GetUserInfoById(adminId)
	if err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	adminInfo.Id = adminId
	adminInfo.Status = 0
	adminInfo.UpdateTime = time.Now().Unix()
	if adminId == 1 {
		this.AjaxMsg("不能删除管理员", MSG_ERR)
	}
	if err = adminInfo.Update(); err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}

	this.AjaxMsg("操作成功", MSG_OK)
}

package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/apiadmin/modules/apiadmin"
)

// 用户控制器
type UserController struct {
	BaseController
}

// 用户资料 - 编辑页面
func (this *UserController) Edit() {
	// 查询管理员信息
	userInfo, _ := apiadmin.GetUserInfoById(this.UserId)

	// 展示模板
	this.Data["admin"] = userInfo
	this.Display()
}

// 用户资料 - 修改信息
func (this *UserController) AjaxSave() {
	// 初始化参数
	adminId, _ := this.GetInt("id")
	resetPwd := this.GetString("reset_pwd")
	userInfo, _ := apiadmin.GetUserInfoById(adminId)
	userInfo.RealName = strings.TrimSpace(this.GetString("real_name"))
	userInfo.Phone = strings.TrimSpace(this.GetString("phone"))
	userInfo.Email = strings.TrimSpace(this.GetString("email"))
	userInfo.UpdateTime = time.Now().Unix()
	userInfo.UpdateId = adminId

	// 验证填写的信息
	if len(userInfo.RealName) == 0 {
		this.AjaxMsg("真实姓名不能为空", MSG_ERR)
	}
	if len(userInfo.Phone) == 0 {
		this.AjaxMsg("手机号码不能为空", MSG_ERR)
	}
	if len(userInfo.Email) == 0 {
		this.AjaxMsg("电子邮箱不能为空", MSG_ERR)
	}

	// 修改密码
	if resetPwd == "1" {
		// 验证旧密码是否正确
		passwdOld := strings.TrimSpace(this.GetString("password_old"))
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(passwdOld + userInfo.Salt))
		md5Str := hex.EncodeToString(md5Ctx.Sum(nil))
		if md5Str != userInfo.Password {
			this.AjaxMsg("原始密码不正确", MSG_ERR)
		}

		// 验证重复密码是否一样
		passwordNew1 := strings.TrimSpace(this.GetString("password_new1"))
		passwordNew2 := strings.TrimSpace(this.GetString("password_new2"))
		if passwordNew1 != passwordNew2 {
			this.AjaxMsg("两次密码不一致", MSG_ERR)
		}

		// 验证密码长度
		if len(passwordNew1) < 6 {
			this.AjaxMsg("密码长度不能小于6位", MSG_ERR)
		}

		// 设置新密码
		md5Ctx = md5.New()
		md5Ctx.Write([]byte(passwordNew1 + userInfo.Salt))
		newPwd := hex.EncodeToString(md5Ctx.Sum(nil))
		userInfo.Password = newPwd
	}

	// 修改数据信息
	if err := userInfo.Update(); err != nil {
		this.AjaxMsg(err.Error(), MSG_ERR)
	}
	this.AjaxMsg("", MSG_OK)
}

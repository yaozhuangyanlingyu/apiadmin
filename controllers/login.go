package controllers

import (
	"strings"

	"github.com/apiadmin/modules/apiadmin"
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

// 登录接口
func (this *LoginController) LoginIn() {
	// POST提交
	if this.IsPost() {
		// 获取参数
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))

		// 验证参数
		errMsg := ""
		if len(username) == 0 {
			errMsg = "账号不能为空"
		}
		if len(password) == 0 {
			errMsg = "密码不能为空"
		}
		flash := beego.NewFlash()
		if errMsg != "" {
			flash.Error(errMsg)
		}

		// 验证并登录
		err := apiadmin.Login(&this.Controller, username, password)
		if err != nil {
			flash.Error(err.Error())
		}

		// 跳转处理
		this.Redirect(beego.URLFor("HomeController.Index"), 302)
		return
	}

	// 渲染模板
	this.TplName = "login/login.html"
}

// 退出登录
func (this *LoginController) LoginOut() {
	// cookie失效
	this.Ctx.SetCookie("auth", "", 0)

	// 跳转到登陆页
	this.Redirect(beego.URLFor("LoginController.LoginIn"), 302)
}

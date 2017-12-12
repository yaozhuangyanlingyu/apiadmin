package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/apiadmin/models"
	"github.com/apiadmin/modules/apiadmin"
	"github.com/astaxie/beego"
)

const (
	MSG_OK  = 0
	MSG_ERR = 1001
)

// 公共控制器
type BaseController struct {
	beego.Controller
	UserInfo       *models.Admin
	ControllerName string // 控制器名称
	ActionName     string // 方法名称
	UserId         int    // 用户ID
}

// 初始化
func (this *BaseController) Prepare() {
	// 自身变量初始化
	this.initVal()

	// 权限验证
	this.auth()
}

// 自身变量初始化
func (this *BaseController) initVal() {
	controllerName, actionName := this.GetControllerAndAction()
	this.ControllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.ActionName = strings.ToLower(actionName)

	// Title
	this.Data["siteName"] = beego.AppConfig.String("site.name")
}

// 权限验证
func (this *BaseController) auth() {
	// 获取token
	auth := this.Ctx.GetCookie("auth")
	authArr := strings.Split(auth, "|")
	this.UserId = 0

	// 验证token
	if len(authArr) == 2 {
		// 根据token获取用户数据
		idstr, token := authArr[0], authArr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := apiadmin.GetUserInfoById(userId)
			if err == nil {
				// 验证token
				md5Ctx := md5.New()
				tmpRemoteAddr := strings.Split(this.Ctx.Request.RemoteAddr, ":")
				remoteAddr := tmpRemoteAddr[0]
				md5Ctx.Write([]byte(user.Password + remoteAddr + user.Salt))
				md5sum := hex.EncodeToString(md5Ctx.Sum(nil))
				if md5sum == token {
					this.UserInfo = user
					this.UserId = userId
					this.AdminAuth()
				}
			}
		}
	}

	// 登录无效
	if this.UserId == 0 && this.ControllerName != "login" && this.ActionName != "loginin" {
		this.Redirect(beego.URLFor("LoginController.LoginIn"), 302)
		return
	}
}

// 菜单权限
func (this *BaseController) AdminAuth() {
	// 查询权限列表
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	authObj := models.Auth{}
	result, _ := authObj.AuthGetList(1, 1000, filters...)
	list1 := make([]map[string]interface{}, len(result))
	list2 := make([]map[string]interface{}, len(result))
	i, j := 0, 0
	for _, v := range result {
		row := make(map[string]interface{})
		if v.Id > 0 && v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list1[i] = row
			i++
		}
		if v.Id > 0 && v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}

	// 分配变量
	this.Data["SiteMenu1"] = list1[:i]
	this.Data["SiteMenu2"] = list2[:j]
}

// 判断是否post请求
func (this *BaseController) IsPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 显示模板
func (this *BaseController) Display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = this.ControllerName + "/" + this.ActionName + ".html"
	}
	this.Layout = "public/layout.html"
	this.TplName = tplname
}

// AJAX返回提示信息
func (this *BaseController) AjaxMsg(message interface{}, status int) {
	out := make(map[string]interface{})
	out["status"] = status
	out["message"] = message
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// Ajax返回列表数据
func (this *BaseController) AjaxList(message interface{}, code int, data interface{}, count int64) {
	out := make(map[string]interface{})
	out["msg"] = message
	out["code"] = code
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

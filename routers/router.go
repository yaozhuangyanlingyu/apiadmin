package routers

import (
	"github.com/apiadmin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 登陆退出
	// -- 登录
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	// -- 退出
	beego.Router("/logout", &controllers.LoginController{}, "*:LoginOut")

	// 用户相关
	// -- 资料修改 - 编辑页面
	beego.Router("/user/edit", &controllers.UserController{}, "*:Edit")
	// -- 资料修改 - 保存信息接口
	beego.Router("/user/save", &controllers.UserController{}, "*:AjaxSave")

	// 管理员相关
	// -- 管理员列表
	beego.Router("/admin/list", &controllers.AdminController{}, "*:List")
	// -- 管理员列表数据
	beego.Router("/admin/table", &controllers.AdminController{}, "*:Table")
	// -- 管理员信息编辑
	beego.Router("/admin/edit", &controllers.AdminController{}, "*:Edit")
	// -- 管理员信息保存
	beego.Router("/admin/save", &controllers.AdminController{}, "*:AjaxSave")

	// - 系统首页
	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")

	// - 控制面板
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
}

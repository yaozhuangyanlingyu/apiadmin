package controllers

type HomeController struct {
	BaseController
}

// 系统首页
func (this *HomeController) Index() {
	this.Data["pageTitle"] = "系统首页"

	this.TplName = "public/main.html"
}

// 控制面板
func (this *HomeController) Start() {
	this.Data["pageTitle"] = "控制面板"

	this.TplName = "home/start.html"
}

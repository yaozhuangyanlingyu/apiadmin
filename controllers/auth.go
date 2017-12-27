package controllers

// 权限控制器
type AuthController struct {
	BaseController
}

// 权限列表页面
func (this *AuthController) List() {

	// 显示模板
	this.Display()
}

// 获取权限数据
func (this *AuthController) Getnodes() {

}

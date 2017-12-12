package main

import (
	_ "github.com/apiadmin/routers"

	"github.com/apiadmin/setting"

	"github.com/astaxie/beego"
)

func main() {
	// 初始化配置
	err := setting.Init()
	if err != nil {
		panic(err)
		return
	}

	// 运行框架
	beego.Run()
}

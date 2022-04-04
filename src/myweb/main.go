package main

import (
	"myweb/controllers"
	_ "myweb/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/look", "download1") //左边路由，右边路径
	beego.SetStaticPath("/img", "static/img")
	beego.Router("/upload", &controllers.FileUploadController{}) //左边路由，右边控制器
	beego.Run()
}

package main

import (
	_ "myweb/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down1", "download1") //左边路由，右边路径
	beego.SetStaticPath("/down2", "static/img")
	beego.Run()
}

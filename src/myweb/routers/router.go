package routers

import (
	"myweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/hh", &controllers.MainController{})
	beego.Router("/h1", &controllers.Main2Controller{})
}

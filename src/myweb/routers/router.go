package routers

import (
	"myweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/h", &controllers.Main2Controller{})
}

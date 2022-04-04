package routers

import (
	"myweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hh", &controllers.MainController{})
	beego.Router("/h1", &controllers.Main2Controller{})
	beego.Router("/upload", &controllers.FileUploadController{})
}

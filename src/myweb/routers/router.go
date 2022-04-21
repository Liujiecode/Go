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
	beego.Router("/register", &controllers.RegisterController{})
	// # /welcome  为路由名
	// # WelcomeController{} 为控制器
	// # POST:Add  为 POST请求方式 Add 方法
	// # @@@@ 注意：
	// # 1.控制器名使用驼峰命名法
	// # 2.方法名首字符大写  如果小写是私有的方法 大写的是共有的方法
	// beego.Router("/welcome", &controllers.WelcomeController{},POST:Add);

}

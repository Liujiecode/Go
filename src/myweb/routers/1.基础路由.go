package routers

/*基础路由：
beego提供的常见的7种http请求类型方法的路由方案：get，post，head，options，delete等
*/
// func init() {
// 	beego.Get("/get", func(context *context.Context) {
// 		beego.Info("基础路由中的get请求！")
// 		context.Output.Body([]byte("基础路由中的get请求 get method"))
// 	})
// 	beego.Get("/getUserInfo", func(context *context.Context) {
// 		beego.Info("获取用户信息")
// 		context.Output.Body([]byte("获取用户信息请求 "))
// 	})
// 	beego.Post("/post", func(this *context.Context) {
// 		beego.Info("基础路由终得post请求")
// 		this.Output.Body([]byte("基础路由终得post请求"))
// 	})
// 	beego.Delete("/deleteInfo", func(context *context.Context) {
// 		beego.Info("基础路由终得delete方法")
// 		context.Output.Body([]byte("delete方法请求"))
// 	})

// }

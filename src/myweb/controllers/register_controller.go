package controllers

import (
	"fmt"
	"myweb/models"
	"time"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}
func (r *RegisterController) Post() {

	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")
	fmt.Println(username, password, repassword)
	// log.Info(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		r.Data["json"] = map[string]interface{}{"code": 0, "messqge": "用户名已经存在"}
		r.ServeJSON()
		return
	}

	// password=utils.MD5(password)
	user := models.User{Id: 0, Username: username, Password: password, Status: 0, Createtime: int64(time.Now().Local().Nanosecond())}
	_, err := models.InsertUser(user)
	if err != nil {
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		r.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	r.ServeJSON()
}

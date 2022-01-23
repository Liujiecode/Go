package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Main2Controller struct {
	beego.Controller
}

func (c *Main2Controller) Get() {
	c.Data["Website"] = "hahaha"
	c.Data["Email"] = "hahahah@gmail.com"
	c.TplName = "index2.tpl"
}

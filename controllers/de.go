package controllers

import (
	"github.com/astaxie/beego"
)

type DeController struct {
	beego.Controller
}

func (c *DeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "bla.tpl"
}

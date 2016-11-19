package controllers

import (
	"github.com/astaxie/beego"
)

type BlaController struct {
	beego.Controller
}

func (c *BlaController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "bla.tpl"
}

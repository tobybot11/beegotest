package routers

import (
	"github.com/astaxie/beego"
	"github.com/tobybot11/beegotest/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/bla", &controllers.BlaController{})
	beego.Router("/de", &controllers.DeController{})
	beego.SetStaticPath("/down1", "download1")
	beego.SetStaticPath("/down2", "download2")
}

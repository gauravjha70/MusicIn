package routers

import (
	"MusicIn/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get,post:Get")
}

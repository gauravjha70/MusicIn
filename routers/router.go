package routers

import (
	"MusicIn/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get,post:Get")

	beego.Router("/signup", &controllers.MainController{}, "get:SignUp")
	beego.Router("/login", &controllers.MainController{}, "get:Login")
	beego.Router("/verify-user", &controllers.MainController{}, "post:VerifyUser")
	beego.Router("/create-user", &controllers.MainController{}, "post:CreateUser")

}

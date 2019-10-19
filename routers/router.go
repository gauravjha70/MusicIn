package routers

import (
	"MusicIn/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Login")                  //Login and Dashboard Page
	beego.Router("/signup", &controllers.MainController{}, "get:SignUp")           //SignUp Page
	beego.Router("/logout", &controllers.MainController{}, "get:Logout")           //Logout user
	beego.Router("/verify-user", &controllers.MainController{}, "post:VerifyUser") //Verify a user using login credentials
	beego.Router("/create-user", &controllers.MainController{}, "post:CreateUser") //Create a new user
	beego.Router("/home", &controllers.MainController{}, "get:HomePage")           //Home page after login

}

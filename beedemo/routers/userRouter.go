package routers

import (
	"beedemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &controllers.UserController{}, "get:Index")
	beego.Router("/user/list", &controllers.UserController{}, "get:List")
	beego.Router("/user/add", &controllers.UserController{}, "post:AddUser")
	beego.Router("/user/edit", &controllers.UserController{}, "put:EditUser")
	beego.Router("/user/delete", &controllers.UserController{}, "delete:DeleteUser")
}

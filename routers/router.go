package routers

import (
    "doolp-web/controllers"
    "github.com/astaxie/beego"
)

func init() {
    //注册 路由
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/console", &controllers.ConsoleController{})
}

package controllers

import (
    "github.com/astaxie/beego"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    this.TplNames = "login.html"
}

func (this *LoginController) Post() {
    uname := this.Input().Get("uname")
    pwd := this.Input().Get("pwd")
    autologin := this.Input().Get("autologin") == "on"

    if "mick" == uname && "mick" == pwd {
        maxAge := 0
        if autologin {
            maxAge = 1<<31 - 1
        }
        this.Ctx.SetCookie("uname", uname, maxAge, "/")
        this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
        this.Redirect("/console", 301)
        this.Data["IsLogin"] = true
        return
    } else {
        this.Redirect("/error", 301)
        return
    }

}

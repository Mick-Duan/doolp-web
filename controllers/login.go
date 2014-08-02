package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

const (
    Secret = "dsfadffasdfjdsajfsdjfa;sjdfsdjf"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    isExit := this.Input().Get("exit") == "true"
    if isExit {
        this.Ctx.SetCookie("uname", "", -1, "/")
        this.Ctx.SetCookie("pwd", "", -1, "/")
        this.Redirect("/", 301)
        return
    }
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
        this.Ctx.SetSecureCookie(Secret, "uname", uname, maxAge, "/")
        this.Ctx.SetSecureCookie(Secret, "pwd", pwd, maxAge, "/")

    }
    this.Redirect("/", 301)
    return
}

func checkAccount(ctx *context.Context) bool {
    ck, ok := ctx.GetSecureCookie(Secret, "uname")
    if !ok {
        return false
    }
    uname := ck

    ck, ok = ctx.GetSecureCookie(Secret, "pwd")
    if !ok {
        return false
    }
    pwd := ck

    return "mick" == uname && "mick" == pwd
}

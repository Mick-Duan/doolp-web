package controllers

import (
    "doolp-web/models"
    "github.com/astaxie/beego"
)

type ConsoleController struct {
    beego.Controller
}

func (this *ConsoleController) Get() {
    if !checkAccount(this.Ctx) {
        this.Redirect("/login", 301)
    }
    this.Data["IsConsole"] = true

    op := this.Input().Get("op")

    switch op {
    case "add":
        sname := this.Input().Get("sname")
        if len(sname) == 0 {
            break
        }

        err := models.AddServer(sname)
        if err != nil {
            beego.Error(err)
        }

        this.Redirect("/console", 301)
        return

    case "del":
        id := this.Input().Get("id")
        if len(id) == 0 {
            break
        }
        return
    }
    this.TplNames = "console.html"
    var err error
    this.Data["Servers"], err = models.GetAllServerNmaes()

    if err != nil {
        beego.Error(err)
    }
}

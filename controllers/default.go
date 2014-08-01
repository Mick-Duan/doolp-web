package controllers

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Data["Website"] = "doolp.com"
    this.Data["Email"] = "reidxy@gmail.com"
    this.TplNames = "index.tpl"
}

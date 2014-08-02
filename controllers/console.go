package controllers

import (
    "github.com/astaxie/beego"
)

type ConsoleController struct {
    beego.Controller
}

func (this *ConsoleController) Get() {
    this.TplNames = "console.html"
}

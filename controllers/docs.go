package controllers

import (
    "github.com/astaxie/beego"
)

type DocsController struct {
    beego.Controller
}

func (this *DocsController) Get() {
    this.Data["IsDocs"] = true
    this.TplNames = "docs.html"
}

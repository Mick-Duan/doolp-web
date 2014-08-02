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
    this.Data["TrueCond"] = true
    this.Data["FalseCond"] = false

    type u struct {
        Name string
        Age  int
        Sex  string
    }

    user := &u{
        Name: "joe",
        Age:  20,
        Sex:  "male",
    }

    this.Data["User"] = user

    nums := []int{1, 3, 3, 4, 5, 6}
    this.Data["Number"] = nums

    this.Data["TplVar"] = "hi hello"

    this.Data["Html"] = "<div> hello mick</div>"

    this.Data["Pipe"] = "<div> hello mick 2 </div>"
}

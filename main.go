package main

import (
    _ "doolp-web/routers"
    "github.com/astaxie/beego"
)

func main() {
    beego.StaticDir["/static"] = "static"
    beego.Run()
}

package main

import (
    _ "doolp-web/routers"
    "github.com/astaxie/beego"
)

func main() {
    StaticDir["/static"] = "static"
    beego.Run()
}

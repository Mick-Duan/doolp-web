package main

import (
    "doolp-web/models"
    _ "doolp-web/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

func init() {
    //register DB
    models.RegisterDB()
}

func main() {
    orm.Debug = true
    orm.RunSyncdb("default", false, true)
    beego.StaticDir["/static"] = "static"
    beego.Run()
}

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
    // 开启 ORM 请求测试
    orm.Debug = true
    // 自动建表
    orm.RunSyncdb("default", false, true)

    //加载静态目录
    beego.StaticDir["/static"] = "static"

    beego.Run()
}

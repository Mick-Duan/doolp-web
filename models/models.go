package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

const (
    _DBAlias  = "default"
    _DBDriver = "mysql"
    _DBConn   = "doolp:PASSWORD@tcp(127.0.0.1:3306)/doolpweb?charset=utf8"
)

type Category struct {
    Id              int64
    Title           string
    Created         time.Time `orm:"index"`
    Views           int64     `orm:"index"`
    TopicTime       time.Time `orm:"index"`
    TopicCount      int64
    TopicLastUserID int64
}

type Topic struct {
    Id              int64
    Uid             int64
    Title           string
    Content         string `orm:"size(5000)"`
    Attachment      string
    Created         time.Time `orm:"index"`
    Updated         time.Time `orm:"index"`
    Views           int64     `orm:"index"`
    Author          string
    ReplyTime       time.Time `orm:"index"`
    ReplyCount      int64
    ReplyLastUserId int64
}

func RegisterDB() {
    orm.RegisterModel(new(Category), new(Topic))
}

func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)
    // 参数1        数据库的别名，用来在ORM中切换数据库使用
    // 参数2        driverName
    // 参数3        对应的链接字符串
    // 参数4(可选)  设置最大空闲连接
    // 参数5(可选)  设置最大数据库连接 (go >= 1.2)
    // doc: https://github.com/go-sql-driver/mysql
    maxIdle := 30
    maxConn := 30
    orm.RegisterDataBase(_DBAlias, _DBDriver, _DBConn, maxIdle, maxConn)
}

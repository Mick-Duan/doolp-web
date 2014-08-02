package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

const (
    _DB_NAME        = "data/doolp.db"
    _SQLITE3_DRIVER = "sqlite3"
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

}

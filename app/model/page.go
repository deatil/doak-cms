package model

// 单页
type Page struct {
    Id          int64  `xorm:"int(11)"`
    UserId      int64  `xorm:"int(11)"`
    Slug        string `xorm:"varchar(150)"`
    Title       string `xorm:"varchar(200)"`
    Keywords    string `xorm:"varchar(100)"`
    Description string `xorm:"varchar(200)"`
    Content     string `xorm:"longtext"`
    Tpl         string `xorm:"varchar(200)"`
    Status      int    `xorm:"tinyint(1)"`
    AddTime     int64  `xorm:"created"`
    AddIp       string `xorm:"varchar(50)"`
}

package model

// 文章
type Art struct {
    Id          int64  `xorm:"int(11)"`
    Uuid        string `xorm:"char(32)"`
    CateId      int64  `xorm:"int(11)"`
    UserId      int64  `xorm:"int(11)"`
    Cover       string `xorm:"char(32)"`
    Title       string `xorm:"varchar(150)"`
    Keywords    string `xorm:"varchar(100)"`
    Description string `xorm:"varchar(200)"`
    Content     string `xorm:"longtext"`
    Tags        string `xorm:"varchar(250)"`
    From        string `xorm:"varchar(200)"`
    IsTop       int    `xorm:"tinyint(1)"`
    Status      int    `xorm:"tinyint(1)"`
    AddTime     int64  `xorm:"created"`
    AddIp       string `xorm:"varchar(50)"`
}

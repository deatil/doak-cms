package model

// 文章
type Art struct {
    Id int64 `xorm:"char(32)"`
    Title string `xorm:"varchar(150)"`
    Keywords string `xorm:"varchar(100)"`
    Description string `xorm:"varchar(200)"`
    Content string `xorm:"longtext"`
    Tags string `xorm:"varchar(250)"`
    UserId string `xorm:"int(11)"`
    From string `xorm:"varchar(200)"`
    Status string `xorm:"tinyint(1)"`
    AddTime int64 `xorm:"created"`
    AddIp string `xorm:"varchar(50)"`
}

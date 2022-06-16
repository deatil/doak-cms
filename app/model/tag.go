package model

// 标签
type Tag struct {
    Id int64 `xorm:"int(11)"`
    Name string `xorm:"varchar(50)"`
    Desc string `xorm:"varchar(200)"`
    AddTime int64 `xorm:"created"`
    AddIp string `xorm:"varchar(50)"`
}

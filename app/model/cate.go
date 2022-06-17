package model

// 分类
type Cate struct {
    Id int64 `xorm:"int(11)"`
    Pid int64 `xorm:"int(11)"`
    Name string `xorm:"varchar(50)"`
    Slug string `xorm:"varchar(50)"`
    Desc string `xorm:"varchar(200)"`
    Status int `xorm:"tinyint(1)"`
    AddTime int64 `xorm:"created"`
    AddIp string `xorm:"varchar(50)"`
}

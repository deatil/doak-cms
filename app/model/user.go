package model

// 用户
type User struct {
    Id       int64 `xorm:"int(11)"`
    Username string `xorm:"varchar(50)"`
    Password string `xorm:"char(32)"`
    Nickname string `xorm:"varchar(100)"`
    Sign     string `xorm:"varchar(200)"`
    Status   int `xorm:"tinyint(1)"`
    AddTime  int64 `xorm:"created"`
    AddIp    string `xorm:"varchar(50)"`
}

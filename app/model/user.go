package model

// 用户
type User struct {
    Id       int64  `xorm:"int(10)" json:"id"`
    Username string `xorm:"varchar(50)" json:"username"`
    Password string `xorm:"char(32)" json:"password"`
    Nickname string `xorm:"varchar(100)" json:"nickname"`
    Avatar   string `xorm:"char(32)" json:"avatar"`
    Sign     string `xorm:"varchar(200)" json:"sign"`
    Status   int    `xorm:"tinyint(1)" json:"status"`
    AddTime  int64  `xorm:"created" json:"add_time"`
    AddIp    string `xorm:"varchar(50)" json:"add_ip"`
}

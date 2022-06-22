package model

// 标签
type Tag struct {
    Id      int64  `xorm:"int(10)" json:"id"`
    Name    string `xorm:"varchar(50)" json:"name"`
    Desc    string `xorm:"varchar(200)" json:"desc"`
    Sort    int    `xorm:"int(5)" json:"sort"`
    Status  int    `xorm:"tinyint(1)" json:"status"`
    AddTime int64  `xorm:"created" json:"add_time"`
    AddIp   string `xorm:"varchar(50)" json:"add_ip"`
}

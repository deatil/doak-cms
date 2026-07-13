package model

// 附件
type Attach struct {
    Id      string `xorm:"char(32)" json:"id"`
    Name    string `xorm:"text" json:"name"`
    Path    string `xorm:"varchar(250)" json:"path"`
    Ext     string `xorm:"varchar(10)" json:"ext"`
    Size    int64  `xorm:"int(10)" json:"size"`
    Md5     string `xorm:"char(32)" json:"md5"`
    Status  int    `xorm:"tinyint(1)" json:"status"`
    AddTime int64  `xorm:"created" json:"add_time"`
    AddIp   string `xorm:"varchar(50)" json:"add_ip"`
}

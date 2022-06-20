package model

// 附件
type Attach struct {
    Id      string `xorm:"char(32)"`
    Name    string `xorm:"text"`
    Path    string `xorm:"varchar(250)"`
    Ext     string `xorm:"varchar(10)"`
    Size    int64  `xorm:"int(10)"`
    Md5     string `xorm:"char(32)"`
    Status  int    `xorm:"tinyint(1)"`
    AddTime int64  `xorm:"created"`
    AddIp   string `xorm:"varchar(50)"`
}

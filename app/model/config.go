package model

// 配置
type Config struct {
    Id    int64  `xorm:"int(11)"`
    Key   string `xorm:"varchar(100)"`
    Value string `xorm:"text"`
    Desc  string `xorm:"varchar(200)"`
}

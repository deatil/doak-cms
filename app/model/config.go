package model

// 配置
type Config struct {
    Id    int64  `xorm:"int(10)" json:"id"`
    Key   string `xorm:"varchar(100)" json:"key"`
    Value string `xorm:"text" json:"value"`
    Desc  string `xorm:"varchar(200)" json:"desc"`
}

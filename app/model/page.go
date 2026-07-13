package model

// 单页
type Page struct {
    Id          int64  `xorm:"int(10)" json:"id"`
    UserId      int64  `xorm:"int(10)" json:"user_id"`
    Slug        string `xorm:"varchar(150)" json:"slug"`
    Title       string `xorm:"varchar(200)" json:"title"`
    Keywords    string `xorm:"varchar(100)" json:"keywords"`
    Description string `xorm:"varchar(200)" json:"description"`
    Content     string `xorm:"longtext" json:"content"`
    Tpl         string `xorm:"varchar(200)" json:"tpl"`
    Status      int    `xorm:"tinyint(1)" json:"status"`
    AddTime     int64  `xorm:"created" json:"add_time"`
    AddIp       string `xorm:"varchar(50)" json:"add_ip"`
}
